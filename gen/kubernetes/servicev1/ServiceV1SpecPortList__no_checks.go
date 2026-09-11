//go:build no_runtime_type_checking

package servicev1

// Building without runtime type checking enabled, so all the below just return nil

func (s *jsiiProxy_ServiceV1SpecPortList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (s *jsiiProxy_ServiceV1SpecPortList) validateGetParameters(index *float64) error {
	return nil
}

func (s *jsiiProxy_ServiceV1SpecPortList) validateResolveParameters(context cdktn.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_ServiceV1SpecPortList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_ServiceV1SpecPortList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_ServiceV1SpecPortList) validateSetTerraformResourceParameters(val cdktn.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_ServiceV1SpecPortList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewServiceV1SpecPortListParameters(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

