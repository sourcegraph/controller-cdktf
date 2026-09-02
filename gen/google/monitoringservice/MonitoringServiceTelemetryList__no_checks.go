//go:build no_runtime_type_checking

package monitoringservice

// Building without runtime type checking enabled, so all the below just return nil

func (m *jsiiProxy_MonitoringServiceTelemetryList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (m *jsiiProxy_MonitoringServiceTelemetryList) validateGetParameters(index *float64) error {
	return nil
}

func (m *jsiiProxy_MonitoringServiceTelemetryList) validateResolveParameters(context cdktn.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_MonitoringServiceTelemetryList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_MonitoringServiceTelemetryList) validateSetTerraformResourceParameters(val cdktn.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_MonitoringServiceTelemetryList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewMonitoringServiceTelemetryListParameters(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

