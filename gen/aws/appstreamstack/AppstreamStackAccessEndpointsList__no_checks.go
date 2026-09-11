//go:build no_runtime_type_checking

package appstreamstack

// Building without runtime type checking enabled, so all the below just return nil

func (a *jsiiProxy_AppstreamStackAccessEndpointsList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (a *jsiiProxy_AppstreamStackAccessEndpointsList) validateGetParameters(index *float64) error {
	return nil
}

func (a *jsiiProxy_AppstreamStackAccessEndpointsList) validateResolveParameters(context cdktn.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_AppstreamStackAccessEndpointsList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_AppstreamStackAccessEndpointsList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_AppstreamStackAccessEndpointsList) validateSetTerraformResourceParameters(val cdktn.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_AppstreamStackAccessEndpointsList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewAppstreamStackAccessEndpointsListParameters(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

