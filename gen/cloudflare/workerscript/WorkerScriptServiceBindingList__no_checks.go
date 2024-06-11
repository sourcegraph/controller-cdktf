//go:build no_runtime_type_checking

package workerscript

// Building without runtime type checking enabled, so all the below just return nil

func (w *jsiiProxy_WorkerScriptServiceBindingList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (w *jsiiProxy_WorkerScriptServiceBindingList) validateGetParameters(index *float64) error {
	return nil
}

func (w *jsiiProxy_WorkerScriptServiceBindingList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_WorkerScriptServiceBindingList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_WorkerScriptServiceBindingList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_WorkerScriptServiceBindingList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_WorkerScriptServiceBindingList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewWorkerScriptServiceBindingListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

