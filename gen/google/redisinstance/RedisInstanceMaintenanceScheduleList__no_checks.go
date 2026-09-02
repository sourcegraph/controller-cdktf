//go:build no_runtime_type_checking

package redisinstance

// Building without runtime type checking enabled, so all the below just return nil

func (r *jsiiProxy_RedisInstanceMaintenanceScheduleList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (r *jsiiProxy_RedisInstanceMaintenanceScheduleList) validateGetParameters(index *float64) error {
	return nil
}

func (r *jsiiProxy_RedisInstanceMaintenanceScheduleList) validateResolveParameters(context cdktn.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_RedisInstanceMaintenanceScheduleList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_RedisInstanceMaintenanceScheduleList) validateSetTerraformResourceParameters(val cdktn.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_RedisInstanceMaintenanceScheduleList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewRedisInstanceMaintenanceScheduleListParameters(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

