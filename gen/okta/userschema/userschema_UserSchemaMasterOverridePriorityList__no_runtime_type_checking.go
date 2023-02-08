//go:build no_runtime_type_checking

package userschema

// Building without runtime type checking enabled, so all the below just return nil

func (u *jsiiProxy_UserSchemaMasterOverridePriorityList) validateGetParameters(index *float64) error {
	return nil
}

func (u *jsiiProxy_UserSchemaMasterOverridePriorityList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_UserSchemaMasterOverridePriorityList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_UserSchemaMasterOverridePriorityList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_UserSchemaMasterOverridePriorityList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_UserSchemaMasterOverridePriorityList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewUserSchemaMasterOverridePriorityListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

