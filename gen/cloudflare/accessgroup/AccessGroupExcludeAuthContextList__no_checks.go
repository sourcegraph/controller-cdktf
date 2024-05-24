//go:build no_runtime_type_checking

package accessgroup

// Building without runtime type checking enabled, so all the below just return nil

func (a *jsiiProxy_AccessGroupExcludeAuthContextList) validateGetParameters(index *float64) error {
	return nil
}

func (a *jsiiProxy_AccessGroupExcludeAuthContextList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_AccessGroupExcludeAuthContextList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_AccessGroupExcludeAuthContextList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_AccessGroupExcludeAuthContextList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_AccessGroupExcludeAuthContextList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewAccessGroupExcludeAuthContextListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

