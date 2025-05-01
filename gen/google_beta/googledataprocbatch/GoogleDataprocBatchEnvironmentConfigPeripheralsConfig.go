package googledataprocbatch


type GoogleDataprocBatchEnvironmentConfigPeripheralsConfig struct {
	// Resource name of an existing Dataproc Metastore service.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.32.0/docs/resources/google_dataproc_batch#metastore_service GoogleDataprocBatch#metastore_service}
	MetastoreService *string `field:"optional" json:"metastoreService" yaml:"metastoreService"`
	// spark_history_server_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.32.0/docs/resources/google_dataproc_batch#spark_history_server_config GoogleDataprocBatch#spark_history_server_config}
	SparkHistoryServerConfig *GoogleDataprocBatchEnvironmentConfigPeripheralsConfigSparkHistoryServerConfig `field:"optional" json:"sparkHistoryServerConfig" yaml:"sparkHistoryServerConfig"`
}

