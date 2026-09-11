//go:build no_runtime_type_checking

package service

// Building without runtime type checking enabled, so all the below just return nil

func (s *jsiiProxy_ServiceStatusList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (s *jsiiProxy_ServiceStatusList) validateGetParameters(index *float64) error {
	return nil
}

func (s *jsiiProxy_ServiceStatusList) validateResolveParameters(context cdktn.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_ServiceStatusList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_ServiceStatusList) validateSetTerraformResourceParameters(val cdktn.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_ServiceStatusList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewServiceStatusListParameters(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

