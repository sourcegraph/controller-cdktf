//go:build no_runtime_type_checking

package appsaml

// Building without runtime type checking enabled, so all the below just return nil

func (a *jsiiProxy_AppSamlKeysList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (a *jsiiProxy_AppSamlKeysList) validateGetParameters(index *float64) error {
	return nil
}

func (a *jsiiProxy_AppSamlKeysList) validateResolveParameters(context cdktn.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_AppSamlKeysList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_AppSamlKeysList) validateSetTerraformResourceParameters(val cdktn.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_AppSamlKeysList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewAppSamlKeysListParameters(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

