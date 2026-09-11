//go:build no_runtime_type_checking

package s3bucket

// Building without runtime type checking enabled, so all the below just return nil

func (s *jsiiProxy_S3BucketGrantList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (s *jsiiProxy_S3BucketGrantList) validateGetParameters(index *float64) error {
	return nil
}

func (s *jsiiProxy_S3BucketGrantList) validateResolveParameters(context cdktn.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_S3BucketGrantList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_S3BucketGrantList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_S3BucketGrantList) validateSetTerraformResourceParameters(val cdktn.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_S3BucketGrantList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewS3BucketGrantListParameters(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

