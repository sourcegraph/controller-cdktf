//go:build no_runtime_type_checking

package emrcluster

// Building without runtime type checking enabled, so all the below just return nil

func (e *jsiiProxy_EmrClusterStepList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (e *jsiiProxy_EmrClusterStepList) validateGetParameters(index *float64) error {
	return nil
}

func (e *jsiiProxy_EmrClusterStepList) validateResolveParameters(context cdktn.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_EmrClusterStepList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_EmrClusterStepList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_EmrClusterStepList) validateSetTerraformResourceParameters(val cdktn.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_EmrClusterStepList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewEmrClusterStepListParameters(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

