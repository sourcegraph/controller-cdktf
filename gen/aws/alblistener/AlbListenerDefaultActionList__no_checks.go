//go:build no_runtime_type_checking

package alblistener

// Building without runtime type checking enabled, so all the below just return nil

func (a *jsiiProxy_AlbListenerDefaultActionList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (a *jsiiProxy_AlbListenerDefaultActionList) validateGetParameters(index *float64) error {
	return nil
}

func (a *jsiiProxy_AlbListenerDefaultActionList) validateResolveParameters(context cdktn.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_AlbListenerDefaultActionList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_AlbListenerDefaultActionList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_AlbListenerDefaultActionList) validateSetTerraformResourceParameters(val cdktn.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_AlbListenerDefaultActionList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewAlbListenerDefaultActionListParameters(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

