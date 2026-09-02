//go:build no_runtime_type_checking

package mskconnectconnector

// Building without runtime type checking enabled, so all the below just return nil

func (m *jsiiProxy_MskconnectConnectorPluginList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (m *jsiiProxy_MskconnectConnectorPluginList) validateGetParameters(index *float64) error {
	return nil
}

func (m *jsiiProxy_MskconnectConnectorPluginList) validateResolveParameters(context cdktn.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_MskconnectConnectorPluginList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_MskconnectConnectorPluginList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_MskconnectConnectorPluginList) validateSetTerraformResourceParameters(val cdktn.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_MskconnectConnectorPluginList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewMskconnectConnectorPluginListParameters(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

