//go:build !no_runtime_type_checking

package cloudsecuritycompliancecloudcontrol

import (
	"fmt"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

func (c *jsiiProxy_CloudSecurityComplianceCloudControlParameterSpecSubParametersValidationAllowedValuesValuesOutputReference) validateGetAnyMapAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (c *jsiiProxy_CloudSecurityComplianceCloudControlParameterSpecSubParametersValidationAllowedValuesValuesOutputReference) validateGetBooleanAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (c *jsiiProxy_CloudSecurityComplianceCloudControlParameterSpecSubParametersValidationAllowedValuesValuesOutputReference) validateGetBooleanMapAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (c *jsiiProxy_CloudSecurityComplianceCloudControlParameterSpecSubParametersValidationAllowedValuesValuesOutputReference) validateGetListAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (c *jsiiProxy_CloudSecurityComplianceCloudControlParameterSpecSubParametersValidationAllowedValuesValuesOutputReference) validateGetNumberAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (c *jsiiProxy_CloudSecurityComplianceCloudControlParameterSpecSubParametersValidationAllowedValuesValuesOutputReference) validateGetNumberListAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (c *jsiiProxy_CloudSecurityComplianceCloudControlParameterSpecSubParametersValidationAllowedValuesValuesOutputReference) validateGetNumberMapAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (c *jsiiProxy_CloudSecurityComplianceCloudControlParameterSpecSubParametersValidationAllowedValuesValuesOutputReference) validateGetStringAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (c *jsiiProxy_CloudSecurityComplianceCloudControlParameterSpecSubParametersValidationAllowedValuesValuesOutputReference) validateGetStringMapAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (c *jsiiProxy_CloudSecurityComplianceCloudControlParameterSpecSubParametersValidationAllowedValuesValuesOutputReference) validateInterpolationForAttributeParameters(property *string) error {
	if property == nil {
		return fmt.Errorf("parameter property is required, but nil was provided")
	}

	return nil
}

func (c *jsiiProxy_CloudSecurityComplianceCloudControlParameterSpecSubParametersValidationAllowedValuesValuesOutputReference) validatePutOneofValueParameters(value *CloudSecurityComplianceCloudControlParameterSpecSubParametersValidationAllowedValuesValuesOneofValue) error {
	if value == nil {
		return fmt.Errorf("parameter value is required, but nil was provided")
	}
	if err := _jsii_.ValidateStruct(value, func() string { return "parameter value" }); err != nil {
		return err
	}

	return nil
}

func (c *jsiiProxy_CloudSecurityComplianceCloudControlParameterSpecSubParametersValidationAllowedValuesValuesOutputReference) validatePutStringListValueParameters(value *CloudSecurityComplianceCloudControlParameterSpecSubParametersValidationAllowedValuesValuesStringListValue) error {
	if value == nil {
		return fmt.Errorf("parameter value is required, but nil was provided")
	}
	if err := _jsii_.ValidateStruct(value, func() string { return "parameter value" }); err != nil {
		return err
	}

	return nil
}

func (c *jsiiProxy_CloudSecurityComplianceCloudControlParameterSpecSubParametersValidationAllowedValuesValuesOutputReference) validateResolveParameters(_context cdktf.IResolveContext) error {
	if _context == nil {
		return fmt.Errorf("parameter _context is required, but nil was provided")
	}

	return nil
}

func (j *jsiiProxy_CloudSecurityComplianceCloudControlParameterSpecSubParametersValidationAllowedValuesValuesOutputReference) validateSetBoolValueParameters(val interface{}) error {
	if val == nil {
		return fmt.Errorf("parameter val is required, but nil was provided")
	}
	switch val.(type) {
	case *bool:
		// ok
	case bool:
		// ok
	case cdktf.IResolvable:
		// ok
	default:
		if !_jsii_.IsAnonymousProxy(val) {
			return fmt.Errorf("parameter val must be one of the allowed types: *bool, cdktf.IResolvable; received %#v (a %T)", val, val)
		}
	}

	return nil
}

func (j *jsiiProxy_CloudSecurityComplianceCloudControlParameterSpecSubParametersValidationAllowedValuesValuesOutputReference) validateSetComplexObjectIndexParameters(val interface{}) error {
	switch val.(type) {
	case *string:
		// ok
	case string:
		// ok
	case *float64:
		// ok
	case float64:
		// ok
	case *int:
		// ok
	case int:
		// ok
	case *uint:
		// ok
	case uint:
		// ok
	case *int8:
		// ok
	case int8:
		// ok
	case *int16:
		// ok
	case int16:
		// ok
	case *int32:
		// ok
	case int32:
		// ok
	case *int64:
		// ok
	case int64:
		// ok
	case *uint8:
		// ok
	case uint8:
		// ok
	case *uint16:
		// ok
	case uint16:
		// ok
	case *uint32:
		// ok
	case uint32:
		// ok
	case *uint64:
		// ok
	case uint64:
		// ok
	default:
		return fmt.Errorf("parameter val must be one of the allowed types: *string, *float64; received %#v (a %T)", val, val)
	}

	return nil
}

func (j *jsiiProxy_CloudSecurityComplianceCloudControlParameterSpecSubParametersValidationAllowedValuesValuesOutputReference) validateSetComplexObjectIsFromSetParameters(val *bool) error {
	if val == nil {
		return fmt.Errorf("parameter val is required, but nil was provided")
	}

	return nil
}

func (j *jsiiProxy_CloudSecurityComplianceCloudControlParameterSpecSubParametersValidationAllowedValuesValuesOutputReference) validateSetInternalValueParameters(val interface{}) error {
	switch val.(type) {
	case cdktf.IResolvable:
		// ok
	case *CloudSecurityComplianceCloudControlParameterSpecSubParametersValidationAllowedValuesValues:
		val := val.(*CloudSecurityComplianceCloudControlParameterSpecSubParametersValidationAllowedValuesValues)
		if err := _jsii_.ValidateStruct(val, func() string { return "parameter val" }); err != nil {
			return err
		}
	case CloudSecurityComplianceCloudControlParameterSpecSubParametersValidationAllowedValuesValues:
		val_ := val.(CloudSecurityComplianceCloudControlParameterSpecSubParametersValidationAllowedValuesValues)
		val := &val_
		if err := _jsii_.ValidateStruct(val, func() string { return "parameter val" }); err != nil {
			return err
		}
	default:
		if !_jsii_.IsAnonymousProxy(val) {
			return fmt.Errorf("parameter val must be one of the allowed types: cdktf.IResolvable, *CloudSecurityComplianceCloudControlParameterSpecSubParametersValidationAllowedValuesValues; received %#v (a %T)", val, val)
		}
	}

	return nil
}

func (j *jsiiProxy_CloudSecurityComplianceCloudControlParameterSpecSubParametersValidationAllowedValuesValuesOutputReference) validateSetNumberValueParameters(val *float64) error {
	if val == nil {
		return fmt.Errorf("parameter val is required, but nil was provided")
	}

	return nil
}

func (j *jsiiProxy_CloudSecurityComplianceCloudControlParameterSpecSubParametersValidationAllowedValuesValuesOutputReference) validateSetStringValueParameters(val *string) error {
	if val == nil {
		return fmt.Errorf("parameter val is required, but nil was provided")
	}

	return nil
}

func (j *jsiiProxy_CloudSecurityComplianceCloudControlParameterSpecSubParametersValidationAllowedValuesValuesOutputReference) validateSetTerraformAttributeParameters(val *string) error {
	if val == nil {
		return fmt.Errorf("parameter val is required, but nil was provided")
	}

	return nil
}

func (j *jsiiProxy_CloudSecurityComplianceCloudControlParameterSpecSubParametersValidationAllowedValuesValuesOutputReference) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	if val == nil {
		return fmt.Errorf("parameter val is required, but nil was provided")
	}

	return nil
}

func validateNewCloudSecurityComplianceCloudControlParameterSpecSubParametersValidationAllowedValuesValuesOutputReferenceParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) error {
	if terraformResource == nil {
		return fmt.Errorf("parameter terraformResource is required, but nil was provided")
	}

	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	if complexObjectIndex == nil {
		return fmt.Errorf("parameter complexObjectIndex is required, but nil was provided")
	}

	if complexObjectIsFromSet == nil {
		return fmt.Errorf("parameter complexObjectIsFromSet is required, but nil was provided")
	}

	return nil
}

