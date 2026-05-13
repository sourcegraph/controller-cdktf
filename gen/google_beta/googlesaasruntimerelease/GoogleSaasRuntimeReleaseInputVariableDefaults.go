package googlesaasruntimerelease


type GoogleSaasRuntimeReleaseInputVariableDefaults struct {
	// Name of the variable from actuation configs.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_saas_runtime_release#variable GoogleSaasRuntimeRelease#variable}
	Variable *string `field:"required" json:"variable" yaml:"variable"`
	// Name of a supported variable type. Supported types are STRING, INT, BOOL. Possible values: ["TYPE_UNSPECIFIED", "STRING", "INT", "BOOL"].
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_saas_runtime_release#type GoogleSaasRuntimeRelease#type}
	Type *string `field:"optional" json:"type" yaml:"type"`
	// String encoded value for the variable.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_saas_runtime_release#value GoogleSaasRuntimeRelease#value}
	Value *string `field:"optional" json:"value" yaml:"value"`
}

