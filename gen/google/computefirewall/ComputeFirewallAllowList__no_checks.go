//go:build no_runtime_type_checking

package computefirewall

// Building without runtime type checking enabled, so all the below just return nil

func (c *jsiiProxy_ComputeFirewallAllowList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (c *jsiiProxy_ComputeFirewallAllowList) validateGetParameters(index *float64) error {
	return nil
}

func (c *jsiiProxy_ComputeFirewallAllowList) validateResolveParameters(context cdktn.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_ComputeFirewallAllowList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_ComputeFirewallAllowList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_ComputeFirewallAllowList) validateSetTerraformResourceParameters(val cdktn.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_ComputeFirewallAllowList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewComputeFirewallAllowListParameters(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

