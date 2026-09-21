package firecracker

type bootSource struct {
	KernelImagePath string `json:"kernel_image_path"`
	InitrdPath      string `json:"initrd_path"`
	BootArgs        string `json:"boot_args"`
}

type machineConfig struct {
	CPUCount int `json:"vcpu_count"`
	MemSize  int `json:"mem_size_mib"` // in MB
}

type action struct {
	ActionType string `json:"action_type"`
}
