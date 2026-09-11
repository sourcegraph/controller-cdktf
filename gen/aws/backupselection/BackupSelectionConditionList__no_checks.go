//go:build no_runtime_type_checking

package backupselection

// Building without runtime type checking enabled, so all the below just return nil

func (b *jsiiProxy_BackupSelectionConditionList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (b *jsiiProxy_BackupSelectionConditionList) validateGetParameters(index *float64) error {
	return nil
}

func (b *jsiiProxy_BackupSelectionConditionList) validateResolveParameters(context cdktn.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_BackupSelectionConditionList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_BackupSelectionConditionList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_BackupSelectionConditionList) validateSetTerraformResourceParameters(val cdktn.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_BackupSelectionConditionList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewBackupSelectionConditionListParameters(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

