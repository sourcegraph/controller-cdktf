//go:build no_runtime_type_checking

// executors
package executors

// Building without runtime type checking enabled, so all the below just return nil

func (e *jsiiProxy_Executors) validateAddOverrideParameters(path *string, value interface{}) error {
	return nil
}

func (e *jsiiProxy_Executors) validateAddProviderParameters(provider interface{}) error {
	return nil
}

func (e *jsiiProxy_Executors) validateGetStringParameters(output *string) error {
	return nil
}

func (e *jsiiProxy_Executors) validateInterpolationForOutputParameters(moduleOutput *string) error {
	return nil
}

func (e *jsiiProxy_Executors) validateOverrideLogicalIdParameters(newLogicalId *string) error {
	return nil
}

func validateExecutors_IsConstructParameters(x interface{}) error {
	return nil
}

func (j *jsiiProxy_Executors) validateSetInstanceTagParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_Executors) validateSetMetricsEnvironmentLabelParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_Executors) validateSetNetworkIdParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_Executors) validateSetQueueNameParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_Executors) validateSetSourcegraphExecutorProxyPasswordParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_Executors) validateSetSourcegraphExternalUrlParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_Executors) validateSetSubnetIdParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_Executors) validateSetZoneParameters(val *string) error {
	return nil
}

func validateNewExecutorsParameters(scope constructs.Construct, id *string, options *ExecutorsOptions) error {
	return nil
}

