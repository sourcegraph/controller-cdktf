//go:build no_runtime_type_checking

package policyrulemfa

// Building without runtime type checking enabled, so all the below just return nil

func (p *jsiiProxy_PolicyRuleMfaAppExcludeList) validateGetParameters(index *float64) error {
	return nil
}

func (p *jsiiProxy_PolicyRuleMfaAppExcludeList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_PolicyRuleMfaAppExcludeList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_PolicyRuleMfaAppExcludeList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_PolicyRuleMfaAppExcludeList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_PolicyRuleMfaAppExcludeList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewPolicyRuleMfaAppExcludeListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

