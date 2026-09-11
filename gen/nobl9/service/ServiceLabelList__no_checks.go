//go:build no_runtime_type_checking

package service

// Building without runtime type checking enabled, so all the below just return nil

func (s *jsiiProxy_ServiceLabelList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (s *jsiiProxy_ServiceLabelList) validateGetParameters(index *float64) error {
	return nil
}

func (s *jsiiProxy_ServiceLabelList) validateResolveParameters(context cdktn.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_ServiceLabelList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_ServiceLabelList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_ServiceLabelList) validateSetTerraformResourceParameters(val cdktn.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_ServiceLabelList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewServiceLabelListParameters(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

