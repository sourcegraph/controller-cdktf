//go:build no_runtime_type_checking

package loadbalancerpool

// Building without runtime type checking enabled, so all the below just return nil

func (l *jsiiProxy_LoadBalancerPoolOriginsList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (l *jsiiProxy_LoadBalancerPoolOriginsList) validateGetParameters(index *float64) error {
	return nil
}

func (l *jsiiProxy_LoadBalancerPoolOriginsList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_LoadBalancerPoolOriginsList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_LoadBalancerPoolOriginsList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_LoadBalancerPoolOriginsList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_LoadBalancerPoolOriginsList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewLoadBalancerPoolOriginsListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

