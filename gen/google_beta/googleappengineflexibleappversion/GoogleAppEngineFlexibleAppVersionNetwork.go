package googleappengineflexibleappversion


type GoogleAppEngineFlexibleAppVersionNetwork struct {
	// Google Compute Engine network where the virtual machines are created. Specify the short name, not the resource path.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_app_engine_flexible_app_version#name GoogleAppEngineFlexibleAppVersion#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// List of ports, or port pairs, to forward from the virtual machine to the application container.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_app_engine_flexible_app_version#forwarded_ports GoogleAppEngineFlexibleAppVersion#forwarded_ports}
	ForwardedPorts *[]*string `field:"optional" json:"forwardedPorts" yaml:"forwardedPorts"`
	// Prevent instances from receiving an ephemeral external IP address. Possible values: ["EXTERNAL", "INTERNAL"].
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_app_engine_flexible_app_version#instance_ip_mode GoogleAppEngineFlexibleAppVersion#instance_ip_mode}
	InstanceIpMode *string `field:"optional" json:"instanceIpMode" yaml:"instanceIpMode"`
	// Tag to apply to the instance during creation.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_app_engine_flexible_app_version#instance_tag GoogleAppEngineFlexibleAppVersion#instance_tag}
	InstanceTag *string `field:"optional" json:"instanceTag" yaml:"instanceTag"`
	// Enable session affinity.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_app_engine_flexible_app_version#session_affinity GoogleAppEngineFlexibleAppVersion#session_affinity}
	SessionAffinity interface{} `field:"optional" json:"sessionAffinity" yaml:"sessionAffinity"`
	// Google Cloud Platform sub-network where the virtual machines are created. Specify the short name, not the resource path.
	//
	// If the network that the instance is being created in is a Legacy network, then the IP address is allocated from the IPv4Range.
	// If the network that the instance is being created in is an auto Subnet Mode Network, then only network name should be specified (not the subnetworkName) and the IP address is created from the IPCidrRange of the subnetwork that exists in that zone for that network.
	// If the network that the instance is being created in is a custom Subnet Mode Network, then the subnetworkName must be specified and the IP address is created from the IPCidrRange of the subnetwork.
	// If specified, the subnetwork must exist in the same region as the App Engine flexible environment application.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_app_engine_flexible_app_version#subnetwork GoogleAppEngineFlexibleAppVersion#subnetwork}
	Subnetwork *string `field:"optional" json:"subnetwork" yaml:"subnetwork"`
}

