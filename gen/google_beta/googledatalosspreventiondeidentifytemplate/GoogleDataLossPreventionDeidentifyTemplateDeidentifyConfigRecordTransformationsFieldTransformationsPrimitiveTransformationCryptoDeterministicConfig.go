package googledatalosspreventiondeidentifytemplate


type GoogleDataLossPreventionDeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationCryptoDeterministicConfig struct {
	// context block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_data_loss_prevention_deidentify_template#context GoogleDataLossPreventionDeidentifyTemplate#context}
	Context *GoogleDataLossPreventionDeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationCryptoDeterministicConfigContext `field:"optional" json:"context" yaml:"context"`
	// crypto_key block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_data_loss_prevention_deidentify_template#crypto_key GoogleDataLossPreventionDeidentifyTemplate#crypto_key}
	CryptoKey *GoogleDataLossPreventionDeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationCryptoDeterministicConfigCryptoKey `field:"optional" json:"cryptoKey" yaml:"cryptoKey"`
	// surrogate_info_type block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_data_loss_prevention_deidentify_template#surrogate_info_type GoogleDataLossPreventionDeidentifyTemplate#surrogate_info_type}
	SurrogateInfoType *GoogleDataLossPreventionDeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationCryptoDeterministicConfigSurrogateInfoType `field:"optional" json:"surrogateInfoType" yaml:"surrogateInfoType"`
}

