package googleaccesscontextmanagerserviceperimeter


type GoogleAccessContextManagerServicePerimeterStatusEgressPolicies struct {
	// egress_from block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_access_context_manager_service_perimeter#egress_from GoogleAccessContextManagerServicePerimeter#egress_from}
	EgressFrom *GoogleAccessContextManagerServicePerimeterStatusEgressPoliciesEgressFrom `field:"optional" json:"egressFrom" yaml:"egressFrom"`
	// egress_to block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_access_context_manager_service_perimeter#egress_to GoogleAccessContextManagerServicePerimeter#egress_to}
	EgressTo *GoogleAccessContextManagerServicePerimeterStatusEgressPoliciesEgressTo `field:"optional" json:"egressTo" yaml:"egressTo"`
}

