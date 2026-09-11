//go:build no_runtime_type_checking

package pod

// Building without runtime type checking enabled, so all the below just return nil

func (p *jsiiProxy_PodSpecVolumeProjectedList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (p *jsiiProxy_PodSpecVolumeProjectedList) validateGetParameters(index *float64) error {
	return nil
}

func (p *jsiiProxy_PodSpecVolumeProjectedList) validateResolveParameters(context cdktn.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_PodSpecVolumeProjectedList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_PodSpecVolumeProjectedList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_PodSpecVolumeProjectedList) validateSetTerraformResourceParameters(val cdktn.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_PodSpecVolumeProjectedList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewPodSpecVolumeProjectedListParameters(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

