//go:build no_runtime_type_checking

package provider

// Building without runtime type checking enabled, so all the below just return nil

func (i *jsiiProxy_IncidentProvider) validateAddOverrideParameters(path *string, value interface{}) error {
	return nil
}

func (i *jsiiProxy_IncidentProvider) validateOverrideLogicalIdParameters(newLogicalId *string) error {
	return nil
}

func validateIncidentProvider_GenerateConfigForImportParameters(scope constructs.Construct, importToId *string, importFromId *string) error {
	return nil
}

func validateIncidentProvider_IsConstructParameters(x interface{}) error {
	return nil
}

func validateIncidentProvider_IsTerraformElementParameters(x interface{}) error {
	return nil
}

func validateIncidentProvider_IsTerraformProviderParameters(x interface{}) error {
	return nil
}

func validateNewIncidentProviderParameters(scope constructs.Construct, id *string, config *IncidentProviderConfig) error {
	return nil
}

