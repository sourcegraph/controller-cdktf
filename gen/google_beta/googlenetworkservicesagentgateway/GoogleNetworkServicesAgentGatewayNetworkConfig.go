package googlenetworkservicesagentgateway


type GoogleNetworkServicesAgentGatewayNetworkConfig struct {
	// egress block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_network_services_agent_gateway#egress GoogleNetworkServicesAgentGateway#egress}
	Egress *GoogleNetworkServicesAgentGatewayNetworkConfigEgress `field:"required" json:"egress" yaml:"egress"`
}

