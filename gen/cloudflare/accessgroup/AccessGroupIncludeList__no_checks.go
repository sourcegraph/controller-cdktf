//go:build no_runtime_type_checking

package accessgroup

// Building without runtime type checking enabled, so all the below just return nil

func (a *jsiiProxy_AccessGroupIncludeList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (a *jsiiProxy_AccessGroupIncludeList) validateGetParameters(index *float64) error {
	return nil
}

func (a *jsiiProxy_AccessGroupIncludeList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_AccessGroupIncludeList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_AccessGroupIncludeList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_AccessGroupIncludeList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_AccessGroupIncludeList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewAccessGroupIncludeListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

