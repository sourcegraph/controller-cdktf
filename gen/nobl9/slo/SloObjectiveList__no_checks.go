//go:build no_runtime_type_checking

package slo

// Building without runtime type checking enabled, so all the below just return nil

func (s *jsiiProxy_SloObjectiveList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (s *jsiiProxy_SloObjectiveList) validateGetParameters(index *float64) error {
	return nil
}

func (s *jsiiProxy_SloObjectiveList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_SloObjectiveList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_SloObjectiveList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_SloObjectiveList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_SloObjectiveList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewSloObjectiveListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

