//go:build no_runtime_type_checking

package s3bucket

// Building without runtime type checking enabled, so all the below just return nil

func (s *jsiiProxy_S3BucketLifecycleRuleTransitionList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (s *jsiiProxy_S3BucketLifecycleRuleTransitionList) validateGetParameters(index *float64) error {
	return nil
}

func (s *jsiiProxy_S3BucketLifecycleRuleTransitionList) validateResolveParameters(context cdktn.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_S3BucketLifecycleRuleTransitionList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_S3BucketLifecycleRuleTransitionList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_S3BucketLifecycleRuleTransitionList) validateSetTerraformResourceParameters(val cdktn.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_S3BucketLifecycleRuleTransitionList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewS3BucketLifecycleRuleTransitionListParameters(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

