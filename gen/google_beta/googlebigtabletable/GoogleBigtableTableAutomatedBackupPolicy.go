package googlebigtabletable


type GoogleBigtableTableAutomatedBackupPolicy struct {
	// How frequently automated backups should occur.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_bigtable_table#frequency GoogleBigtableTable#frequency}
	Frequency *string `field:"optional" json:"frequency" yaml:"frequency"`
	// How long the automated backups should be retained.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_bigtable_table#retention_period GoogleBigtableTable#retention_period}
	RetentionPeriod *string `field:"optional" json:"retentionPeriod" yaml:"retentionPeriod"`
}

