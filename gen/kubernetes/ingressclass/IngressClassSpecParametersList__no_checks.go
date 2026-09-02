//go:build no_runtime_type_checking

package ingressclass

// Building without runtime type checking enabled, so all the below just return nil

func (i *jsiiProxy_IngressClassSpecParametersList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (i *jsiiProxy_IngressClassSpecParametersList) validateGetParameters(index *float64) error {
	return nil
}

func (i *jsiiProxy_IngressClassSpecParametersList) validateResolveParameters(context cdktn.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_IngressClassSpecParametersList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_IngressClassSpecParametersList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_IngressClassSpecParametersList) validateSetTerraformResourceParameters(val cdktn.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_IngressClassSpecParametersList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewIngressClassSpecParametersListParameters(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

