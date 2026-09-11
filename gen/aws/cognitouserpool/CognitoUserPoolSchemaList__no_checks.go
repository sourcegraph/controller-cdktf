//go:build no_runtime_type_checking

package cognitouserpool

// Building without runtime type checking enabled, so all the below just return nil

func (c *jsiiProxy_CognitoUserPoolSchemaList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (c *jsiiProxy_CognitoUserPoolSchemaList) validateGetParameters(index *float64) error {
	return nil
}

func (c *jsiiProxy_CognitoUserPoolSchemaList) validateResolveParameters(context cdktn.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_CognitoUserPoolSchemaList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_CognitoUserPoolSchemaList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_CognitoUserPoolSchemaList) validateSetTerraformResourceParameters(val cdktn.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_CognitoUserPoolSchemaList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewCognitoUserPoolSchemaListParameters(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

