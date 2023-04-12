//go:build no_runtime_type_checking

package resourceset

// Building without runtime type checking enabled, so all the below just return nil

func (r *jsiiProxy_ResourceSet) validateAddOverrideParameters(path *string, value interface{}) error {
	return nil
}

func (r *jsiiProxy_ResourceSet) validateGetAnyMapAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (r *jsiiProxy_ResourceSet) validateGetBooleanAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (r *jsiiProxy_ResourceSet) validateGetBooleanMapAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (r *jsiiProxy_ResourceSet) validateGetListAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (r *jsiiProxy_ResourceSet) validateGetNumberAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (r *jsiiProxy_ResourceSet) validateGetNumberListAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (r *jsiiProxy_ResourceSet) validateGetNumberMapAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (r *jsiiProxy_ResourceSet) validateGetStringAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (r *jsiiProxy_ResourceSet) validateGetStringMapAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (r *jsiiProxy_ResourceSet) validateInterpolationForAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (r *jsiiProxy_ResourceSet) validateOverrideLogicalIdParameters(newLogicalId *string) error {
	return nil
}

func validateResourceSet_IsConstructParameters(x interface{}) error {
	return nil
}

func validateResourceSet_IsTerraformElementParameters(x interface{}) error {
	return nil
}

func validateResourceSet_IsTerraformResourceParameters(x interface{}) error {
	return nil
}

func (j *jsiiProxy_ResourceSet) validateSetConnectionParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_ResourceSet) validateSetDescriptionParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_ResourceSet) validateSetIdParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_ResourceSet) validateSetLabelParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_ResourceSet) validateSetLifecycleParameters(val *cdktf.TerraformResourceLifecycle) error {
	return nil
}

func (j *jsiiProxy_ResourceSet) validateSetProvisionersParameters(val *[]interface{}) error {
	return nil
}

func (j *jsiiProxy_ResourceSet) validateSetResourcesParameters(val *[]*string) error {
	return nil
}

func validateNewResourceSetParameters(scope constructs.Construct, id *string, config *ResourceSetConfig) error {
	return nil
}

