//go:build no_runtime_type_checking

package policyrulesignon

// Building without runtime type checking enabled, so all the below just return nil

func (p *jsiiProxy_PolicyRuleSignonFactorSequenceList) validateGetParameters(index *float64) error {
	return nil
}

func (p *jsiiProxy_PolicyRuleSignonFactorSequenceList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_PolicyRuleSignonFactorSequenceList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_PolicyRuleSignonFactorSequenceList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_PolicyRuleSignonFactorSequenceList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_PolicyRuleSignonFactorSequenceList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewPolicyRuleSignonFactorSequenceListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

