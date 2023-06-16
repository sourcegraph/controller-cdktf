//go:build no_runtime_type_checking

package appsharedcredentials

// Building without runtime type checking enabled, so all the below just return nil

func (a *jsiiProxy_AppSharedCredentialsUsersList) validateGetParameters(index *float64) error {
	return nil
}

func (a *jsiiProxy_AppSharedCredentialsUsersList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_AppSharedCredentialsUsersList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_AppSharedCredentialsUsersList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_AppSharedCredentialsUsersList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_AppSharedCredentialsUsersList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewAppSharedCredentialsUsersListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

