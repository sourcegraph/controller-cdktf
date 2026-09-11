//go:build no_runtime_type_checking

package project

// Building without runtime type checking enabled, so all the below just return nil

func (p *jsiiProxy_ProjectLabelList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (p *jsiiProxy_ProjectLabelList) validateGetParameters(index *float64) error {
	return nil
}

func (p *jsiiProxy_ProjectLabelList) validateResolveParameters(context cdktn.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_ProjectLabelList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_ProjectLabelList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_ProjectLabelList) validateSetTerraformResourceParameters(val cdktn.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_ProjectLabelList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewProjectLabelListParameters(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

