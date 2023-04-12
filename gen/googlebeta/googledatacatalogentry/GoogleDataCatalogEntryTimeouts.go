package googledatacatalogentry


type GoogleDataCatalogEntryTimeouts struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_data_catalog_entry#create GoogleDataCatalogEntry#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_data_catalog_entry#delete GoogleDataCatalogEntry#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_data_catalog_entry#update GoogleDataCatalogEntry#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

