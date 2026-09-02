//go:build no_runtime_type_checking

package ami

// Building without runtime type checking enabled, so all the below just return nil

func (a *jsiiProxy_AmiEphemeralBlockDeviceList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (a *jsiiProxy_AmiEphemeralBlockDeviceList) validateGetParameters(index *float64) error {
	return nil
}

func (a *jsiiProxy_AmiEphemeralBlockDeviceList) validateResolveParameters(context cdktn.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_AmiEphemeralBlockDeviceList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_AmiEphemeralBlockDeviceList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_AmiEphemeralBlockDeviceList) validateSetTerraformResourceParameters(val cdktn.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_AmiEphemeralBlockDeviceList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewAmiEphemeralBlockDeviceListParameters(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

