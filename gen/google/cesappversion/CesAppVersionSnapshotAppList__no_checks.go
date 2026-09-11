//go:build no_runtime_type_checking

package cesappversion

// Building without runtime type checking enabled, so all the below just return nil

func (c *jsiiProxy_CesAppVersionSnapshotAppList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (c *jsiiProxy_CesAppVersionSnapshotAppList) validateGetParameters(index *float64) error {
	return nil
}

func (c *jsiiProxy_CesAppVersionSnapshotAppList) validateResolveParameters(context cdktn.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_CesAppVersionSnapshotAppList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_CesAppVersionSnapshotAppList) validateSetTerraformResourceParameters(val cdktn.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_CesAppVersionSnapshotAppList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewCesAppVersionSnapshotAppListParameters(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

