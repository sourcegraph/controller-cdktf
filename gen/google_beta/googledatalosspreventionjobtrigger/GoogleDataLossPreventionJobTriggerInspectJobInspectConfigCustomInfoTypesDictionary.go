package googledatalosspreventionjobtrigger


type GoogleDataLossPreventionJobTriggerInspectJobInspectConfigCustomInfoTypesDictionary struct {
	// cloud_storage_path block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_data_loss_prevention_job_trigger#cloud_storage_path GoogleDataLossPreventionJobTrigger#cloud_storage_path}
	CloudStoragePath *GoogleDataLossPreventionJobTriggerInspectJobInspectConfigCustomInfoTypesDictionaryCloudStoragePath `field:"optional" json:"cloudStoragePath" yaml:"cloudStoragePath"`
	// word_list block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_data_loss_prevention_job_trigger#word_list GoogleDataLossPreventionJobTrigger#word_list}
	WordList *GoogleDataLossPreventionJobTriggerInspectJobInspectConfigCustomInfoTypesDictionaryWordListStruct `field:"optional" json:"wordList" yaml:"wordList"`
}

