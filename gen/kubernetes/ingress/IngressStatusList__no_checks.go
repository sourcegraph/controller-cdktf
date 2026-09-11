//go:build no_runtime_type_checking

package ingress

// Building without runtime type checking enabled, so all the below just return nil

func (i *jsiiProxy_IngressStatusList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (i *jsiiProxy_IngressStatusList) validateGetParameters(index *float64) error {
	return nil
}

func (i *jsiiProxy_IngressStatusList) validateResolveParameters(context cdktn.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_IngressStatusList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_IngressStatusList) validateSetTerraformResourceParameters(val cdktn.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_IngressStatusList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewIngressStatusListParameters(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

