package googlesaasruntimeunitoperation


type GoogleSaasRuntimeUnitOperationUpgradeInputVariables struct {
	// Name of the variable from actuation configs.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_saas_runtime_unit_operation#variable GoogleSaasRuntimeUnitOperation#variable}
	Variable *string `field:"required" json:"variable" yaml:"variable"`
	// Name of a supported variable type. Supported types are string, int, bool. Possible values: STRING INT BOOL.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_saas_runtime_unit_operation#type GoogleSaasRuntimeUnitOperation#type}
	Type *string `field:"optional" json:"type" yaml:"type"`
	// String encoded value for the variable.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_saas_runtime_unit_operation#value GoogleSaasRuntimeUnitOperation#value}
	Value *string `field:"optional" json:"value" yaml:"value"`
}

