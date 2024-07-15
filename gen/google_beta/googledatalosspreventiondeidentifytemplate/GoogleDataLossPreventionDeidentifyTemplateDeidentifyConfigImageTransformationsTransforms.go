package googledatalosspreventiondeidentifytemplate


type GoogleDataLossPreventionDeidentifyTemplateDeidentifyConfigImageTransformationsTransforms struct {
	// all_info_types block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_data_loss_prevention_deidentify_template#all_info_types GoogleDataLossPreventionDeidentifyTemplate#all_info_types}
	AllInfoTypes *GoogleDataLossPreventionDeidentifyTemplateDeidentifyConfigImageTransformationsTransformsAllInfoTypes `field:"optional" json:"allInfoTypes" yaml:"allInfoTypes"`
	// all_text block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_data_loss_prevention_deidentify_template#all_text GoogleDataLossPreventionDeidentifyTemplate#all_text}
	AllText *GoogleDataLossPreventionDeidentifyTemplateDeidentifyConfigImageTransformationsTransformsAllText `field:"optional" json:"allText" yaml:"allText"`
	// redaction_color block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_data_loss_prevention_deidentify_template#redaction_color GoogleDataLossPreventionDeidentifyTemplate#redaction_color}
	RedactionColor *GoogleDataLossPreventionDeidentifyTemplateDeidentifyConfigImageTransformationsTransformsRedactionColor `field:"optional" json:"redactionColor" yaml:"redactionColor"`
	// selected_info_types block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_data_loss_prevention_deidentify_template#selected_info_types GoogleDataLossPreventionDeidentifyTemplate#selected_info_types}
	SelectedInfoTypes *GoogleDataLossPreventionDeidentifyTemplateDeidentifyConfigImageTransformationsTransformsSelectedInfoTypes `field:"optional" json:"selectedInfoTypes" yaml:"selectedInfoTypes"`
}

