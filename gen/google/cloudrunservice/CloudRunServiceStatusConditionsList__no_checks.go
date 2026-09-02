//go:build no_runtime_type_checking

package cloudrunservice

// Building without runtime type checking enabled, so all the below just return nil

func (c *jsiiProxy_CloudRunServiceStatusConditionsList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (c *jsiiProxy_CloudRunServiceStatusConditionsList) validateGetParameters(index *float64) error {
	return nil
}

func (c *jsiiProxy_CloudRunServiceStatusConditionsList) validateResolveParameters(context cdktn.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_CloudRunServiceStatusConditionsList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_CloudRunServiceStatusConditionsList) validateSetTerraformResourceParameters(val cdktn.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_CloudRunServiceStatusConditionsList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewCloudRunServiceStatusConditionsListParameters(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

