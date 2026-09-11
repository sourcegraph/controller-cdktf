//go:build no_runtime_type_checking

package datakubernetesnamespace

// Building without runtime type checking enabled, so all the below just return nil

func (d *jsiiProxy_DataKubernetesNamespaceSpecList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (d *jsiiProxy_DataKubernetesNamespaceSpecList) validateGetParameters(index *float64) error {
	return nil
}

func (d *jsiiProxy_DataKubernetesNamespaceSpecList) validateResolveParameters(context cdktn.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_DataKubernetesNamespaceSpecList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_DataKubernetesNamespaceSpecList) validateSetTerraformResourceParameters(val cdktn.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_DataKubernetesNamespaceSpecList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewDataKubernetesNamespaceSpecListParameters(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

