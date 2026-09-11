//go:build no_runtime_type_checking

package defaultserviceaccount

// Building without runtime type checking enabled, so all the below just return nil

func (d *jsiiProxy_DefaultServiceAccountSecretList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (d *jsiiProxy_DefaultServiceAccountSecretList) validateGetParameters(index *float64) error {
	return nil
}

func (d *jsiiProxy_DefaultServiceAccountSecretList) validateResolveParameters(context cdktn.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_DefaultServiceAccountSecretList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_DefaultServiceAccountSecretList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_DefaultServiceAccountSecretList) validateSetTerraformResourceParameters(val cdktn.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_DefaultServiceAccountSecretList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewDefaultServiceAccountSecretListParameters(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

