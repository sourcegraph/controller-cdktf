//go:build no_runtime_type_checking

package poller

// Building without runtime type checking enabled, so all the below just return nil

func (p *jsiiProxy_PollerHttpTimestampList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (p *jsiiProxy_PollerHttpTimestampList) validateGetParameters(index *float64) error {
	return nil
}

func (p *jsiiProxy_PollerHttpTimestampList) validateResolveParameters(context cdktn.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_PollerHttpTimestampList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_PollerHttpTimestampList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_PollerHttpTimestampList) validateSetTerraformResourceParameters(val cdktn.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_PollerHttpTimestampList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewPollerHttpTimestampListParameters(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

