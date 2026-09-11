//go:build no_runtime_type_checking

package redshiftcluster

// Building without runtime type checking enabled, so all the below just return nil

func (r *jsiiProxy_RedshiftClusterClusterNodesList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (r *jsiiProxy_RedshiftClusterClusterNodesList) validateGetParameters(index *float64) error {
	return nil
}

func (r *jsiiProxy_RedshiftClusterClusterNodesList) validateResolveParameters(context cdktn.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_RedshiftClusterClusterNodesList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_RedshiftClusterClusterNodesList) validateSetTerraformResourceParameters(val cdktn.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_RedshiftClusterClusterNodesList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewRedshiftClusterClusterNodesListParameters(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

