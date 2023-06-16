//go:build no_runtime_type_checking

package datatfeteamaccess

// Building without runtime type checking enabled, so all the below just return nil

func (d *jsiiProxy_DataTfeTeamAccessPermissionsList) validateGetParameters(index *float64) error {
	return nil
}

func (d *jsiiProxy_DataTfeTeamAccessPermissionsList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_DataTfeTeamAccessPermissionsList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_DataTfeTeamAccessPermissionsList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_DataTfeTeamAccessPermissionsList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewDataTfeTeamAccessPermissionsListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

