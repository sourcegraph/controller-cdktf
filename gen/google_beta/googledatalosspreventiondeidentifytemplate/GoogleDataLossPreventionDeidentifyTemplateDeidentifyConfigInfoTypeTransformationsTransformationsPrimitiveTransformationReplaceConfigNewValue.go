package googledatalosspreventiondeidentifytemplate


type GoogleDataLossPreventionDeidentifyTemplateDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformationReplaceConfigNewValue struct {
	// A boolean value.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_data_loss_prevention_deidentify_template#boolean_value GoogleDataLossPreventionDeidentifyTemplate#boolean_value}
	BooleanValue interface{} `field:"optional" json:"booleanValue" yaml:"booleanValue"`
	// date_value block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_data_loss_prevention_deidentify_template#date_value GoogleDataLossPreventionDeidentifyTemplate#date_value}
	DateValue *GoogleDataLossPreventionDeidentifyTemplateDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformationReplaceConfigNewValueDateValue `field:"optional" json:"dateValue" yaml:"dateValue"`
	// Represents a day of the week. Possible values: ["MONDAY", "TUESDAY", "WEDNESDAY", "THURSDAY", "FRIDAY", "SATURDAY", "SUNDAY"].
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_data_loss_prevention_deidentify_template#day_of_week_value GoogleDataLossPreventionDeidentifyTemplate#day_of_week_value}
	DayOfWeekValue *string `field:"optional" json:"dayOfWeekValue" yaml:"dayOfWeekValue"`
	// A float value.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_data_loss_prevention_deidentify_template#float_value GoogleDataLossPreventionDeidentifyTemplate#float_value}
	FloatValue *float64 `field:"optional" json:"floatValue" yaml:"floatValue"`
	// An integer value.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_data_loss_prevention_deidentify_template#integer_value GoogleDataLossPreventionDeidentifyTemplate#integer_value}
	IntegerValue *float64 `field:"optional" json:"integerValue" yaml:"integerValue"`
	// A string value.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_data_loss_prevention_deidentify_template#string_value GoogleDataLossPreventionDeidentifyTemplate#string_value}
	StringValue *string `field:"optional" json:"stringValue" yaml:"stringValue"`
	// A timestamp in RFC3339 UTC "Zulu" format, with nanosecond resolution and up to nine fractional digits. Examples: "2014-10-02T15:01:23Z" and "2014-10-02T15:01:23.045123456Z".
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_data_loss_prevention_deidentify_template#timestamp_value GoogleDataLossPreventionDeidentifyTemplate#timestamp_value}
	TimestampValue *string `field:"optional" json:"timestampValue" yaml:"timestampValue"`
	// time_value block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_data_loss_prevention_deidentify_template#time_value GoogleDataLossPreventionDeidentifyTemplate#time_value}
	TimeValue *GoogleDataLossPreventionDeidentifyTemplateDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformationReplaceConfigNewValueTimeValue `field:"optional" json:"timeValue" yaml:"timeValue"`
}

