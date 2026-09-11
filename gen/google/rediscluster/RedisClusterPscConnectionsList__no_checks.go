//go:build no_runtime_type_checking

package rediscluster

// Building without runtime type checking enabled, so all the below just return nil

func (r *jsiiProxy_RedisClusterPscConnectionsList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (r *jsiiProxy_RedisClusterPscConnectionsList) validateGetParameters(index *float64) error {
	return nil
}

func (r *jsiiProxy_RedisClusterPscConnectionsList) validateResolveParameters(context cdktn.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_RedisClusterPscConnectionsList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_RedisClusterPscConnectionsList) validateSetTerraformResourceParameters(val cdktn.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_RedisClusterPscConnectionsList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewRedisClusterPscConnectionsListParameters(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

