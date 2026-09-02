//go:build no_runtime_type_checking

package securitygroup

// Building without runtime type checking enabled, so all the below just return nil

func (s *jsiiProxy_SecurityGroupEgressList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (s *jsiiProxy_SecurityGroupEgressList) validateGetParameters(index *float64) error {
	return nil
}

func (s *jsiiProxy_SecurityGroupEgressList) validateResolveParameters(context cdktn.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_SecurityGroupEgressList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_SecurityGroupEgressList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_SecurityGroupEgressList) validateSetTerraformResourceParameters(val cdktn.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_SecurityGroupEgressList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewSecurityGroupEgressListParameters(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

