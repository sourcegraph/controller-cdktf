//go:build no_runtime_type_checking

package dbproxy

// Building without runtime type checking enabled, so all the below just return nil

func (d *jsiiProxy_DbProxyAuthList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (d *jsiiProxy_DbProxyAuthList) validateGetParameters(index *float64) error {
	return nil
}

func (d *jsiiProxy_DbProxyAuthList) validateResolveParameters(context cdktn.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_DbProxyAuthList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_DbProxyAuthList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_DbProxyAuthList) validateSetTerraformResourceParameters(val cdktn.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_DbProxyAuthList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewDbProxyAuthListParameters(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

