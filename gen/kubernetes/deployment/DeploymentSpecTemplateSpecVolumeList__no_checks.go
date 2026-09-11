//go:build no_runtime_type_checking

package deployment

// Building without runtime type checking enabled, so all the below just return nil

func (d *jsiiProxy_DeploymentSpecTemplateSpecVolumeList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (d *jsiiProxy_DeploymentSpecTemplateSpecVolumeList) validateGetParameters(index *float64) error {
	return nil
}

func (d *jsiiProxy_DeploymentSpecTemplateSpecVolumeList) validateResolveParameters(context cdktn.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_DeploymentSpecTemplateSpecVolumeList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_DeploymentSpecTemplateSpecVolumeList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_DeploymentSpecTemplateSpecVolumeList) validateSetTerraformResourceParameters(val cdktn.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_DeploymentSpecTemplateSpecVolumeList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewDeploymentSpecTemplateSpecVolumeListParameters(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

