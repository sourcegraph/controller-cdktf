//go:build no_runtime_type_checking

package accessgroup

// Building without runtime type checking enabled, so all the below just return nil

func (a *jsiiProxy_AccessGroupRequireSamlList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (a *jsiiProxy_AccessGroupRequireSamlList) validateGetParameters(index *float64) error {
	return nil
}

func (a *jsiiProxy_AccessGroupRequireSamlList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_AccessGroupRequireSamlList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_AccessGroupRequireSamlList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_AccessGroupRequireSamlList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_AccessGroupRequireSamlList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewAccessGroupRequireSamlListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

