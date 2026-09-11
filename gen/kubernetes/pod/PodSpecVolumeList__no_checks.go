//go:build no_runtime_type_checking

package pod

// Building without runtime type checking enabled, so all the below just return nil

func (p *jsiiProxy_PodSpecVolumeList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (p *jsiiProxy_PodSpecVolumeList) validateGetParameters(index *float64) error {
	return nil
}

func (p *jsiiProxy_PodSpecVolumeList) validateResolveParameters(context cdktn.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_PodSpecVolumeList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_PodSpecVolumeList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_PodSpecVolumeList) validateSetTerraformResourceParameters(val cdktn.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_PodSpecVolumeList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewPodSpecVolumeListParameters(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

