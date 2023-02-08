//go:build no_runtime_type_checking

package userschema

// Building without runtime type checking enabled, so all the below just return nil

func (u *jsiiProxy_UserSchemaArrayOneOfList) validateGetParameters(index *float64) error {
	return nil
}

func (u *jsiiProxy_UserSchemaArrayOneOfList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_UserSchemaArrayOneOfList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_UserSchemaArrayOneOfList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_UserSchemaArrayOneOfList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_UserSchemaArrayOneOfList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewUserSchemaArrayOneOfListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

