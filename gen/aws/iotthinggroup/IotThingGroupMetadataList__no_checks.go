//go:build no_runtime_type_checking

package iotthinggroup

// Building without runtime type checking enabled, so all the below just return nil

func (i *jsiiProxy_IotThingGroupMetadataList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (i *jsiiProxy_IotThingGroupMetadataList) validateGetParameters(index *float64) error {
	return nil
}

func (i *jsiiProxy_IotThingGroupMetadataList) validateResolveParameters(context cdktn.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_IotThingGroupMetadataList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_IotThingGroupMetadataList) validateSetTerraformResourceParameters(val cdktn.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_IotThingGroupMetadataList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewIotThingGroupMetadataListParameters(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

