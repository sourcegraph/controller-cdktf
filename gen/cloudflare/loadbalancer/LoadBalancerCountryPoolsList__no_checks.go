//go:build no_runtime_type_checking

package loadbalancer

// Building without runtime type checking enabled, so all the below just return nil

func (l *jsiiProxy_LoadBalancerCountryPoolsList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (l *jsiiProxy_LoadBalancerCountryPoolsList) validateGetParameters(index *float64) error {
	return nil
}

func (l *jsiiProxy_LoadBalancerCountryPoolsList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_LoadBalancerCountryPoolsList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_LoadBalancerCountryPoolsList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_LoadBalancerCountryPoolsList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_LoadBalancerCountryPoolsList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewLoadBalancerCountryPoolsListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

