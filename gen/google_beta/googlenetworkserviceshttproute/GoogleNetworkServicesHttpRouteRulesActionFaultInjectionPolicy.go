package googlenetworkserviceshttproute


type GoogleNetworkServicesHttpRouteRulesActionFaultInjectionPolicy struct {
	// abort block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_network_services_http_route#abort GoogleNetworkServicesHttpRoute#abort}
	Abort *GoogleNetworkServicesHttpRouteRulesActionFaultInjectionPolicyAbort `field:"optional" json:"abort" yaml:"abort"`
	// delay block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_network_services_http_route#delay GoogleNetworkServicesHttpRoute#delay}
	Delay *GoogleNetworkServicesHttpRouteRulesActionFaultInjectionPolicyDelay `field:"optional" json:"delay" yaml:"delay"`
}

