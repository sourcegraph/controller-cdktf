//go:build no_runtime_type_checking

package computefirewall

// Building without runtime type checking enabled, so all the below just return nil

func (c *jsiiProxy_ComputeFirewallDenyList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (c *jsiiProxy_ComputeFirewallDenyList) validateGetParameters(index *float64) error {
	return nil
}

func (c *jsiiProxy_ComputeFirewallDenyList) validateResolveParameters(context cdktn.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_ComputeFirewallDenyList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_ComputeFirewallDenyList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_ComputeFirewallDenyList) validateSetTerraformResourceParameters(val cdktn.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_ComputeFirewallDenyList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewComputeFirewallDenyListParameters(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

