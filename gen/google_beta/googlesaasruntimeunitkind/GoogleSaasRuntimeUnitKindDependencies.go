package googlesaasruntimeunitkind


type GoogleSaasRuntimeUnitKindDependencies struct {
	// An alias for the dependency. Used for input variable mapping.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_saas_runtime_unit_kind#alias GoogleSaasRuntimeUnitKind#alias}
	Alias *string `field:"required" json:"alias" yaml:"alias"`
	// The unit kind of the dependency.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_saas_runtime_unit_kind#unit_kind GoogleSaasRuntimeUnitKind#unit_kind}
	UnitKind *string `field:"required" json:"unitKind" yaml:"unitKind"`
}

