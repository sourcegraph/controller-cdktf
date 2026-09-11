//go:build no_runtime_type_checking

package lexslottype

// Building without runtime type checking enabled, so all the below just return nil

func (l *jsiiProxy_LexSlotTypeEnumerationValueList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (l *jsiiProxy_LexSlotTypeEnumerationValueList) validateGetParameters(index *float64) error {
	return nil
}

func (l *jsiiProxy_LexSlotTypeEnumerationValueList) validateResolveParameters(context cdktn.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_LexSlotTypeEnumerationValueList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_LexSlotTypeEnumerationValueList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_LexSlotTypeEnumerationValueList) validateSetTerraformResourceParameters(val cdktn.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_LexSlotTypeEnumerationValueList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewLexSlotTypeEnumerationValueListParameters(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

