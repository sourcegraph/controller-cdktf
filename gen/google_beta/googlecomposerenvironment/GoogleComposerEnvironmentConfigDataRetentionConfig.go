package googlecomposerenvironment


type GoogleComposerEnvironmentConfigDataRetentionConfig struct {
	// airflow_metadata_retention_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.30.0/docs/resources/google_composer_environment#airflow_metadata_retention_config GoogleComposerEnvironment#airflow_metadata_retention_config}
	AirflowMetadataRetentionConfig interface{} `field:"optional" json:"airflowMetadataRetentionConfig" yaml:"airflowMetadataRetentionConfig"`
	// task_logs_retention_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.30.0/docs/resources/google_composer_environment#task_logs_retention_config GoogleComposerEnvironment#task_logs_retention_config}
	TaskLogsRetentionConfig interface{} `field:"optional" json:"taskLogsRetentionConfig" yaml:"taskLogsRetentionConfig"`
}

