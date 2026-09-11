//go:build no_runtime_type_checking

package pod

// Building without runtime type checking enabled, so all the below just return nil

func (p *jsiiProxy_PodSpecInitContainerList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (p *jsiiProxy_PodSpecInitContainerList) validateGetParameters(index *float64) error {
	return nil
}

func (p *jsiiProxy_PodSpecInitContainerList) validateResolveParameters(context cdktn.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_PodSpecInitContainerList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_PodSpecInitContainerList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_PodSpecInitContainerList) validateSetTerraformResourceParameters(val cdktn.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_PodSpecInitContainerList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewPodSpecInitContainerListParameters(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

