package googledatalosspreventiondeidentifytemplate


type GoogleDataLossPreventionDeidentifyTemplateDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformation struct {
	// character_mask_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/4.65.2/docs/resources/google_data_loss_prevention_deidentify_template#character_mask_config GoogleDataLossPreventionDeidentifyTemplate#character_mask_config}
	CharacterMaskConfig *GoogleDataLossPreventionDeidentifyTemplateDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformationCharacterMaskConfig `field:"optional" json:"characterMaskConfig" yaml:"characterMaskConfig"`
	// crypto_deterministic_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/4.65.2/docs/resources/google_data_loss_prevention_deidentify_template#crypto_deterministic_config GoogleDataLossPreventionDeidentifyTemplate#crypto_deterministic_config}
	CryptoDeterministicConfig *GoogleDataLossPreventionDeidentifyTemplateDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformationCryptoDeterministicConfig `field:"optional" json:"cryptoDeterministicConfig" yaml:"cryptoDeterministicConfig"`
	// crypto_replace_ffx_fpe_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/4.65.2/docs/resources/google_data_loss_prevention_deidentify_template#crypto_replace_ffx_fpe_config GoogleDataLossPreventionDeidentifyTemplate#crypto_replace_ffx_fpe_config}
	CryptoReplaceFfxFpeConfig *GoogleDataLossPreventionDeidentifyTemplateDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformationCryptoReplaceFfxFpeConfig `field:"optional" json:"cryptoReplaceFfxFpeConfig" yaml:"cryptoReplaceFfxFpeConfig"`
	// replace_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/4.65.2/docs/resources/google_data_loss_prevention_deidentify_template#replace_config GoogleDataLossPreventionDeidentifyTemplate#replace_config}
	ReplaceConfig *GoogleDataLossPreventionDeidentifyTemplateDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformationReplaceConfig `field:"optional" json:"replaceConfig" yaml:"replaceConfig"`
	// replace_dictionary_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/4.65.2/docs/resources/google_data_loss_prevention_deidentify_template#replace_dictionary_config GoogleDataLossPreventionDeidentifyTemplate#replace_dictionary_config}
	ReplaceDictionaryConfig *GoogleDataLossPreventionDeidentifyTemplateDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformationReplaceDictionaryConfig `field:"optional" json:"replaceDictionaryConfig" yaml:"replaceDictionaryConfig"`
	// Replace each matching finding with the name of the info type.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/4.65.2/docs/resources/google_data_loss_prevention_deidentify_template#replace_with_info_type_config GoogleDataLossPreventionDeidentifyTemplate#replace_with_info_type_config}
	ReplaceWithInfoTypeConfig interface{} `field:"optional" json:"replaceWithInfoTypeConfig" yaml:"replaceWithInfoTypeConfig"`
}

