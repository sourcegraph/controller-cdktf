//go:build !no_runtime_type_checking

package wafv2rulegroup

import (
	"fmt"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

func (w *jsiiProxy_Wafv2RuleGroupRuleStatementRateBasedStatementScopeDownStatementAndStatementStatementAndStatementStatementSqliMatchStatementTextTransformationOutputReference) validateGetAnyMapAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (w *jsiiProxy_Wafv2RuleGroupRuleStatementRateBasedStatementScopeDownStatementAndStatementStatementAndStatementStatementSqliMatchStatementTextTransformationOutputReference) validateGetBooleanAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (w *jsiiProxy_Wafv2RuleGroupRuleStatementRateBasedStatementScopeDownStatementAndStatementStatementAndStatementStatementSqliMatchStatementTextTransformationOutputReference) validateGetBooleanMapAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (w *jsiiProxy_Wafv2RuleGroupRuleStatementRateBasedStatementScopeDownStatementAndStatementStatementAndStatementStatementSqliMatchStatementTextTransformationOutputReference) validateGetListAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (w *jsiiProxy_Wafv2RuleGroupRuleStatementRateBasedStatementScopeDownStatementAndStatementStatementAndStatementStatementSqliMatchStatementTextTransformationOutputReference) validateGetNumberAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (w *jsiiProxy_Wafv2RuleGroupRuleStatementRateBasedStatementScopeDownStatementAndStatementStatementAndStatementStatementSqliMatchStatementTextTransformationOutputReference) validateGetNumberListAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (w *jsiiProxy_Wafv2RuleGroupRuleStatementRateBasedStatementScopeDownStatementAndStatementStatementAndStatementStatementSqliMatchStatementTextTransformationOutputReference) validateGetNumberMapAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (w *jsiiProxy_Wafv2RuleGroupRuleStatementRateBasedStatementScopeDownStatementAndStatementStatementAndStatementStatementSqliMatchStatementTextTransformationOutputReference) validateGetStringAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (w *jsiiProxy_Wafv2RuleGroupRuleStatementRateBasedStatementScopeDownStatementAndStatementStatementAndStatementStatementSqliMatchStatementTextTransformationOutputReference) validateGetStringMapAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (w *jsiiProxy_Wafv2RuleGroupRuleStatementRateBasedStatementScopeDownStatementAndStatementStatementAndStatementStatementSqliMatchStatementTextTransformationOutputReference) validateInterpolationForAttributeParameters(property *string) error {
	if property == nil {
		return fmt.Errorf("parameter property is required, but nil was provided")
	}

	return nil
}

func (w *jsiiProxy_Wafv2RuleGroupRuleStatementRateBasedStatementScopeDownStatementAndStatementStatementAndStatementStatementSqliMatchStatementTextTransformationOutputReference) validateResolveParameters(_context cdktf.IResolveContext) error {
	if _context == nil {
		return fmt.Errorf("parameter _context is required, but nil was provided")
	}

	return nil
}

func (j *jsiiProxy_Wafv2RuleGroupRuleStatementRateBasedStatementScopeDownStatementAndStatementStatementAndStatementStatementSqliMatchStatementTextTransformationOutputReference) validateSetComplexObjectIndexParameters(val interface{}) error {
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

func (j *jsiiProxy_Wafv2RuleGroupRuleStatementRateBasedStatementScopeDownStatementAndStatementStatementAndStatementStatementSqliMatchStatementTextTransformationOutputReference) validateSetComplexObjectIsFromSetParameters(val *bool) error {
	if val == nil {
		return fmt.Errorf("parameter val is required, but nil was provided")
	}

	return nil
}

func (j *jsiiProxy_Wafv2RuleGroupRuleStatementRateBasedStatementScopeDownStatementAndStatementStatementAndStatementStatementSqliMatchStatementTextTransformationOutputReference) validateSetInternalValueParameters(val interface{}) error {
	switch val.(type) {
	case *Wafv2RuleGroupRuleStatementRateBasedStatementScopeDownStatementAndStatementStatementAndStatementStatementSqliMatchStatementTextTransformation:
		val := val.(*Wafv2RuleGroupRuleStatementRateBasedStatementScopeDownStatementAndStatementStatementAndStatementStatementSqliMatchStatementTextTransformation)
		if err := _jsii_.ValidateStruct(val, func() string { return "parameter val" }); err != nil {
			return err
		}
	case Wafv2RuleGroupRuleStatementRateBasedStatementScopeDownStatementAndStatementStatementAndStatementStatementSqliMatchStatementTextTransformation:
		val_ := val.(Wafv2RuleGroupRuleStatementRateBasedStatementScopeDownStatementAndStatementStatementAndStatementStatementSqliMatchStatementTextTransformation)
		val := &val_
		if err := _jsii_.ValidateStruct(val, func() string { return "parameter val" }); err != nil {
			return err
		}
	case cdktf.IResolvable:
		// ok
	default:
		if !_jsii_.IsAnonymousProxy(val) {
			return fmt.Errorf("parameter val must be one of the allowed types: *Wafv2RuleGroupRuleStatementRateBasedStatementScopeDownStatementAndStatementStatementAndStatementStatementSqliMatchStatementTextTransformation, cdktf.IResolvable; received %#v (a %T)", val, val)
		}
	}

	return nil
}

func (j *jsiiProxy_Wafv2RuleGroupRuleStatementRateBasedStatementScopeDownStatementAndStatementStatementAndStatementStatementSqliMatchStatementTextTransformationOutputReference) validateSetPriorityParameters(val *float64) error {
	if val == nil {
		return fmt.Errorf("parameter val is required, but nil was provided")
	}

	return nil
}

func (j *jsiiProxy_Wafv2RuleGroupRuleStatementRateBasedStatementScopeDownStatementAndStatementStatementAndStatementStatementSqliMatchStatementTextTransformationOutputReference) validateSetTerraformAttributeParameters(val *string) error {
	if val == nil {
		return fmt.Errorf("parameter val is required, but nil was provided")
	}

	return nil
}

func (j *jsiiProxy_Wafv2RuleGroupRuleStatementRateBasedStatementScopeDownStatementAndStatementStatementAndStatementStatementSqliMatchStatementTextTransformationOutputReference) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	if val == nil {
		return fmt.Errorf("parameter val is required, but nil was provided")
	}

	return nil
}

func (j *jsiiProxy_Wafv2RuleGroupRuleStatementRateBasedStatementScopeDownStatementAndStatementStatementAndStatementStatementSqliMatchStatementTextTransformationOutputReference) validateSetTypeParameters(val *string) error {
	if val == nil {
		return fmt.Errorf("parameter val is required, but nil was provided")
	}

	return nil
}

func validateNewWafv2RuleGroupRuleStatementRateBasedStatementScopeDownStatementAndStatementStatementAndStatementStatementSqliMatchStatementTextTransformationOutputReferenceParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) error {
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

