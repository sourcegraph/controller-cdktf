//go:build no_runtime_type_checking

package appsaml

// Building without runtime type checking enabled, so all the below just return nil

func (a *jsiiProxy_AppSamlAttributeStatementsList) validateGetParameters(index *float64) error {
	return nil
}

func (a *jsiiProxy_AppSamlAttributeStatementsList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_AppSamlAttributeStatementsList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_AppSamlAttributeStatementsList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_AppSamlAttributeStatementsList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_AppSamlAttributeStatementsList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewAppSamlAttributeStatementsListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

