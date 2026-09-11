//go:build no_runtime_type_checking

package eksnodegroup

// Building without runtime type checking enabled, so all the below just return nil

func (e *jsiiProxy_EksNodeGroupResourcesList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (e *jsiiProxy_EksNodeGroupResourcesList) validateGetParameters(index *float64) error {
	return nil
}

func (e *jsiiProxy_EksNodeGroupResourcesList) validateResolveParameters(context cdktn.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_EksNodeGroupResourcesList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_EksNodeGroupResourcesList) validateSetTerraformResourceParameters(val cdktn.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_EksNodeGroupResourcesList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewEksNodeGroupResourcesListParameters(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

