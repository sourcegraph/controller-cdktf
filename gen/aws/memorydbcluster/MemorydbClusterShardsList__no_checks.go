//go:build no_runtime_type_checking

package memorydbcluster

// Building without runtime type checking enabled, so all the below just return nil

func (m *jsiiProxy_MemorydbClusterShardsList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (m *jsiiProxy_MemorydbClusterShardsList) validateGetParameters(index *float64) error {
	return nil
}

func (m *jsiiProxy_MemorydbClusterShardsList) validateResolveParameters(context cdktn.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_MemorydbClusterShardsList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_MemorydbClusterShardsList) validateSetTerraformResourceParameters(val cdktn.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_MemorydbClusterShardsList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewMemorydbClusterShardsListParameters(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

