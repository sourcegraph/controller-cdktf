//go:build no_runtime_type_checking

package domain

// Building without runtime type checking enabled, so all the below just return nil

func (d *jsiiProxy_DomainDnsRecordsList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (d *jsiiProxy_DomainDnsRecordsList) validateGetParameters(index *float64) error {
	return nil
}

func (d *jsiiProxy_DomainDnsRecordsList) validateResolveParameters(context cdktn.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_DomainDnsRecordsList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_DomainDnsRecordsList) validateSetTerraformResourceParameters(val cdktn.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_DomainDnsRecordsList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewDomainDnsRecordsListParameters(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

