//go:build no_runtime_type_checking

package awsvpc

// Building without runtime type checking enabled, so all the below just return nil

func (a *jsiiProxy_Awsvpc) validateAddOverrideParameters(path *string, value interface{}) error {
	return nil
}

func (a *jsiiProxy_Awsvpc) validateAddProviderParameters(provider interface{}) error {
	return nil
}

func (a *jsiiProxy_Awsvpc) validateGetStringParameters(output *string) error {
	return nil
}

func (a *jsiiProxy_Awsvpc) validateInterpolationForOutputParameters(moduleOutput *string) error {
	return nil
}

func (a *jsiiProxy_Awsvpc) validateOverrideLogicalIdParameters(newLogicalId *string) error {
	return nil
}

func validateAwsvpc_IsConstructParameters(x interface{}) error {
	return nil
}

func validateAwsvpc_IsTerraformElementParameters(x interface{}) error {
	return nil
}

func validateNewAwsvpcParameters(scope constructs.Construct, id *string, config *AwsvpcConfig) error {
	return nil
}

