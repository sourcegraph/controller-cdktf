//go:build no_runtime_type_checking

package ruleset

// Building without runtime type checking enabled, so all the below just return nil

func (r *jsiiProxy_RulesetRulesActionParametersFromValueList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (r *jsiiProxy_RulesetRulesActionParametersFromValueList) validateGetParameters(index *float64) error {
	return nil
}

func (r *jsiiProxy_RulesetRulesActionParametersFromValueList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_RulesetRulesActionParametersFromValueList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_RulesetRulesActionParametersFromValueList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_RulesetRulesActionParametersFromValueList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_RulesetRulesActionParametersFromValueList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewRulesetRulesActionParametersFromValueListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

