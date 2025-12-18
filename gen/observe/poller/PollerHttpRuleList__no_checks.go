//go:build no_runtime_type_checking

package poller

// Building without runtime type checking enabled, so all the below just return nil

func (p *jsiiProxy_PollerHttpRuleList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (p *jsiiProxy_PollerHttpRuleList) validateGetParameters(index *float64) error {
	return nil
}

func (p *jsiiProxy_PollerHttpRuleList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_PollerHttpRuleList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_PollerHttpRuleList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_PollerHttpRuleList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_PollerHttpRuleList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewPollerHttpRuleListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

