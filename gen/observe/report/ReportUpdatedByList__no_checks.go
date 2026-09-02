//go:build no_runtime_type_checking

package report

// Building without runtime type checking enabled, so all the below just return nil

func (r *jsiiProxy_ReportUpdatedByList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (r *jsiiProxy_ReportUpdatedByList) validateGetParameters(index *float64) error {
	return nil
}

func (r *jsiiProxy_ReportUpdatedByList) validateResolveParameters(context cdktn.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_ReportUpdatedByList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_ReportUpdatedByList) validateSetTerraformResourceParameters(val cdktn.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_ReportUpdatedByList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewReportUpdatedByListParameters(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

