//go:build no_runtime_type_checking

package accesspolicy

// Building without runtime type checking enabled, so all the below just return nil

func (a *jsiiProxy_AccessPolicyRequireSamlList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (a *jsiiProxy_AccessPolicyRequireSamlList) validateGetParameters(index *float64) error {
	return nil
}

func (a *jsiiProxy_AccessPolicyRequireSamlList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_AccessPolicyRequireSamlList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_AccessPolicyRequireSamlList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_AccessPolicyRequireSamlList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_AccessPolicyRequireSamlList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewAccessPolicyRequireSamlListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

