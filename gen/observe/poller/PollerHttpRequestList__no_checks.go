//go:build no_runtime_type_checking

package poller

// Building without runtime type checking enabled, so all the below just return nil

func (p *jsiiProxy_PollerHttpRequestList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (p *jsiiProxy_PollerHttpRequestList) validateGetParameters(index *float64) error {
	return nil
}

func (p *jsiiProxy_PollerHttpRequestList) validateResolveParameters(context cdktn.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_PollerHttpRequestList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_PollerHttpRequestList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_PollerHttpRequestList) validateSetTerraformResourceParameters(val cdktn.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_PollerHttpRequestList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewPollerHttpRequestListParameters(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

