//go:build no_runtime_type_checking

package s3bucketacl

// Building without runtime type checking enabled, so all the below just return nil

func (s *jsiiProxy_S3BucketAclAccessControlPolicyGrantList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (s *jsiiProxy_S3BucketAclAccessControlPolicyGrantList) validateGetParameters(index *float64) error {
	return nil
}

func (s *jsiiProxy_S3BucketAclAccessControlPolicyGrantList) validateResolveParameters(context cdktn.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_S3BucketAclAccessControlPolicyGrantList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_S3BucketAclAccessControlPolicyGrantList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_S3BucketAclAccessControlPolicyGrantList) validateSetTerraformResourceParameters(val cdktn.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_S3BucketAclAccessControlPolicyGrantList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewS3BucketAclAccessControlPolicyGrantListParameters(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

