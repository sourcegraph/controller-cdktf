//go:build no_runtime_type_checking

package extension

// Building without runtime type checking enabled, so all the below just return nil

func (e *jsiiProxy_Extension) validateAddOverrideParameters(path *string, value interface{}) error {
	return nil
}

func (e *jsiiProxy_Extension) validateGetAnyMapAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (e *jsiiProxy_Extension) validateGetBooleanAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (e *jsiiProxy_Extension) validateGetBooleanMapAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (e *jsiiProxy_Extension) validateGetListAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (e *jsiiProxy_Extension) validateGetNumberAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (e *jsiiProxy_Extension) validateGetNumberListAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (e *jsiiProxy_Extension) validateGetNumberMapAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (e *jsiiProxy_Extension) validateGetStringAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (e *jsiiProxy_Extension) validateGetStringMapAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (e *jsiiProxy_Extension) validateInterpolationForAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (e *jsiiProxy_Extension) validateOverrideLogicalIdParameters(newLogicalId *string) error {
	return nil
}

func validateExtension_IsConstructParameters(x interface{}) error {
	return nil
}

func validateExtension_IsTerraformElementParameters(x interface{}) error {
	return nil
}

func validateExtension_IsTerraformResourceParameters(x interface{}) error {
	return nil
}

func (j *jsiiProxy_Extension) validateSetConnectionParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_Extension) validateSetCreateCascadeParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_Extension) validateSetDatabaseParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_Extension) validateSetDropCascadeParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_Extension) validateSetIdParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_Extension) validateSetLifecycleParameters(val *cdktf.TerraformResourceLifecycle) error {
	return nil
}

func (j *jsiiProxy_Extension) validateSetNameParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_Extension) validateSetProvisionersParameters(val *[]interface{}) error {
	return nil
}

func (j *jsiiProxy_Extension) validateSetSchemaParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_Extension) validateSetVersionParameters(val *string) error {
	return nil
}

func validateNewExtensionParameters(scope constructs.Construct, id *string, config *ExtensionConfig) error {
	return nil
}

