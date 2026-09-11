//go:build no_runtime_type_checking

package instance

// Building without runtime type checking enabled, so all the below just return nil

func (i *jsiiProxy_InstanceEbsBlockDeviceList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (i *jsiiProxy_InstanceEbsBlockDeviceList) validateGetParameters(index *float64) error {
	return nil
}

func (i *jsiiProxy_InstanceEbsBlockDeviceList) validateResolveParameters(context cdktn.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_InstanceEbsBlockDeviceList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_InstanceEbsBlockDeviceList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_InstanceEbsBlockDeviceList) validateSetTerraformResourceParameters(val cdktn.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_InstanceEbsBlockDeviceList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewInstanceEbsBlockDeviceListParameters(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

