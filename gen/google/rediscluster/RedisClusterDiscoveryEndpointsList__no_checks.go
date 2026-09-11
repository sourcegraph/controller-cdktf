//go:build no_runtime_type_checking

package rediscluster

// Building without runtime type checking enabled, so all the below just return nil

func (r *jsiiProxy_RedisClusterDiscoveryEndpointsList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (r *jsiiProxy_RedisClusterDiscoveryEndpointsList) validateGetParameters(index *float64) error {
	return nil
}

func (r *jsiiProxy_RedisClusterDiscoveryEndpointsList) validateResolveParameters(context cdktn.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_RedisClusterDiscoveryEndpointsList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_RedisClusterDiscoveryEndpointsList) validateSetTerraformResourceParameters(val cdktn.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_RedisClusterDiscoveryEndpointsList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewRedisClusterDiscoveryEndpointsListParameters(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

