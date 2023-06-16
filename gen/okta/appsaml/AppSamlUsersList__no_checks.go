//go:build no_runtime_type_checking

package appsaml

// Building without runtime type checking enabled, so all the below just return nil

func (a *jsiiProxy_AppSamlUsersList) validateGetParameters(index *float64) error {
	return nil
}

func (a *jsiiProxy_AppSamlUsersList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_AppSamlUsersList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_AppSamlUsersList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_AppSamlUsersList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_AppSamlUsersList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewAppSamlUsersListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

