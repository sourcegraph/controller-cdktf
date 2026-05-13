package googledataplexdatascan


type GoogleDataplexDatascanDataDiscoverySpecBigqueryPublishingConfig struct {
	// The BigQuery connection used to create BigLake tables. Must be in the form 'projects/{projectId}/locations/{locationId}/connections/{connection_id}'.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_dataplex_datascan#connection GoogleDataplexDatascan#connection}
	Connection *string `field:"optional" json:"connection" yaml:"connection"`
	// The location of the BigQuery dataset to publish BigLake external or non-BigLake external tables to.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_dataplex_datascan#location GoogleDataplexDatascan#location}
	Location *string `field:"optional" json:"location" yaml:"location"`
	// The project of the BigQuery dataset to publish BigLake external or non-BigLake external tables to.
	//
	// If not specified, the project of the Cloud Storage bucket will be used. The format is "projects/{project_id_or_number}".
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_dataplex_datascan#project GoogleDataplexDatascan#project}
	Project *string `field:"optional" json:"project" yaml:"project"`
	// Determines whether to publish discovered tables as BigLake external tables or non-BigLake external tables. Possible values: ["TABLE_TYPE_UNSPECIFIED", "EXTERNAL", "BIGLAKE"].
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_dataplex_datascan#table_type GoogleDataplexDatascan#table_type}
	TableType *string `field:"optional" json:"tableType" yaml:"tableType"`
}

