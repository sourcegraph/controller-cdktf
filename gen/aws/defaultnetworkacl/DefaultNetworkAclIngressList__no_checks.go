//go:build no_runtime_type_checking

package defaultnetworkacl

// Building without runtime type checking enabled, so all the below just return nil

func (d *jsiiProxy_DefaultNetworkAclIngressList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (d *jsiiProxy_DefaultNetworkAclIngressList) validateGetParameters(index *float64) error {
	return nil
}

func (d *jsiiProxy_DefaultNetworkAclIngressList) validateResolveParameters(context cdktn.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_DefaultNetworkAclIngressList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_DefaultNetworkAclIngressList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_DefaultNetworkAclIngressList) validateSetTerraformResourceParameters(val cdktn.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_DefaultNetworkAclIngressList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewDefaultNetworkAclIngressListParameters(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

