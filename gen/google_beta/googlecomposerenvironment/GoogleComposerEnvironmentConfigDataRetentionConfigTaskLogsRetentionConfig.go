package googlecomposerenvironment


type GoogleComposerEnvironmentConfigDataRetentionConfigTaskLogsRetentionConfig struct {
	// Whether logs in cloud logging only is enabled or not.
	//
	// This field is supported for Cloud Composer environments in versions composer-2.0.32-airflow-2.1.4 and newer but not in composer-3*
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.32.0/docs/resources/google_composer_environment#storage_mode GoogleComposerEnvironment#storage_mode}
	StorageMode *string `field:"optional" json:"storageMode" yaml:"storageMode"`
}

