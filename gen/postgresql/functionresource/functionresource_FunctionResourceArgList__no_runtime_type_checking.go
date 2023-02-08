//go:build no_runtime_type_checking

package functionresource

// Building without runtime type checking enabled, so all the below just return nil

func (f *jsiiProxy_FunctionResourceArgList) validateGetParameters(index *float64) error {
	return nil
}

func (f *jsiiProxy_FunctionResourceArgList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_FunctionResourceArgList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_FunctionResourceArgList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_FunctionResourceArgList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_FunctionResourceArgList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewFunctionResourceArgListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

