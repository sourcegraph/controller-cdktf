package googlecomposerenvironment


type GoogleComposerEnvironmentConfigRecoveryConfigScheduledSnapshotsConfig struct {
	// When enabled, Cloud Composer periodically saves snapshots of your environment to a Cloud Storage bucket.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_composer_environment#enabled GoogleComposerEnvironment#enabled}
	Enabled interface{} `field:"required" json:"enabled" yaml:"enabled"`
	// Snapshot schedule, in the unix-cron format.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_composer_environment#snapshot_creation_schedule GoogleComposerEnvironment#snapshot_creation_schedule}
	SnapshotCreationSchedule *string `field:"optional" json:"snapshotCreationSchedule" yaml:"snapshotCreationSchedule"`
	// the URI of a bucket folder where to save the snapshot.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_composer_environment#snapshot_location GoogleComposerEnvironment#snapshot_location}
	SnapshotLocation *string `field:"optional" json:"snapshotLocation" yaml:"snapshotLocation"`
	// A time zone for the schedule.
	//
	// This value is a time offset and does not take into account daylight saving time changes. Valid values are from UTC-12 to UTC+12. Examples: UTC, UTC-01, UTC+03.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_composer_environment#time_zone GoogleComposerEnvironment#time_zone}
	TimeZone *string `field:"optional" json:"timeZone" yaml:"timeZone"`
}

