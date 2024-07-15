package googledatalosspreventionstoredinfotype


type GoogleDataLossPreventionStoredInfoTypeLargeCustomDictionaryBigQueryField struct {
	// field block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_data_loss_prevention_stored_info_type#field GoogleDataLossPreventionStoredInfoType#field}
	Field *GoogleDataLossPreventionStoredInfoTypeLargeCustomDictionaryBigQueryFieldField `field:"required" json:"field" yaml:"field"`
	// table block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_data_loss_prevention_stored_info_type#table GoogleDataLossPreventionStoredInfoType#table}
	Table *GoogleDataLossPreventionStoredInfoTypeLargeCustomDictionaryBigQueryFieldTable `field:"required" json:"table" yaml:"table"`
}

