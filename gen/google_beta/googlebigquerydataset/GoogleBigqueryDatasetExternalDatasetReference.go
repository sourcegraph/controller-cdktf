package googlebigquerydataset


type GoogleBigqueryDatasetExternalDatasetReference struct {
	// The connection id that is used to access the externalSource. Format: projects/{projectId}/locations/{locationId}/connections/{connectionId}.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_bigquery_dataset#connection GoogleBigqueryDataset#connection}
	Connection *string `field:"required" json:"connection" yaml:"connection"`
	// External source that backs this dataset.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_bigquery_dataset#external_source GoogleBigqueryDataset#external_source}
	ExternalSource *string `field:"required" json:"externalSource" yaml:"externalSource"`
}

