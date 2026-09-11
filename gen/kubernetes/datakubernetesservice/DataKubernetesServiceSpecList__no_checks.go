//go:build no_runtime_type_checking

package datakubernetesservice

// Building without runtime type checking enabled, so all the below just return nil

func (d *jsiiProxy_DataKubernetesServiceSpecList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (d *jsiiProxy_DataKubernetesServiceSpecList) validateGetParameters(index *float64) error {
	return nil
}

func (d *jsiiProxy_DataKubernetesServiceSpecList) validateResolveParameters(context cdktn.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_DataKubernetesServiceSpecList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_DataKubernetesServiceSpecList) validateSetTerraformResourceParameters(val cdktn.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_DataKubernetesServiceSpecList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewDataKubernetesServiceSpecListParameters(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

