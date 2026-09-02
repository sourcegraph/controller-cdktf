//go:build no_runtime_type_checking

package containercluster

// Building without runtime type checking enabled, so all the below just return nil

func (c *jsiiProxy_ContainerClusterNodePoolList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (c *jsiiProxy_ContainerClusterNodePoolList) validateGetParameters(index *float64) error {
	return nil
}

func (c *jsiiProxy_ContainerClusterNodePoolList) validateResolveParameters(context cdktn.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_ContainerClusterNodePoolList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_ContainerClusterNodePoolList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_ContainerClusterNodePoolList) validateSetTerraformResourceParameters(val cdktn.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_ContainerClusterNodePoolList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewContainerClusterNodePoolListParameters(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

