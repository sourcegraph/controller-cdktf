package googledatalosspreventiondeidentifytemplate


type GoogleDataLossPreventionDeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformationCryptoHashConfigCryptoKeyKmsWrapped struct {
	// The resource name of the KMS CryptoKey to use for unwrapping.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_data_loss_prevention_deidentify_template#crypto_key_name GoogleDataLossPreventionDeidentifyTemplate#crypto_key_name}
	CryptoKeyName *string `field:"required" json:"cryptoKeyName" yaml:"cryptoKeyName"`
	// The wrapped data crypto key.
	//
	// A base64-encoded string.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_data_loss_prevention_deidentify_template#wrapped_key GoogleDataLossPreventionDeidentifyTemplate#wrapped_key}
	WrappedKey *string `field:"required" json:"wrappedKey" yaml:"wrappedKey"`
}

