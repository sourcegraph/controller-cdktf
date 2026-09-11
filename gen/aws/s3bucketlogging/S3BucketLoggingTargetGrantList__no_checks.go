//go:build no_runtime_type_checking

package s3bucketlogging

// Building without runtime type checking enabled, so all the below just return nil

func (s *jsiiProxy_S3BucketLoggingTargetGrantList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (s *jsiiProxy_S3BucketLoggingTargetGrantList) validateGetParameters(index *float64) error {
	return nil
}

func (s *jsiiProxy_S3BucketLoggingTargetGrantList) validateResolveParameters(context cdktn.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_S3BucketLoggingTargetGrantList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_S3BucketLoggingTargetGrantList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_S3BucketLoggingTargetGrantList) validateSetTerraformResourceParameters(val cdktn.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_S3BucketLoggingTargetGrantList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewS3BucketLoggingTargetGrantListParameters(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

