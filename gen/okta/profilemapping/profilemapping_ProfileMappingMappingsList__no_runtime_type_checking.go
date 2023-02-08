//go:build no_runtime_type_checking

package profilemapping

// Building without runtime type checking enabled, so all the below just return nil

func (p *jsiiProxy_ProfileMappingMappingsList) validateGetParameters(index *float64) error {
	return nil
}

func (p *jsiiProxy_ProfileMappingMappingsList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_ProfileMappingMappingsList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_ProfileMappingMappingsList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_ProfileMappingMappingsList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_ProfileMappingMappingsList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewProfileMappingMappingsListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

