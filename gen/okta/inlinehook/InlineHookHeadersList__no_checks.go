//go:build no_runtime_type_checking

package inlinehook

// Building without runtime type checking enabled, so all the below just return nil

func (i *jsiiProxy_InlineHookHeadersList) validateGetParameters(index *float64) error {
	return nil
}

func (i *jsiiProxy_InlineHookHeadersList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_InlineHookHeadersList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_InlineHookHeadersList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_InlineHookHeadersList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_InlineHookHeadersList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewInlineHookHeadersListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

