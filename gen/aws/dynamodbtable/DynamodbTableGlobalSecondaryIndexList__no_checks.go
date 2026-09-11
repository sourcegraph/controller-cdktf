//go:build no_runtime_type_checking

package dynamodbtable

// Building without runtime type checking enabled, so all the below just return nil

func (d *jsiiProxy_DynamodbTableGlobalSecondaryIndexList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (d *jsiiProxy_DynamodbTableGlobalSecondaryIndexList) validateGetParameters(index *float64) error {
	return nil
}

func (d *jsiiProxy_DynamodbTableGlobalSecondaryIndexList) validateResolveParameters(context cdktn.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_DynamodbTableGlobalSecondaryIndexList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_DynamodbTableGlobalSecondaryIndexList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_DynamodbTableGlobalSecondaryIndexList) validateSetTerraformResourceParameters(val cdktn.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_DynamodbTableGlobalSecondaryIndexList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewDynamodbTableGlobalSecondaryIndexListParameters(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

