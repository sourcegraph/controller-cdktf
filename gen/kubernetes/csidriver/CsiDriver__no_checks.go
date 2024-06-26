//go:build no_runtime_type_checking

package csidriver

// Building without runtime type checking enabled, so all the below just return nil

func (c *jsiiProxy_CsiDriver) validateAddMoveTargetParameters(moveTarget *string) error {
	return nil
}

func (c *jsiiProxy_CsiDriver) validateAddOverrideParameters(path *string, value interface{}) error {
	return nil
}

func (c *jsiiProxy_CsiDriver) validateGetAnyMapAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (c *jsiiProxy_CsiDriver) validateGetBooleanAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (c *jsiiProxy_CsiDriver) validateGetBooleanMapAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (c *jsiiProxy_CsiDriver) validateGetListAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (c *jsiiProxy_CsiDriver) validateGetNumberAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (c *jsiiProxy_CsiDriver) validateGetNumberListAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (c *jsiiProxy_CsiDriver) validateGetNumberMapAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (c *jsiiProxy_CsiDriver) validateGetStringAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (c *jsiiProxy_CsiDriver) validateGetStringMapAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (c *jsiiProxy_CsiDriver) validateImportFromParameters(id *string) error {
	return nil
}

func (c *jsiiProxy_CsiDriver) validateInterpolationForAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (c *jsiiProxy_CsiDriver) validateMoveFromIdParameters(id *string) error {
	return nil
}

func (c *jsiiProxy_CsiDriver) validateMoveToParameters(moveTarget *string, index interface{}) error {
	return nil
}

func (c *jsiiProxy_CsiDriver) validateMoveToIdParameters(id *string) error {
	return nil
}

func (c *jsiiProxy_CsiDriver) validateOverrideLogicalIdParameters(newLogicalId *string) error {
	return nil
}

func (c *jsiiProxy_CsiDriver) validatePutMetadataParameters(value *CsiDriverMetadata) error {
	return nil
}

func (c *jsiiProxy_CsiDriver) validatePutSpecParameters(value *CsiDriverSpec) error {
	return nil
}

func validateCsiDriver_GenerateConfigForImportParameters(scope constructs.Construct, importToId *string, importFromId *string) error {
	return nil
}

func validateCsiDriver_IsConstructParameters(x interface{}) error {
	return nil
}

func validateCsiDriver_IsTerraformElementParameters(x interface{}) error {
	return nil
}

func validateCsiDriver_IsTerraformResourceParameters(x interface{}) error {
	return nil
}

func (j *jsiiProxy_CsiDriver) validateSetConnectionParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_CsiDriver) validateSetCountParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_CsiDriver) validateSetIdParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_CsiDriver) validateSetLifecycleParameters(val *cdktf.TerraformResourceLifecycle) error {
	return nil
}

func (j *jsiiProxy_CsiDriver) validateSetProvisionersParameters(val *[]interface{}) error {
	return nil
}

func validateNewCsiDriverParameters(scope constructs.Construct, id *string, config *CsiDriverConfig) error {
	return nil
}

