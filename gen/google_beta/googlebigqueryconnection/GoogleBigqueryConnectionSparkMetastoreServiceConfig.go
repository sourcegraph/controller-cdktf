package googlebigqueryconnection


type GoogleBigqueryConnectionSparkMetastoreServiceConfig struct {
	// Resource name of an existing Dataproc Metastore service in the form of projects/[projectId]/locations/[region]/services/[serviceId].
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_bigquery_connection#metastore_service GoogleBigqueryConnection#metastore_service}
	MetastoreService *string `field:"optional" json:"metastoreService" yaml:"metastoreService"`
}

