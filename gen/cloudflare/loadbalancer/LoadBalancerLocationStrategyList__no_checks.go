//go:build no_runtime_type_checking

package loadbalancer

// Building without runtime type checking enabled, so all the below just return nil

func (l *jsiiProxy_LoadBalancerLocationStrategyList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (l *jsiiProxy_LoadBalancerLocationStrategyList) validateGetParameters(index *float64) error {
	return nil
}

func (l *jsiiProxy_LoadBalancerLocationStrategyList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_LoadBalancerLocationStrategyList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_LoadBalancerLocationStrategyList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_LoadBalancerLocationStrategyList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_LoadBalancerLocationStrategyList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewLoadBalancerLocationStrategyListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

