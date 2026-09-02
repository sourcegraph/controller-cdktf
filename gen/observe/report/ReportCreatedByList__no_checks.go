//go:build no_runtime_type_checking

package report

// Building without runtime type checking enabled, so all the below just return nil

func (r *jsiiProxy_ReportCreatedByList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (r *jsiiProxy_ReportCreatedByList) validateGetParameters(index *float64) error {
	return nil
}

func (r *jsiiProxy_ReportCreatedByList) validateResolveParameters(context cdktn.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_ReportCreatedByList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_ReportCreatedByList) validateSetTerraformResourceParameters(val cdktn.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_ReportCreatedByList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewReportCreatedByListParameters(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

