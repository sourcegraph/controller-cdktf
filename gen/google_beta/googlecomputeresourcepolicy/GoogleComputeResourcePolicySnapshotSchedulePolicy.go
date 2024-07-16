package googlecomputeresourcepolicy


type GoogleComputeResourcePolicySnapshotSchedulePolicy struct {
	// schedule block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_compute_resource_policy#schedule GoogleComputeResourcePolicy#schedule}
	Schedule *GoogleComputeResourcePolicySnapshotSchedulePolicySchedule `field:"required" json:"schedule" yaml:"schedule"`
	// retention_policy block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_compute_resource_policy#retention_policy GoogleComputeResourcePolicy#retention_policy}
	RetentionPolicy *GoogleComputeResourcePolicySnapshotSchedulePolicyRetentionPolicy `field:"optional" json:"retentionPolicy" yaml:"retentionPolicy"`
	// snapshot_properties block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_compute_resource_policy#snapshot_properties GoogleComputeResourcePolicy#snapshot_properties}
	SnapshotProperties *GoogleComputeResourcePolicySnapshotSchedulePolicySnapshotProperties `field:"optional" json:"snapshotProperties" yaml:"snapshotProperties"`
}

