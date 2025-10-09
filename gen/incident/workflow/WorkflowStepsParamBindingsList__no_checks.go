//go:build no_runtime_type_checking

package workflow

// Building without runtime type checking enabled, so all the below just return nil

func (w *jsiiProxy_WorkflowStepsParamBindingsList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (w *jsiiProxy_WorkflowStepsParamBindingsList) validateGetParameters(index *float64) error {
	return nil
}

func (w *jsiiProxy_WorkflowStepsParamBindingsList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_WorkflowStepsParamBindingsList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_WorkflowStepsParamBindingsList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_WorkflowStepsParamBindingsList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_WorkflowStepsParamBindingsList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewWorkflowStepsParamBindingsListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

