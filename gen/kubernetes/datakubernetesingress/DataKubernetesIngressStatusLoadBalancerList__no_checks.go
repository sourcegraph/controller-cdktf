//go:build no_runtime_type_checking

package datakubernetesingress

// Building without runtime type checking enabled, so all the below just return nil

func (d *jsiiProxy_DataKubernetesIngressStatusLoadBalancerList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (d *jsiiProxy_DataKubernetesIngressStatusLoadBalancerList) validateGetParameters(index *float64) error {
	return nil
}

func (d *jsiiProxy_DataKubernetesIngressStatusLoadBalancerList) validateResolveParameters(context cdktn.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_DataKubernetesIngressStatusLoadBalancerList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_DataKubernetesIngressStatusLoadBalancerList) validateSetTerraformResourceParameters(val cdktn.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_DataKubernetesIngressStatusLoadBalancerList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewDataKubernetesIngressStatusLoadBalancerListParameters(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

