//go:build no_runtime_type_checking

package autologinapp

// Building without runtime type checking enabled, so all the below just return nil

func (a *jsiiProxy_AutoLoginAppUsersList) validateGetParameters(index *float64) error {
	return nil
}

func (a *jsiiProxy_AutoLoginAppUsersList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_AutoLoginAppUsersList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_AutoLoginAppUsersList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_AutoLoginAppUsersList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_AutoLoginAppUsersList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewAutoLoginAppUsersListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

