package googledataformrepositoryworkflowconfig


type GoogleDataformRepositoryWorkflowConfigInvocationConfigIncludedTargets struct {
	// The action's database (Google Cloud project ID).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_dataform_repository_workflow_config#database GoogleDataformRepositoryWorkflowConfig#database}
	Database *string `field:"optional" json:"database" yaml:"database"`
	// The action's name, within database and schema.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_dataform_repository_workflow_config#name GoogleDataformRepositoryWorkflowConfig#name}
	Name *string `field:"optional" json:"name" yaml:"name"`
	// The action's schema (BigQuery dataset ID), within database.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_dataform_repository_workflow_config#schema GoogleDataformRepositoryWorkflowConfig#schema}
	Schema *string `field:"optional" json:"schema" yaml:"schema"`
}

