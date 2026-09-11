//go:build no_runtime_type_checking

package kmsgrant

// Building without runtime type checking enabled, so all the below just return nil

func (k *jsiiProxy_KmsGrantConstraintsList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (k *jsiiProxy_KmsGrantConstraintsList) validateGetParameters(index *float64) error {
	return nil
}

func (k *jsiiProxy_KmsGrantConstraintsList) validateResolveParameters(context cdktn.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_KmsGrantConstraintsList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_KmsGrantConstraintsList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_KmsGrantConstraintsList) validateSetTerraformResourceParameters(val cdktn.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_KmsGrantConstraintsList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewKmsGrantConstraintsListParameters(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

