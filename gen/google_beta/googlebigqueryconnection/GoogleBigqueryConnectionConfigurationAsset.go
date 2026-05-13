package googlebigqueryconnection


type GoogleBigqueryConnectionConfigurationAsset struct {
	// The name of the database.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_bigquery_connection#database GoogleBigqueryConnection#database}
	Database *string `field:"optional" json:"database" yaml:"database"`
	// The full resource name of the Google Cloud resource. For AlloyDB, this is in the format of '//alloydb.googleapis.com/projects/{project}/locations/{region}/clusters/{cluster}/instances/{instance}'.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_bigquery_connection#google_cloud_resource GoogleBigqueryConnection#google_cloud_resource}
	GoogleCloudResource *string `field:"optional" json:"googleCloudResource" yaml:"googleCloudResource"`
}

