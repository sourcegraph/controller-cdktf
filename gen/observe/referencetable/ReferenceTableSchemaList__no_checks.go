//go:build no_runtime_type_checking

package referencetable

// Building without runtime type checking enabled, so all the below just return nil

func (r *jsiiProxy_ReferenceTableSchemaList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (r *jsiiProxy_ReferenceTableSchemaList) validateGetParameters(index *float64) error {
	return nil
}

func (r *jsiiProxy_ReferenceTableSchemaList) validateResolveParameters(context cdktn.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_ReferenceTableSchemaList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_ReferenceTableSchemaList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_ReferenceTableSchemaList) validateSetTerraformResourceParameters(val cdktn.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_ReferenceTableSchemaList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewReferenceTableSchemaListParameters(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

