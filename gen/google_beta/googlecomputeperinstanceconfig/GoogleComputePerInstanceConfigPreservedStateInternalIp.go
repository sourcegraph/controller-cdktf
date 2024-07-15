package googlecomputeperinstanceconfig


type GoogleComputePerInstanceConfigPreservedStateInternalIp struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_compute_per_instance_config#interface_name GoogleComputePerInstanceConfig#interface_name}.
	InterfaceName *string `field:"required" json:"interfaceName" yaml:"interfaceName"`
	// These stateful IPs will never be released during autohealing, update or VM instance recreate operations.
	//
	// This flag is used to configure if the IP reservation should be deleted after it is no longer used by the group, e.g. when the given instance or the whole group is deleted. Default value: "NEVER" Possible values: ["NEVER", "ON_PERMANENT_INSTANCE_DELETION"]
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_compute_per_instance_config#auto_delete GoogleComputePerInstanceConfig#auto_delete}
	AutoDelete *string `field:"optional" json:"autoDelete" yaml:"autoDelete"`
	// ip_address block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_compute_per_instance_config#ip_address GoogleComputePerInstanceConfig#ip_address}
	IpAddress *GoogleComputePerInstanceConfigPreservedStateInternalIpIpAddress `field:"optional" json:"ipAddress" yaml:"ipAddress"`
}

