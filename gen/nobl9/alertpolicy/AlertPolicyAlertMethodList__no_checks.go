//go:build no_runtime_type_checking

package alertpolicy

// Building without runtime type checking enabled, so all the below just return nil

func (a *jsiiProxy_AlertPolicyAlertMethodList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (a *jsiiProxy_AlertPolicyAlertMethodList) validateGetParameters(index *float64) error {
	return nil
}

func (a *jsiiProxy_AlertPolicyAlertMethodList) validateResolveParameters(context cdktn.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_AlertPolicyAlertMethodList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_AlertPolicyAlertMethodList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_AlertPolicyAlertMethodList) validateSetTerraformResourceParameters(val cdktn.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_AlertPolicyAlertMethodList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewAlertPolicyAlertMethodListParameters(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

