package googleaccesscontextmanagerserviceperimeters


type GoogleAccessContextManagerServicePerimetersServicePerimetersSpecEgressPolicies struct {
	// egress_from block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.30.0/docs/resources/google_access_context_manager_service_perimeters#egress_from GoogleAccessContextManagerServicePerimeters#egress_from}
	EgressFrom *GoogleAccessContextManagerServicePerimetersServicePerimetersSpecEgressPoliciesEgressFrom `field:"optional" json:"egressFrom" yaml:"egressFrom"`
	// egress_to block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.30.0/docs/resources/google_access_context_manager_service_perimeters#egress_to GoogleAccessContextManagerServicePerimeters#egress_to}
	EgressTo *GoogleAccessContextManagerServicePerimetersServicePerimetersSpecEgressPoliciesEgressTo `field:"optional" json:"egressTo" yaml:"egressTo"`
	// Human readable title. Must be unique within the perimeter. Does not affect behavior.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.30.0/docs/resources/google_access_context_manager_service_perimeters#title GoogleAccessContextManagerServicePerimeters#title}
	Title *string `field:"optional" json:"title" yaml:"title"`
}

