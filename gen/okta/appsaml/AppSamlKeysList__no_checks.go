//go:build no_runtime_type_checking

package appsaml

// Building without runtime type checking enabled, so all the below just return nil

func (a *jsiiProxy_AppSamlKeysList) validateGetParameters(index *float64) error {
	return nil
}

func (a *jsiiProxy_AppSamlKeysList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_AppSamlKeysList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_AppSamlKeysList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_AppSamlKeysList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewAppSamlKeysListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

