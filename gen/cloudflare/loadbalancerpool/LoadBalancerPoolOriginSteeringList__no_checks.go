//go:build no_runtime_type_checking

package loadbalancerpool

// Building without runtime type checking enabled, so all the below just return nil

func (l *jsiiProxy_LoadBalancerPoolOriginSteeringList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (l *jsiiProxy_LoadBalancerPoolOriginSteeringList) validateGetParameters(index *float64) error {
	return nil
}

func (l *jsiiProxy_LoadBalancerPoolOriginSteeringList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_LoadBalancerPoolOriginSteeringList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_LoadBalancerPoolOriginSteeringList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_LoadBalancerPoolOriginSteeringList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_LoadBalancerPoolOriginSteeringList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewLoadBalancerPoolOriginSteeringListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

