package googlenetworkservicesagentgateway


type GoogleNetworkServicesAgentGatewayGoogleManaged struct {
	// Operating Mode of Agent Gateway. Possible values: ["AGENT_TO_ANYWHERE", "CLIENT_TO_AGENT"].
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_network_services_agent_gateway#governed_access_path GoogleNetworkServicesAgentGateway#governed_access_path}
	GovernedAccessPath *string `field:"required" json:"governedAccessPath" yaml:"governedAccessPath"`
}

