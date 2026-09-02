//go:build no_runtime_type_checking

package appsaml

// Building without runtime type checking enabled, so all the below just return nil

func (a *jsiiProxy_AppSamlAttributeStatementsList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (a *jsiiProxy_AppSamlAttributeStatementsList) validateGetParameters(index *float64) error {
	return nil
}

func (a *jsiiProxy_AppSamlAttributeStatementsList) validateResolveParameters(context cdktn.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_AppSamlAttributeStatementsList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_AppSamlAttributeStatementsList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_AppSamlAttributeStatementsList) validateSetTerraformResourceParameters(val cdktn.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_AppSamlAttributeStatementsList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewAppSamlAttributeStatementsListParameters(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

