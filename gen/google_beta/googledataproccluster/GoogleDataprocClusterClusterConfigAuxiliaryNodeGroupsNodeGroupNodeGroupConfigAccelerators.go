package googledataproccluster


type GoogleDataprocClusterClusterConfigAuxiliaryNodeGroupsNodeGroupNodeGroupConfigAccelerators struct {
	// The number of the accelerator cards of this type exposed to this instance.
	//
	// Often restricted to one of 1, 2, 4, or 8.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_dataproc_cluster#accelerator_count GoogleDataprocCluster#accelerator_count}
	AcceleratorCount *float64 `field:"required" json:"acceleratorCount" yaml:"acceleratorCount"`
	// The short name of the accelerator type to expose to this instance. For example, nvidia-tesla-k80.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_dataproc_cluster#accelerator_type GoogleDataprocCluster#accelerator_type}
	AcceleratorType *string `field:"required" json:"acceleratorType" yaml:"acceleratorType"`
}

