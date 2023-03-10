package computevpngateway


type ComputeVpnGatewayTimeouts struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/compute_vpn_gateway#create ComputeVpnGateway#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/compute_vpn_gateway#delete ComputeVpnGateway#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
}

