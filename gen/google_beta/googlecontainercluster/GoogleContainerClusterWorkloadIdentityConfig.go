package googlecontainercluster


type GoogleContainerClusterWorkloadIdentityConfig struct {
	// The workload pool to attach all Kubernetes service accounts to.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_container_cluster#workload_pool GoogleContainerCluster#workload_pool}
	WorkloadPool *string `field:"optional" json:"workloadPool" yaml:"workloadPool"`
}

