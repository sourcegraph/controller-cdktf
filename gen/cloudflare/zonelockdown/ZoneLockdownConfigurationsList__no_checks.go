//go:build no_runtime_type_checking

package zonelockdown

// Building without runtime type checking enabled, so all the below just return nil

func (z *jsiiProxy_ZoneLockdownConfigurationsList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (z *jsiiProxy_ZoneLockdownConfigurationsList) validateGetParameters(index *float64) error {
	return nil
}

func (z *jsiiProxy_ZoneLockdownConfigurationsList) validateResolveParameters(context cdktn.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_ZoneLockdownConfigurationsList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_ZoneLockdownConfigurationsList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_ZoneLockdownConfigurationsList) validateSetTerraformResourceParameters(val cdktn.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_ZoneLockdownConfigurationsList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewZoneLockdownConfigurationsListParameters(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

