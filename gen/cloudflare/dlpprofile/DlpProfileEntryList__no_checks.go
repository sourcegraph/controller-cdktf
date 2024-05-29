//go:build no_runtime_type_checking

package dlpprofile

// Building without runtime type checking enabled, so all the below just return nil

func (d *jsiiProxy_DlpProfileEntryList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (d *jsiiProxy_DlpProfileEntryList) validateGetParameters(index *float64) error {
	return nil
}

func (d *jsiiProxy_DlpProfileEntryList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_DlpProfileEntryList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_DlpProfileEntryList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_DlpProfileEntryList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_DlpProfileEntryList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewDlpProfileEntryListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

