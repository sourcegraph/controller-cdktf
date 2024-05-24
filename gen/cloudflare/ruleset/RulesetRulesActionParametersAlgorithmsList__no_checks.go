//go:build no_runtime_type_checking

package ruleset

// Building without runtime type checking enabled, so all the below just return nil

func (r *jsiiProxy_RulesetRulesActionParametersAlgorithmsList) validateGetParameters(index *float64) error {
	return nil
}

func (r *jsiiProxy_RulesetRulesActionParametersAlgorithmsList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_RulesetRulesActionParametersAlgorithmsList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_RulesetRulesActionParametersAlgorithmsList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_RulesetRulesActionParametersAlgorithmsList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_RulesetRulesActionParametersAlgorithmsList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewRulesetRulesActionParametersAlgorithmsListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

