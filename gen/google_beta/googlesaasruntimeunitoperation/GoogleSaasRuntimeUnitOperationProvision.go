package googlesaasruntimeunitoperation


type GoogleSaasRuntimeUnitOperationProvision struct {
	// input_variables block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_saas_runtime_unit_operation#input_variables GoogleSaasRuntimeUnitOperation#input_variables}
	InputVariables interface{} `field:"optional" json:"inputVariables" yaml:"inputVariables"`
	// Reference to the Release object to use for the Unit. (optional).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_saas_runtime_unit_operation#release GoogleSaasRuntimeUnitOperation#release}
	Release *string `field:"optional" json:"release" yaml:"release"`
}

