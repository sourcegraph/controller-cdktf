//go:build no_runtime_type_checking

package appconfigextension

// Building without runtime type checking enabled, so all the below just return nil

func (a *jsiiProxy_AppconfigExtensionParameterList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (a *jsiiProxy_AppconfigExtensionParameterList) validateGetParameters(index *float64) error {
	return nil
}

func (a *jsiiProxy_AppconfigExtensionParameterList) validateResolveParameters(context cdktn.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_AppconfigExtensionParameterList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_AppconfigExtensionParameterList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_AppconfigExtensionParameterList) validateSetTerraformResourceParameters(val cdktn.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_AppconfigExtensionParameterList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewAppconfigExtensionParameterListParameters(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

