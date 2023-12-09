//go:build no_runtime_type_checking

package waitingroom

// Building without runtime type checking enabled, so all the below just return nil

func (w *jsiiProxy_WaitingRoomAdditionalRoutesList) validateGetParameters(index *float64) error {
	return nil
}

func (w *jsiiProxy_WaitingRoomAdditionalRoutesList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_WaitingRoomAdditionalRoutesList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_WaitingRoomAdditionalRoutesList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_WaitingRoomAdditionalRoutesList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_WaitingRoomAdditionalRoutesList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewWaitingRoomAdditionalRoutesListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

