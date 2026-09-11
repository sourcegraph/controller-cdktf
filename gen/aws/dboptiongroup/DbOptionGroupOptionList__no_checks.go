//go:build no_runtime_type_checking

package dboptiongroup

// Building without runtime type checking enabled, so all the below just return nil

func (d *jsiiProxy_DbOptionGroupOptionList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (d *jsiiProxy_DbOptionGroupOptionList) validateGetParameters(index *float64) error {
	return nil
}

func (d *jsiiProxy_DbOptionGroupOptionList) validateResolveParameters(context cdktn.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_DbOptionGroupOptionList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_DbOptionGroupOptionList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_DbOptionGroupOptionList) validateSetTerraformResourceParameters(val cdktn.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_DbOptionGroupOptionList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewDbOptionGroupOptionListParameters(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

