//go:build no_runtime_type_checking

package pod

// Building without runtime type checking enabled, so all the below just return nil

func (p *jsiiProxy_PodSpecContainerList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (p *jsiiProxy_PodSpecContainerList) validateGetParameters(index *float64) error {
	return nil
}

func (p *jsiiProxy_PodSpecContainerList) validateResolveParameters(context cdktn.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_PodSpecContainerList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_PodSpecContainerList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_PodSpecContainerList) validateSetTerraformResourceParameters(val cdktn.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_PodSpecContainerList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewPodSpecContainerListParameters(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

