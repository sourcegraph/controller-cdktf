package googlesaasruntimeunitkind


type GoogleSaasRuntimeUnitKindOutputVariableMappingsTo struct {
	// Alias of the dependency that the inputVariable will pass its value to.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_saas_runtime_unit_kind#dependency GoogleSaasRuntimeUnitKind#dependency}
	Dependency *string `field:"required" json:"dependency" yaml:"dependency"`
	// Name of the inputVariable on the dependency.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_saas_runtime_unit_kind#input_variable GoogleSaasRuntimeUnitKind#input_variable}
	InputVariable *string `field:"required" json:"inputVariable" yaml:"inputVariable"`
	// Tells App Lifecycle Manager if this mapping should be used during lookup or not.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_saas_runtime_unit_kind#ignore_for_lookup GoogleSaasRuntimeUnitKind#ignore_for_lookup}
	IgnoreForLookup interface{} `field:"optional" json:"ignoreForLookup" yaml:"ignoreForLookup"`
}

