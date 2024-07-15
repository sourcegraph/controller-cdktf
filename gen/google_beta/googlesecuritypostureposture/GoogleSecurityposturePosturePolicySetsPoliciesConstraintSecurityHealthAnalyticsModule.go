package googlesecuritypostureposture


type GoogleSecurityposturePosturePolicySetsPoliciesConstraintSecurityHealthAnalyticsModule struct {
	// The name of the module eg: BIGQUERY_TABLE_CMEK_DISABLED.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_securityposture_posture#module_name GoogleSecurityposturePosture#module_name}
	ModuleName *string `field:"required" json:"moduleName" yaml:"moduleName"`
	// The state of enablement for the module at its level of the resource hierarchy. Possible values: ["ENABLEMENT_STATE_UNSPECIFIED", "ENABLED", "DISABLED"].
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_securityposture_posture#module_enablement_state GoogleSecurityposturePosture#module_enablement_state}
	ModuleEnablementState *string `field:"optional" json:"moduleEnablementState" yaml:"moduleEnablementState"`
}

