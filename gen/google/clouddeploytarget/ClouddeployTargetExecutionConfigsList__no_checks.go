//go:build no_runtime_type_checking

package clouddeploytarget

// Building without runtime type checking enabled, so all the below just return nil

func (c *jsiiProxy_ClouddeployTargetExecutionConfigsList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (c *jsiiProxy_ClouddeployTargetExecutionConfigsList) validateGetParameters(index *float64) error {
	return nil
}

func (c *jsiiProxy_ClouddeployTargetExecutionConfigsList) validateResolveParameters(context cdktn.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_ClouddeployTargetExecutionConfigsList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_ClouddeployTargetExecutionConfigsList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_ClouddeployTargetExecutionConfigsList) validateSetTerraformResourceParameters(val cdktn.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_ClouddeployTargetExecutionConfigsList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewClouddeployTargetExecutionConfigsListParameters(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

