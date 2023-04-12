package googledatalosspreventionstoredinfotype


type GoogleDataLossPreventionStoredInfoTypeLargeCustomDictionaryBigQueryField struct {
	// field block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_data_loss_prevention_stored_info_type#field GoogleDataLossPreventionStoredInfoType#field}
	Field *GoogleDataLossPreventionStoredInfoTypeLargeCustomDictionaryBigQueryFieldField `field:"required" json:"field" yaml:"field"`
	// table block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_data_loss_prevention_stored_info_type#table GoogleDataLossPreventionStoredInfoType#table}
	Table *GoogleDataLossPreventionStoredInfoTypeLargeCustomDictionaryBigQueryFieldTable `field:"required" json:"table" yaml:"table"`
}

