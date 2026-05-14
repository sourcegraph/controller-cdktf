//go:build no_runtime_type_checking

package netappvolume

// Building without runtime type checking enabled, so all the below just return nil

func (n *jsiiProxy_NetappVolumeBlockDevicesList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (n *jsiiProxy_NetappVolumeBlockDevicesList) validateGetParameters(index *float64) error {
	return nil
}

func (n *jsiiProxy_NetappVolumeBlockDevicesList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_NetappVolumeBlockDevicesList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_NetappVolumeBlockDevicesList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_NetappVolumeBlockDevicesList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_NetappVolumeBlockDevicesList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewNetappVolumeBlockDevicesListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

