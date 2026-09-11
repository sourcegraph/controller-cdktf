//go:build no_runtime_type_checking

package env

// Building without runtime type checking enabled, so all the below just return nil

func (e *jsiiProxy_EnvEnvList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (e *jsiiProxy_EnvEnvList) validateGetParameters(index *float64) error {
	return nil
}

func (e *jsiiProxy_EnvEnvList) validateResolveParameters(context cdktn.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_EnvEnvList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_EnvEnvList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_EnvEnvList) validateSetTerraformResourceParameters(val cdktn.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_EnvEnvList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewEnvEnvListParameters(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

