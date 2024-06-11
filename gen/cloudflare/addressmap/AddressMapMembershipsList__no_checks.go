//go:build no_runtime_type_checking

package addressmap

// Building without runtime type checking enabled, so all the below just return nil

func (a *jsiiProxy_AddressMapMembershipsList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (a *jsiiProxy_AddressMapMembershipsList) validateGetParameters(index *float64) error {
	return nil
}

func (a *jsiiProxy_AddressMapMembershipsList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_AddressMapMembershipsList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_AddressMapMembershipsList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_AddressMapMembershipsList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_AddressMapMembershipsList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewAddressMapMembershipsListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

