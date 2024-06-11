//go:build no_runtime_type_checking

package ruleset

// Building without runtime type checking enabled, so all the below just return nil

func (r *jsiiProxy_RulesetRulesList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (r *jsiiProxy_RulesetRulesList) validateGetParameters(index *float64) error {
	return nil
}

func (r *jsiiProxy_RulesetRulesList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_RulesetRulesList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_RulesetRulesList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_RulesetRulesList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_RulesetRulesList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewRulesetRulesListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

