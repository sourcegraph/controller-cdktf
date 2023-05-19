//go:build no_runtime_type_checking

package appswa

// Building without runtime type checking enabled, so all the below just return nil

func (a *jsiiProxy_AppSwaUsersList) validateGetParameters(index *float64) error {
	return nil
}

func (a *jsiiProxy_AppSwaUsersList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_AppSwaUsersList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_AppSwaUsersList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_AppSwaUsersList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_AppSwaUsersList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewAppSwaUsersListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

