//go:build no_runtime_type_checking

package link

// Building without runtime type checking enabled, so all the below just return nil

func (l *jsiiProxy_Link) validateAddMoveTargetParameters(moveTarget *string) error {
	return nil
}

func (l *jsiiProxy_Link) validateAddOverrideParameters(path *string, value interface{}) error {
	return nil
}

func (l *jsiiProxy_Link) validateGetAnyMapAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (l *jsiiProxy_Link) validateGetBooleanAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (l *jsiiProxy_Link) validateGetBooleanMapAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (l *jsiiProxy_Link) validateGetListAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (l *jsiiProxy_Link) validateGetNumberAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (l *jsiiProxy_Link) validateGetNumberListAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (l *jsiiProxy_Link) validateGetNumberMapAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (l *jsiiProxy_Link) validateGetStringAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (l *jsiiProxy_Link) validateGetStringMapAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (l *jsiiProxy_Link) validateImportFromParameters(id *string) error {
	return nil
}

func (l *jsiiProxy_Link) validateInterpolationForAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (l *jsiiProxy_Link) validateMarkWriteOnlyAttributeParameters(value interface{}) error {
	return nil
}

func (l *jsiiProxy_Link) validateMoveFromIdParameters(id *string) error {
	return nil
}

func (l *jsiiProxy_Link) validateMoveToParameters(moveTarget *string, index interface{}) error {
	return nil
}

func (l *jsiiProxy_Link) validateMoveToIdParameters(id *string) error {
	return nil
}

func (l *jsiiProxy_Link) validateOverrideLogicalIdParameters(newLogicalId *string) error {
	return nil
}

func (l *jsiiProxy_Link) validateRegisterProviderFeatureUsageParameters(feature cdktn.ProviderFeature) error {
	return nil
}

func validateLink_GenerateConfigForImportParameters(scope constructs.Construct, importToId *string, importFromId *string) error {
	return nil
}

func validateLink_IsConstructParameters(x interface{}) error {
	return nil
}

func validateLink_IsTerraformElementParameters(x interface{}) error {
	return nil
}

func validateLink_IsTerraformResourceParameters(x interface{}) error {
	return nil
}

func (j *jsiiProxy_Link) validateSetConnectionParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_Link) validateSetCountParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_Link) validateSetFieldsParameters(val *[]*string) error {
	return nil
}

func (j *jsiiProxy_Link) validateSetIdParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_Link) validateSetLabelParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_Link) validateSetLifecycleParameters(val *cdktn.TerraformResourceLifecycle) error {
	return nil
}

func (j *jsiiProxy_Link) validateSetProvisionersParameters(val *[]interface{}) error {
	return nil
}

func (j *jsiiProxy_Link) validateSetSourceParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_Link) validateSetTargetParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_Link) validateSetWorkspaceParameters(val *string) error {
	return nil
}

func validateNewLinkParameters(scope constructs.Construct, id *string, config *LinkConfig) error {
	return nil
}

