//go:build no_runtime_type_checking

package sourcedataset

// Building without runtime type checking enabled, so all the below just return nil

func (s *jsiiProxy_SourceDatasetFieldList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (s *jsiiProxy_SourceDatasetFieldList) validateGetParameters(index *float64) error {
	return nil
}

func (s *jsiiProxy_SourceDatasetFieldList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_SourceDatasetFieldList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_SourceDatasetFieldList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_SourceDatasetFieldList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_SourceDatasetFieldList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewSourceDatasetFieldListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

