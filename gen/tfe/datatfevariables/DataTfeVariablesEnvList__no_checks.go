//go:build no_runtime_type_checking

package datatfevariables

// Building without runtime type checking enabled, so all the below just return nil

func (d *jsiiProxy_DataTfeVariablesEnvList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (d *jsiiProxy_DataTfeVariablesEnvList) validateGetParameters(index *float64) error {
	return nil
}

func (d *jsiiProxy_DataTfeVariablesEnvList) validateResolveParameters(context cdktn.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_DataTfeVariablesEnvList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_DataTfeVariablesEnvList) validateSetTerraformResourceParameters(val cdktn.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_DataTfeVariablesEnvList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewDataTfeVariablesEnvListParameters(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

