//go:build no_runtime_type_checking

package alertsource

// Building without runtime type checking enabled, so all the below just return nil

func (a *jsiiProxy_AlertSourceTemplateExpressionsList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (a *jsiiProxy_AlertSourceTemplateExpressionsList) validateGetParameters(index *float64) error {
	return nil
}

func (a *jsiiProxy_AlertSourceTemplateExpressionsList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_AlertSourceTemplateExpressionsList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_AlertSourceTemplateExpressionsList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_AlertSourceTemplateExpressionsList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_AlertSourceTemplateExpressionsList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewAlertSourceTemplateExpressionsListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

