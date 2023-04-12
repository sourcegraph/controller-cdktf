//go:build no_runtime_type_checking

package dataoktabehaviors

// Building without runtime type checking enabled, so all the below just return nil

func (d *jsiiProxy_DataOktaBehaviorsBehaviorsList) validateGetParameters(index *float64) error {
	return nil
}

func (d *jsiiProxy_DataOktaBehaviorsBehaviorsList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_DataOktaBehaviorsBehaviorsList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_DataOktaBehaviorsBehaviorsList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_DataOktaBehaviorsBehaviorsList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewDataOktaBehaviorsBehaviorsListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

