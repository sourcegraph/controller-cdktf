//go:build no_runtime_type_checking

package loadbalancer

// Building without runtime type checking enabled, so all the below just return nil

func (l *jsiiProxy_LoadBalancerRandomSteeringList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (l *jsiiProxy_LoadBalancerRandomSteeringList) validateGetParameters(index *float64) error {
	return nil
}

func (l *jsiiProxy_LoadBalancerRandomSteeringList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_LoadBalancerRandomSteeringList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_LoadBalancerRandomSteeringList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_LoadBalancerRandomSteeringList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_LoadBalancerRandomSteeringList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewLoadBalancerRandomSteeringListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

