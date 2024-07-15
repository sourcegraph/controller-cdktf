package googledatalosspreventionstoredinfotype


type GoogleDataLossPreventionStoredInfoTypeLargeCustomDictionaryBigQueryFieldTable struct {
	// The dataset ID of the table.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_data_loss_prevention_stored_info_type#dataset_id GoogleDataLossPreventionStoredInfoType#dataset_id}
	DatasetId *string `field:"required" json:"datasetId" yaml:"datasetId"`
	// The Google Cloud Platform project ID of the project containing the table.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_data_loss_prevention_stored_info_type#project_id GoogleDataLossPreventionStoredInfoType#project_id}
	ProjectId *string `field:"required" json:"projectId" yaml:"projectId"`
	// The name of the table.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_data_loss_prevention_stored_info_type#table_id GoogleDataLossPreventionStoredInfoType#table_id}
	TableId *string `field:"required" json:"tableId" yaml:"tableId"`
}

