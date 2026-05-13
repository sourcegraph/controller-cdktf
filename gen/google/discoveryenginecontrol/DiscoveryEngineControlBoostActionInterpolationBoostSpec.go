package discoveryenginecontrol


type DiscoveryEngineControlBoostActionInterpolationBoostSpec struct {
	// The attribute type to be used to determine the boost amount. Possible values: ["NUMERICAL", "FRESHNESS"].
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/discovery_engine_control#attribute_type DiscoveryEngineControl#attribute_type}
	AttributeType *string `field:"optional" json:"attributeType" yaml:"attributeType"`
	// control_point block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/discovery_engine_control#control_point DiscoveryEngineControl#control_point}
	ControlPoint *DiscoveryEngineControlBoostActionInterpolationBoostSpecControlPoint `field:"optional" json:"controlPoint" yaml:"controlPoint"`
	// The name of the field whose value will be used to determine the boost amount.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/discovery_engine_control#field_name DiscoveryEngineControl#field_name}
	FieldName *string `field:"optional" json:"fieldName" yaml:"fieldName"`
	// The interpolation type to be applied to connect the control points. Possible values: ["LINEAR"].
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/discovery_engine_control#interpolation_type DiscoveryEngineControl#interpolation_type}
	InterpolationType *string `field:"optional" json:"interpolationType" yaml:"interpolationType"`
}

