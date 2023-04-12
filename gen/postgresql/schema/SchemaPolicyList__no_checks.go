//go:build no_runtime_type_checking

package schema

// Building without runtime type checking enabled, so all the below just return nil

func (s *jsiiProxy_SchemaPolicyList) validateGetParameters(index *float64) error {
	return nil
}

func (s *jsiiProxy_SchemaPolicyList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_SchemaPolicyList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_SchemaPolicyList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_SchemaPolicyList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_SchemaPolicyList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewSchemaPolicyListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

