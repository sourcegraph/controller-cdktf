package googledatalosspreventionstoredinfotype


type GoogleDataLossPreventionStoredInfoTypeLargeCustomDictionaryCloudStorageFileSet struct {
	// The url, in the format 'gs://<bucket>/<path>'. Trailing wildcard in the path is allowed.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_data_loss_prevention_stored_info_type#url GoogleDataLossPreventionStoredInfoType#url}
	Url *string `field:"required" json:"url" yaml:"url"`
}

