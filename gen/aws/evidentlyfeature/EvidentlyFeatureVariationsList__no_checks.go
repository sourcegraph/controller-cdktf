//go:build no_runtime_type_checking

package evidentlyfeature

// Building without runtime type checking enabled, so all the below just return nil

func (e *jsiiProxy_EvidentlyFeatureVariationsList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (e *jsiiProxy_EvidentlyFeatureVariationsList) validateGetParameters(index *float64) error {
	return nil
}

func (e *jsiiProxy_EvidentlyFeatureVariationsList) validateResolveParameters(context cdktn.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_EvidentlyFeatureVariationsList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_EvidentlyFeatureVariationsList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_EvidentlyFeatureVariationsList) validateSetTerraformResourceParameters(val cdktn.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_EvidentlyFeatureVariationsList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewEvidentlyFeatureVariationsListParameters(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

