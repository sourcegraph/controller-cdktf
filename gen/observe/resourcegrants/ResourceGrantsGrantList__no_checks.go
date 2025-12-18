//go:build no_runtime_type_checking

package resourcegrants

// Building without runtime type checking enabled, so all the below just return nil

func (r *jsiiProxy_ResourceGrantsGrantList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (r *jsiiProxy_ResourceGrantsGrantList) validateGetParameters(index *float64) error {
	return nil
}

func (r *jsiiProxy_ResourceGrantsGrantList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_ResourceGrantsGrantList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_ResourceGrantsGrantList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_ResourceGrantsGrantList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_ResourceGrantsGrantList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewResourceGrantsGrantListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

