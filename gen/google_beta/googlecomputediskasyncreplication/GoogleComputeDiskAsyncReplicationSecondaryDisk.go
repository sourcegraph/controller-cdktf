package googlecomputediskasyncreplication


type GoogleComputeDiskAsyncReplicationSecondaryDisk struct {
	// Secondary disk for asynchronous replication.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_compute_disk_async_replication#disk GoogleComputeDiskAsyncReplication#disk}
	Disk *string `field:"required" json:"disk" yaml:"disk"`
}

