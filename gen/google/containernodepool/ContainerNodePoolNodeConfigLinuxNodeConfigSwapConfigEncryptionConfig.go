package containernodepool

type ContainerNodePoolNodeConfigLinuxNodeConfigSwapConfigEncryptionConfig struct {
	// If true, swap space will not be encrypted. Defaults to false (encrypted).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.32.0/docs/resources/container_node_pool#disabled ContainerNodePool#disabled}
	Disabled any `field:"optional" json:"disabled" yaml:"disabled"`
}
