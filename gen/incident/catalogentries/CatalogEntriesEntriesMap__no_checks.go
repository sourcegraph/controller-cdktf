//go:build no_runtime_type_checking

package catalogentries

// Building without runtime type checking enabled, so all the below just return nil

func (c *jsiiProxy_CatalogEntriesEntriesMap) validateGetParameters(key *string) error {
	return nil
}

func (c *jsiiProxy_CatalogEntriesEntriesMap) validateInterpolationForAttributeParameters(property *string) error {
	return nil
}

func (c *jsiiProxy_CatalogEntriesEntriesMap) validateResolveParameters(context cdktn.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_CatalogEntriesEntriesMap) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_CatalogEntriesEntriesMap) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_CatalogEntriesEntriesMap) validateSetTerraformResourceParameters(val cdktn.IInterpolatingParent) error {
	return nil
}

func validateNewCatalogEntriesEntriesMapParameters(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) error {
	return nil
}

