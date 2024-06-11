//go:build no_runtime_type_checking

package workerscript

// Building without runtime type checking enabled, so all the below just return nil

func (w *jsiiProxy_WorkerScriptSecretTextBindingList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (w *jsiiProxy_WorkerScriptSecretTextBindingList) validateGetParameters(index *float64) error {
	return nil
}

func (w *jsiiProxy_WorkerScriptSecretTextBindingList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_WorkerScriptSecretTextBindingList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_WorkerScriptSecretTextBindingList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_WorkerScriptSecretTextBindingList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_WorkerScriptSecretTextBindingList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewWorkerScriptSecretTextBindingListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

