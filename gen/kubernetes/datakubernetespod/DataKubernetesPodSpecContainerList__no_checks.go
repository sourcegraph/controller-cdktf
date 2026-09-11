//go:build no_runtime_type_checking

package datakubernetespod

// Building without runtime type checking enabled, so all the below just return nil

func (d *jsiiProxy_DataKubernetesPodSpecContainerList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (d *jsiiProxy_DataKubernetesPodSpecContainerList) validateGetParameters(index *float64) error {
	return nil
}

func (d *jsiiProxy_DataKubernetesPodSpecContainerList) validateResolveParameters(context cdktn.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_DataKubernetesPodSpecContainerList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_DataKubernetesPodSpecContainerList) validateSetTerraformResourceParameters(val cdktn.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_DataKubernetesPodSpecContainerList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewDataKubernetesPodSpecContainerListParameters(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

