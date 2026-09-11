//go:build no_runtime_type_checking

package apphubservice

// Building without runtime type checking enabled, so all the below just return nil

func (a *jsiiProxy_ApphubServiceServiceReferenceList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (a *jsiiProxy_ApphubServiceServiceReferenceList) validateGetParameters(index *float64) error {
	return nil
}

func (a *jsiiProxy_ApphubServiceServiceReferenceList) validateResolveParameters(context cdktn.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_ApphubServiceServiceReferenceList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_ApphubServiceServiceReferenceList) validateSetTerraformResourceParameters(val cdktn.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_ApphubServiceServiceReferenceList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewApphubServiceServiceReferenceListParameters(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

