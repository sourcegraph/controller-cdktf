package googlehealthcaredicomstore


type GoogleHealthcareDicomStoreStreamConfigsBigqueryDestination struct {
	// a fully qualified BigQuery table URI where DICOM instance metadata will be streamed.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_healthcare_dicom_store#table_uri GoogleHealthcareDicomStore#table_uri}
	TableUri *string `field:"required" json:"tableUri" yaml:"tableUri"`
}

