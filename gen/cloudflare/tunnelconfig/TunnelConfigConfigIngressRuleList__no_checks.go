//go:build no_runtime_type_checking

package tunnelconfig

// Building without runtime type checking enabled, so all the below just return nil

func (t *jsiiProxy_TunnelConfigConfigIngressRuleList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (t *jsiiProxy_TunnelConfigConfigIngressRuleList) validateGetParameters(index *float64) error {
	return nil
}

func (t *jsiiProxy_TunnelConfigConfigIngressRuleList) validateResolveParameters(context cdktn.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_TunnelConfigConfigIngressRuleList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_TunnelConfigConfigIngressRuleList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_TunnelConfigConfigIngressRuleList) validateSetTerraformResourceParameters(val cdktn.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_TunnelConfigConfigIngressRuleList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewTunnelConfigConfigIngressRuleListParameters(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

