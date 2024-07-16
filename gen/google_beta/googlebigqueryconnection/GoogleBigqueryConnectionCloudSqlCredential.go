package googlebigqueryconnection


type GoogleBigqueryConnectionCloudSqlCredential struct {
	// Password for database.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_bigquery_connection#password GoogleBigqueryConnection#password}
	Password *string `field:"required" json:"password" yaml:"password"`
	// Username for database.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_bigquery_connection#username GoogleBigqueryConnection#username}
	Username *string `field:"required" json:"username" yaml:"username"`
}

