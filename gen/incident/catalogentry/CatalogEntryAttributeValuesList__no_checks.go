//go:build no_runtime_type_checking

package catalogentry

// Building without runtime type checking enabled, so all the below just return nil

func (c *jsiiProxy_CatalogEntryAttributeValuesList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (c *jsiiProxy_CatalogEntryAttributeValuesList) validateGetParameters(index *float64) error {
	return nil
}

func (c *jsiiProxy_CatalogEntryAttributeValuesList) validateResolveParameters(context cdktn.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_CatalogEntryAttributeValuesList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_CatalogEntryAttributeValuesList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_CatalogEntryAttributeValuesList) validateSetTerraformResourceParameters(val cdktn.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_CatalogEntryAttributeValuesList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewCatalogEntryAttributeValuesListParameters(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

