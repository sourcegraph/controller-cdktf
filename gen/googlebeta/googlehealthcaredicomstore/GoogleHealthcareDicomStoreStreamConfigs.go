package googlehealthcaredicomstore


type GoogleHealthcareDicomStoreStreamConfigs struct {
	// bigquery_destination block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_healthcare_dicom_store#bigquery_destination GoogleHealthcareDicomStore#bigquery_destination}
	BigqueryDestination *GoogleHealthcareDicomStoreStreamConfigsBigqueryDestination `field:"required" json:"bigqueryDestination" yaml:"bigqueryDestination"`
}

