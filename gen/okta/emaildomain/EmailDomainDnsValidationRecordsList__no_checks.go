//go:build no_runtime_type_checking

package emaildomain

// Building without runtime type checking enabled, so all the below just return nil

func (e *jsiiProxy_EmailDomainDnsValidationRecordsList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (e *jsiiProxy_EmailDomainDnsValidationRecordsList) validateGetParameters(index *float64) error {
	return nil
}

func (e *jsiiProxy_EmailDomainDnsValidationRecordsList) validateResolveParameters(context cdktn.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_EmailDomainDnsValidationRecordsList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_EmailDomainDnsValidationRecordsList) validateSetTerraformResourceParameters(val cdktn.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_EmailDomainDnsValidationRecordsList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewEmailDomainDnsValidationRecordsListParameters(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

