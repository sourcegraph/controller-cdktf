//go:build no_runtime_type_checking

package report

// Building without runtime type checking enabled, so all the below just return nil

func (r *jsiiProxy_Report) validateAddMoveTargetParameters(moveTarget *string) error {
	return nil
}

func (r *jsiiProxy_Report) validateAddOverrideParameters(path *string, value interface{}) error {
	return nil
}

func (r *jsiiProxy_Report) validateGetAnyMapAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (r *jsiiProxy_Report) validateGetBooleanAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (r *jsiiProxy_Report) validateGetBooleanMapAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (r *jsiiProxy_Report) validateGetListAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (r *jsiiProxy_Report) validateGetNumberAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (r *jsiiProxy_Report) validateGetNumberListAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (r *jsiiProxy_Report) validateGetNumberMapAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (r *jsiiProxy_Report) validateGetStringAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (r *jsiiProxy_Report) validateGetStringMapAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (r *jsiiProxy_Report) validateImportFromParameters(id *string) error {
	return nil
}

func (r *jsiiProxy_Report) validateInterpolationForAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (r *jsiiProxy_Report) validateMarkWriteOnlyAttributeParameters(value interface{}) error {
	return nil
}

func (r *jsiiProxy_Report) validateMoveFromIdParameters(id *string) error {
	return nil
}

func (r *jsiiProxy_Report) validateMoveToParameters(moveTarget *string, index interface{}) error {
	return nil
}

func (r *jsiiProxy_Report) validateMoveToIdParameters(id *string) error {
	return nil
}

func (r *jsiiProxy_Report) validateOverrideLogicalIdParameters(newLogicalId *string) error {
	return nil
}

func (r *jsiiProxy_Report) validatePutDashboardParameters(value *ReportDashboard) error {
	return nil
}

func (r *jsiiProxy_Report) validatePutScheduleParameters(value *ReportSchedule) error {
	return nil
}

func (r *jsiiProxy_Report) validateRegisterProviderFeatureUsageParameters(feature cdktn.ProviderFeature) error {
	return nil
}

func validateReport_GenerateConfigForImportParameters(scope constructs.Construct, importToId *string, importFromId *string) error {
	return nil
}

func validateReport_IsConstructParameters(x interface{}) error {
	return nil
}

func validateReport_IsTerraformElementParameters(x interface{}) error {
	return nil
}

func validateReport_IsTerraformResourceParameters(x interface{}) error {
	return nil
}

func (j *jsiiProxy_Report) validateSetConnectionParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_Report) validateSetCountParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_Report) validateSetEmailBccRecipientsParameters(val *[]*string) error {
	return nil
}

func (j *jsiiProxy_Report) validateSetEmailBodyParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_Report) validateSetEmailRecipientsParameters(val *[]*string) error {
	return nil
}

func (j *jsiiProxy_Report) validateSetEmailSubjectParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_Report) validateSetEnabledParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_Report) validateSetIdParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_Report) validateSetLabelParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_Report) validateSetLifecycleParameters(val *cdktn.TerraformResourceLifecycle) error {
	return nil
}

func (j *jsiiProxy_Report) validateSetProvisionersParameters(val *[]interface{}) error {
	return nil
}

func validateNewReportParameters(scope constructs.Construct, id *string, config *ReportConfig) error {
	return nil
}

