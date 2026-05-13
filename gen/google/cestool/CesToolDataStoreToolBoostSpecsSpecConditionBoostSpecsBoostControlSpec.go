package cestool


type CesToolDataStoreToolBoostSpecsSpecConditionBoostSpecsBoostControlSpec struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/ces_tool#attribute_type CesTool#attribute_type}
	AttributeType *string `field:"optional" json:"attributeType" yaml:"attributeType"`
	// control_points block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/ces_tool#control_points CesTool#control_points}
	ControlPoints interface{} `field:"optional" json:"controlPoints" yaml:"controlPoints"`
	// The name of the field whose value will be used to determine the boost amount.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/ces_tool#field_name CesTool#field_name}
	FieldName *string `field:"optional" json:"fieldName" yaml:"fieldName"`
	// The interpolation type to be applied to connect the control points listed below. Possible values: LINEAR.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/ces_tool#interpolation_type CesTool#interpolation_type}
	InterpolationType *string `field:"optional" json:"interpolationType" yaml:"interpolationType"`
}

