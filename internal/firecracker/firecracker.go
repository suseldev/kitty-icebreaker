package firecracker

import (
	"fmt"
	"os"
	"os/exec"
	"time"
)

const socketPath = "/tmp/kittyice-firecracker.sock"
const firecrackerBinary = "build/firecracker"

func cleanup(cmd *exec.Cmd) {
	if cmd.Process != nil && cmd.ProcessState == nil {
		_ = cmd.Process.Kill()
	}
	_ = os.Remove(socketPath)
}

func waitForSocket(path string, timeout time.Duration) error {
	deadline := time.Now().Add(timeout)
	for time.Now().Before(deadline) {
		if _, err := os.Stat(path); err == nil {
			return nil
		}
		time.Sleep(50 * time.Millisecond)
	}
	return fmt.Errorf("wait for Firecracker socket: timeout")
}

func Run() error {
	if _, err := os.Stat(socketPath); err == nil {
		//lint:ignore ST1005 Firecracker should preserve capitalization
		return fmt.Errorf("Firecracker socket already exists: %s", socketPath)
	} else if !os.IsNotExist(err) {
		return fmt.Errorf("could not check Firecracker socket: %w", err)
	}

	cmd := exec.Command(
		firecrackerBinary,
		"--api-sock",
		socketPath,
	)

	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	fmt.Println("kittyice: starting Firecracker")
	if err := cmd.Start(); err != nil {
		return fmt.Errorf("could not start firecracker: %w", err)
	}

	defer cleanup(cmd)

	if err := waitForSocket(socketPath, 10*time.Second); err != nil {
		return fmt.Errorf("could not start Firecracker API: %w", err)
	}

	fmt.Println("kittyice: Firecracker API ready")

	fc := newClient(socketPath)

	// Configure boot source
	// TODO: Make paths configurable
	bs := bootSource{
		KernelImagePath: "build/vmlinux-6.18.48",
		InitrdPath:      "build/initramfs.cpio.gz",
		BootArgs:        "console=ttyS0 reboot=k panic=1 pci=off",
	}

	if err := fc.put("/boot-source", bs); err != nil {
		return fmt.Errorf("could not configure boot source: %w", err)
	}

	fmt.Println("kittyice: boot source configured")

	// Configure machine
	mc := machineConfig{
		CPUCount: 1,
		MemSize:  128,
	}

	if err := fc.put("/machine-config", mc); err != nil {
		return fmt.Errorf("could not configure machine: %w", err)
	}

	fmt.Println("kittyice: machine configured")

	if err := fc.put("/actions", action{
		ActionType: "InstanceStart",
	}); err != nil {
		return fmt.Errorf("could not start instance: %w", err)
	}

	fmt.Println("kittyice: instance started")

	if err := cmd.Wait(); err != nil {
		//lint:ignore ST1005 Firecracker should preserve capitalization
		return fmt.Errorf("Firecracker exited: %w", err)
	}

	return nil
}
