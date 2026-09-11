//go:build no_runtime_type_checking

package alloydbcluster

// Building without runtime type checking enabled, so all the below just return nil

func (a *jsiiProxy_AlloydbClusterBackupSourceList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (a *jsiiProxy_AlloydbClusterBackupSourceList) validateGetParameters(index *float64) error {
	return nil
}

func (a *jsiiProxy_AlloydbClusterBackupSourceList) validateResolveParameters(context cdktn.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_AlloydbClusterBackupSourceList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_AlloydbClusterBackupSourceList) validateSetTerraformResourceParameters(val cdktn.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_AlloydbClusterBackupSourceList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewAlloydbClusterBackupSourceListParameters(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

