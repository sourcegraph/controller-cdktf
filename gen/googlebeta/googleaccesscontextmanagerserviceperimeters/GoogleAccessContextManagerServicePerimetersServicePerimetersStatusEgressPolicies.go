package googleaccesscontextmanagerserviceperimeters


type GoogleAccessContextManagerServicePerimetersServicePerimetersStatusEgressPolicies struct {
	// egress_from block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_access_context_manager_service_perimeters#egress_from GoogleAccessContextManagerServicePerimeters#egress_from}
	EgressFrom *GoogleAccessContextManagerServicePerimetersServicePerimetersStatusEgressPoliciesEgressFrom `field:"optional" json:"egressFrom" yaml:"egressFrom"`
	// egress_to block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_access_context_manager_service_perimeters#egress_to GoogleAccessContextManagerServicePerimeters#egress_to}
	EgressTo *GoogleAccessContextManagerServicePerimetersServicePerimetersStatusEgressPoliciesEgressTo `field:"optional" json:"egressTo" yaml:"egressTo"`
}

