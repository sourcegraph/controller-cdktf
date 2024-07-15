package googleaccesscontextmanagerserviceperimeters


type GoogleAccessContextManagerServicePerimetersServicePerimetersStatusEgressPoliciesEgressFromSources struct {
	// An AccessLevel resource name that allows resources outside the ServicePerimeter to be accessed from the inside.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_access_context_manager_service_perimeters#access_level GoogleAccessContextManagerServicePerimeters#access_level}
	AccessLevel *string `field:"optional" json:"accessLevel" yaml:"accessLevel"`
}

