//go:build no_runtime_type_checking

package healthcheck

// Building without runtime type checking enabled, so all the below just return nil

func (h *jsiiProxy_HealthcheckHeaderList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (h *jsiiProxy_HealthcheckHeaderList) validateGetParameters(index *float64) error {
	return nil
}

func (h *jsiiProxy_HealthcheckHeaderList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_HealthcheckHeaderList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_HealthcheckHeaderList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_HealthcheckHeaderList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_HealthcheckHeaderList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewHealthcheckHeaderListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

