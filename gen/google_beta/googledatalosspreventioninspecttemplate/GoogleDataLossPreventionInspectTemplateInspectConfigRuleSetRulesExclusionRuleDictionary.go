package googledatalosspreventioninspecttemplate


type GoogleDataLossPreventionInspectTemplateInspectConfigRuleSetRulesExclusionRuleDictionary struct {
	// cloud_storage_path block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_data_loss_prevention_inspect_template#cloud_storage_path GoogleDataLossPreventionInspectTemplate#cloud_storage_path}
	CloudStoragePath *GoogleDataLossPreventionInspectTemplateInspectConfigRuleSetRulesExclusionRuleDictionaryCloudStoragePath `field:"optional" json:"cloudStoragePath" yaml:"cloudStoragePath"`
	// word_list block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_data_loss_prevention_inspect_template#word_list GoogleDataLossPreventionInspectTemplate#word_list}
	WordList *GoogleDataLossPreventionInspectTemplateInspectConfigRuleSetRulesExclusionRuleDictionaryWordListStruct `field:"optional" json:"wordList" yaml:"wordList"`
}

