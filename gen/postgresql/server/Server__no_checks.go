//go:build no_runtime_type_checking

package server

// Building without runtime type checking enabled, so all the below just return nil

func (s *jsiiProxy_Server) validateAddMoveTargetParameters(moveTarget *string) error {
	return nil
}

func (s *jsiiProxy_Server) validateAddOverrideParameters(path *string, value interface{}) error {
	return nil
}

func (s *jsiiProxy_Server) validateGetAnyMapAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (s *jsiiProxy_Server) validateGetBooleanAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (s *jsiiProxy_Server) validateGetBooleanMapAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (s *jsiiProxy_Server) validateGetListAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (s *jsiiProxy_Server) validateGetNumberAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (s *jsiiProxy_Server) validateGetNumberListAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (s *jsiiProxy_Server) validateGetNumberMapAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (s *jsiiProxy_Server) validateGetStringAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (s *jsiiProxy_Server) validateGetStringMapAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (s *jsiiProxy_Server) validateImportFromParameters(id *string) error {
	return nil
}

func (s *jsiiProxy_Server) validateInterpolationForAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (s *jsiiProxy_Server) validateMoveFromIdParameters(id *string) error {
	return nil
}

func (s *jsiiProxy_Server) validateMoveToParameters(moveTarget *string, index interface{}) error {
	return nil
}

func (s *jsiiProxy_Server) validateMoveToIdParameters(id *string) error {
	return nil
}

func (s *jsiiProxy_Server) validateOverrideLogicalIdParameters(newLogicalId *string) error {
	return nil
}

func validateServer_GenerateConfigForImportParameters(scope constructs.Construct, importToId *string, importFromId *string) error {
	return nil
}

func validateServer_IsConstructParameters(x interface{}) error {
	return nil
}

func validateServer_IsTerraformElementParameters(x interface{}) error {
	return nil
}

func validateServer_IsTerraformResourceParameters(x interface{}) error {
	return nil
}

func (j *jsiiProxy_Server) validateSetConnectionParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_Server) validateSetCountParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_Server) validateSetDropCascadeParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_Server) validateSetFdwNameParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_Server) validateSetIdParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_Server) validateSetLifecycleParameters(val *cdktf.TerraformResourceLifecycle) error {
	return nil
}

func (j *jsiiProxy_Server) validateSetOptionsParameters(val *map[string]*string) error {
	return nil
}

func (j *jsiiProxy_Server) validateSetProvisionersParameters(val *[]interface{}) error {
	return nil
}

func (j *jsiiProxy_Server) validateSetServerNameParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_Server) validateSetServerOwnerParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_Server) validateSetServerTypeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_Server) validateSetServerVersionParameters(val *string) error {
	return nil
}

func validateNewServerParameters(scope constructs.Construct, id *string, config *ServerConfig) error {
	return nil
}

