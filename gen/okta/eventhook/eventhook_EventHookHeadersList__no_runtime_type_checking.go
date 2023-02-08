//go:build no_runtime_type_checking

package eventhook

// Building without runtime type checking enabled, so all the below just return nil

func (e *jsiiProxy_EventHookHeadersList) validateGetParameters(index *float64) error {
	return nil
}

func (e *jsiiProxy_EventHookHeadersList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_EventHookHeadersList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_EventHookHeadersList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_EventHookHeadersList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_EventHookHeadersList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewEventHookHeadersListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

