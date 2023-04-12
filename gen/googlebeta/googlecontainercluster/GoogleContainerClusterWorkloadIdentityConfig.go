package googlecontainercluster


type GoogleContainerClusterWorkloadIdentityConfig struct {
	// The workload pool to attach all Kubernetes service accounts to.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_container_cluster#workload_pool GoogleContainerCluster#workload_pool}
	WorkloadPool *string `field:"optional" json:"workloadPool" yaml:"workloadPool"`
}

