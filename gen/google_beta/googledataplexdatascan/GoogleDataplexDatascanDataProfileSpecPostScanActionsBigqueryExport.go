package googledataplexdatascan


type GoogleDataplexDatascanDataProfileSpecPostScanActionsBigqueryExport struct {
	// The BigQuery table to export DataProfileScan results to. Format://bigquery.googleapis.com/projects/PROJECT_ID/datasets/DATASET_ID/tables/TABLE_ID.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_dataplex_datascan#results_table GoogleDataplexDatascan#results_table}
	ResultsTable *string `field:"optional" json:"resultsTable" yaml:"resultsTable"`
}

