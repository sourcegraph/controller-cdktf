package accesscontextmanagerserviceperimeter


type AccessContextManagerServicePerimeterStatusEgressPoliciesEgressFromSources struct {
	// An AccessLevel resource name that allows resources outside the ServicePerimeter to be accessed from the inside.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/5.38.0/docs/resources/access_context_manager_service_perimeter#access_level AccessContextManagerServicePerimeter#access_level}
	AccessLevel *string `field:"optional" json:"accessLevel" yaml:"accessLevel"`
}

