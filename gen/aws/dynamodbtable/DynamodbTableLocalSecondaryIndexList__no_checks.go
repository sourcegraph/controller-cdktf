//go:build no_runtime_type_checking

package dynamodbtable

// Building without runtime type checking enabled, so all the below just return nil

func (d *jsiiProxy_DynamodbTableLocalSecondaryIndexList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (d *jsiiProxy_DynamodbTableLocalSecondaryIndexList) validateGetParameters(index *float64) error {
	return nil
}

func (d *jsiiProxy_DynamodbTableLocalSecondaryIndexList) validateResolveParameters(context cdktn.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_DynamodbTableLocalSecondaryIndexList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_DynamodbTableLocalSecondaryIndexList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_DynamodbTableLocalSecondaryIndexList) validateSetTerraformResourceParameters(val cdktn.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_DynamodbTableLocalSecondaryIndexList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewDynamodbTableLocalSecondaryIndexListParameters(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

