//go:build no_runtime_type_checking

package samlapp

// Building without runtime type checking enabled, so all the below just return nil

func (s *jsiiProxy_SamlAppUsersList) validateGetParameters(index *float64) error {
	return nil
}

func (s *jsiiProxy_SamlAppUsersList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_SamlAppUsersList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_SamlAppUsersList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_SamlAppUsersList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_SamlAppUsersList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewSamlAppUsersListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

