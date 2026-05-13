package googledataprocsessiontemplate


type GoogleDataprocSessionTemplateEnvironmentConfigPeripheralsConfig struct {
	// Resource name of an existing Dataproc Metastore service.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_dataproc_session_template#metastore_service GoogleDataprocSessionTemplate#metastore_service}
	MetastoreService *string `field:"optional" json:"metastoreService" yaml:"metastoreService"`
	// spark_history_server_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_dataproc_session_template#spark_history_server_config GoogleDataprocSessionTemplate#spark_history_server_config}
	SparkHistoryServerConfig *GoogleDataprocSessionTemplateEnvironmentConfigPeripheralsConfigSparkHistoryServerConfig `field:"optional" json:"sparkHistoryServerConfig" yaml:"sparkHistoryServerConfig"`
}

