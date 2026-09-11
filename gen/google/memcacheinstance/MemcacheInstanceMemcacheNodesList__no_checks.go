//go:build no_runtime_type_checking

package memcacheinstance

// Building without runtime type checking enabled, so all the below just return nil

func (m *jsiiProxy_MemcacheInstanceMemcacheNodesList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (m *jsiiProxy_MemcacheInstanceMemcacheNodesList) validateGetParameters(index *float64) error {
	return nil
}

func (m *jsiiProxy_MemcacheInstanceMemcacheNodesList) validateResolveParameters(context cdktn.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_MemcacheInstanceMemcacheNodesList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_MemcacheInstanceMemcacheNodesList) validateSetTerraformResourceParameters(val cdktn.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_MemcacheInstanceMemcacheNodesList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewMemcacheInstanceMemcacheNodesListParameters(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

