//go:build no_runtime_type_checking

package ecrrepository

// Building without runtime type checking enabled, so all the below just return nil

func (e *jsiiProxy_EcrRepositoryEncryptionConfigurationList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (e *jsiiProxy_EcrRepositoryEncryptionConfigurationList) validateGetParameters(index *float64) error {
	return nil
}

func (e *jsiiProxy_EcrRepositoryEncryptionConfigurationList) validateResolveParameters(context cdktn.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_EcrRepositoryEncryptionConfigurationList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_EcrRepositoryEncryptionConfigurationList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_EcrRepositoryEncryptionConfigurationList) validateSetTerraformResourceParameters(val cdktn.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_EcrRepositoryEncryptionConfigurationList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewEcrRepositoryEncryptionConfigurationListParameters(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

