package googlecontainercluster


type GoogleContainerClusterNodePoolNodeConfigAdvancedMachineFeatures struct {
	// The number of threads per physical core.
	//
	// To disable simultaneous multithreading (SMT) set this to 1. If unset, the maximum number of threads supported per core by the underlying processor is assumed.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_container_cluster#threads_per_core GoogleContainerCluster#threads_per_core}
	ThreadsPerCore *float64 `field:"required" json:"threadsPerCore" yaml:"threadsPerCore"`
	// Whether the node should have nested virtualization enabled.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_container_cluster#enable_nested_virtualization GoogleContainerCluster#enable_nested_virtualization}
	EnableNestedVirtualization interface{} `field:"optional" json:"enableNestedVirtualization" yaml:"enableNestedVirtualization"`
}

