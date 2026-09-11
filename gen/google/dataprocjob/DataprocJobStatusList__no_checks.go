//go:build no_runtime_type_checking

package dataprocjob

// Building without runtime type checking enabled, so all the below just return nil

func (d *jsiiProxy_DataprocJobStatusList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (d *jsiiProxy_DataprocJobStatusList) validateGetParameters(index *float64) error {
	return nil
}

func (d *jsiiProxy_DataprocJobStatusList) validateResolveParameters(context cdktn.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_DataprocJobStatusList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_DataprocJobStatusList) validateSetTerraformResourceParameters(val cdktn.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_DataprocJobStatusList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewDataprocJobStatusListParameters(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

