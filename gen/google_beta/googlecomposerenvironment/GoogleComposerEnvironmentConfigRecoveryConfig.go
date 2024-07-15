package googlecomposerenvironment


type GoogleComposerEnvironmentConfigRecoveryConfig struct {
	// scheduled_snapshots_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_composer_environment#scheduled_snapshots_config GoogleComposerEnvironment#scheduled_snapshots_config}
	ScheduledSnapshotsConfig *GoogleComposerEnvironmentConfigRecoveryConfigScheduledSnapshotsConfig `field:"optional" json:"scheduledSnapshotsConfig" yaml:"scheduledSnapshotsConfig"`
}

