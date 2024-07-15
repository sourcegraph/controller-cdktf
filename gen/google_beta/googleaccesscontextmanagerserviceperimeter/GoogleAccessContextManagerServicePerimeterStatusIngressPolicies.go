package googleaccesscontextmanagerserviceperimeter


type GoogleAccessContextManagerServicePerimeterStatusIngressPolicies struct {
	// ingress_from block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_access_context_manager_service_perimeter#ingress_from GoogleAccessContextManagerServicePerimeter#ingress_from}
	IngressFrom *GoogleAccessContextManagerServicePerimeterStatusIngressPoliciesIngressFrom `field:"optional" json:"ingressFrom" yaml:"ingressFrom"`
	// ingress_to block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_access_context_manager_service_perimeter#ingress_to GoogleAccessContextManagerServicePerimeter#ingress_to}
	IngressTo *GoogleAccessContextManagerServicePerimeterStatusIngressPoliciesIngressTo `field:"optional" json:"ingressTo" yaml:"ingressTo"`
}

