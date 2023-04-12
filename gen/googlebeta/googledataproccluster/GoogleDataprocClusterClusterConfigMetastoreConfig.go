package googledataproccluster


type GoogleDataprocClusterClusterConfigMetastoreConfig struct {
	// Resource name of an existing Dataproc Metastore service.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_dataproc_cluster#dataproc_metastore_service GoogleDataprocCluster#dataproc_metastore_service}
	DataprocMetastoreService *string `field:"required" json:"dataprocMetastoreService" yaml:"dataprocMetastoreService"`
}

