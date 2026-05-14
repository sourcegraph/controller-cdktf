//go:build !no_runtime_type_checking

package computeorganizationsecuritypolicyrule

import (
	"fmt"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

func (c *jsiiProxy_ComputeOrganizationSecurityPolicyRulePreconfiguredWafConfigExclusionOutputReference) validateGetAnyMapAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (c *jsiiProxy_ComputeOrganizationSecurityPolicyRulePreconfiguredWafConfigExclusionOutputReference) validateGetBooleanAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (c *jsiiProxy_ComputeOrganizationSecurityPolicyRulePreconfiguredWafConfigExclusionOutputReference) validateGetBooleanMapAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (c *jsiiProxy_ComputeOrganizationSecurityPolicyRulePreconfiguredWafConfigExclusionOutputReference) validateGetListAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (c *jsiiProxy_ComputeOrganizationSecurityPolicyRulePreconfiguredWafConfigExclusionOutputReference) validateGetNumberAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (c *jsiiProxy_ComputeOrganizationSecurityPolicyRulePreconfiguredWafConfigExclusionOutputReference) validateGetNumberListAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (c *jsiiProxy_ComputeOrganizationSecurityPolicyRulePreconfiguredWafConfigExclusionOutputReference) validateGetNumberMapAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (c *jsiiProxy_ComputeOrganizationSecurityPolicyRulePreconfiguredWafConfigExclusionOutputReference) validateGetStringAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (c *jsiiProxy_ComputeOrganizationSecurityPolicyRulePreconfiguredWafConfigExclusionOutputReference) validateGetStringMapAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (c *jsiiProxy_ComputeOrganizationSecurityPolicyRulePreconfiguredWafConfigExclusionOutputReference) validateInterpolationForAttributeParameters(property *string) error {
	if property == nil {
		return fmt.Errorf("parameter property is required, but nil was provided")
	}

	return nil
}

func (c *jsiiProxy_ComputeOrganizationSecurityPolicyRulePreconfiguredWafConfigExclusionOutputReference) validatePutRequestCookieParameters(value interface{}) error {
	if value == nil {
		return fmt.Errorf("parameter value is required, but nil was provided")
	}
	switch value.(type) {
	case cdktf.IResolvable:
		// ok
	case *[]*ComputeOrganizationSecurityPolicyRulePreconfiguredWafConfigExclusionRequestCookie:
		value := value.(*[]*ComputeOrganizationSecurityPolicyRulePreconfiguredWafConfigExclusionRequestCookie)
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	case []*ComputeOrganizationSecurityPolicyRulePreconfiguredWafConfigExclusionRequestCookie:
		value_ := value.([]*ComputeOrganizationSecurityPolicyRulePreconfiguredWafConfigExclusionRequestCookie)
		value := &value_
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	default:
		if !_jsii_.IsAnonymousProxy(value) {
			return fmt.Errorf("parameter value must be one of the allowed types: cdktf.IResolvable, *[]*ComputeOrganizationSecurityPolicyRulePreconfiguredWafConfigExclusionRequestCookie; received %#v (a %T)", value, value)
		}
	}

	return nil
}

func (c *jsiiProxy_ComputeOrganizationSecurityPolicyRulePreconfiguredWafConfigExclusionOutputReference) validatePutRequestHeaderParameters(value interface{}) error {
	if value == nil {
		return fmt.Errorf("parameter value is required, but nil was provided")
	}
	switch value.(type) {
	case cdktf.IResolvable:
		// ok
	case *[]*ComputeOrganizationSecurityPolicyRulePreconfiguredWafConfigExclusionRequestHeader:
		value := value.(*[]*ComputeOrganizationSecurityPolicyRulePreconfiguredWafConfigExclusionRequestHeader)
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	case []*ComputeOrganizationSecurityPolicyRulePreconfiguredWafConfigExclusionRequestHeader:
		value_ := value.([]*ComputeOrganizationSecurityPolicyRulePreconfiguredWafConfigExclusionRequestHeader)
		value := &value_
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	default:
		if !_jsii_.IsAnonymousProxy(value) {
			return fmt.Errorf("parameter value must be one of the allowed types: cdktf.IResolvable, *[]*ComputeOrganizationSecurityPolicyRulePreconfiguredWafConfigExclusionRequestHeader; received %#v (a %T)", value, value)
		}
	}

	return nil
}

func (c *jsiiProxy_ComputeOrganizationSecurityPolicyRulePreconfiguredWafConfigExclusionOutputReference) validatePutRequestQueryParamParameters(value interface{}) error {
	if value == nil {
		return fmt.Errorf("parameter value is required, but nil was provided")
	}
	switch value.(type) {
	case cdktf.IResolvable:
		// ok
	case *[]*ComputeOrganizationSecurityPolicyRulePreconfiguredWafConfigExclusionRequestQueryParam:
		value := value.(*[]*ComputeOrganizationSecurityPolicyRulePreconfiguredWafConfigExclusionRequestQueryParam)
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	case []*ComputeOrganizationSecurityPolicyRulePreconfiguredWafConfigExclusionRequestQueryParam:
		value_ := value.([]*ComputeOrganizationSecurityPolicyRulePreconfiguredWafConfigExclusionRequestQueryParam)
		value := &value_
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	default:
		if !_jsii_.IsAnonymousProxy(value) {
			return fmt.Errorf("parameter value must be one of the allowed types: cdktf.IResolvable, *[]*ComputeOrganizationSecurityPolicyRulePreconfiguredWafConfigExclusionRequestQueryParam; received %#v (a %T)", value, value)
		}
	}

	return nil
}

func (c *jsiiProxy_ComputeOrganizationSecurityPolicyRulePreconfiguredWafConfigExclusionOutputReference) validatePutRequestUriParameters(value interface{}) error {
	if value == nil {
		return fmt.Errorf("parameter value is required, but nil was provided")
	}
	switch value.(type) {
	case cdktf.IResolvable:
		// ok
	case *[]*ComputeOrganizationSecurityPolicyRulePreconfiguredWafConfigExclusionRequestUri:
		value := value.(*[]*ComputeOrganizationSecurityPolicyRulePreconfiguredWafConfigExclusionRequestUri)
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	case []*ComputeOrganizationSecurityPolicyRulePreconfiguredWafConfigExclusionRequestUri:
		value_ := value.([]*ComputeOrganizationSecurityPolicyRulePreconfiguredWafConfigExclusionRequestUri)
		value := &value_
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	default:
		if !_jsii_.IsAnonymousProxy(value) {
			return fmt.Errorf("parameter value must be one of the allowed types: cdktf.IResolvable, *[]*ComputeOrganizationSecurityPolicyRulePreconfiguredWafConfigExclusionRequestUri; received %#v (a %T)", value, value)
		}
	}

	return nil
}

func (c *jsiiProxy_ComputeOrganizationSecurityPolicyRulePreconfiguredWafConfigExclusionOutputReference) validateResolveParameters(_context cdktf.IResolveContext) error {
	if _context == nil {
		return fmt.Errorf("parameter _context is required, but nil was provided")
	}

	return nil
}

func (j *jsiiProxy_ComputeOrganizationSecurityPolicyRulePreconfiguredWafConfigExclusionOutputReference) validateSetComplexObjectIndexParameters(val interface{}) error {
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

func (j *jsiiProxy_ComputeOrganizationSecurityPolicyRulePreconfiguredWafConfigExclusionOutputReference) validateSetComplexObjectIsFromSetParameters(val *bool) error {
	if val == nil {
		return fmt.Errorf("parameter val is required, but nil was provided")
	}

	return nil
}

func (j *jsiiProxy_ComputeOrganizationSecurityPolicyRulePreconfiguredWafConfigExclusionOutputReference) validateSetInternalValueParameters(val interface{}) error {
	switch val.(type) {
	case cdktf.IResolvable:
		// ok
	case *ComputeOrganizationSecurityPolicyRulePreconfiguredWafConfigExclusion:
		val := val.(*ComputeOrganizationSecurityPolicyRulePreconfiguredWafConfigExclusion)
		if err := _jsii_.ValidateStruct(val, func() string { return "parameter val" }); err != nil {
			return err
		}
	case ComputeOrganizationSecurityPolicyRulePreconfiguredWafConfigExclusion:
		val_ := val.(ComputeOrganizationSecurityPolicyRulePreconfiguredWafConfigExclusion)
		val := &val_
		if err := _jsii_.ValidateStruct(val, func() string { return "parameter val" }); err != nil {
			return err
		}
	default:
		if !_jsii_.IsAnonymousProxy(val) {
			return fmt.Errorf("parameter val must be one of the allowed types: cdktf.IResolvable, *ComputeOrganizationSecurityPolicyRulePreconfiguredWafConfigExclusion; received %#v (a %T)", val, val)
		}
	}

	return nil
}

func (j *jsiiProxy_ComputeOrganizationSecurityPolicyRulePreconfiguredWafConfigExclusionOutputReference) validateSetTargetRuleIdsParameters(val *[]*string) error {
	if val == nil {
		return fmt.Errorf("parameter val is required, but nil was provided")
	}

	return nil
}

func (j *jsiiProxy_ComputeOrganizationSecurityPolicyRulePreconfiguredWafConfigExclusionOutputReference) validateSetTargetRuleSetParameters(val *string) error {
	if val == nil {
		return fmt.Errorf("parameter val is required, but nil was provided")
	}

	return nil
}

func (j *jsiiProxy_ComputeOrganizationSecurityPolicyRulePreconfiguredWafConfigExclusionOutputReference) validateSetTerraformAttributeParameters(val *string) error {
	if val == nil {
		return fmt.Errorf("parameter val is required, but nil was provided")
	}

	return nil
}

func (j *jsiiProxy_ComputeOrganizationSecurityPolicyRulePreconfiguredWafConfigExclusionOutputReference) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	if val == nil {
		return fmt.Errorf("parameter val is required, but nil was provided")
	}

	return nil
}

func validateNewComputeOrganizationSecurityPolicyRulePreconfiguredWafConfigExclusionOutputReferenceParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) error {
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

