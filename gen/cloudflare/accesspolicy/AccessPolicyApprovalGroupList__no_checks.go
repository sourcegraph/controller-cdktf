//go:build no_runtime_type_checking

package accesspolicy

// Building without runtime type checking enabled, so all the below just return nil

func (a *jsiiProxy_AccessPolicyApprovalGroupList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (a *jsiiProxy_AccessPolicyApprovalGroupList) validateGetParameters(index *float64) error {
	return nil
}

func (a *jsiiProxy_AccessPolicyApprovalGroupList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_AccessPolicyApprovalGroupList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_AccessPolicyApprovalGroupList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_AccessPolicyApprovalGroupList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_AccessPolicyApprovalGroupList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewAccessPolicyApprovalGroupListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

