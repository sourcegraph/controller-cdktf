//go:build no_runtime_type_checking

package ingress

// Building without runtime type checking enabled, so all the below just return nil

func (i *jsiiProxy_IngressSpecRuleHttpPathList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (i *jsiiProxy_IngressSpecRuleHttpPathList) validateGetParameters(index *float64) error {
	return nil
}

func (i *jsiiProxy_IngressSpecRuleHttpPathList) validateResolveParameters(context cdktn.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_IngressSpecRuleHttpPathList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_IngressSpecRuleHttpPathList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_IngressSpecRuleHttpPathList) validateSetTerraformResourceParameters(val cdktn.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_IngressSpecRuleHttpPathList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewIngressSpecRuleHttpPathListParameters(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

