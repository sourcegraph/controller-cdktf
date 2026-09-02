//go:build no_runtime_type_checking

package provider

// Building without runtime type checking enabled, so all the below just return nil

func (o *jsiiProxy_ObserveProvider) validateAddOverrideParameters(path *string, value interface{}) error {
	return nil
}

func (o *jsiiProxy_ObserveProvider) validateOverrideLogicalIdParameters(newLogicalId *string) error {
	return nil
}

func (o *jsiiProxy_ObserveProvider) validateRegisterProviderFeatureUsageParameters(feature cdktn.ProviderFeature) error {
	return nil
}

func validateObserveProvider_GenerateConfigForImportParameters(scope constructs.Construct, importToId *string, importFromId *string) error {
	return nil
}

func validateObserveProvider_IsConstructParameters(x interface{}) error {
	return nil
}

func validateObserveProvider_IsTerraformElementParameters(x interface{}) error {
	return nil
}

func validateObserveProvider_IsTerraformProviderParameters(x interface{}) error {
	return nil
}

func (j *jsiiProxy_ObserveProvider) validateSetExportObjectBindingsParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_ObserveProvider) validateSetInsecureParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_ObserveProvider) validateSetSkipDatasetDryRunsParameters(val interface{}) error {
	return nil
}

func validateNewObserveProviderParameters(scope constructs.Construct, id *string, config *ObserveProviderConfig) error {
	return nil
}

