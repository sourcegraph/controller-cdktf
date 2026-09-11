//go:build no_runtime_type_checking

package fallbackdomain

// Building without runtime type checking enabled, so all the below just return nil

func (f *jsiiProxy_FallbackDomainDomainsList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (f *jsiiProxy_FallbackDomainDomainsList) validateGetParameters(index *float64) error {
	return nil
}

func (f *jsiiProxy_FallbackDomainDomainsList) validateResolveParameters(context cdktn.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_FallbackDomainDomainsList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_FallbackDomainDomainsList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_FallbackDomainDomainsList) validateSetTerraformResourceParameters(val cdktn.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_FallbackDomainDomainsList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewFallbackDomainDomainsListParameters(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

