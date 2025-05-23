package dataprocmetastoreservice


type DataprocMetastoreServiceMetadataIntegration struct {
	// data_catalog_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.32.0/docs/resources/dataproc_metastore_service#data_catalog_config DataprocMetastoreService#data_catalog_config}
	DataCatalogConfig *DataprocMetastoreServiceMetadataIntegrationDataCatalogConfig `field:"required" json:"dataCatalogConfig" yaml:"dataCatalogConfig"`
}

