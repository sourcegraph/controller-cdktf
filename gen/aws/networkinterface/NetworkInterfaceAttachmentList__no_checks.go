//go:build no_runtime_type_checking

package networkinterface

// Building without runtime type checking enabled, so all the below just return nil

func (n *jsiiProxy_NetworkInterfaceAttachmentList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (n *jsiiProxy_NetworkInterfaceAttachmentList) validateGetParameters(index *float64) error {
	return nil
}

func (n *jsiiProxy_NetworkInterfaceAttachmentList) validateResolveParameters(context cdktn.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_NetworkInterfaceAttachmentList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_NetworkInterfaceAttachmentList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_NetworkInterfaceAttachmentList) validateSetTerraformResourceParameters(val cdktn.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_NetworkInterfaceAttachmentList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewNetworkInterfaceAttachmentListParameters(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

