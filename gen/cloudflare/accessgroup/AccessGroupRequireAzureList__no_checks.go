//go:build no_runtime_type_checking

package accessgroup

// Building without runtime type checking enabled, so all the below just return nil

func (a *jsiiProxy_AccessGroupRequireAzureList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (a *jsiiProxy_AccessGroupRequireAzureList) validateGetParameters(index *float64) error {
	return nil
}

func (a *jsiiProxy_AccessGroupRequireAzureList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_AccessGroupRequireAzureList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_AccessGroupRequireAzureList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_AccessGroupRequireAzureList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_AccessGroupRequireAzureList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewAccessGroupRequireAzureListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

