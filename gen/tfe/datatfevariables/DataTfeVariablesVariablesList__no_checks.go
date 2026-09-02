//go:build no_runtime_type_checking

package datatfevariables

// Building without runtime type checking enabled, so all the below just return nil

func (d *jsiiProxy_DataTfeVariablesVariablesList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (d *jsiiProxy_DataTfeVariablesVariablesList) validateGetParameters(index *float64) error {
	return nil
}

func (d *jsiiProxy_DataTfeVariablesVariablesList) validateResolveParameters(context cdktn.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_DataTfeVariablesVariablesList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_DataTfeVariablesVariablesList) validateSetTerraformResourceParameters(val cdktn.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_DataTfeVariablesVariablesList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewDataTfeVariablesVariablesListParameters(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

