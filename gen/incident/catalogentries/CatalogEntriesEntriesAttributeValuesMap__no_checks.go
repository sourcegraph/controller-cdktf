//go:build no_runtime_type_checking

package catalogentries

// Building without runtime type checking enabled, so all the below just return nil

func (c *jsiiProxy_CatalogEntriesEntriesAttributeValuesMap) validateGetParameters(key *string) error {
	return nil
}

func (c *jsiiProxy_CatalogEntriesEntriesAttributeValuesMap) validateInterpolationForAttributeParameters(property *string) error {
	return nil
}

func (c *jsiiProxy_CatalogEntriesEntriesAttributeValuesMap) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_CatalogEntriesEntriesAttributeValuesMap) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_CatalogEntriesEntriesAttributeValuesMap) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_CatalogEntriesEntriesAttributeValuesMap) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func validateNewCatalogEntriesEntriesAttributeValuesMapParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) error {
	return nil
}

