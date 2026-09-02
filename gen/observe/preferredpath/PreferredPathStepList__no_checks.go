//go:build no_runtime_type_checking

package preferredpath

// Building without runtime type checking enabled, so all the below just return nil

func (p *jsiiProxy_PreferredPathStepList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (p *jsiiProxy_PreferredPathStepList) validateGetParameters(index *float64) error {
	return nil
}

func (p *jsiiProxy_PreferredPathStepList) validateResolveParameters(context cdktn.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_PreferredPathStepList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_PreferredPathStepList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_PreferredPathStepList) validateSetTerraformResourceParameters(val cdktn.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_PreferredPathStepList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewPreferredPathStepListParameters(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

