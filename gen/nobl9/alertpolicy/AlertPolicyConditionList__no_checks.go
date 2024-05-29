//go:build no_runtime_type_checking

package alertpolicy

// Building without runtime type checking enabled, so all the below just return nil

func (a *jsiiProxy_AlertPolicyConditionList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (a *jsiiProxy_AlertPolicyConditionList) validateGetParameters(index *float64) error {
	return nil
}

func (a *jsiiProxy_AlertPolicyConditionList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_AlertPolicyConditionList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_AlertPolicyConditionList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_AlertPolicyConditionList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_AlertPolicyConditionList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewAlertPolicyConditionListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

