//go:build no_runtime_type_checking

package apitoken

// Building without runtime type checking enabled, so all the below just return nil

func (a *jsiiProxy_ApiTokenPolicyList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (a *jsiiProxy_ApiTokenPolicyList) validateGetParameters(index *float64) error {
	return nil
}

func (a *jsiiProxy_ApiTokenPolicyList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_ApiTokenPolicyList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_ApiTokenPolicyList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_ApiTokenPolicyList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_ApiTokenPolicyList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewApiTokenPolicyListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

