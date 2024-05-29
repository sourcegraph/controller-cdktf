//go:build no_runtime_type_checking

package slo

// Building without runtime type checking enabled, so all the below just return nil

func (s *jsiiProxy_SloLabelList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (s *jsiiProxy_SloLabelList) validateGetParameters(index *float64) error {
	return nil
}

func (s *jsiiProxy_SloLabelList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_SloLabelList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_SloLabelList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_SloLabelList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_SloLabelList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewSloLabelListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

