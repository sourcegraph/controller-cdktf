package googlecestool


type GoogleCesToolDataStoreToolBoostSpecsSpecConditionBoostSpecsBoostControlSpec struct {
	// The attribute type to be used to determine the boost amount.
	//
	// The
	// attribute value can be derived from the field value of the specified
	// field_name. In the case of numerical it is straightforward i.e.
	// attribute_value = numerical_field_value. In the case of freshness
	// however, attribute_value = (time.now() - datetime_field_value).
	// Possible values:
	// NUMERICAL
	// FRESHNESS
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_ces_tool#attribute_type GoogleCesTool#attribute_type}
	AttributeType *string `field:"optional" json:"attributeType" yaml:"attributeType"`
	// control_points block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_ces_tool#control_points GoogleCesTool#control_points}
	ControlPoints interface{} `field:"optional" json:"controlPoints" yaml:"controlPoints"`
	// The name of the field whose value will be used to determine the boost amount.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_ces_tool#field_name GoogleCesTool#field_name}
	FieldName *string `field:"optional" json:"fieldName" yaml:"fieldName"`
	// The interpolation type to be applied to connect the control points listed below. Possible values: LINEAR.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_ces_tool#interpolation_type GoogleCesTool#interpolation_type}
	InterpolationType *string `field:"optional" json:"interpolationType" yaml:"interpolationType"`
}

