//go:build no_runtime_type_checking

package apishield

// Building without runtime type checking enabled, so all the below just return nil

func (a *jsiiProxy_ApiShieldAuthIdCharacteristicsList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (a *jsiiProxy_ApiShieldAuthIdCharacteristicsList) validateGetParameters(index *float64) error {
	return nil
}

func (a *jsiiProxy_ApiShieldAuthIdCharacteristicsList) validateResolveParameters(context cdktn.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_ApiShieldAuthIdCharacteristicsList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_ApiShieldAuthIdCharacteristicsList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_ApiShieldAuthIdCharacteristicsList) validateSetTerraformResourceParameters(val cdktn.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_ApiShieldAuthIdCharacteristicsList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewApiShieldAuthIdCharacteristicsListParameters(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

