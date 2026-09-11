//go:build no_runtime_type_checking

package instance

// Building without runtime type checking enabled, so all the below just return nil

func (i *jsiiProxy_InstanceNetworkInterfaceList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (i *jsiiProxy_InstanceNetworkInterfaceList) validateGetParameters(index *float64) error {
	return nil
}

func (i *jsiiProxy_InstanceNetworkInterfaceList) validateResolveParameters(context cdktn.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_InstanceNetworkInterfaceList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_InstanceNetworkInterfaceList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_InstanceNetworkInterfaceList) validateSetTerraformResourceParameters(val cdktn.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_InstanceNetworkInterfaceList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewInstanceNetworkInterfaceListParameters(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

