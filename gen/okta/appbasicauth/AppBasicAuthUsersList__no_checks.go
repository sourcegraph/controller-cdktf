//go:build no_runtime_type_checking

package appbasicauth

// Building without runtime type checking enabled, so all the below just return nil

func (a *jsiiProxy_AppBasicAuthUsersList) validateGetParameters(index *float64) error {
	return nil
}

func (a *jsiiProxy_AppBasicAuthUsersList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_AppBasicAuthUsersList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_AppBasicAuthUsersList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_AppBasicAuthUsersList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_AppBasicAuthUsersList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewAppBasicAuthUsersListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

