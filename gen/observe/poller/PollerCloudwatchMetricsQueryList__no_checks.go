//go:build no_runtime_type_checking

package poller

// Building without runtime type checking enabled, so all the below just return nil

func (p *jsiiProxy_PollerCloudwatchMetricsQueryList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (p *jsiiProxy_PollerCloudwatchMetricsQueryList) validateGetParameters(index *float64) error {
	return nil
}

func (p *jsiiProxy_PollerCloudwatchMetricsQueryList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_PollerCloudwatchMetricsQueryList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_PollerCloudwatchMetricsQueryList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_PollerCloudwatchMetricsQueryList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_PollerCloudwatchMetricsQueryList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewPollerCloudwatchMetricsQueryListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

