//go:build no_runtime_type_checking

package limitrange

// Building without runtime type checking enabled, so all the below just return nil

func (l *jsiiProxy_LimitRangeSpecLimitList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (l *jsiiProxy_LimitRangeSpecLimitList) validateGetParameters(index *float64) error {
	return nil
}

func (l *jsiiProxy_LimitRangeSpecLimitList) validateResolveParameters(context cdktn.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_LimitRangeSpecLimitList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_LimitRangeSpecLimitList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_LimitRangeSpecLimitList) validateSetTerraformResourceParameters(val cdktn.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_LimitRangeSpecLimitList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewLimitRangeSpecLimitListParameters(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

