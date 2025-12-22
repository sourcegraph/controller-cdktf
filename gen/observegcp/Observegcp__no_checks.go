//go:build no_runtime_type_checking

package observegcp

// Building without runtime type checking enabled, so all the below just return nil

func (o *jsiiProxy_Observegcp) validateAddOverrideParameters(path *string, value interface{}) error {
	return nil
}

func (o *jsiiProxy_Observegcp) validateAddProviderParameters(provider interface{}) error {
	return nil
}

func (o *jsiiProxy_Observegcp) validateGetStringParameters(output *string) error {
	return nil
}

func (o *jsiiProxy_Observegcp) validateInterpolationForOutputParameters(moduleOutput *string) error {
	return nil
}

func (o *jsiiProxy_Observegcp) validateOverrideLogicalIdParameters(newLogicalId *string) error {
	return nil
}

func validateObservegcp_IsConstructParameters(x interface{}) error {
	return nil
}

func validateObservegcp_IsTerraformElementParameters(x interface{}) error {
	return nil
}

func (j *jsiiProxy_Observegcp) validateSetLoggingExclusionsParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_Observegcp) validateSetResourceParameters(val *string) error {
	return nil
}

func validateNewObservegcpParameters(scope constructs.Construct, id *string, config *ObservegcpConfig) error {
	return nil
}

