package googledataprocmetastoreservice


type GoogleDataprocMetastoreServiceMetadataIntegrationDataCatalogConfig struct {
	// Defines whether the metastore metadata should be synced to Data Catalog.
	//
	// The default value is to disable syncing metastore metadata to Data Catalog.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_dataproc_metastore_service#enabled GoogleDataprocMetastoreService#enabled}
	Enabled interface{} `field:"required" json:"enabled" yaml:"enabled"`
}

