//go:build no_runtime_type_checking

package appstreamstack

// Building without runtime type checking enabled, so all the below just return nil

func (a *jsiiProxy_AppstreamStackUserSettingsList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (a *jsiiProxy_AppstreamStackUserSettingsList) validateGetParameters(index *float64) error {
	return nil
}

func (a *jsiiProxy_AppstreamStackUserSettingsList) validateResolveParameters(context cdktn.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_AppstreamStackUserSettingsList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_AppstreamStackUserSettingsList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_AppstreamStackUserSettingsList) validateSetTerraformResourceParameters(val cdktn.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_AppstreamStackUserSettingsList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewAppstreamStackUserSettingsListParameters(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

