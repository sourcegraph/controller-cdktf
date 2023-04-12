package googleaccesscontextmanagerserviceperimeters


type GoogleAccessContextManagerServicePerimetersServicePerimetersSpecIngressPolicies struct {
	// ingress_from block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_access_context_manager_service_perimeters#ingress_from GoogleAccessContextManagerServicePerimeters#ingress_from}
	IngressFrom *GoogleAccessContextManagerServicePerimetersServicePerimetersSpecIngressPoliciesIngressFrom `field:"optional" json:"ingressFrom" yaml:"ingressFrom"`
	// ingress_to block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_access_context_manager_service_perimeters#ingress_to GoogleAccessContextManagerServicePerimeters#ingress_to}
	IngressTo *GoogleAccessContextManagerServicePerimetersServicePerimetersSpecIngressPoliciesIngressTo `field:"optional" json:"ingressTo" yaml:"ingressTo"`
}

