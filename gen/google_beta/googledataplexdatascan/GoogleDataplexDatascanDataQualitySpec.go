package googledataplexdatascan


type GoogleDataplexDatascanDataQualitySpec struct {
	// If set, the latest DataScan job result will be published to Dataplex Catalog.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_dataplex_datascan#catalog_publishing_enabled GoogleDataplexDatascan#catalog_publishing_enabled}
	CatalogPublishingEnabled interface{} `field:"optional" json:"catalogPublishingEnabled" yaml:"catalogPublishingEnabled"`
	// If set to true, the scan will retrieve rules defined in Data Catalog for the resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_dataplex_datascan#enable_catalog_based_rules GoogleDataplexDatascan#enable_catalog_based_rules}
	EnableCatalogBasedRules interface{} `field:"optional" json:"enableCatalogBasedRules" yaml:"enableCatalogBasedRules"`
	// A filter to selectively run a subset of rules.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_dataplex_datascan#filter GoogleDataplexDatascan#filter}
	Filter *string `field:"optional" json:"filter" yaml:"filter"`
	// post_scan_actions block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_dataplex_datascan#post_scan_actions GoogleDataplexDatascan#post_scan_actions}
	PostScanActions *GoogleDataplexDatascanDataQualitySpecPostScanActions `field:"optional" json:"postScanActions" yaml:"postScanActions"`
	// A filter applied to all rows in a single DataScan job.
	//
	// The filter needs to be a valid SQL expression for a WHERE clause in BigQuery standard SQL syntax. Example: col1 >= 0 AND col2 < 10
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_dataplex_datascan#row_filter GoogleDataplexDatascan#row_filter}
	RowFilter *string `field:"optional" json:"rowFilter" yaml:"rowFilter"`
	// rules block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_dataplex_datascan#rules GoogleDataplexDatascan#rules}
	Rules interface{} `field:"optional" json:"rules" yaml:"rules"`
	// The percentage of the records to be selected from the dataset for DataScan.
	//
	// Value can range between 0.0 and 100.0 with up to 3 significant decimal digits.
	// Sampling is not applied if 'sampling_percent' is not specified, 0 or 100.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_dataplex_datascan#sampling_percent GoogleDataplexDatascan#sampling_percent}
	SamplingPercent *float64 `field:"optional" json:"samplingPercent" yaml:"samplingPercent"`
}

