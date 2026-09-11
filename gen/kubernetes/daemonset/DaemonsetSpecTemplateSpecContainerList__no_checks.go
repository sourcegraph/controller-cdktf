//go:build no_runtime_type_checking

package daemonset

// Building without runtime type checking enabled, so all the below just return nil

func (d *jsiiProxy_DaemonsetSpecTemplateSpecContainerList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (d *jsiiProxy_DaemonsetSpecTemplateSpecContainerList) validateGetParameters(index *float64) error {
	return nil
}

func (d *jsiiProxy_DaemonsetSpecTemplateSpecContainerList) validateResolveParameters(context cdktn.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_DaemonsetSpecTemplateSpecContainerList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_DaemonsetSpecTemplateSpecContainerList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_DaemonsetSpecTemplateSpecContainerList) validateSetTerraformResourceParameters(val cdktn.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_DaemonsetSpecTemplateSpecContainerList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewDaemonsetSpecTemplateSpecContainerListParameters(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

