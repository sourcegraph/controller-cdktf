//go:build !no_runtime_type_checking

package googleaccesscontextmanagerserviceperimeters

import (
	"fmt"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

func (g *jsiiProxy_GoogleAccessContextManagerServicePerimetersServicePerimetersStatusEgressPoliciesEgressToOperationsOutputReference) validateGetAnyMapAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (g *jsiiProxy_GoogleAccessContextManagerServicePerimetersServicePerimetersStatusEgressPoliciesEgressToOperationsOutputReference) validateGetBooleanAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (g *jsiiProxy_GoogleAccessContextManagerServicePerimetersServicePerimetersStatusEgressPoliciesEgressToOperationsOutputReference) validateGetBooleanMapAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (g *jsiiProxy_GoogleAccessContextManagerServicePerimetersServicePerimetersStatusEgressPoliciesEgressToOperationsOutputReference) validateGetListAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (g *jsiiProxy_GoogleAccessContextManagerServicePerimetersServicePerimetersStatusEgressPoliciesEgressToOperationsOutputReference) validateGetNumberAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (g *jsiiProxy_GoogleAccessContextManagerServicePerimetersServicePerimetersStatusEgressPoliciesEgressToOperationsOutputReference) validateGetNumberListAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (g *jsiiProxy_GoogleAccessContextManagerServicePerimetersServicePerimetersStatusEgressPoliciesEgressToOperationsOutputReference) validateGetNumberMapAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (g *jsiiProxy_GoogleAccessContextManagerServicePerimetersServicePerimetersStatusEgressPoliciesEgressToOperationsOutputReference) validateGetStringAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (g *jsiiProxy_GoogleAccessContextManagerServicePerimetersServicePerimetersStatusEgressPoliciesEgressToOperationsOutputReference) validateGetStringMapAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (g *jsiiProxy_GoogleAccessContextManagerServicePerimetersServicePerimetersStatusEgressPoliciesEgressToOperationsOutputReference) validateInterpolationForAttributeParameters(property *string) error {
	if property == nil {
		return fmt.Errorf("parameter property is required, but nil was provided")
	}

	return nil
}

func (g *jsiiProxy_GoogleAccessContextManagerServicePerimetersServicePerimetersStatusEgressPoliciesEgressToOperationsOutputReference) validatePutMethodSelectorsParameters(value interface{}) error {
	if value == nil {
		return fmt.Errorf("parameter value is required, but nil was provided")
	}
	switch value.(type) {
	case cdktf.IResolvable:
		// ok
	case *[]*GoogleAccessContextManagerServicePerimetersServicePerimetersStatusEgressPoliciesEgressToOperationsMethodSelectors:
		value := value.(*[]*GoogleAccessContextManagerServicePerimetersServicePerimetersStatusEgressPoliciesEgressToOperationsMethodSelectors)
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	case []*GoogleAccessContextManagerServicePerimetersServicePerimetersStatusEgressPoliciesEgressToOperationsMethodSelectors:
		value_ := value.([]*GoogleAccessContextManagerServicePerimetersServicePerimetersStatusEgressPoliciesEgressToOperationsMethodSelectors)
		value := &value_
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	default:
		if !_jsii_.IsAnonymousProxy(value) {
			return fmt.Errorf("parameter value must be one of the allowed types: cdktf.IResolvable, *[]*GoogleAccessContextManagerServicePerimetersServicePerimetersStatusEgressPoliciesEgressToOperationsMethodSelectors; received %#v (a %T)", value, value)
		}
	}

	return nil
}

func (g *jsiiProxy_GoogleAccessContextManagerServicePerimetersServicePerimetersStatusEgressPoliciesEgressToOperationsOutputReference) validateResolveParameters(_context cdktf.IResolveContext) error {
	if _context == nil {
		return fmt.Errorf("parameter _context is required, but nil was provided")
	}

	return nil
}

func (j *jsiiProxy_GoogleAccessContextManagerServicePerimetersServicePerimetersStatusEgressPoliciesEgressToOperationsOutputReference) validateSetComplexObjectIndexParameters(val interface{}) error {
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

func (j *jsiiProxy_GoogleAccessContextManagerServicePerimetersServicePerimetersStatusEgressPoliciesEgressToOperationsOutputReference) validateSetComplexObjectIsFromSetParameters(val *bool) error {
	if val == nil {
		return fmt.Errorf("parameter val is required, but nil was provided")
	}

	return nil
}

func (j *jsiiProxy_GoogleAccessContextManagerServicePerimetersServicePerimetersStatusEgressPoliciesEgressToOperationsOutputReference) validateSetInternalValueParameters(val interface{}) error {
	switch val.(type) {
	case cdktf.IResolvable:
		// ok
	case *GoogleAccessContextManagerServicePerimetersServicePerimetersStatusEgressPoliciesEgressToOperations:
		val := val.(*GoogleAccessContextManagerServicePerimetersServicePerimetersStatusEgressPoliciesEgressToOperations)
		if err := _jsii_.ValidateStruct(val, func() string { return "parameter val" }); err != nil {
			return err
		}
	case GoogleAccessContextManagerServicePerimetersServicePerimetersStatusEgressPoliciesEgressToOperations:
		val_ := val.(GoogleAccessContextManagerServicePerimetersServicePerimetersStatusEgressPoliciesEgressToOperations)
		val := &val_
		if err := _jsii_.ValidateStruct(val, func() string { return "parameter val" }); err != nil {
			return err
		}
	default:
		if !_jsii_.IsAnonymousProxy(val) {
			return fmt.Errorf("parameter val must be one of the allowed types: cdktf.IResolvable, *GoogleAccessContextManagerServicePerimetersServicePerimetersStatusEgressPoliciesEgressToOperations; received %#v (a %T)", val, val)
		}
	}

	return nil
}

func (j *jsiiProxy_GoogleAccessContextManagerServicePerimetersServicePerimetersStatusEgressPoliciesEgressToOperationsOutputReference) validateSetServiceNameParameters(val *string) error {
	if val == nil {
		return fmt.Errorf("parameter val is required, but nil was provided")
	}

	return nil
}

func (j *jsiiProxy_GoogleAccessContextManagerServicePerimetersServicePerimetersStatusEgressPoliciesEgressToOperationsOutputReference) validateSetTerraformAttributeParameters(val *string) error {
	if val == nil {
		return fmt.Errorf("parameter val is required, but nil was provided")
	}

	return nil
}

func (j *jsiiProxy_GoogleAccessContextManagerServicePerimetersServicePerimetersStatusEgressPoliciesEgressToOperationsOutputReference) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	if val == nil {
		return fmt.Errorf("parameter val is required, but nil was provided")
	}

	return nil
}

func validateNewGoogleAccessContextManagerServicePerimetersServicePerimetersStatusEgressPoliciesEgressToOperationsOutputReferenceParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) error {
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

