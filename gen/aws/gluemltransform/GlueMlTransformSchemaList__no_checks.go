//go:build no_runtime_type_checking

package gluemltransform

// Building without runtime type checking enabled, so all the below just return nil

func (g *jsiiProxy_GlueMlTransformSchemaList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (g *jsiiProxy_GlueMlTransformSchemaList) validateGetParameters(index *float64) error {
	return nil
}

func (g *jsiiProxy_GlueMlTransformSchemaList) validateResolveParameters(context cdktn.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_GlueMlTransformSchemaList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_GlueMlTransformSchemaList) validateSetTerraformResourceParameters(val cdktn.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_GlueMlTransformSchemaList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewGlueMlTransformSchemaListParameters(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

