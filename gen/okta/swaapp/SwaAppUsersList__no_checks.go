//go:build no_runtime_type_checking

package swaapp

// Building without runtime type checking enabled, so all the below just return nil

func (s *jsiiProxy_SwaAppUsersList) validateGetParameters(index *float64) error {
	return nil
}

func (s *jsiiProxy_SwaAppUsersList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_SwaAppUsersList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_SwaAppUsersList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_SwaAppUsersList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_SwaAppUsersList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewSwaAppUsersListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

