//go:build no_runtime_type_checking

package monitor

// Building without runtime type checking enabled, so all the below just return nil

func (m *jsiiProxy_MonitorStageList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (m *jsiiProxy_MonitorStageList) validateGetParameters(index *float64) error {
	return nil
}

func (m *jsiiProxy_MonitorStageList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_MonitorStageList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_MonitorStageList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_MonitorStageList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_MonitorStageList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewMonitorStageListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

