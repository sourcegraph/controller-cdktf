//go:build no_runtime_type_checking

package configconformancepack

// Building without runtime type checking enabled, so all the below just return nil

func (c *jsiiProxy_ConfigConformancePackInputParameterList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (c *jsiiProxy_ConfigConformancePackInputParameterList) validateGetParameters(index *float64) error {
	return nil
}

func (c *jsiiProxy_ConfigConformancePackInputParameterList) validateResolveParameters(context cdktn.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_ConfigConformancePackInputParameterList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_ConfigConformancePackInputParameterList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_ConfigConformancePackInputParameterList) validateSetTerraformResourceParameters(val cdktn.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_ConfigConformancePackInputParameterList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewConfigConformancePackInputParameterListParameters(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

