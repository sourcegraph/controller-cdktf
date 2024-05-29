//go:build no_runtime_type_checking

package certificatepack

// Building without runtime type checking enabled, so all the below just return nil

func (c *jsiiProxy_CertificatePackValidationRecordsList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (c *jsiiProxy_CertificatePackValidationRecordsList) validateGetParameters(index *float64) error {
	return nil
}

func (c *jsiiProxy_CertificatePackValidationRecordsList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_CertificatePackValidationRecordsList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_CertificatePackValidationRecordsList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_CertificatePackValidationRecordsList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_CertificatePackValidationRecordsList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewCertificatePackValidationRecordsListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

