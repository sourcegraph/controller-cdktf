//go:build no_runtime_type_checking

package provider

// Building without runtime type checking enabled, so all the below just return nil

func (t *jsiiProxy_TlsProvider) validateAddOverrideParameters(path *string, value interface{}) error {
	return nil
}

func (t *jsiiProxy_TlsProvider) validateOverrideLogicalIdParameters(newLogicalId *string) error {
	return nil
}

func validateTlsProvider_IsConstructParameters(x interface{}) error {
	return nil
}

func validateNewTlsProviderParameters(scope constructs.Construct, id *string, config *TlsProviderConfig) error {
	return nil
}

