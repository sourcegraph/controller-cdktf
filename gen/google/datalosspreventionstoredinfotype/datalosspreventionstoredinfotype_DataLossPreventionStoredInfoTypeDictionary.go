package datalosspreventionstoredinfotype


type DataLossPreventionStoredInfoTypeDictionary struct {
	// cloud_storage_path block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/data_loss_prevention_stored_info_type#cloud_storage_path DataLossPreventionStoredInfoType#cloud_storage_path}
	CloudStoragePath *DataLossPreventionStoredInfoTypeDictionaryCloudStoragePath `field:"optional" json:"cloudStoragePath" yaml:"cloudStoragePath"`
	// word_list block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/data_loss_prevention_stored_info_type#word_list DataLossPreventionStoredInfoType#word_list}
	WordList *DataLossPreventionStoredInfoTypeDictionaryWordList `field:"optional" json:"wordList" yaml:"wordList"`
}

