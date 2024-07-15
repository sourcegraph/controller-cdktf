package googledatalosspreventiondeidentifytemplate


type GoogleDataLossPreventionDeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformation struct {
	// bucketing_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_data_loss_prevention_deidentify_template#bucketing_config GoogleDataLossPreventionDeidentifyTemplate#bucketing_config}
	BucketingConfig *GoogleDataLossPreventionDeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationBucketingConfig `field:"optional" json:"bucketingConfig" yaml:"bucketingConfig"`
	// character_mask_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_data_loss_prevention_deidentify_template#character_mask_config GoogleDataLossPreventionDeidentifyTemplate#character_mask_config}
	CharacterMaskConfig *GoogleDataLossPreventionDeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationCharacterMaskConfig `field:"optional" json:"characterMaskConfig" yaml:"characterMaskConfig"`
	// crypto_deterministic_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_data_loss_prevention_deidentify_template#crypto_deterministic_config GoogleDataLossPreventionDeidentifyTemplate#crypto_deterministic_config}
	CryptoDeterministicConfig *GoogleDataLossPreventionDeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationCryptoDeterministicConfig `field:"optional" json:"cryptoDeterministicConfig" yaml:"cryptoDeterministicConfig"`
	// crypto_hash_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_data_loss_prevention_deidentify_template#crypto_hash_config GoogleDataLossPreventionDeidentifyTemplate#crypto_hash_config}
	CryptoHashConfig *GoogleDataLossPreventionDeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationCryptoHashConfig `field:"optional" json:"cryptoHashConfig" yaml:"cryptoHashConfig"`
	// crypto_replace_ffx_fpe_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_data_loss_prevention_deidentify_template#crypto_replace_ffx_fpe_config GoogleDataLossPreventionDeidentifyTemplate#crypto_replace_ffx_fpe_config}
	CryptoReplaceFfxFpeConfig *GoogleDataLossPreventionDeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationCryptoReplaceFfxFpeConfig `field:"optional" json:"cryptoReplaceFfxFpeConfig" yaml:"cryptoReplaceFfxFpeConfig"`
	// date_shift_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_data_loss_prevention_deidentify_template#date_shift_config GoogleDataLossPreventionDeidentifyTemplate#date_shift_config}
	DateShiftConfig *GoogleDataLossPreventionDeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationDateShiftConfig `field:"optional" json:"dateShiftConfig" yaml:"dateShiftConfig"`
	// fixed_size_bucketing_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_data_loss_prevention_deidentify_template#fixed_size_bucketing_config GoogleDataLossPreventionDeidentifyTemplate#fixed_size_bucketing_config}
	FixedSizeBucketingConfig *GoogleDataLossPreventionDeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationFixedSizeBucketingConfig `field:"optional" json:"fixedSizeBucketingConfig" yaml:"fixedSizeBucketingConfig"`
	// redact_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_data_loss_prevention_deidentify_template#redact_config GoogleDataLossPreventionDeidentifyTemplate#redact_config}
	RedactConfig *GoogleDataLossPreventionDeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationRedactConfig `field:"optional" json:"redactConfig" yaml:"redactConfig"`
	// replace_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_data_loss_prevention_deidentify_template#replace_config GoogleDataLossPreventionDeidentifyTemplate#replace_config}
	ReplaceConfig *GoogleDataLossPreventionDeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationReplaceConfig `field:"optional" json:"replaceConfig" yaml:"replaceConfig"`
	// replace_dictionary_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_data_loss_prevention_deidentify_template#replace_dictionary_config GoogleDataLossPreventionDeidentifyTemplate#replace_dictionary_config}
	ReplaceDictionaryConfig *GoogleDataLossPreventionDeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationReplaceDictionaryConfig `field:"optional" json:"replaceDictionaryConfig" yaml:"replaceDictionaryConfig"`
	// time_part_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_data_loss_prevention_deidentify_template#time_part_config GoogleDataLossPreventionDeidentifyTemplate#time_part_config}
	TimePartConfig *GoogleDataLossPreventionDeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationTimePartConfig `field:"optional" json:"timePartConfig" yaml:"timePartConfig"`
}

