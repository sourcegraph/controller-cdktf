//go:build no_runtime_type_checking

package report

// Building without runtime type checking enabled, so all the below just return nil

func (r *jsiiProxy_ReportDashboardParametersList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (r *jsiiProxy_ReportDashboardParametersList) validateGetParameters(index *float64) error {
	return nil
}

func (r *jsiiProxy_ReportDashboardParametersList) validateResolveParameters(context cdktn.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_ReportDashboardParametersList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_ReportDashboardParametersList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_ReportDashboardParametersList) validateSetTerraformResourceParameters(val cdktn.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_ReportDashboardParametersList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewReportDashboardParametersListParameters(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

