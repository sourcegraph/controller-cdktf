package googledatalosspreventiondeidentifytemplate


type GoogleDataLossPreventionDeidentifyTemplateDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformationCryptoDeterministicConfigSurrogateInfoTypeSensitivityScore struct {
	// The sensitivity score applied to the resource. Possible values: ["SENSITIVITY_LOW", "SENSITIVITY_MODERATE", "SENSITIVITY_HIGH"].
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_data_loss_prevention_deidentify_template#score GoogleDataLossPreventionDeidentifyTemplate#score}
	Score *string `field:"required" json:"score" yaml:"score"`
}

