//go:build no_runtime_type_checking

package memorystoreinstance

// Building without runtime type checking enabled, so all the below just return nil

func (m *jsiiProxy_MemorystoreInstanceStateInfoList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (m *jsiiProxy_MemorystoreInstanceStateInfoList) validateGetParameters(index *float64) error {
	return nil
}

func (m *jsiiProxy_MemorystoreInstanceStateInfoList) validateResolveParameters(context cdktn.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_MemorystoreInstanceStateInfoList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_MemorystoreInstanceStateInfoList) validateSetTerraformResourceParameters(val cdktn.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_MemorystoreInstanceStateInfoList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewMemorystoreInstanceStateInfoListParameters(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

