//go:build no_runtime_type_checking

package serviceaccount

// Building without runtime type checking enabled, so all the below just return nil

func (s *jsiiProxy_ServiceAccountSecretList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (s *jsiiProxy_ServiceAccountSecretList) validateGetParameters(index *float64) error {
	return nil
}

func (s *jsiiProxy_ServiceAccountSecretList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_ServiceAccountSecretList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_ServiceAccountSecretList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_ServiceAccountSecretList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_ServiceAccountSecretList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewServiceAccountSecretListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

