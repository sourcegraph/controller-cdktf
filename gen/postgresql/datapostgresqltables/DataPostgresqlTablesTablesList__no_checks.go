//go:build no_runtime_type_checking

package datapostgresqltables

// Building without runtime type checking enabled, so all the below just return nil

func (d *jsiiProxy_DataPostgresqlTablesTablesList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (d *jsiiProxy_DataPostgresqlTablesTablesList) validateGetParameters(index *float64) error {
	return nil
}

func (d *jsiiProxy_DataPostgresqlTablesTablesList) validateResolveParameters(context cdktn.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_DataPostgresqlTablesTablesList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_DataPostgresqlTablesTablesList) validateSetTerraformResourceParameters(val cdktn.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_DataPostgresqlTablesTablesList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewDataPostgresqlTablesTablesListParameters(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

