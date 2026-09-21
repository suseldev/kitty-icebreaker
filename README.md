# Kitty Icebreaker
Kitty Icebreaker is a local workload runner for untrusted or potentially unsafe workloads built around Firecracker microVMs.

The goal is to execute untrusted workloads in isolated, disposable microVMs. 

## Components
- `kittyice` - host-side CLI and runtime
- `kittyinit` - minimal guest init process responsible for launching workloads
- `initramfs` - minimal guest initramfs containing `kittyinit` as `/init`

## Build
`make` - builds the `kittyice` and `kittyinit` components

Built binaries are written to the `build/` directory.

### Initramfs
`make -C system/initramfs`

Builds a minimal initramfs with `kittyinit` as the guest init.

The resulting archive is written to `build/initramfs.cpio.gz`

## Run

Before running Kitty Icebreaker, download a Firecracker binary and a compatible Linux kernel from the official Firecracker CI artifacts (see: [Firecracker Getting Started Guide](https://github.com/firecracker-microvm/firecracker/blob/main/docs/getting-started.md)).

Place the kernel image in the `build/`directory.

Then start the microVM:
`./build/kittyice run`

Kitty Icebreaker starts a Firecracker microVM using the kernel and initramfs from the `build/` directory.

> [NOTE]
> The Firecracker binary, kernel, and initramfs paths are currently hardcoded and are not configurable yet.
