package googlesecuritypostureposture


type GoogleSecurityposturePosturePolicySetsPoliciesConstraintSecurityHealthAnalyticsCustomModule struct {
	// config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_securityposture_posture#config GoogleSecurityposturePosture#config}
	Config *GoogleSecurityposturePosturePolicySetsPoliciesConstraintSecurityHealthAnalyticsCustomModuleConfig `field:"required" json:"config" yaml:"config"`
	// The display name of the Security Health Analytics custom module.
	//
	// This
	// display name becomes the finding category for all findings that are
	// returned by this custom module.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_securityposture_posture#display_name GoogleSecurityposturePosture#display_name}
	DisplayName *string `field:"optional" json:"displayName" yaml:"displayName"`
	// The state of enablement for the module at its level of the resource hierarchy. Possible values: ["ENABLEMENT_STATE_UNSPECIFIED", "ENABLED", "DISABLED"].
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_securityposture_posture#module_enablement_state GoogleSecurityposturePosture#module_enablement_state}
	ModuleEnablementState *string `field:"optional" json:"moduleEnablementState" yaml:"moduleEnablementState"`
}

