package googlesqldatabaseinstance


type GoogleSqlDatabaseInstanceSettingsSqlServerAuditConfig struct {
	// The name of the destination bucket (e.g., gs://mybucket).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_sql_database_instance#bucket GoogleSqlDatabaseInstance#bucket}
	Bucket *string `field:"optional" json:"bucket" yaml:"bucket"`
	// How long to keep generated audit files.
	//
	// A duration in seconds with up to nine fractional digits, terminated by 's'. Example: "3.5s"..
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_sql_database_instance#retention_interval GoogleSqlDatabaseInstance#retention_interval}
	RetentionInterval *string `field:"optional" json:"retentionInterval" yaml:"retentionInterval"`
	// How often to upload generated audit files.
	//
	// A duration in seconds with up to nine fractional digits, terminated by 's'. Example: "3.5s".
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_sql_database_instance#upload_interval GoogleSqlDatabaseInstance#upload_interval}
	UploadInterval *string `field:"optional" json:"uploadInterval" yaml:"uploadInterval"`
}

