//go:build no_runtime_type_checking

package provider

// Building without runtime type checking enabled, so all the below just return nil

func (o *jsiiProxy_OktaProvider) validateAddOverrideParameters(path *string, value interface{}) error {
	return nil
}

func (o *jsiiProxy_OktaProvider) validateOverrideLogicalIdParameters(newLogicalId *string) error {
	return nil
}

func validateOktaProvider_IsConstructParameters(x interface{}) error {
	return nil
}

func (j *jsiiProxy_OktaProvider) validateSetBackoffParameters(val interface{}) error {
	return nil
}

func validateNewOktaProviderParameters(scope constructs.Construct, id *string, config *OktaProviderConfig) error {
	return nil
}

