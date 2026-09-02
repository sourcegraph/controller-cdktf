//go:build no_runtime_type_checking

package cesagent

// Building without runtime type checking enabled, so all the below just return nil

func (c *jsiiProxy_CesAgentToolsetsList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (c *jsiiProxy_CesAgentToolsetsList) validateGetParameters(index *float64) error {
	return nil
}

func (c *jsiiProxy_CesAgentToolsetsList) validateResolveParameters(context cdktn.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_CesAgentToolsetsList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_CesAgentToolsetsList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_CesAgentToolsetsList) validateSetTerraformResourceParameters(val cdktn.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_CesAgentToolsetsList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewCesAgentToolsetsListParameters(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

