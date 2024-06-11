//go:build no_runtime_type_checking

package loadbalancerpool

// Building without runtime type checking enabled, so all the below just return nil

func (l *jsiiProxy_LoadBalancerPoolLoadSheddingList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (l *jsiiProxy_LoadBalancerPoolLoadSheddingList) validateGetParameters(index *float64) error {
	return nil
}

func (l *jsiiProxy_LoadBalancerPoolLoadSheddingList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_LoadBalancerPoolLoadSheddingList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_LoadBalancerPoolLoadSheddingList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_LoadBalancerPoolLoadSheddingList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_LoadBalancerPoolLoadSheddingList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewLoadBalancerPoolLoadSheddingListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

