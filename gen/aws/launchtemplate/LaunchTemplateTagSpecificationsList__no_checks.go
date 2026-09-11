//go:build no_runtime_type_checking

package launchtemplate

// Building without runtime type checking enabled, so all the below just return nil

func (l *jsiiProxy_LaunchTemplateTagSpecificationsList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (l *jsiiProxy_LaunchTemplateTagSpecificationsList) validateGetParameters(index *float64) error {
	return nil
}

func (l *jsiiProxy_LaunchTemplateTagSpecificationsList) validateResolveParameters(context cdktn.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_LaunchTemplateTagSpecificationsList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_LaunchTemplateTagSpecificationsList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_LaunchTemplateTagSpecificationsList) validateSetTerraformResourceParameters(val cdktn.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_LaunchTemplateTagSpecificationsList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewLaunchTemplateTagSpecificationsListParameters(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

