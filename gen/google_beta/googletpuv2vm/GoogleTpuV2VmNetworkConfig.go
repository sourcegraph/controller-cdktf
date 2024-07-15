package googletpuv2vm


type GoogleTpuV2VmNetworkConfig struct {
	// Allows the TPU node to send and receive packets with non-matching destination or source IPs.
	//
	// This is required if you plan to use the TPU workers to forward routes.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_tpu_v2_vm#can_ip_forward GoogleTpuV2Vm#can_ip_forward}
	CanIpForward interface{} `field:"optional" json:"canIpForward" yaml:"canIpForward"`
	// Indicates that external IP addresses would be associated with the TPU workers.
	//
	// If set to
	// false, the specified subnetwork or network should have Private Google Access enabled.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_tpu_v2_vm#enable_external_ips GoogleTpuV2Vm#enable_external_ips}
	EnableExternalIps interface{} `field:"optional" json:"enableExternalIps" yaml:"enableExternalIps"`
	// The name of the network for the TPU node.
	//
	// It must be a preexisting Google Compute Engine
	// network. If both network and subnetwork are specified, the given subnetwork must belong
	// to the given network. If network is not specified, it will be looked up from the
	// subnetwork if one is provided, or otherwise use "default".
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_tpu_v2_vm#network GoogleTpuV2Vm#network}
	Network *string `field:"optional" json:"network" yaml:"network"`
	// The name of the subnetwork for the TPU node.
	//
	// It must be a preexisting Google Compute
	// Engine subnetwork. If both network and subnetwork are specified, the given subnetwork
	// must belong to the given network. If subnetwork is not specified, the subnetwork with the
	// same name as the network will be used.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_tpu_v2_vm#subnetwork GoogleTpuV2Vm#subnetwork}
	Subnetwork *string `field:"optional" json:"subnetwork" yaml:"subnetwork"`
}

