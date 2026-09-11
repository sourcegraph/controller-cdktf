//go:build no_runtime_type_checking

package rediscluster

// Building without runtime type checking enabled, so all the below just return nil

func (r *jsiiProxy_RedisClusterMaintenanceScheduleList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (r *jsiiProxy_RedisClusterMaintenanceScheduleList) validateGetParameters(index *float64) error {
	return nil
}

func (r *jsiiProxy_RedisClusterMaintenanceScheduleList) validateResolveParameters(context cdktn.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_RedisClusterMaintenanceScheduleList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_RedisClusterMaintenanceScheduleList) validateSetTerraformResourceParameters(val cdktn.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_RedisClusterMaintenanceScheduleList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewRedisClusterMaintenanceScheduleListParameters(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

