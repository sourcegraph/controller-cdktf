package googleaccesscontextmanagerserviceperimeterdryrunegresspolicy


type GoogleAccessContextManagerServicePerimeterDryRunEgressPolicyEgressFromSources struct {
	// An AccessLevel resource name that allows resources outside the ServicePerimeter to be accessed from the inside.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.45.0/docs/resources/google_access_context_manager_service_perimeter_dry_run_egress_policy#access_level GoogleAccessContextManagerServicePerimeterDryRunEgressPolicy#access_level}
	AccessLevel *string `field:"optional" json:"accessLevel" yaml:"accessLevel"`
}

