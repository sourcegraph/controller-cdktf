package googledataplexdatascan


type GoogleDataplexDatascanData struct {
	// The Dataplex entity that represents the data source(e.g. BigQuery table) for Datascan.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_dataplex_datascan#entity GoogleDataplexDatascan#entity}
	Entity *string `field:"optional" json:"entity" yaml:"entity"`
	// The service-qualified full resource name of the cloud resource for a DataScan job to scan against.
	//
	// The field could be:
	// Cloud Storage bucket (//storage.googleapis.com/projects/PROJECT_ID/buckets/BUCKET_ID) for DataDiscoveryScan OR BigQuery table of type "TABLE" (/bigquery.googleapis.com/projects/PROJECT_ID/datasets/DATASET_ID/tables/TABLE_ID) for DataProfileScan/DataQualityScan.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_dataplex_datascan#resource GoogleDataplexDatascan#resource}
	Resource *string `field:"optional" json:"resource" yaml:"resource"`
}

