//go:build no_runtime_type_checking

package cestool

// Building without runtime type checking enabled, so all the below just return nil

func (c *jsiiProxy_CesToolSystemToolList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (c *jsiiProxy_CesToolSystemToolList) validateGetParameters(index *float64) error {
	return nil
}

func (c *jsiiProxy_CesToolSystemToolList) validateResolveParameters(context cdktn.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_CesToolSystemToolList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_CesToolSystemToolList) validateSetTerraformResourceParameters(val cdktn.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_CesToolSystemToolList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewCesToolSystemToolListParameters(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

