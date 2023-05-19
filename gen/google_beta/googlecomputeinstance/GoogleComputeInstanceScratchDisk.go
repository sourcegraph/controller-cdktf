package googlecomputeinstance


type GoogleComputeInstanceScratchDisk struct {
	// The disk interface used for attaching this disk. One of SCSI or NVME.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/4.65.2/docs/resources/google_compute_instance#interface GoogleComputeInstance#interface}
	Interface *string `field:"required" json:"interface" yaml:"interface"`
	// The size of the disk in gigabytes. One of 375 or 3000.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/4.65.2/docs/resources/google_compute_instance#size GoogleComputeInstance#size}
	Size *float64 `field:"optional" json:"size" yaml:"size"`
}

