//go:build no_runtime_type_checking

package slo

// Building without runtime type checking enabled, so all the below just return nil

func (s *jsiiProxy_SloAttachmentList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (s *jsiiProxy_SloAttachmentList) validateGetParameters(index *float64) error {
	return nil
}

func (s *jsiiProxy_SloAttachmentList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_SloAttachmentList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_SloAttachmentList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_SloAttachmentList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_SloAttachmentList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewSloAttachmentListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

