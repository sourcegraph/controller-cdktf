//go:build no_runtime_type_checking

package budget

// Building without runtime type checking enabled, so all the below just return nil

func (b *jsiiProxy_Budget) validateAddOverrideParameters(path *string, value interface{}) error {
	return nil
}

func (b *jsiiProxy_Budget) validateAddProviderParameters(provider interface{}) error {
	return nil
}

func (b *jsiiProxy_Budget) validateGetStringParameters(output *string) error {
	return nil
}

func (b *jsiiProxy_Budget) validateInterpolationForOutputParameters(moduleOutput *string) error {
	return nil
}

func (b *jsiiProxy_Budget) validateOverrideLogicalIdParameters(newLogicalId *string) error {
	return nil
}

func validateBudget_IsConstructParameters(x interface{}) error {
	return nil
}

func validateBudget_IsTerraformElementParameters(x interface{}) error {
	return nil
}

func (j *jsiiProxy_Budget) validateSetAmountParameters(val *float64) error {
	return nil
}

func (j *jsiiProxy_Budget) validateSetBillingAccountParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_Budget) validateSetProjectsParameters(val *[]*string) error {
	return nil
}

func validateNewBudgetParameters(scope constructs.Construct, id *string, config *BudgetConfig) error {
	return nil
}

