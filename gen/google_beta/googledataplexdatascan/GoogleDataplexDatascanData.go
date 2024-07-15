package googledataplexdatascan


type GoogleDataplexDatascanData struct {
	// The Dataplex entity that represents the data source(e.g. BigQuery table) for Datascan.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_dataplex_datascan#entity GoogleDataplexDatascan#entity}
	Entity *string `field:"optional" json:"entity" yaml:"entity"`
	// The service-qualified full resource name of the cloud resource for a DataScan job to scan against.
	//
	// The field could be:
	// (Cloud Storage bucket for DataDiscoveryScan)BigQuery table of type "TABLE" for DataProfileScan/DataQualityScan.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_dataplex_datascan#resource GoogleDataplexDatascan#resource}
	Resource *string `field:"optional" json:"resource" yaml:"resource"`
}

