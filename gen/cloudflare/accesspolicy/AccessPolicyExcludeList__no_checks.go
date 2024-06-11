//go:build no_runtime_type_checking

package accesspolicy

// Building without runtime type checking enabled, so all the below just return nil

func (a *jsiiProxy_AccessPolicyExcludeList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (a *jsiiProxy_AccessPolicyExcludeList) validateGetParameters(index *float64) error {
	return nil
}

func (a *jsiiProxy_AccessPolicyExcludeList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_AccessPolicyExcludeList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_AccessPolicyExcludeList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_AccessPolicyExcludeList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_AccessPolicyExcludeList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewAccessPolicyExcludeListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

