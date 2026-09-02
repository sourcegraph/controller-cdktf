//go:build no_runtime_type_checking

package serviceaccountv1

// Building without runtime type checking enabled, so all the below just return nil

func (s *jsiiProxy_ServiceAccountV1SecretList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (s *jsiiProxy_ServiceAccountV1SecretList) validateGetParameters(index *float64) error {
	return nil
}

func (s *jsiiProxy_ServiceAccountV1SecretList) validateResolveParameters(context cdktn.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_ServiceAccountV1SecretList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_ServiceAccountV1SecretList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_ServiceAccountV1SecretList) validateSetTerraformResourceParameters(val cdktn.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_ServiceAccountV1SecretList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewServiceAccountV1SecretListParameters(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

