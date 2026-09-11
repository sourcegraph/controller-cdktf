//go:build no_runtime_type_checking

package codepipeline

// Building without runtime type checking enabled, so all the below just return nil

func (c *jsiiProxy_CodepipelineStageList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (c *jsiiProxy_CodepipelineStageList) validateGetParameters(index *float64) error {
	return nil
}

func (c *jsiiProxy_CodepipelineStageList) validateResolveParameters(context cdktn.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_CodepipelineStageList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_CodepipelineStageList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_CodepipelineStageList) validateSetTerraformResourceParameters(val cdktn.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_CodepipelineStageList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewCodepipelineStageListParameters(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

