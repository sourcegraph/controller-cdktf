//go:build no_runtime_type_checking

package computestoragepool

// Building without runtime type checking enabled, so all the below just return nil

func (c *jsiiProxy_ComputeStoragePoolResourceStatusList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (c *jsiiProxy_ComputeStoragePoolResourceStatusList) validateGetParameters(index *float64) error {
	return nil
}

func (c *jsiiProxy_ComputeStoragePoolResourceStatusList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_ComputeStoragePoolResourceStatusList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_ComputeStoragePoolResourceStatusList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_ComputeStoragePoolResourceStatusList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewComputeStoragePoolResourceStatusListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

