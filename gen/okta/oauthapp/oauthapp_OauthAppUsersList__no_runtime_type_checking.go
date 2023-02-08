//go:build no_runtime_type_checking

package oauthapp

// Building without runtime type checking enabled, so all the below just return nil

func (o *jsiiProxy_OauthAppUsersList) validateGetParameters(index *float64) error {
	return nil
}

func (o *jsiiProxy_OauthAppUsersList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_OauthAppUsersList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_OauthAppUsersList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_OauthAppUsersList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_OauthAppUsersList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewOauthAppUsersListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

