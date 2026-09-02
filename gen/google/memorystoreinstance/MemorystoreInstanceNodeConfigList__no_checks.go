//go:build no_runtime_type_checking

package memorystoreinstance

// Building without runtime type checking enabled, so all the below just return nil

func (m *jsiiProxy_MemorystoreInstanceNodeConfigList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (m *jsiiProxy_MemorystoreInstanceNodeConfigList) validateGetParameters(index *float64) error {
	return nil
}

func (m *jsiiProxy_MemorystoreInstanceNodeConfigList) validateResolveParameters(context cdktn.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_MemorystoreInstanceNodeConfigList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_MemorystoreInstanceNodeConfigList) validateSetTerraformResourceParameters(val cdktn.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_MemorystoreInstanceNodeConfigList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewMemorystoreInstanceNodeConfigListParameters(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

