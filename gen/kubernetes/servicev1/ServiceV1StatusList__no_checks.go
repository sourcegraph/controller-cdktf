//go:build no_runtime_type_checking

package servicev1

// Building without runtime type checking enabled, so all the below just return nil

func (s *jsiiProxy_ServiceV1StatusList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (s *jsiiProxy_ServiceV1StatusList) validateGetParameters(index *float64) error {
	return nil
}

func (s *jsiiProxy_ServiceV1StatusList) validateResolveParameters(context cdktn.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_ServiceV1StatusList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_ServiceV1StatusList) validateSetTerraformResourceParameters(val cdktn.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_ServiceV1StatusList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewServiceV1StatusListParameters(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

