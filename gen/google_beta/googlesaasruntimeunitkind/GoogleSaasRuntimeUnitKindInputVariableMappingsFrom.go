package googlesaasruntimeunitkind


type GoogleSaasRuntimeUnitKindInputVariableMappingsFrom struct {
	// Alias of the dependency that the outputVariable will pass its value to.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_saas_runtime_unit_kind#dependency GoogleSaasRuntimeUnitKind#dependency}
	Dependency *string `field:"required" json:"dependency" yaml:"dependency"`
	// Name of the outputVariable on the dependency.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_saas_runtime_unit_kind#output_variable GoogleSaasRuntimeUnitKind#output_variable}
	OutputVariable *string `field:"required" json:"outputVariable" yaml:"outputVariable"`
}

