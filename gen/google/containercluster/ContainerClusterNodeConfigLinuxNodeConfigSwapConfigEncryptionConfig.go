package containercluster


type ContainerClusterNodeConfigLinuxNodeConfigSwapConfigEncryptionConfig struct {
	// If true, swap space will not be encrypted. Defaults to false (encrypted).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/container_cluster#disabled ContainerCluster#disabled}
	Disabled interface{} `field:"optional" json:"disabled" yaml:"disabled"`
}

