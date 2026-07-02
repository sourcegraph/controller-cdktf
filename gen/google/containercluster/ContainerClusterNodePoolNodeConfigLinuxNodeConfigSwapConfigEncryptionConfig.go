package containercluster

type ContainerClusterNodePoolNodeConfigLinuxNodeConfigSwapConfigEncryptionConfig struct {
	// If true, swap space will not be encrypted. Defaults to false (encrypted).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.32.0/docs/resources/container_cluster#disabled ContainerCluster#disabled}
	Disabled any `field:"optional" json:"disabled" yaml:"disabled"`
}
