package googledataproccluster


type GoogleDataprocClusterClusterConfigMetastoreConfig struct {
	// Resource name of an existing Dataproc Metastore service.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_dataproc_cluster#dataproc_metastore_service GoogleDataprocCluster#dataproc_metastore_service}
	DataprocMetastoreService *string `field:"required" json:"dataprocMetastoreService" yaml:"dataprocMetastoreService"`
}

