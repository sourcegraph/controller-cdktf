//go:build no_runtime_type_checking

package deployment

// Building without runtime type checking enabled, so all the below just return nil

func (d *jsiiProxy_DeploymentSpecSelectorMatchExpressionsList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (d *jsiiProxy_DeploymentSpecSelectorMatchExpressionsList) validateGetParameters(index *float64) error {
	return nil
}

func (d *jsiiProxy_DeploymentSpecSelectorMatchExpressionsList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_DeploymentSpecSelectorMatchExpressionsList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_DeploymentSpecSelectorMatchExpressionsList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_DeploymentSpecSelectorMatchExpressionsList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_DeploymentSpecSelectorMatchExpressionsList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewDeploymentSpecSelectorMatchExpressionsListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

