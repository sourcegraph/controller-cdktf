//go:build no_runtime_type_checking

package backupplan

// Building without runtime type checking enabled, so all the below just return nil

func (b *jsiiProxy_BackupPlanAdvancedBackupSettingList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (b *jsiiProxy_BackupPlanAdvancedBackupSettingList) validateGetParameters(index *float64) error {
	return nil
}

func (b *jsiiProxy_BackupPlanAdvancedBackupSettingList) validateResolveParameters(context cdktn.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_BackupPlanAdvancedBackupSettingList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_BackupPlanAdvancedBackupSettingList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_BackupPlanAdvancedBackupSettingList) validateSetTerraformResourceParameters(val cdktn.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_BackupPlanAdvancedBackupSettingList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewBackupPlanAdvancedBackupSettingListParameters(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

