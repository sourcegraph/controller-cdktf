package googledatalosspreventiondeidentifytemplate


type GoogleDataLossPreventionDeidentifyTemplateDeidentifyConfig struct {
	// image_transformations block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_data_loss_prevention_deidentify_template#image_transformations GoogleDataLossPreventionDeidentifyTemplate#image_transformations}
	ImageTransformations *GoogleDataLossPreventionDeidentifyTemplateDeidentifyConfigImageTransformations `field:"optional" json:"imageTransformations" yaml:"imageTransformations"`
	// info_type_transformations block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_data_loss_prevention_deidentify_template#info_type_transformations GoogleDataLossPreventionDeidentifyTemplate#info_type_transformations}
	InfoTypeTransformations *GoogleDataLossPreventionDeidentifyTemplateDeidentifyConfigInfoTypeTransformations `field:"optional" json:"infoTypeTransformations" yaml:"infoTypeTransformations"`
	// record_transformations block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_data_loss_prevention_deidentify_template#record_transformations GoogleDataLossPreventionDeidentifyTemplate#record_transformations}
	RecordTransformations *GoogleDataLossPreventionDeidentifyTemplateDeidentifyConfigRecordTransformations `field:"optional" json:"recordTransformations" yaml:"recordTransformations"`
}

