//go:build no_runtime_type_checking

package workerscript

// Building without runtime type checking enabled, so all the below just return nil

func (w *jsiiProxy_WorkerScriptKvNamespaceBindingList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (w *jsiiProxy_WorkerScriptKvNamespaceBindingList) validateGetParameters(index *float64) error {
	return nil
}

func (w *jsiiProxy_WorkerScriptKvNamespaceBindingList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_WorkerScriptKvNamespaceBindingList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_WorkerScriptKvNamespaceBindingList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_WorkerScriptKvNamespaceBindingList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_WorkerScriptKvNamespaceBindingList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewWorkerScriptKvNamespaceBindingListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

