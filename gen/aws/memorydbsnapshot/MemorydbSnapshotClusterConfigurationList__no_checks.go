//go:build no_runtime_type_checking

package memorydbsnapshot

// Building without runtime type checking enabled, so all the below just return nil

func (m *jsiiProxy_MemorydbSnapshotClusterConfigurationList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (m *jsiiProxy_MemorydbSnapshotClusterConfigurationList) validateGetParameters(index *float64) error {
	return nil
}

func (m *jsiiProxy_MemorydbSnapshotClusterConfigurationList) validateResolveParameters(context cdktn.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_MemorydbSnapshotClusterConfigurationList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_MemorydbSnapshotClusterConfigurationList) validateSetTerraformResourceParameters(val cdktn.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_MemorydbSnapshotClusterConfigurationList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewMemorydbSnapshotClusterConfigurationListParameters(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

