package googlecomputedisk


type GoogleComputeDiskAsyncPrimaryDisk struct {
	// Primary disk for asynchronous disk replication.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_compute_disk#disk GoogleComputeDisk#disk}
	Disk *string `field:"required" json:"disk" yaml:"disk"`
}

