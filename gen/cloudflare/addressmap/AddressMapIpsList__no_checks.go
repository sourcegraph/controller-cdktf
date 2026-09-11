//go:build no_runtime_type_checking

package addressmap

// Building without runtime type checking enabled, so all the below just return nil

func (a *jsiiProxy_AddressMapIpsList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (a *jsiiProxy_AddressMapIpsList) validateGetParameters(index *float64) error {
	return nil
}

func (a *jsiiProxy_AddressMapIpsList) validateResolveParameters(context cdktn.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_AddressMapIpsList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_AddressMapIpsList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_AddressMapIpsList) validateSetTerraformResourceParameters(val cdktn.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_AddressMapIpsList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewAddressMapIpsListParameters(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

