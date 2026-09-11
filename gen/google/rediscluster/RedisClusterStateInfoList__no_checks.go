//go:build no_runtime_type_checking

package rediscluster

// Building without runtime type checking enabled, so all the below just return nil

func (r *jsiiProxy_RedisClusterStateInfoList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (r *jsiiProxy_RedisClusterStateInfoList) validateGetParameters(index *float64) error {
	return nil
}

func (r *jsiiProxy_RedisClusterStateInfoList) validateResolveParameters(context cdktn.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_RedisClusterStateInfoList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_RedisClusterStateInfoList) validateSetTerraformResourceParameters(val cdktn.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_RedisClusterStateInfoList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewRedisClusterStateInfoListParameters(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

