//go:build no_runtime_type_checking

package teamaccess

// Building without runtime type checking enabled, so all the below just return nil

func (t *jsiiProxy_TeamAccessPermissionsList) validateGetParameters(index *float64) error {
	return nil
}

func (t *jsiiProxy_TeamAccessPermissionsList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_TeamAccessPermissionsList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_TeamAccessPermissionsList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_TeamAccessPermissionsList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_TeamAccessPermissionsList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewTeamAccessPermissionsListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

