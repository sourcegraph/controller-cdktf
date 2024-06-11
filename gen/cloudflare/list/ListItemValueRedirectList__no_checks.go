//go:build no_runtime_type_checking

package list

// Building without runtime type checking enabled, so all the below just return nil

func (l *jsiiProxy_ListItemValueRedirectList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (l *jsiiProxy_ListItemValueRedirectList) validateGetParameters(index *float64) error {
	return nil
}

func (l *jsiiProxy_ListItemValueRedirectList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_ListItemValueRedirectList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_ListItemValueRedirectList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_ListItemValueRedirectList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_ListItemValueRedirectList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewListItemValueRedirectListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

