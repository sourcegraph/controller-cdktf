//go:build no_runtime_type_checking

package certificatepack

// Building without runtime type checking enabled, so all the below just return nil

func (c *jsiiProxy_CertificatePackValidationErrorsList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (c *jsiiProxy_CertificatePackValidationErrorsList) validateGetParameters(index *float64) error {
	return nil
}

func (c *jsiiProxy_CertificatePackValidationErrorsList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_CertificatePackValidationErrorsList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_CertificatePackValidationErrorsList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_CertificatePackValidationErrorsList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_CertificatePackValidationErrorsList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewCertificatePackValidationErrorsListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

