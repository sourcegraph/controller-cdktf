package googledatalosspreventiondeidentifytemplate


type GoogleDataLossPreventionDeidentifyTemplateDeidentifyConfigRecordTransformations struct {
	// field_transformations block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_data_loss_prevention_deidentify_template#field_transformations GoogleDataLossPreventionDeidentifyTemplate#field_transformations}
	FieldTransformations interface{} `field:"optional" json:"fieldTransformations" yaml:"fieldTransformations"`
	// record_suppressions block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_data_loss_prevention_deidentify_template#record_suppressions GoogleDataLossPreventionDeidentifyTemplate#record_suppressions}
	RecordSuppressions interface{} `field:"optional" json:"recordSuppressions" yaml:"recordSuppressions"`
}

