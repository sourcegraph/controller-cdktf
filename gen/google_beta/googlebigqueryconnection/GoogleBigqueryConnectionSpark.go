package googlebigqueryconnection


type GoogleBigqueryConnectionSpark struct {
	// metastore_service_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_bigquery_connection#metastore_service_config GoogleBigqueryConnection#metastore_service_config}
	MetastoreServiceConfig *GoogleBigqueryConnectionSparkMetastoreServiceConfig `field:"optional" json:"metastoreServiceConfig" yaml:"metastoreServiceConfig"`
	// spark_history_server_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_bigquery_connection#spark_history_server_config GoogleBigqueryConnection#spark_history_server_config}
	SparkHistoryServerConfig *GoogleBigqueryConnectionSparkSparkHistoryServerConfig `field:"optional" json:"sparkHistoryServerConfig" yaml:"sparkHistoryServerConfig"`
}

