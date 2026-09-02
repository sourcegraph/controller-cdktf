//go:build no_runtime_type_checking

package route53zone

// Building without runtime type checking enabled, so all the below just return nil

func (r *jsiiProxy_Route53ZoneVpcList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (r *jsiiProxy_Route53ZoneVpcList) validateGetParameters(index *float64) error {
	return nil
}

func (r *jsiiProxy_Route53ZoneVpcList) validateResolveParameters(context cdktn.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_Route53ZoneVpcList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_Route53ZoneVpcList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_Route53ZoneVpcList) validateSetTerraformResourceParameters(val cdktn.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_Route53ZoneVpcList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewRoute53ZoneVpcListParameters(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

