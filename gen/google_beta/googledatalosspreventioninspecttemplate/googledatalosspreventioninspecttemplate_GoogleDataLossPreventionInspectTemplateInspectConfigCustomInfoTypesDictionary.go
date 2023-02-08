package googledatalosspreventioninspecttemplate


type GoogleDataLossPreventionInspectTemplateInspectConfigCustomInfoTypesDictionary struct {
	// cloud_storage_path block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_data_loss_prevention_inspect_template#cloud_storage_path GoogleDataLossPreventionInspectTemplate#cloud_storage_path}
	CloudStoragePath *GoogleDataLossPreventionInspectTemplateInspectConfigCustomInfoTypesDictionaryCloudStoragePath `field:"optional" json:"cloudStoragePath" yaml:"cloudStoragePath"`
	// word_list block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_data_loss_prevention_inspect_template#word_list GoogleDataLossPreventionInspectTemplate#word_list}
	WordList *GoogleDataLossPreventionInspectTemplateInspectConfigCustomInfoTypesDictionaryWordList `field:"optional" json:"wordList" yaml:"wordList"`
}

