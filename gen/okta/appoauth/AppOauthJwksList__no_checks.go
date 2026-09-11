//go:build no_runtime_type_checking

package appoauth

// Building without runtime type checking enabled, so all the below just return nil

func (a *jsiiProxy_AppOauthJwksList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (a *jsiiProxy_AppOauthJwksList) validateGetParameters(index *float64) error {
	return nil
}

func (a *jsiiProxy_AppOauthJwksList) validateResolveParameters(context cdktn.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_AppOauthJwksList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_AppOauthJwksList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_AppOauthJwksList) validateSetTerraformResourceParameters(val cdktn.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_AppOauthJwksList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewAppOauthJwksListParameters(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

