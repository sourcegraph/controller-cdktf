//go:build no_runtime_type_checking

package workbenchinstance

// Building without runtime type checking enabled, so all the below just return nil

func (w *jsiiProxy_WorkbenchInstanceUpgradeHistoryList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (w *jsiiProxy_WorkbenchInstanceUpgradeHistoryList) validateGetParameters(index *float64) error {
	return nil
}

func (w *jsiiProxy_WorkbenchInstanceUpgradeHistoryList) validateResolveParameters(context cdktn.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_WorkbenchInstanceUpgradeHistoryList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_WorkbenchInstanceUpgradeHistoryList) validateSetTerraformResourceParameters(val cdktn.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_WorkbenchInstanceUpgradeHistoryList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewWorkbenchInstanceUpgradeHistoryListParameters(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

