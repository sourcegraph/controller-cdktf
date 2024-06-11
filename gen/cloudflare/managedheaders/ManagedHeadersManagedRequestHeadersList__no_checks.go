//go:build no_runtime_type_checking

package managedheaders

// Building without runtime type checking enabled, so all the below just return nil

func (m *jsiiProxy_ManagedHeadersManagedRequestHeadersList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (m *jsiiProxy_ManagedHeadersManagedRequestHeadersList) validateGetParameters(index *float64) error {
	return nil
}

func (m *jsiiProxy_ManagedHeadersManagedRequestHeadersList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_ManagedHeadersManagedRequestHeadersList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_ManagedHeadersManagedRequestHeadersList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_ManagedHeadersManagedRequestHeadersList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_ManagedHeadersManagedRequestHeadersList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewManagedHeadersManagedRequestHeadersListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

