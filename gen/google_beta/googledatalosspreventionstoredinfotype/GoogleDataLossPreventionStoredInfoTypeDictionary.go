package googledatalosspreventionstoredinfotype


type GoogleDataLossPreventionStoredInfoTypeDictionary struct {
	// cloud_storage_path block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_data_loss_prevention_stored_info_type#cloud_storage_path GoogleDataLossPreventionStoredInfoType#cloud_storage_path}
	CloudStoragePath *GoogleDataLossPreventionStoredInfoTypeDictionaryCloudStoragePath `field:"optional" json:"cloudStoragePath" yaml:"cloudStoragePath"`
	// word_list block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_data_loss_prevention_stored_info_type#word_list GoogleDataLossPreventionStoredInfoType#word_list}
	WordList *GoogleDataLossPreventionStoredInfoTypeDictionaryWordListStruct `field:"optional" json:"wordList" yaml:"wordList"`
}

