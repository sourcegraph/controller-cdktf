//go:build no_runtime_type_checking

package alertroute

// Building without runtime type checking enabled, so all the below just return nil

func (a *jsiiProxy_AlertRouteChannelConfigList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (a *jsiiProxy_AlertRouteChannelConfigList) validateGetParameters(index *float64) error {
	return nil
}

func (a *jsiiProxy_AlertRouteChannelConfigList) validateResolveParameters(context cdktn.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_AlertRouteChannelConfigList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_AlertRouteChannelConfigList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_AlertRouteChannelConfigList) validateSetTerraformResourceParameters(val cdktn.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_AlertRouteChannelConfigList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewAlertRouteChannelConfigListParameters(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

