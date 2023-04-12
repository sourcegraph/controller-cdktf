package googledatalosspreventionstoredinfotype


type GoogleDataLossPreventionStoredInfoTypeDictionary struct {
	// cloud_storage_path block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_data_loss_prevention_stored_info_type#cloud_storage_path GoogleDataLossPreventionStoredInfoType#cloud_storage_path}
	CloudStoragePath *GoogleDataLossPreventionStoredInfoTypeDictionaryCloudStoragePath `field:"optional" json:"cloudStoragePath" yaml:"cloudStoragePath"`
	// word_list block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_data_loss_prevention_stored_info_type#word_list GoogleDataLossPreventionStoredInfoType#word_list}
	WordList *GoogleDataLossPreventionStoredInfoTypeDictionaryWordList `field:"optional" json:"wordList" yaml:"wordList"`
}

