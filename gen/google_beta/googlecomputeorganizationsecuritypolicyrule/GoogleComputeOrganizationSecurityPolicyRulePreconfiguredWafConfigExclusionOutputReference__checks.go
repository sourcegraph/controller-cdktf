//go:build !no_runtime_type_checking

package googlecomputeorganizationsecuritypolicyrule

import (
	"fmt"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

func (g *jsiiProxy_GoogleComputeOrganizationSecurityPolicyRulePreconfiguredWafConfigExclusionOutputReference) validateGetAnyMapAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (g *jsiiProxy_GoogleComputeOrganizationSecurityPolicyRulePreconfiguredWafConfigExclusionOutputReference) validateGetBooleanAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (g *jsiiProxy_GoogleComputeOrganizationSecurityPolicyRulePreconfiguredWafConfigExclusionOutputReference) validateGetBooleanMapAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (g *jsiiProxy_GoogleComputeOrganizationSecurityPolicyRulePreconfiguredWafConfigExclusionOutputReference) validateGetListAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (g *jsiiProxy_GoogleComputeOrganizationSecurityPolicyRulePreconfiguredWafConfigExclusionOutputReference) validateGetNumberAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (g *jsiiProxy_GoogleComputeOrganizationSecurityPolicyRulePreconfiguredWafConfigExclusionOutputReference) validateGetNumberListAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (g *jsiiProxy_GoogleComputeOrganizationSecurityPolicyRulePreconfiguredWafConfigExclusionOutputReference) validateGetNumberMapAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (g *jsiiProxy_GoogleComputeOrganizationSecurityPolicyRulePreconfiguredWafConfigExclusionOutputReference) validateGetStringAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (g *jsiiProxy_GoogleComputeOrganizationSecurityPolicyRulePreconfiguredWafConfigExclusionOutputReference) validateGetStringMapAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (g *jsiiProxy_GoogleComputeOrganizationSecurityPolicyRulePreconfiguredWafConfigExclusionOutputReference) validateInterpolationForAttributeParameters(property *string) error {
	if property == nil {
		return fmt.Errorf("parameter property is required, but nil was provided")
	}

	return nil
}

func (g *jsiiProxy_GoogleComputeOrganizationSecurityPolicyRulePreconfiguredWafConfigExclusionOutputReference) validatePutRequestCookieParameters(value interface{}) error {
	if value == nil {
		return fmt.Errorf("parameter value is required, but nil was provided")
	}
	switch value.(type) {
	case cdktf.IResolvable:
		// ok
	case *[]*GoogleComputeOrganizationSecurityPolicyRulePreconfiguredWafConfigExclusionRequestCookie:
		value := value.(*[]*GoogleComputeOrganizationSecurityPolicyRulePreconfiguredWafConfigExclusionRequestCookie)
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	case []*GoogleComputeOrganizationSecurityPolicyRulePreconfiguredWafConfigExclusionRequestCookie:
		value_ := value.([]*GoogleComputeOrganizationSecurityPolicyRulePreconfiguredWafConfigExclusionRequestCookie)
		value := &value_
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	default:
		if !_jsii_.IsAnonymousProxy(value) {
			return fmt.Errorf("parameter value must be one of the allowed types: cdktf.IResolvable, *[]*GoogleComputeOrganizationSecurityPolicyRulePreconfiguredWafConfigExclusionRequestCookie; received %#v (a %T)", value, value)
		}
	}

	return nil
}

func (g *jsiiProxy_GoogleComputeOrganizationSecurityPolicyRulePreconfiguredWafConfigExclusionOutputReference) validatePutRequestHeaderParameters(value interface{}) error {
	if value == nil {
		return fmt.Errorf("parameter value is required, but nil was provided")
	}
	switch value.(type) {
	case cdktf.IResolvable:
		// ok
	case *[]*GoogleComputeOrganizationSecurityPolicyRulePreconfiguredWafConfigExclusionRequestHeader:
		value := value.(*[]*GoogleComputeOrganizationSecurityPolicyRulePreconfiguredWafConfigExclusionRequestHeader)
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	case []*GoogleComputeOrganizationSecurityPolicyRulePreconfiguredWafConfigExclusionRequestHeader:
		value_ := value.([]*GoogleComputeOrganizationSecurityPolicyRulePreconfiguredWafConfigExclusionRequestHeader)
		value := &value_
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	default:
		if !_jsii_.IsAnonymousProxy(value) {
			return fmt.Errorf("parameter value must be one of the allowed types: cdktf.IResolvable, *[]*GoogleComputeOrganizationSecurityPolicyRulePreconfiguredWafConfigExclusionRequestHeader; received %#v (a %T)", value, value)
		}
	}

	return nil
}

func (g *jsiiProxy_GoogleComputeOrganizationSecurityPolicyRulePreconfiguredWafConfigExclusionOutputReference) validatePutRequestQueryParamParameters(value interface{}) error {
	if value == nil {
		return fmt.Errorf("parameter value is required, but nil was provided")
	}
	switch value.(type) {
	case cdktf.IResolvable:
		// ok
	case *[]*GoogleComputeOrganizationSecurityPolicyRulePreconfiguredWafConfigExclusionRequestQueryParam:
		value := value.(*[]*GoogleComputeOrganizationSecurityPolicyRulePreconfiguredWafConfigExclusionRequestQueryParam)
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	case []*GoogleComputeOrganizationSecurityPolicyRulePreconfiguredWafConfigExclusionRequestQueryParam:
		value_ := value.([]*GoogleComputeOrganizationSecurityPolicyRulePreconfiguredWafConfigExclusionRequestQueryParam)
		value := &value_
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	default:
		if !_jsii_.IsAnonymousProxy(value) {
			return fmt.Errorf("parameter value must be one of the allowed types: cdktf.IResolvable, *[]*GoogleComputeOrganizationSecurityPolicyRulePreconfiguredWafConfigExclusionRequestQueryParam; received %#v (a %T)", value, value)
		}
	}

	return nil
}

func (g *jsiiProxy_GoogleComputeOrganizationSecurityPolicyRulePreconfiguredWafConfigExclusionOutputReference) validatePutRequestUriParameters(value interface{}) error {
	if value == nil {
		return fmt.Errorf("parameter value is required, but nil was provided")
	}
	switch value.(type) {
	case cdktf.IResolvable:
		// ok
	case *[]*GoogleComputeOrganizationSecurityPolicyRulePreconfiguredWafConfigExclusionRequestUri:
		value := value.(*[]*GoogleComputeOrganizationSecurityPolicyRulePreconfiguredWafConfigExclusionRequestUri)
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	case []*GoogleComputeOrganizationSecurityPolicyRulePreconfiguredWafConfigExclusionRequestUri:
		value_ := value.([]*GoogleComputeOrganizationSecurityPolicyRulePreconfiguredWafConfigExclusionRequestUri)
		value := &value_
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	default:
		if !_jsii_.IsAnonymousProxy(value) {
			return fmt.Errorf("parameter value must be one of the allowed types: cdktf.IResolvable, *[]*GoogleComputeOrganizationSecurityPolicyRulePreconfiguredWafConfigExclusionRequestUri; received %#v (a %T)", value, value)
		}
	}

	return nil
}

func (g *jsiiProxy_GoogleComputeOrganizationSecurityPolicyRulePreconfiguredWafConfigExclusionOutputReference) validateResolveParameters(_context cdktf.IResolveContext) error {
	if _context == nil {
		return fmt.Errorf("parameter _context is required, but nil was provided")
	}

	return nil
}

func (j *jsiiProxy_GoogleComputeOrganizationSecurityPolicyRulePreconfiguredWafConfigExclusionOutputReference) validateSetComplexObjectIndexParameters(val interface{}) error {
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

func (j *jsiiProxy_GoogleComputeOrganizationSecurityPolicyRulePreconfiguredWafConfigExclusionOutputReference) validateSetComplexObjectIsFromSetParameters(val *bool) error {
	if val == nil {
		return fmt.Errorf("parameter val is required, but nil was provided")
	}

	return nil
}

func (j *jsiiProxy_GoogleComputeOrganizationSecurityPolicyRulePreconfiguredWafConfigExclusionOutputReference) validateSetInternalValueParameters(val interface{}) error {
	switch val.(type) {
	case cdktf.IResolvable:
		// ok
	case *GoogleComputeOrganizationSecurityPolicyRulePreconfiguredWafConfigExclusion:
		val := val.(*GoogleComputeOrganizationSecurityPolicyRulePreconfiguredWafConfigExclusion)
		if err := _jsii_.ValidateStruct(val, func() string { return "parameter val" }); err != nil {
			return err
		}
	case GoogleComputeOrganizationSecurityPolicyRulePreconfiguredWafConfigExclusion:
		val_ := val.(GoogleComputeOrganizationSecurityPolicyRulePreconfiguredWafConfigExclusion)
		val := &val_
		if err := _jsii_.ValidateStruct(val, func() string { return "parameter val" }); err != nil {
			return err
		}
	default:
		if !_jsii_.IsAnonymousProxy(val) {
			return fmt.Errorf("parameter val must be one of the allowed types: cdktf.IResolvable, *GoogleComputeOrganizationSecurityPolicyRulePreconfiguredWafConfigExclusion; received %#v (a %T)", val, val)
		}
	}

	return nil
}

func (j *jsiiProxy_GoogleComputeOrganizationSecurityPolicyRulePreconfiguredWafConfigExclusionOutputReference) validateSetTargetRuleIdsParameters(val *[]*string) error {
	if val == nil {
		return fmt.Errorf("parameter val is required, but nil was provided")
	}

	return nil
}

func (j *jsiiProxy_GoogleComputeOrganizationSecurityPolicyRulePreconfiguredWafConfigExclusionOutputReference) validateSetTargetRuleSetParameters(val *string) error {
	if val == nil {
		return fmt.Errorf("parameter val is required, but nil was provided")
	}

	return nil
}

func (j *jsiiProxy_GoogleComputeOrganizationSecurityPolicyRulePreconfiguredWafConfigExclusionOutputReference) validateSetTerraformAttributeParameters(val *string) error {
	if val == nil {
		return fmt.Errorf("parameter val is required, but nil was provided")
	}

	return nil
}

func (j *jsiiProxy_GoogleComputeOrganizationSecurityPolicyRulePreconfiguredWafConfigExclusionOutputReference) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	if val == nil {
		return fmt.Errorf("parameter val is required, but nil was provided")
	}

	return nil
}

func validateNewGoogleComputeOrganizationSecurityPolicyRulePreconfiguredWafConfigExclusionOutputReferenceParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) error {
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

