package googlecomputenodetemplate


type GoogleComputeNodeTemplateAccelerators struct {
	// The number of the guest accelerator cards exposed to this node template.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.30.0/docs/resources/google_compute_node_template#accelerator_count GoogleComputeNodeTemplate#accelerator_count}
	AcceleratorCount *float64 `field:"optional" json:"acceleratorCount" yaml:"acceleratorCount"`
	// Full or partial URL of the accelerator type resource to expose to this node template.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.30.0/docs/resources/google_compute_node_template#accelerator_type GoogleComputeNodeTemplate#accelerator_type}
	AcceleratorType *string `field:"optional" json:"acceleratorType" yaml:"acceleratorType"`
}

