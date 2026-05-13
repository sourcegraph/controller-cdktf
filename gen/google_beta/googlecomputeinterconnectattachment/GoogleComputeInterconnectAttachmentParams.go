package googlecomputeinterconnectattachment


type GoogleComputeInterconnectAttachmentParams struct {
	// Resource manager tags to be bound to the interconnect attachment.
	//
	// Tag keys and values have the
	// same definition as resource manager tags. Keys must be in the format tagKeys/{tag_key_id},
	// and values are in the format tagValues/456.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_compute_interconnect_attachment#resource_manager_tags GoogleComputeInterconnectAttachment#resource_manager_tags}
	ResourceManagerTags *map[string]*string `field:"optional" json:"resourceManagerTags" yaml:"resourceManagerTags"`
}

