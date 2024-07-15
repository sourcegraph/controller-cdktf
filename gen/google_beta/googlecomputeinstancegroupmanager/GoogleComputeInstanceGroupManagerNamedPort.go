package googlecomputeinstancegroupmanager


type GoogleComputeInstanceGroupManagerNamedPort struct {
	// The name of the port.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_compute_instance_group_manager#name GoogleComputeInstanceGroupManager#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// The port number.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_compute_instance_group_manager#port GoogleComputeInstanceGroupManager#port}
	Port *float64 `field:"required" json:"port" yaml:"port"`
}

