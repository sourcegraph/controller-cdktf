package googlecomposerenvironment


type GoogleComposerEnvironmentConfigDataRetentionConfigAirflowMetadataRetentionConfig struct {
	// How many days data should be retained for.
	//
	// This field is supported for Cloud Composer environments in composer 3 and newer.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.32.0/docs/resources/google_composer_environment#retention_days GoogleComposerEnvironment#retention_days}
	RetentionDays *float64 `field:"optional" json:"retentionDays" yaml:"retentionDays"`
	// Whether database retention is enabled or not.
	//
	// This field is supported for Cloud Composer environments in composer 3 and newer.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.32.0/docs/resources/google_composer_environment#retention_mode GoogleComposerEnvironment#retention_mode}
	RetentionMode *string `field:"optional" json:"retentionMode" yaml:"retentionMode"`
}

