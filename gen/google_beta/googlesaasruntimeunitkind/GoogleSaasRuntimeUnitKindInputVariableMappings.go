package googlesaasruntimeunitkind


type GoogleSaasRuntimeUnitKindInputVariableMappings struct {
	// name of the variable.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_saas_runtime_unit_kind#variable GoogleSaasRuntimeUnitKind#variable}
	Variable *string `field:"required" json:"variable" yaml:"variable"`
	// from block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_saas_runtime_unit_kind#from GoogleSaasRuntimeUnitKind#from}
	From *GoogleSaasRuntimeUnitKindInputVariableMappingsFrom `field:"optional" json:"from" yaml:"from"`
	// to block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_saas_runtime_unit_kind#to GoogleSaasRuntimeUnitKind#to}
	To *GoogleSaasRuntimeUnitKindInputVariableMappingsTo `field:"optional" json:"to" yaml:"to"`
}

