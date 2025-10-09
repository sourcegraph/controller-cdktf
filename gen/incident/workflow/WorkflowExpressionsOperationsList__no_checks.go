//go:build no_runtime_type_checking

package workflow

// Building without runtime type checking enabled, so all the below just return nil

func (w *jsiiProxy_WorkflowExpressionsOperationsList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (w *jsiiProxy_WorkflowExpressionsOperationsList) validateGetParameters(index *float64) error {
	return nil
}

func (w *jsiiProxy_WorkflowExpressionsOperationsList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_WorkflowExpressionsOperationsList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_WorkflowExpressionsOperationsList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_WorkflowExpressionsOperationsList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_WorkflowExpressionsOperationsList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewWorkflowExpressionsOperationsListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

