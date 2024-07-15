package googlecomputenodetemplate


type GoogleComputeNodeTemplateNodeTypeFlexibility struct {
	// Number of virtual CPUs to use.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_compute_node_template#cpus GoogleComputeNodeTemplate#cpus}
	Cpus *string `field:"optional" json:"cpus" yaml:"cpus"`
	// Physical memory available to the node, defined in MB.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_compute_node_template#memory GoogleComputeNodeTemplate#memory}
	Memory *string `field:"optional" json:"memory" yaml:"memory"`
}

