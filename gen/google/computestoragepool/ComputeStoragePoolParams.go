package computestoragepool


type ComputeStoragePoolParams struct {
	// Resource manager tags to be bound to the storage pool.
	//
	// Tag keys and values have the
	// same definition as resource manager tags. Keys and values can be either in numeric format,
	// such as tagKeys/{tag_key_id} and tagValues/{tag_value_id} or in namespaced format such as
	// {org_id|projectId}/{tag_key_short_name} and {tag_value_short_name}. The field is ignored when empty.
	// The field is immutable and causes resource replacement when mutated. This field is only
	// set at create time and modifying this field after creation will trigger recreation.
	// To apply tags to an existing resource, see the google_tags_tag_binding resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/compute_storage_pool#resource_manager_tags ComputeStoragePool#resource_manager_tags}
	ResourceManagerTags *map[string]*string `field:"optional" json:"resourceManagerTags" yaml:"resourceManagerTags"`
}

