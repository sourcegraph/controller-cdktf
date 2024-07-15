package googledatalosspreventiondeidentifytemplate


type GoogleDataLossPreventionDeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationDateShiftConfigCryptoKey struct {
	// kms_wrapped block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_data_loss_prevention_deidentify_template#kms_wrapped GoogleDataLossPreventionDeidentifyTemplate#kms_wrapped}
	KmsWrapped *GoogleDataLossPreventionDeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationDateShiftConfigCryptoKeyKmsWrapped `field:"optional" json:"kmsWrapped" yaml:"kmsWrapped"`
	// transient block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_data_loss_prevention_deidentify_template#transient GoogleDataLossPreventionDeidentifyTemplate#transient}
	Transient *GoogleDataLossPreventionDeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationDateShiftConfigCryptoKeyTransient `field:"optional" json:"transient" yaml:"transient"`
	// unwrapped block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_data_loss_prevention_deidentify_template#unwrapped GoogleDataLossPreventionDeidentifyTemplate#unwrapped}
	Unwrapped *GoogleDataLossPreventionDeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationDateShiftConfigCryptoKeyUnwrapped `field:"optional" json:"unwrapped" yaml:"unwrapped"`
}

