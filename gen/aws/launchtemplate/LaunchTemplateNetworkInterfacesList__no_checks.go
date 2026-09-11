//go:build no_runtime_type_checking

package launchtemplate

// Building without runtime type checking enabled, so all the below just return nil

func (l *jsiiProxy_LaunchTemplateNetworkInterfacesList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (l *jsiiProxy_LaunchTemplateNetworkInterfacesList) validateGetParameters(index *float64) error {
	return nil
}

func (l *jsiiProxy_LaunchTemplateNetworkInterfacesList) validateResolveParameters(context cdktn.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_LaunchTemplateNetworkInterfacesList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_LaunchTemplateNetworkInterfacesList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_LaunchTemplateNetworkInterfacesList) validateSetTerraformResourceParameters(val cdktn.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_LaunchTemplateNetworkInterfacesList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewLaunchTemplateNetworkInterfacesListParameters(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

