package googledataproccluster


type GoogleDataprocClusterVirtualClusterConfigAuxiliaryServicesConfig struct {
	// metastore_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_dataproc_cluster#metastore_config GoogleDataprocCluster#metastore_config}
	MetastoreConfig *GoogleDataprocClusterVirtualClusterConfigAuxiliaryServicesConfigMetastoreConfig `field:"optional" json:"metastoreConfig" yaml:"metastoreConfig"`
	// spark_history_server_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_dataproc_cluster#spark_history_server_config GoogleDataprocCluster#spark_history_server_config}
	SparkHistoryServerConfig *GoogleDataprocClusterVirtualClusterConfigAuxiliaryServicesConfigSparkHistoryServerConfig `field:"optional" json:"sparkHistoryServerConfig" yaml:"sparkHistoryServerConfig"`
}

