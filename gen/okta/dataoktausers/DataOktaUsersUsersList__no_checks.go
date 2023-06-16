//go:build no_runtime_type_checking

package dataoktausers

// Building without runtime type checking enabled, so all the below just return nil

func (d *jsiiProxy_DataOktaUsersUsersList) validateGetParameters(index *float64) error {
	return nil
}

func (d *jsiiProxy_DataOktaUsersUsersList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_DataOktaUsersUsersList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_DataOktaUsersUsersList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_DataOktaUsersUsersList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewDataOktaUsersUsersListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

