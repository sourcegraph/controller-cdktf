package googlecomputeinstancegroupmanager


type GoogleComputeInstanceGroupManagerAllInstancesConfig struct {
	// The label key-value pairs that you want to patch onto the instance,.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_compute_instance_group_manager#labels GoogleComputeInstanceGroupManager#labels}
	Labels *map[string]*string `field:"optional" json:"labels" yaml:"labels"`
	// The metadata key-value pairs that you want to patch onto the instance.
	//
	// For more information, see Project and instance metadata,
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_compute_instance_group_manager#metadata GoogleComputeInstanceGroupManager#metadata}
	Metadata *map[string]*string `field:"optional" json:"metadata" yaml:"metadata"`
}

