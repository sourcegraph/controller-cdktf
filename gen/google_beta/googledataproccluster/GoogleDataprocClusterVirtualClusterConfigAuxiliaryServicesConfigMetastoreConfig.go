package googledataproccluster


type GoogleDataprocClusterVirtualClusterConfigAuxiliaryServicesConfigMetastoreConfig struct {
	// The Hive Metastore configuration for this workload.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_dataproc_cluster#dataproc_metastore_service GoogleDataprocCluster#dataproc_metastore_service}
	DataprocMetastoreService *string `field:"optional" json:"dataprocMetastoreService" yaml:"dataprocMetastoreService"`
}

