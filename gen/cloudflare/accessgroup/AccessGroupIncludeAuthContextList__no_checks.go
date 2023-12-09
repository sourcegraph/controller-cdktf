//go:build no_runtime_type_checking

package accessgroup

// Building without runtime type checking enabled, so all the below just return nil

func (a *jsiiProxy_AccessGroupIncludeAuthContextList) validateGetParameters(index *float64) error {
	return nil
}

func (a *jsiiProxy_AccessGroupIncludeAuthContextList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_AccessGroupIncludeAuthContextList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_AccessGroupIncludeAuthContextList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_AccessGroupIncludeAuthContextList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_AccessGroupIncludeAuthContextList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewAccessGroupIncludeAuthContextListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

