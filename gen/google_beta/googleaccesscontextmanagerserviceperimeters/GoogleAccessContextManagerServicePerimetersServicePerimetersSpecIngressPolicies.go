package googleaccesscontextmanagerserviceperimeters


type GoogleAccessContextManagerServicePerimetersServicePerimetersSpecIngressPolicies struct {
	// ingress_from block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.30.0/docs/resources/google_access_context_manager_service_perimeters#ingress_from GoogleAccessContextManagerServicePerimeters#ingress_from}
	IngressFrom *GoogleAccessContextManagerServicePerimetersServicePerimetersSpecIngressPoliciesIngressFrom `field:"optional" json:"ingressFrom" yaml:"ingressFrom"`
	// ingress_to block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.30.0/docs/resources/google_access_context_manager_service_perimeters#ingress_to GoogleAccessContextManagerServicePerimeters#ingress_to}
	IngressTo *GoogleAccessContextManagerServicePerimetersServicePerimetersSpecIngressPoliciesIngressTo `field:"optional" json:"ingressTo" yaml:"ingressTo"`
	// Human readable title. Must be unique within the perimeter. Does not affect behavior.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.30.0/docs/resources/google_access_context_manager_service_perimeters#title GoogleAccessContextManagerServicePerimeters#title}
	Title *string `field:"optional" json:"title" yaml:"title"`
}

