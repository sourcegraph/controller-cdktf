//go:build no_runtime_type_checking

package ruleset

// Building without runtime type checking enabled, so all the below just return nil

func (r *jsiiProxy_RulesetRulesActionParametersList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (r *jsiiProxy_RulesetRulesActionParametersList) validateGetParameters(index *float64) error {
	return nil
}

func (r *jsiiProxy_RulesetRulesActionParametersList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_RulesetRulesActionParametersList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_RulesetRulesActionParametersList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_RulesetRulesActionParametersList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_RulesetRulesActionParametersList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewRulesetRulesActionParametersListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

