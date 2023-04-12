package googlecontainercluster


type GoogleContainerClusterNodeConfigTaint struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_container_cluster#effect GoogleContainerCluster#effect}.
	Effect *string `field:"optional" json:"effect" yaml:"effect"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_container_cluster#key GoogleContainerCluster#key}.
	Key *string `field:"optional" json:"key" yaml:"key"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_container_cluster#value GoogleContainerCluster#value}.
	Value *string `field:"optional" json:"value" yaml:"value"`
}

