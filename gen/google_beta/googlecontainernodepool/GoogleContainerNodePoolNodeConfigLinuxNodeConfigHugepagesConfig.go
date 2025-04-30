package googlecontainernodepool


type GoogleContainerNodePoolNodeConfigLinuxNodeConfigHugepagesConfig struct {
	// Amount of 1G hugepages.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.32.0/docs/resources/google_container_node_pool#hugepage_size_1g GoogleContainerNodePool#hugepage_size_1g}
	HugepageSize1G *float64 `field:"optional" json:"hugepageSize1G" yaml:"hugepageSize1G"`
	// Amount of 2M hugepages.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.32.0/docs/resources/google_container_node_pool#hugepage_size_2m GoogleContainerNodePool#hugepage_size_2m}
	HugepageSize2M *float64 `field:"optional" json:"hugepageSize2M" yaml:"hugepageSize2M"`
}

