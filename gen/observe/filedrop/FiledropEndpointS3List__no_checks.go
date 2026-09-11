//go:build no_runtime_type_checking

package filedrop

// Building without runtime type checking enabled, so all the below just return nil

func (f *jsiiProxy_FiledropEndpointS3List) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (f *jsiiProxy_FiledropEndpointS3List) validateGetParameters(index *float64) error {
	return nil
}

func (f *jsiiProxy_FiledropEndpointS3List) validateResolveParameters(context cdktn.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_FiledropEndpointS3List) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_FiledropEndpointS3List) validateSetTerraformResourceParameters(val cdktn.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_FiledropEndpointS3List) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewFiledropEndpointS3ListParameters(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

