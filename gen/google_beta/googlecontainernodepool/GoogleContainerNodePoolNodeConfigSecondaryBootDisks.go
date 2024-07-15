package googlecontainernodepool


type GoogleContainerNodePoolNodeConfigSecondaryBootDisks struct {
	// Disk image to create the secondary boot disk from.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_container_node_pool#disk_image GoogleContainerNodePool#disk_image}
	DiskImage *string `field:"required" json:"diskImage" yaml:"diskImage"`
	// Mode for how the secondary boot disk is used.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_container_node_pool#mode GoogleContainerNodePool#mode}
	Mode *string `field:"optional" json:"mode" yaml:"mode"`
}

