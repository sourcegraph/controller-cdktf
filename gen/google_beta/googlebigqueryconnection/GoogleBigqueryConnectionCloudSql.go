package googlebigqueryconnection


type GoogleBigqueryConnectionCloudSql struct {
	// credential block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_bigquery_connection#credential GoogleBigqueryConnection#credential}
	Credential *GoogleBigqueryConnectionCloudSqlCredential `field:"required" json:"credential" yaml:"credential"`
	// Database name.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_bigquery_connection#database GoogleBigqueryConnection#database}
	Database *string `field:"required" json:"database" yaml:"database"`
	// Cloud SQL instance ID in the form project:location:instance.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_bigquery_connection#instance_id GoogleBigqueryConnection#instance_id}
	InstanceId *string `field:"required" json:"instanceId" yaml:"instanceId"`
	// Type of the Cloud SQL database. Possible values: ["DATABASE_TYPE_UNSPECIFIED", "POSTGRES", "MYSQL"].
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_bigquery_connection#type GoogleBigqueryConnection#type}
	Type *string `field:"required" json:"type" yaml:"type"`
}

