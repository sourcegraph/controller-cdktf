//go:build no_runtime_type_checking

package loadbalancer

// Building without runtime type checking enabled, so all the below just return nil

func (l *jsiiProxy_LoadBalancerRulesList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (l *jsiiProxy_LoadBalancerRulesList) validateGetParameters(index *float64) error {
	return nil
}

func (l *jsiiProxy_LoadBalancerRulesList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_LoadBalancerRulesList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_LoadBalancerRulesList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_LoadBalancerRulesList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_LoadBalancerRulesList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewLoadBalancerRulesListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

