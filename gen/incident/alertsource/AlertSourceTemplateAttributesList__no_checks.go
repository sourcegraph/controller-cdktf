//go:build no_runtime_type_checking

package alertsource

// Building without runtime type checking enabled, so all the below just return nil

func (a *jsiiProxy_AlertSourceTemplateAttributesList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (a *jsiiProxy_AlertSourceTemplateAttributesList) validateGetParameters(index *float64) error {
	return nil
}

func (a *jsiiProxy_AlertSourceTemplateAttributesList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_AlertSourceTemplateAttributesList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_AlertSourceTemplateAttributesList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_AlertSourceTemplateAttributesList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_AlertSourceTemplateAttributesList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewAlertSourceTemplateAttributesListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

