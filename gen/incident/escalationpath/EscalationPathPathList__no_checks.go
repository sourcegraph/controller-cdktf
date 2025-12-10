//go:build no_runtime_type_checking

package escalationpath

// Building without runtime type checking enabled, so all the below just return nil

func (e *jsiiProxy_EscalationPathPathList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (e *jsiiProxy_EscalationPathPathList) validateGetParameters(index *float64) error {
	return nil
}

func (e *jsiiProxy_EscalationPathPathList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_EscalationPathPathList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_EscalationPathPathList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_EscalationPathPathList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_EscalationPathPathList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewEscalationPathPathListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

