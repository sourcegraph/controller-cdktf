package googlecomputeinstance


type GoogleComputeInstanceScratchDisk struct {
	// The disk interface used for attaching this disk. One of SCSI or NVME.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_compute_instance#interface GoogleComputeInstance#interface}
	Interface *string `field:"required" json:"interface" yaml:"interface"`
}

