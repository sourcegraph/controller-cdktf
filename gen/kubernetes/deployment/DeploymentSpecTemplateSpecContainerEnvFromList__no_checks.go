//go:build no_runtime_type_checking

package deployment

// Building without runtime type checking enabled, so all the below just return nil

func (d *jsiiProxy_DeploymentSpecTemplateSpecContainerEnvFromList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (d *jsiiProxy_DeploymentSpecTemplateSpecContainerEnvFromList) validateGetParameters(index *float64) error {
	return nil
}

func (d *jsiiProxy_DeploymentSpecTemplateSpecContainerEnvFromList) validateResolveParameters(context cdktn.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_DeploymentSpecTemplateSpecContainerEnvFromList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_DeploymentSpecTemplateSpecContainerEnvFromList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_DeploymentSpecTemplateSpecContainerEnvFromList) validateSetTerraformResourceParameters(val cdktn.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_DeploymentSpecTemplateSpecContainerEnvFromList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewDeploymentSpecTemplateSpecContainerEnvFromListParameters(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

