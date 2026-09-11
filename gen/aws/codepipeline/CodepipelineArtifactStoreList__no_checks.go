//go:build no_runtime_type_checking

package codepipeline

// Building without runtime type checking enabled, so all the below just return nil

func (c *jsiiProxy_CodepipelineArtifactStoreList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (c *jsiiProxy_CodepipelineArtifactStoreList) validateGetParameters(index *float64) error {
	return nil
}

func (c *jsiiProxy_CodepipelineArtifactStoreList) validateResolveParameters(context cdktn.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_CodepipelineArtifactStoreList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_CodepipelineArtifactStoreList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_CodepipelineArtifactStoreList) validateSetTerraformResourceParameters(val cdktn.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_CodepipelineArtifactStoreList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewCodepipelineArtifactStoreListParameters(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

