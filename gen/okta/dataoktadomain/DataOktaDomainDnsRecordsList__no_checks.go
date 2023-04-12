//go:build no_runtime_type_checking

package dataoktadomain

// Building without runtime type checking enabled, so all the below just return nil

func (d *jsiiProxy_DataOktaDomainDnsRecordsList) validateGetParameters(index *float64) error {
	return nil
}

func (d *jsiiProxy_DataOktaDomainDnsRecordsList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_DataOktaDomainDnsRecordsList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_DataOktaDomainDnsRecordsList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_DataOktaDomainDnsRecordsList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewDataOktaDomainDnsRecordsListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

