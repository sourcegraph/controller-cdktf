package googlealloydbcluster


type GoogleAlloydbClusterAutomatedBackupPolicyTimeBasedRetention struct {
	// The retention period. A duration in seconds with up to nine fractional digits, terminated by 's'. Example: "3.5s".
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_alloydb_cluster#retention_period GoogleAlloydbCluster#retention_period}
	RetentionPeriod *string `field:"optional" json:"retentionPeriod" yaml:"retentionPeriod"`
}

