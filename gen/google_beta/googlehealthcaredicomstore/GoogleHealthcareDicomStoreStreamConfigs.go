package googlehealthcaredicomstore


type GoogleHealthcareDicomStoreStreamConfigs struct {
	// bigquery_destination block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_healthcare_dicom_store#bigquery_destination GoogleHealthcareDicomStore#bigquery_destination}
	BigqueryDestination *GoogleHealthcareDicomStoreStreamConfigsBigqueryDestination `field:"required" json:"bigqueryDestination" yaml:"bigqueryDestination"`
}

