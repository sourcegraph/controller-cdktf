//go:build no_runtime_type_checking

package backupframework

// Building without runtime type checking enabled, so all the below just return nil

func (b *jsiiProxy_BackupFrameworkControlList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (b *jsiiProxy_BackupFrameworkControlList) validateGetParameters(index *float64) error {
	return nil
}

func (b *jsiiProxy_BackupFrameworkControlList) validateResolveParameters(context cdktn.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_BackupFrameworkControlList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_BackupFrameworkControlList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_BackupFrameworkControlList) validateSetTerraformResourceParameters(val cdktn.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_BackupFrameworkControlList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewBackupFrameworkControlListParameters(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

