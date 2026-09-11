//go:build no_runtime_type_checking

package accesspolicy

// Building without runtime type checking enabled, so all the below just return nil

func (a *jsiiProxy_AccessPolicyExcludeGsuiteList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (a *jsiiProxy_AccessPolicyExcludeGsuiteList) validateGetParameters(index *float64) error {
	return nil
}

func (a *jsiiProxy_AccessPolicyExcludeGsuiteList) validateResolveParameters(context cdktn.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_AccessPolicyExcludeGsuiteList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_AccessPolicyExcludeGsuiteList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_AccessPolicyExcludeGsuiteList) validateSetTerraformResourceParameters(val cdktn.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_AccessPolicyExcludeGsuiteList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewAccessPolicyExcludeGsuiteListParameters(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

