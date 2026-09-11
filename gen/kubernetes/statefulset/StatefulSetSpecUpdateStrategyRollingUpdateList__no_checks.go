//go:build no_runtime_type_checking

package statefulset

// Building without runtime type checking enabled, so all the below just return nil

func (s *jsiiProxy_StatefulSetSpecUpdateStrategyRollingUpdateList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (s *jsiiProxy_StatefulSetSpecUpdateStrategyRollingUpdateList) validateGetParameters(index *float64) error {
	return nil
}

func (s *jsiiProxy_StatefulSetSpecUpdateStrategyRollingUpdateList) validateResolveParameters(context cdktn.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_StatefulSetSpecUpdateStrategyRollingUpdateList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_StatefulSetSpecUpdateStrategyRollingUpdateList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_StatefulSetSpecUpdateStrategyRollingUpdateList) validateSetTerraformResourceParameters(val cdktn.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_StatefulSetSpecUpdateStrategyRollingUpdateList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewStatefulSetSpecUpdateStrategyRollingUpdateListParameters(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

