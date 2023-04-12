package googledatalosspreventiondeidentifytemplate


type GoogleDataLossPreventionDeidentifyTemplateDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformationCryptoReplaceFfxFpeConfigCryptoKeyUnwrapped struct {
	// A 128/192/256 bit key.
	//
	// A base64-encoded string.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_data_loss_prevention_deidentify_template#key GoogleDataLossPreventionDeidentifyTemplate#key}
	Key *string `field:"required" json:"key" yaml:"key"`
}

