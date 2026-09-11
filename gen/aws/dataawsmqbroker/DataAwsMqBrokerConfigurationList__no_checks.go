//go:build no_runtime_type_checking

package dataawsmqbroker

// Building without runtime type checking enabled, so all the below just return nil

func (d *jsiiProxy_DataAwsMqBrokerConfigurationList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (d *jsiiProxy_DataAwsMqBrokerConfigurationList) validateGetParameters(index *float64) error {
	return nil
}

func (d *jsiiProxy_DataAwsMqBrokerConfigurationList) validateResolveParameters(context cdktn.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_DataAwsMqBrokerConfigurationList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_DataAwsMqBrokerConfigurationList) validateSetTerraformResourceParameters(val cdktn.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_DataAwsMqBrokerConfigurationList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewDataAwsMqBrokerConfigurationListParameters(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

