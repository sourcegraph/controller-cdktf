package googledataformrepositoryworkflowconfig


type GoogleDataformRepositoryWorkflowConfigInvocationConfig struct {
	// Optional. When set to true, any incremental tables will be fully refreshed.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_dataform_repository_workflow_config#fully_refresh_incremental_tables_enabled GoogleDataformRepositoryWorkflowConfig#fully_refresh_incremental_tables_enabled}
	FullyRefreshIncrementalTablesEnabled interface{} `field:"optional" json:"fullyRefreshIncrementalTablesEnabled" yaml:"fullyRefreshIncrementalTablesEnabled"`
	// Optional. The set of tags to include.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_dataform_repository_workflow_config#included_tags GoogleDataformRepositoryWorkflowConfig#included_tags}
	IncludedTags *[]*string `field:"optional" json:"includedTags" yaml:"includedTags"`
	// included_targets block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_dataform_repository_workflow_config#included_targets GoogleDataformRepositoryWorkflowConfig#included_targets}
	IncludedTargets interface{} `field:"optional" json:"includedTargets" yaml:"includedTargets"`
	// Optional. The service account to run workflow invocations under.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_dataform_repository_workflow_config#service_account GoogleDataformRepositoryWorkflowConfig#service_account}
	ServiceAccount *string `field:"optional" json:"serviceAccount" yaml:"serviceAccount"`
	// Optional. When set to true, transitive dependencies of included actions will be executed.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_dataform_repository_workflow_config#transitive_dependencies_included GoogleDataformRepositoryWorkflowConfig#transitive_dependencies_included}
	TransitiveDependenciesIncluded interface{} `field:"optional" json:"transitiveDependenciesIncluded" yaml:"transitiveDependenciesIncluded"`
	// Optional. When set to true, transitive dependents of included actions will be executed.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_dataform_repository_workflow_config#transitive_dependents_included GoogleDataformRepositoryWorkflowConfig#transitive_dependents_included}
	TransitiveDependentsIncluded interface{} `field:"optional" json:"transitiveDependentsIncluded" yaml:"transitiveDependentsIncluded"`
}

