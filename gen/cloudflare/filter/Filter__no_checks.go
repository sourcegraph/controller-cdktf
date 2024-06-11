//go:build no_runtime_type_checking

package filter

// Building without runtime type checking enabled, so all the below just return nil

func (f *jsiiProxy_Filter) validateAddMoveTargetParameters(moveTarget *string) error {
	return nil
}

func (f *jsiiProxy_Filter) validateAddOverrideParameters(path *string, value interface{}) error {
	return nil
}

func (f *jsiiProxy_Filter) validateGetAnyMapAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (f *jsiiProxy_Filter) validateGetBooleanAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (f *jsiiProxy_Filter) validateGetBooleanMapAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (f *jsiiProxy_Filter) validateGetListAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (f *jsiiProxy_Filter) validateGetNumberAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (f *jsiiProxy_Filter) validateGetNumberListAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (f *jsiiProxy_Filter) validateGetNumberMapAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (f *jsiiProxy_Filter) validateGetStringAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (f *jsiiProxy_Filter) validateGetStringMapAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (f *jsiiProxy_Filter) validateImportFromParameters(id *string) error {
	return nil
}

func (f *jsiiProxy_Filter) validateInterpolationForAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (f *jsiiProxy_Filter) validateMoveFromIdParameters(id *string) error {
	return nil
}

func (f *jsiiProxy_Filter) validateMoveToParameters(moveTarget *string, index interface{}) error {
	return nil
}

func (f *jsiiProxy_Filter) validateMoveToIdParameters(id *string) error {
	return nil
}

func (f *jsiiProxy_Filter) validateOverrideLogicalIdParameters(newLogicalId *string) error {
	return nil
}

func validateFilter_GenerateConfigForImportParameters(scope constructs.Construct, importToId *string, importFromId *string) error {
	return nil
}

func validateFilter_IsConstructParameters(x interface{}) error {
	return nil
}

func validateFilter_IsTerraformElementParameters(x interface{}) error {
	return nil
}

func validateFilter_IsTerraformResourceParameters(x interface{}) error {
	return nil
}

func (j *jsiiProxy_Filter) validateSetConnectionParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_Filter) validateSetCountParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_Filter) validateSetDescriptionParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_Filter) validateSetExpressionParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_Filter) validateSetIdParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_Filter) validateSetLifecycleParameters(val *cdktf.TerraformResourceLifecycle) error {
	return nil
}

func (j *jsiiProxy_Filter) validateSetPausedParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_Filter) validateSetProvisionersParameters(val *[]interface{}) error {
	return nil
}

func (j *jsiiProxy_Filter) validateSetRefParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_Filter) validateSetZoneIdParameters(val *string) error {
	return nil
}

func validateNewFilterParameters(scope constructs.Construct, id *string, config *FilterConfig) error {
	return nil
}

