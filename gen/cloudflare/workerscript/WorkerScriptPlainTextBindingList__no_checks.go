//go:build no_runtime_type_checking

package workerscript

// Building without runtime type checking enabled, so all the below just return nil

func (w *jsiiProxy_WorkerScriptPlainTextBindingList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (w *jsiiProxy_WorkerScriptPlainTextBindingList) validateGetParameters(index *float64) error {
	return nil
}

func (w *jsiiProxy_WorkerScriptPlainTextBindingList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_WorkerScriptPlainTextBindingList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_WorkerScriptPlainTextBindingList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_WorkerScriptPlainTextBindingList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_WorkerScriptPlainTextBindingList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewWorkerScriptPlainTextBindingListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

