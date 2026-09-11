//go:build no_runtime_type_checking

package backupplan

// Building without runtime type checking enabled, so all the below just return nil

func (b *jsiiProxy_BackupPlanRuleCopyActionList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (b *jsiiProxy_BackupPlanRuleCopyActionList) validateGetParameters(index *float64) error {
	return nil
}

func (b *jsiiProxy_BackupPlanRuleCopyActionList) validateResolveParameters(context cdktn.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_BackupPlanRuleCopyActionList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_BackupPlanRuleCopyActionList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_BackupPlanRuleCopyActionList) validateSetTerraformResourceParameters(val cdktn.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_BackupPlanRuleCopyActionList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewBackupPlanRuleCopyActionListParameters(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

