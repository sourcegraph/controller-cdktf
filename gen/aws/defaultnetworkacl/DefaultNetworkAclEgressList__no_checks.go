//go:build no_runtime_type_checking

package defaultnetworkacl

// Building without runtime type checking enabled, so all the below just return nil

func (d *jsiiProxy_DefaultNetworkAclEgressList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (d *jsiiProxy_DefaultNetworkAclEgressList) validateGetParameters(index *float64) error {
	return nil
}

func (d *jsiiProxy_DefaultNetworkAclEgressList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_DefaultNetworkAclEgressList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_DefaultNetworkAclEgressList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_DefaultNetworkAclEgressList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_DefaultNetworkAclEgressList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewDefaultNetworkAclEgressListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

