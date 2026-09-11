//go:build no_runtime_type_checking

package rediscluster

// Building without runtime type checking enabled, so all the below just return nil

func (r *jsiiProxy_RedisClusterStateInfoUpdateInfoList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (r *jsiiProxy_RedisClusterStateInfoUpdateInfoList) validateGetParameters(index *float64) error {
	return nil
}

func (r *jsiiProxy_RedisClusterStateInfoUpdateInfoList) validateResolveParameters(context cdktn.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_RedisClusterStateInfoUpdateInfoList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_RedisClusterStateInfoUpdateInfoList) validateSetTerraformResourceParameters(val cdktn.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_RedisClusterStateInfoUpdateInfoList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewRedisClusterStateInfoUpdateInfoListParameters(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

