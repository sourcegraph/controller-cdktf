//go:build no_runtime_type_checking

package dataplexasset

// Building without runtime type checking enabled, so all the below just return nil

func (d *jsiiProxy_DataplexAssetResourceStatusList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (d *jsiiProxy_DataplexAssetResourceStatusList) validateGetParameters(index *float64) error {
	return nil
}

func (d *jsiiProxy_DataplexAssetResourceStatusList) validateResolveParameters(context cdktn.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_DataplexAssetResourceStatusList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_DataplexAssetResourceStatusList) validateSetTerraformResourceParameters(val cdktn.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_DataplexAssetResourceStatusList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewDataplexAssetResourceStatusListParameters(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

