//go:build no_runtime_type_checking

package list

// Building without runtime type checking enabled, so all the below just return nil

func (l *jsiiProxy_ListItemList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (l *jsiiProxy_ListItemList) validateGetParameters(index *float64) error {
	return nil
}

func (l *jsiiProxy_ListItemList) validateResolveParameters(context cdktn.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_ListItemList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_ListItemList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_ListItemList) validateSetTerraformResourceParameters(val cdktn.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_ListItemList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewListItemListParameters(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

