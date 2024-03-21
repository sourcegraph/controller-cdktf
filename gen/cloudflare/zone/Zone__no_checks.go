//go:build no_runtime_type_checking

package zone

// Building without runtime type checking enabled, so all the below just return nil

func (z *jsiiProxy_Zone) validateAddMoveTargetParameters(moveTarget *string) error {
	return nil
}

func (z *jsiiProxy_Zone) validateAddOverrideParameters(path *string, value interface{}) error {
	return nil
}

func (z *jsiiProxy_Zone) validateGetAnyMapAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (z *jsiiProxy_Zone) validateGetBooleanAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (z *jsiiProxy_Zone) validateGetBooleanMapAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (z *jsiiProxy_Zone) validateGetListAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (z *jsiiProxy_Zone) validateGetNumberAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (z *jsiiProxy_Zone) validateGetNumberListAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (z *jsiiProxy_Zone) validateGetNumberMapAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (z *jsiiProxy_Zone) validateGetStringAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (z *jsiiProxy_Zone) validateGetStringMapAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (z *jsiiProxy_Zone) validateImportFromParameters(id *string) error {
	return nil
}

func (z *jsiiProxy_Zone) validateInterpolationForAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (z *jsiiProxy_Zone) validateMoveFromIdParameters(id *string) error {
	return nil
}

func (z *jsiiProxy_Zone) validateMoveToParameters(moveTarget *string, index interface{}) error {
	return nil
}

func (z *jsiiProxy_Zone) validateMoveToIdParameters(id *string) error {
	return nil
}

func (z *jsiiProxy_Zone) validateOverrideLogicalIdParameters(newLogicalId *string) error {
	return nil
}

func validateZone_GenerateConfigForImportParameters(scope constructs.Construct, importToId *string, importFromId *string) error {
	return nil
}

func validateZone_IsConstructParameters(x interface{}) error {
	return nil
}

func validateZone_IsTerraformElementParameters(x interface{}) error {
	return nil
}

func validateZone_IsTerraformResourceParameters(x interface{}) error {
	return nil
}

func (j *jsiiProxy_Zone) validateSetAccountIdParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_Zone) validateSetConnectionParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_Zone) validateSetCountParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_Zone) validateSetIdParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_Zone) validateSetJumpStartParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_Zone) validateSetLifecycleParameters(val *cdktf.TerraformResourceLifecycle) error {
	return nil
}

func (j *jsiiProxy_Zone) validateSetPausedParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_Zone) validateSetPlanParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_Zone) validateSetProvisionersParameters(val *[]interface{}) error {
	return nil
}

func (j *jsiiProxy_Zone) validateSetTypeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_Zone) validateSetZoneParameters(val *string) error {
	return nil
}

func validateNewZoneParameters(scope constructs.Construct, id *string, config *ZoneConfig) error {
	return nil
}

