//go:build no_runtime_type_checking

package variable

// Building without runtime type checking enabled, so all the below just return nil

func (v *jsiiProxy_Variable) validateAddOverrideParameters(path *string, value interface{}) error {
	return nil
}

func (v *jsiiProxy_Variable) validateGetAnyMapAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (v *jsiiProxy_Variable) validateGetBooleanAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (v *jsiiProxy_Variable) validateGetBooleanMapAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (v *jsiiProxy_Variable) validateGetListAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (v *jsiiProxy_Variable) validateGetNumberAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (v *jsiiProxy_Variable) validateGetNumberListAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (v *jsiiProxy_Variable) validateGetNumberMapAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (v *jsiiProxy_Variable) validateGetStringAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (v *jsiiProxy_Variable) validateGetStringMapAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (v *jsiiProxy_Variable) validateInterpolationForAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (v *jsiiProxy_Variable) validateOverrideLogicalIdParameters(newLogicalId *string) error {
	return nil
}

func validateVariable_IsConstructParameters(x interface{}) error {
	return nil
}

func (j *jsiiProxy_Variable) validateSetCategoryParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_Variable) validateSetConnectionParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_Variable) validateSetDescriptionParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_Variable) validateSetHclParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_Variable) validateSetIdParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_Variable) validateSetKeyParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_Variable) validateSetLifecycleParameters(val *cdktf.TerraformResourceLifecycle) error {
	return nil
}

func (j *jsiiProxy_Variable) validateSetProvisionersParameters(val *[]interface{}) error {
	return nil
}

func (j *jsiiProxy_Variable) validateSetSensitiveParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_Variable) validateSetValueParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_Variable) validateSetVariableSetIdParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_Variable) validateSetWorkspaceIdParameters(val *string) error {
	return nil
}

func validateNewVariableParameters(scope constructs.Construct, id *string, config *VariableConfig) error {
	return nil
}

