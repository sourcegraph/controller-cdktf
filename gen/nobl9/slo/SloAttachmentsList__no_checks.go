//go:build no_runtime_type_checking

package slo

// Building without runtime type checking enabled, so all the below just return nil

func (s *jsiiProxy_SloAttachmentsList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (s *jsiiProxy_SloAttachmentsList) validateGetParameters(index *float64) error {
	return nil
}

func (s *jsiiProxy_SloAttachmentsList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_SloAttachmentsList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_SloAttachmentsList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_SloAttachmentsList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_SloAttachmentsList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewSloAttachmentsListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

