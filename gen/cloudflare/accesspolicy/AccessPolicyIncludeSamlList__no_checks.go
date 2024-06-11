//go:build no_runtime_type_checking

package accesspolicy

// Building without runtime type checking enabled, so all the below just return nil

func (a *jsiiProxy_AccessPolicyIncludeSamlList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (a *jsiiProxy_AccessPolicyIncludeSamlList) validateGetParameters(index *float64) error {
	return nil
}

func (a *jsiiProxy_AccessPolicyIncludeSamlList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_AccessPolicyIncludeSamlList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_AccessPolicyIncludeSamlList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_AccessPolicyIncludeSamlList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_AccessPolicyIncludeSamlList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewAccessPolicyIncludeSamlListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

