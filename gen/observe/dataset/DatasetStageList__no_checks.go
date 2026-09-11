//go:build no_runtime_type_checking

package dataset

// Building without runtime type checking enabled, so all the below just return nil

func (d *jsiiProxy_DatasetStageList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (d *jsiiProxy_DatasetStageList) validateGetParameters(index *float64) error {
	return nil
}

func (d *jsiiProxy_DatasetStageList) validateResolveParameters(context cdktn.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_DatasetStageList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_DatasetStageList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_DatasetStageList) validateSetTerraformResourceParameters(val cdktn.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_DatasetStageList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewDatasetStageListParameters(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

