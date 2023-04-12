//go:build no_runtime_type_checking

package samlapp

// Building without runtime type checking enabled, so all the below just return nil

func (s *jsiiProxy_SamlAppAttributeStatementsList) validateGetParameters(index *float64) error {
	return nil
}

func (s *jsiiProxy_SamlAppAttributeStatementsList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_SamlAppAttributeStatementsList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_SamlAppAttributeStatementsList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_SamlAppAttributeStatementsList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_SamlAppAttributeStatementsList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewSamlAppAttributeStatementsListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

