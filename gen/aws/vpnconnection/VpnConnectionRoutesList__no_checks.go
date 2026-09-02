//go:build no_runtime_type_checking

package vpnconnection

// Building without runtime type checking enabled, so all the below just return nil

func (v *jsiiProxy_VpnConnectionRoutesList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (v *jsiiProxy_VpnConnectionRoutesList) validateGetParameters(index *float64) error {
	return nil
}

func (v *jsiiProxy_VpnConnectionRoutesList) validateResolveParameters(context cdktn.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_VpnConnectionRoutesList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_VpnConnectionRoutesList) validateSetTerraformResourceParameters(val cdktn.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_VpnConnectionRoutesList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewVpnConnectionRoutesListParameters(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

