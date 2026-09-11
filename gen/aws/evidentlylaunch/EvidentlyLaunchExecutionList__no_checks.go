//go:build no_runtime_type_checking

package evidentlylaunch

// Building without runtime type checking enabled, so all the below just return nil

func (e *jsiiProxy_EvidentlyLaunchExecutionList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (e *jsiiProxy_EvidentlyLaunchExecutionList) validateGetParameters(index *float64) error {
	return nil
}

func (e *jsiiProxy_EvidentlyLaunchExecutionList) validateResolveParameters(context cdktn.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_EvidentlyLaunchExecutionList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_EvidentlyLaunchExecutionList) validateSetTerraformResourceParameters(val cdktn.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_EvidentlyLaunchExecutionList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewEvidentlyLaunchExecutionListParameters(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

