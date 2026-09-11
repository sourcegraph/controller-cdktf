//go:build no_runtime_type_checking

package endpointsservice

// Building without runtime type checking enabled, so all the below just return nil

func (e *jsiiProxy_EndpointsServiceApisList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (e *jsiiProxy_EndpointsServiceApisList) validateGetParameters(index *float64) error {
	return nil
}

func (e *jsiiProxy_EndpointsServiceApisList) validateResolveParameters(context cdktn.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_EndpointsServiceApisList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_EndpointsServiceApisList) validateSetTerraformResourceParameters(val cdktn.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_EndpointsServiceApisList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewEndpointsServiceApisListParameters(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

