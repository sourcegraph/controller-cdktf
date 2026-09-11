//go:build no_runtime_type_checking

package filedrop

// Building without runtime type checking enabled, so all the below just return nil

func (f *jsiiProxy_FiledropEndpointList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (f *jsiiProxy_FiledropEndpointList) validateGetParameters(index *float64) error {
	return nil
}

func (f *jsiiProxy_FiledropEndpointList) validateResolveParameters(context cdktn.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_FiledropEndpointList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_FiledropEndpointList) validateSetTerraformResourceParameters(val cdktn.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_FiledropEndpointList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewFiledropEndpointListParameters(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

