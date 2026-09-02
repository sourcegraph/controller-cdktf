//go:build no_runtime_type_checking

package dataawsdbproxy

// Building without runtime type checking enabled, so all the below just return nil

func (d *jsiiProxy_DataAwsDbProxyAuthList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (d *jsiiProxy_DataAwsDbProxyAuthList) validateGetParameters(index *float64) error {
	return nil
}

func (d *jsiiProxy_DataAwsDbProxyAuthList) validateResolveParameters(context cdktn.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_DataAwsDbProxyAuthList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_DataAwsDbProxyAuthList) validateSetTerraformResourceParameters(val cdktn.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_DataAwsDbProxyAuthList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewDataAwsDbProxyAuthListParameters(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

