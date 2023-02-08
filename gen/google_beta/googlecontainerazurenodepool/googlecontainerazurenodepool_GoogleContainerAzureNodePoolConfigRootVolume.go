package googlecontainerazurenodepool


type GoogleContainerAzureNodePoolConfigRootVolume struct {
	// Optional.
	//
	// The size of the disk, in GiBs. When unspecified, a default value is provided. See the specific reference in the parent resource.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_container_azure_node_pool#size_gib GoogleContainerAzureNodePool#size_gib}
	SizeGib *float64 `field:"optional" json:"sizeGib" yaml:"sizeGib"`
}

