package googlealloydbcluster


type GoogleAlloydbClusterAutomatedBackupPolicy struct {
	// The length of the time window during which a backup can be taken.
	//
	// If a backup does not succeed within this time window, it will be canceled and considered failed.
	//
	// The backup window must be at least 5 minutes long. There is no upper bound on the window. If not set, it will default to 1 hour.
	//
	// A duration in seconds with up to nine fractional digits, terminated by 's'. Example: "3.5s".
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_alloydb_cluster#backup_window GoogleAlloydbCluster#backup_window}
	BackupWindow *string `field:"optional" json:"backupWindow" yaml:"backupWindow"`
	// Whether automated backups are enabled.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_alloydb_cluster#enabled GoogleAlloydbCluster#enabled}
	Enabled interface{} `field:"optional" json:"enabled" yaml:"enabled"`
	// encryption_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_alloydb_cluster#encryption_config GoogleAlloydbCluster#encryption_config}
	EncryptionConfig *GoogleAlloydbClusterAutomatedBackupPolicyEncryptionConfig `field:"optional" json:"encryptionConfig" yaml:"encryptionConfig"`
	// Labels to apply to backups created using this configuration.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_alloydb_cluster#labels GoogleAlloydbCluster#labels}
	Labels *map[string]*string `field:"optional" json:"labels" yaml:"labels"`
	// The location where the backup will be stored.
	//
	// Currently, the only supported option is to store the backup in the same region as the cluster.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_alloydb_cluster#location GoogleAlloydbCluster#location}
	Location *string `field:"optional" json:"location" yaml:"location"`
	// quantity_based_retention block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_alloydb_cluster#quantity_based_retention GoogleAlloydbCluster#quantity_based_retention}
	QuantityBasedRetention *GoogleAlloydbClusterAutomatedBackupPolicyQuantityBasedRetention `field:"optional" json:"quantityBasedRetention" yaml:"quantityBasedRetention"`
	// time_based_retention block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_alloydb_cluster#time_based_retention GoogleAlloydbCluster#time_based_retention}
	TimeBasedRetention *GoogleAlloydbClusterAutomatedBackupPolicyTimeBasedRetention `field:"optional" json:"timeBasedRetention" yaml:"timeBasedRetention"`
	// weekly_schedule block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_alloydb_cluster#weekly_schedule GoogleAlloydbCluster#weekly_schedule}
	WeeklySchedule *GoogleAlloydbClusterAutomatedBackupPolicyWeeklySchedule `field:"optional" json:"weeklySchedule" yaml:"weeklySchedule"`
}

