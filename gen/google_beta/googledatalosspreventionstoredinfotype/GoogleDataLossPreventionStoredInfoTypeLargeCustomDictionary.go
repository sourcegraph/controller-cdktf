package googledatalosspreventionstoredinfotype


type GoogleDataLossPreventionStoredInfoTypeLargeCustomDictionary struct {
	// output_path block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_data_loss_prevention_stored_info_type#output_path GoogleDataLossPreventionStoredInfoType#output_path}
	OutputPath *GoogleDataLossPreventionStoredInfoTypeLargeCustomDictionaryOutputPath `field:"required" json:"outputPath" yaml:"outputPath"`
	// big_query_field block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_data_loss_prevention_stored_info_type#big_query_field GoogleDataLossPreventionStoredInfoType#big_query_field}
	BigQueryField *GoogleDataLossPreventionStoredInfoTypeLargeCustomDictionaryBigQueryField `field:"optional" json:"bigQueryField" yaml:"bigQueryField"`
	// cloud_storage_file_set block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_data_loss_prevention_stored_info_type#cloud_storage_file_set GoogleDataLossPreventionStoredInfoType#cloud_storage_file_set}
	CloudStorageFileSet *GoogleDataLossPreventionStoredInfoTypeLargeCustomDictionaryCloudStorageFileSet `field:"optional" json:"cloudStorageFileSet" yaml:"cloudStorageFileSet"`
}

