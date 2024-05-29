//go:build no_runtime_type_checking

package loadbalancer

// Building without runtime type checking enabled, so all the below just return nil

func (l *jsiiProxy_LoadBalancerRulesOverridesList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (l *jsiiProxy_LoadBalancerRulesOverridesList) validateGetParameters(index *float64) error {
	return nil
}

func (l *jsiiProxy_LoadBalancerRulesOverridesList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_LoadBalancerRulesOverridesList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_LoadBalancerRulesOverridesList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_LoadBalancerRulesOverridesList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_LoadBalancerRulesOverridesList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewLoadBalancerRulesOverridesListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

