//go:build no_runtime_type_checking

package memorydbcluster

// Building without runtime type checking enabled, so all the below just return nil

func (m *jsiiProxy_MemorydbClusterShardsNodesList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (m *jsiiProxy_MemorydbClusterShardsNodesList) validateGetParameters(index *float64) error {
	return nil
}

func (m *jsiiProxy_MemorydbClusterShardsNodesList) validateResolveParameters(context cdktn.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_MemorydbClusterShardsNodesList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_MemorydbClusterShardsNodesList) validateSetTerraformResourceParameters(val cdktn.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_MemorydbClusterShardsNodesList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewMemorydbClusterShardsNodesListParameters(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

