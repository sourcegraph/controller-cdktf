package googlecomputeinstancegroup


type GoogleComputeInstanceGroupNamedPort struct {
	// The name which the port will be mapped to.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_compute_instance_group#name GoogleComputeInstanceGroup#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// The port number to map the name to.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_compute_instance_group#port GoogleComputeInstanceGroup#port}
	Port *float64 `field:"required" json:"port" yaml:"port"`
}

