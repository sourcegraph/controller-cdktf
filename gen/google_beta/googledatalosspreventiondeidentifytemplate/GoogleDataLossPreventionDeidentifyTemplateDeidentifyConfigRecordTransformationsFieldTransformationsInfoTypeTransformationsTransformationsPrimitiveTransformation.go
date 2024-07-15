package googledatalosspreventiondeidentifytemplate


type GoogleDataLossPreventionDeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformation struct {
	// bucketing_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_data_loss_prevention_deidentify_template#bucketing_config GoogleDataLossPreventionDeidentifyTemplate#bucketing_config}
	BucketingConfig *GoogleDataLossPreventionDeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformationBucketingConfig `field:"optional" json:"bucketingConfig" yaml:"bucketingConfig"`
	// character_mask_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_data_loss_prevention_deidentify_template#character_mask_config GoogleDataLossPreventionDeidentifyTemplate#character_mask_config}
	CharacterMaskConfig *GoogleDataLossPreventionDeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformationCharacterMaskConfig `field:"optional" json:"characterMaskConfig" yaml:"characterMaskConfig"`
	// crypto_deterministic_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_data_loss_prevention_deidentify_template#crypto_deterministic_config GoogleDataLossPreventionDeidentifyTemplate#crypto_deterministic_config}
	CryptoDeterministicConfig *GoogleDataLossPreventionDeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformationCryptoDeterministicConfig `field:"optional" json:"cryptoDeterministicConfig" yaml:"cryptoDeterministicConfig"`
	// crypto_hash_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_data_loss_prevention_deidentify_template#crypto_hash_config GoogleDataLossPreventionDeidentifyTemplate#crypto_hash_config}
	CryptoHashConfig *GoogleDataLossPreventionDeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformationCryptoHashConfig `field:"optional" json:"cryptoHashConfig" yaml:"cryptoHashConfig"`
	// crypto_replace_ffx_fpe_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_data_loss_prevention_deidentify_template#crypto_replace_ffx_fpe_config GoogleDataLossPreventionDeidentifyTemplate#crypto_replace_ffx_fpe_config}
	CryptoReplaceFfxFpeConfig *GoogleDataLossPreventionDeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformationCryptoReplaceFfxFpeConfig `field:"optional" json:"cryptoReplaceFfxFpeConfig" yaml:"cryptoReplaceFfxFpeConfig"`
	// date_shift_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_data_loss_prevention_deidentify_template#date_shift_config GoogleDataLossPreventionDeidentifyTemplate#date_shift_config}
	DateShiftConfig *GoogleDataLossPreventionDeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformationDateShiftConfig `field:"optional" json:"dateShiftConfig" yaml:"dateShiftConfig"`
	// fixed_size_bucketing_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_data_loss_prevention_deidentify_template#fixed_size_bucketing_config GoogleDataLossPreventionDeidentifyTemplate#fixed_size_bucketing_config}
	FixedSizeBucketingConfig *GoogleDataLossPreventionDeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformationFixedSizeBucketingConfig `field:"optional" json:"fixedSizeBucketingConfig" yaml:"fixedSizeBucketingConfig"`
	// redact_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_data_loss_prevention_deidentify_template#redact_config GoogleDataLossPreventionDeidentifyTemplate#redact_config}
	RedactConfig *GoogleDataLossPreventionDeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformationRedactConfig `field:"optional" json:"redactConfig" yaml:"redactConfig"`
	// replace_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_data_loss_prevention_deidentify_template#replace_config GoogleDataLossPreventionDeidentifyTemplate#replace_config}
	ReplaceConfig *GoogleDataLossPreventionDeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformationReplaceConfig `field:"optional" json:"replaceConfig" yaml:"replaceConfig"`
	// replace_dictionary_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_data_loss_prevention_deidentify_template#replace_dictionary_config GoogleDataLossPreventionDeidentifyTemplate#replace_dictionary_config}
	ReplaceDictionaryConfig *GoogleDataLossPreventionDeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformationReplaceDictionaryConfig `field:"optional" json:"replaceDictionaryConfig" yaml:"replaceDictionaryConfig"`
	// replace_with_info_type_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_data_loss_prevention_deidentify_template#replace_with_info_type_config GoogleDataLossPreventionDeidentifyTemplate#replace_with_info_type_config}
	ReplaceWithInfoTypeConfig *GoogleDataLossPreventionDeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformationReplaceWithInfoTypeConfig `field:"optional" json:"replaceWithInfoTypeConfig" yaml:"replaceWithInfoTypeConfig"`
	// time_part_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_data_loss_prevention_deidentify_template#time_part_config GoogleDataLossPreventionDeidentifyTemplate#time_part_config}
	TimePartConfig *GoogleDataLossPreventionDeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformationTimePartConfig `field:"optional" json:"timePartConfig" yaml:"timePartConfig"`
}

