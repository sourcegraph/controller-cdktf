//go:build no_runtime_type_checking

package lexbot

// Building without runtime type checking enabled, so all the below just return nil

func (l *jsiiProxy_LexBotIntentList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (l *jsiiProxy_LexBotIntentList) validateGetParameters(index *float64) error {
	return nil
}

func (l *jsiiProxy_LexBotIntentList) validateResolveParameters(context cdktn.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_LexBotIntentList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_LexBotIntentList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_LexBotIntentList) validateSetTerraformResourceParameters(val cdktn.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_LexBotIntentList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewLexBotIntentListParameters(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

