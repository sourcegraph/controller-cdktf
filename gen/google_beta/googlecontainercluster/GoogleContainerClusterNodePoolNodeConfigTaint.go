package googlecontainercluster


type GoogleContainerClusterNodePoolNodeConfigTaint struct {
	// Effect for taint.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_container_cluster#effect GoogleContainerCluster#effect}
	Effect *string `field:"required" json:"effect" yaml:"effect"`
	// Key for taint.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_container_cluster#key GoogleContainerCluster#key}
	Key *string `field:"required" json:"key" yaml:"key"`
	// Value for taint.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_container_cluster#value GoogleContainerCluster#value}
	Value *string `field:"required" json:"value" yaml:"value"`
}

