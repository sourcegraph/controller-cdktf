//go:build no_runtime_type_checking

package alertroute

// Building without runtime type checking enabled, so all the below just return nil

func (a *jsiiProxy_AlertRouteAlertSourcesList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (a *jsiiProxy_AlertRouteAlertSourcesList) validateGetParameters(index *float64) error {
	return nil
}

func (a *jsiiProxy_AlertRouteAlertSourcesList) validateResolveParameters(context cdktn.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_AlertRouteAlertSourcesList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_AlertRouteAlertSourcesList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_AlertRouteAlertSourcesList) validateSetTerraformResourceParameters(val cdktn.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_AlertRouteAlertSourcesList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewAlertRouteAlertSourcesListParameters(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

