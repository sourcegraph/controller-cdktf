//go:build no_runtime_type_checking

package job

// Building without runtime type checking enabled, so all the below just return nil

func (j *jsiiProxy_JobSpecTemplateSpecContainerList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (j *jsiiProxy_JobSpecTemplateSpecContainerList) validateGetParameters(index *float64) error {
	return nil
}

func (j *jsiiProxy_JobSpecTemplateSpecContainerList) validateResolveParameters(context cdktn.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_JobSpecTemplateSpecContainerList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_JobSpecTemplateSpecContainerList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_JobSpecTemplateSpecContainerList) validateSetTerraformResourceParameters(val cdktn.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_JobSpecTemplateSpecContainerList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewJobSpecTemplateSpecContainerListParameters(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

