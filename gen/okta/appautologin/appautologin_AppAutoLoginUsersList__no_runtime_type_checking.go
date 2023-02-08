//go:build no_runtime_type_checking

package appautologin

// Building without runtime type checking enabled, so all the below just return nil

func (a *jsiiProxy_AppAutoLoginUsersList) validateGetParameters(index *float64) error {
	return nil
}

func (a *jsiiProxy_AppAutoLoginUsersList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_AppAutoLoginUsersList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_AppAutoLoginUsersList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_AppAutoLoginUsersList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_AppAutoLoginUsersList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewAppAutoLoginUsersListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

