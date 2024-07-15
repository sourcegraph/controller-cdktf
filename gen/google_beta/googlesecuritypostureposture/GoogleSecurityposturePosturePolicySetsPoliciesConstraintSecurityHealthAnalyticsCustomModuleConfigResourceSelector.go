package googlesecuritypostureposture


type GoogleSecurityposturePosturePolicySetsPoliciesConstraintSecurityHealthAnalyticsCustomModuleConfigResourceSelector struct {
	// The resource types to run the detector on.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_securityposture_posture#resource_types GoogleSecurityposturePosture#resource_types}
	ResourceTypes *[]*string `field:"required" json:"resourceTypes" yaml:"resourceTypes"`
}

