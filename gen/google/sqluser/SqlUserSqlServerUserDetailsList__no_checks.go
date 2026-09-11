//go:build no_runtime_type_checking

package sqluser

// Building without runtime type checking enabled, so all the below just return nil

func (s *jsiiProxy_SqlUserSqlServerUserDetailsList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (s *jsiiProxy_SqlUserSqlServerUserDetailsList) validateGetParameters(index *float64) error {
	return nil
}

func (s *jsiiProxy_SqlUserSqlServerUserDetailsList) validateResolveParameters(context cdktn.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_SqlUserSqlServerUserDetailsList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_SqlUserSqlServerUserDetailsList) validateSetTerraformResourceParameters(val cdktn.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_SqlUserSqlServerUserDetailsList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewSqlUserSqlServerUserDetailsListParameters(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

