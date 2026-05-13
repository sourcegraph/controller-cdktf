package bigqueryconnection


type BigqueryConnectionConfigurationAsset struct {
	// The name of the database.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/bigquery_connection#database BigqueryConnection#database}
	Database *string `field:"optional" json:"database" yaml:"database"`
	// The full resource name of the Google Cloud resource. For AlloyDB, this is in the format of '//alloydb.googleapis.com/projects/{project}/locations/{region}/clusters/{cluster}/instances/{instance}'.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/bigquery_connection#google_cloud_resource BigqueryConnection#google_cloud_resource}
	GoogleCloudResource *string `field:"optional" json:"googleCloudResource" yaml:"googleCloudResource"`
}

