package googledataprocmetastoreservice


type GoogleDataprocMetastoreServiceMetadataIntegration struct {
	// data_catalog_config block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_dataproc_metastore_service#data_catalog_config GoogleDataprocMetastoreService#data_catalog_config}
	DataCatalogConfig *GoogleDataprocMetastoreServiceMetadataIntegrationDataCatalogConfig `field:"required" json:"dataCatalogConfig" yaml:"dataCatalogConfig"`
}

