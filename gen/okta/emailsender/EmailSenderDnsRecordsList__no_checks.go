//go:build no_runtime_type_checking

package emailsender

// Building without runtime type checking enabled, so all the below just return nil

func (e *jsiiProxy_EmailSenderDnsRecordsList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (e *jsiiProxy_EmailSenderDnsRecordsList) validateGetParameters(index *float64) error {
	return nil
}

func (e *jsiiProxy_EmailSenderDnsRecordsList) validateResolveParameters(context cdktn.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_EmailSenderDnsRecordsList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_EmailSenderDnsRecordsList) validateSetTerraformResourceParameters(val cdktn.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_EmailSenderDnsRecordsList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewEmailSenderDnsRecordsListParameters(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

