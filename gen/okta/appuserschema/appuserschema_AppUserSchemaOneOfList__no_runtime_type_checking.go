//go:build no_runtime_type_checking

package appuserschema

// Building without runtime type checking enabled, so all the below just return nil

func (a *jsiiProxy_AppUserSchemaOneOfList) validateGetParameters(index *float64) error {
	return nil
}

func (a *jsiiProxy_AppUserSchemaOneOfList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_AppUserSchemaOneOfList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_AppUserSchemaOneOfList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_AppUserSchemaOneOfList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_AppUserSchemaOneOfList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewAppUserSchemaOneOfListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

