//go:build no_runtime_type_checking

package endpoints

// Building without runtime type checking enabled, so all the below just return nil

func (e *jsiiProxy_EndpointsSubsetPortList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (e *jsiiProxy_EndpointsSubsetPortList) validateGetParameters(index *float64) error {
	return nil
}

func (e *jsiiProxy_EndpointsSubsetPortList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_EndpointsSubsetPortList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_EndpointsSubsetPortList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_EndpointsSubsetPortList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_EndpointsSubsetPortList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewEndpointsSubsetPortListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

