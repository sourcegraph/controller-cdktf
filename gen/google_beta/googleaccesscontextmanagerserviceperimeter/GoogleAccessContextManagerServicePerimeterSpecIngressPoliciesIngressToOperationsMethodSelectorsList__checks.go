//go:build !no_runtime_type_checking

package googleaccesscontextmanagerserviceperimeter

import (
	"fmt"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

func (g *jsiiProxy_GoogleAccessContextManagerServicePerimeterSpecIngressPoliciesIngressToOperationsMethodSelectorsList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	if mapKeyAttributeName == nil {
		return fmt.Errorf("parameter mapKeyAttributeName is required, but nil was provided")
	}

	return nil
}

func (g *jsiiProxy_GoogleAccessContextManagerServicePerimeterSpecIngressPoliciesIngressToOperationsMethodSelectorsList) validateGetParameters(index *float64) error {
	if index == nil {
		return fmt.Errorf("parameter index is required, but nil was provided")
	}

	return nil
}

func (g *jsiiProxy_GoogleAccessContextManagerServicePerimeterSpecIngressPoliciesIngressToOperationsMethodSelectorsList) validateResolveParameters(_context cdktf.IResolveContext) error {
	if _context == nil {
		return fmt.Errorf("parameter _context is required, but nil was provided")
	}

	return nil
}

func (j *jsiiProxy_GoogleAccessContextManagerServicePerimeterSpecIngressPoliciesIngressToOperationsMethodSelectorsList) validateSetInternalValueParameters(val interface{}) error {
	switch val.(type) {
	case cdktf.IResolvable:
		// ok
	case *[]*GoogleAccessContextManagerServicePerimeterSpecIngressPoliciesIngressToOperationsMethodSelectors:
		val := val.(*[]*GoogleAccessContextManagerServicePerimeterSpecIngressPoliciesIngressToOperationsMethodSelectors)
		for idx_97dfc6, v := range *val {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter val[%#v]", idx_97dfc6) }); err != nil {
				return err
			}
		}
	case []*GoogleAccessContextManagerServicePerimeterSpecIngressPoliciesIngressToOperationsMethodSelectors:
		val_ := val.([]*GoogleAccessContextManagerServicePerimeterSpecIngressPoliciesIngressToOperationsMethodSelectors)
		val := &val_
		for idx_97dfc6, v := range *val {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter val[%#v]", idx_97dfc6) }); err != nil {
				return err
			}
		}
	default:
		if !_jsii_.IsAnonymousProxy(val) {
			return fmt.Errorf("parameter val must be one of the allowed types: cdktf.IResolvable, *[]*GoogleAccessContextManagerServicePerimeterSpecIngressPoliciesIngressToOperationsMethodSelectors; received %#v (a %T)", val, val)
		}
	}

	return nil
}

func (j *jsiiProxy_GoogleAccessContextManagerServicePerimeterSpecIngressPoliciesIngressToOperationsMethodSelectorsList) validateSetTerraformAttributeParameters(val *string) error {
	if val == nil {
		return fmt.Errorf("parameter val is required, but nil was provided")
	}

	return nil
}

func (j *jsiiProxy_GoogleAccessContextManagerServicePerimeterSpecIngressPoliciesIngressToOperationsMethodSelectorsList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	if val == nil {
		return fmt.Errorf("parameter val is required, but nil was provided")
	}

	return nil
}

func (j *jsiiProxy_GoogleAccessContextManagerServicePerimeterSpecIngressPoliciesIngressToOperationsMethodSelectorsList) validateSetWrapsSetParameters(val *bool) error {
	if val == nil {
		return fmt.Errorf("parameter val is required, but nil was provided")
	}

	return nil
}

func validateNewGoogleAccessContextManagerServicePerimeterSpecIngressPoliciesIngressToOperationsMethodSelectorsListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	if terraformResource == nil {
		return fmt.Errorf("parameter terraformResource is required, but nil was provided")
	}

	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	if wrapsSet == nil {
		return fmt.Errorf("parameter wrapsSet is required, but nil was provided")
	}

	return nil
}

