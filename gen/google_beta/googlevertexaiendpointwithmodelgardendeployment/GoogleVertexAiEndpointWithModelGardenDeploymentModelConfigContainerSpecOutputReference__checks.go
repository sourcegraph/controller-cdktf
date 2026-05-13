//go:build !no_runtime_type_checking

package googlevertexaiendpointwithmodelgardendeployment

import (
	"fmt"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

func (g *jsiiProxy_GoogleVertexAiEndpointWithModelGardenDeploymentModelConfigContainerSpecOutputReference) validateGetAnyMapAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (g *jsiiProxy_GoogleVertexAiEndpointWithModelGardenDeploymentModelConfigContainerSpecOutputReference) validateGetBooleanAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (g *jsiiProxy_GoogleVertexAiEndpointWithModelGardenDeploymentModelConfigContainerSpecOutputReference) validateGetBooleanMapAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (g *jsiiProxy_GoogleVertexAiEndpointWithModelGardenDeploymentModelConfigContainerSpecOutputReference) validateGetListAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (g *jsiiProxy_GoogleVertexAiEndpointWithModelGardenDeploymentModelConfigContainerSpecOutputReference) validateGetNumberAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (g *jsiiProxy_GoogleVertexAiEndpointWithModelGardenDeploymentModelConfigContainerSpecOutputReference) validateGetNumberListAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (g *jsiiProxy_GoogleVertexAiEndpointWithModelGardenDeploymentModelConfigContainerSpecOutputReference) validateGetNumberMapAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (g *jsiiProxy_GoogleVertexAiEndpointWithModelGardenDeploymentModelConfigContainerSpecOutputReference) validateGetStringAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (g *jsiiProxy_GoogleVertexAiEndpointWithModelGardenDeploymentModelConfigContainerSpecOutputReference) validateGetStringMapAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (g *jsiiProxy_GoogleVertexAiEndpointWithModelGardenDeploymentModelConfigContainerSpecOutputReference) validateInterpolationForAttributeParameters(property *string) error {
	if property == nil {
		return fmt.Errorf("parameter property is required, but nil was provided")
	}

	return nil
}

func (g *jsiiProxy_GoogleVertexAiEndpointWithModelGardenDeploymentModelConfigContainerSpecOutputReference) validatePutEnvParameters(value interface{}) error {
	if value == nil {
		return fmt.Errorf("parameter value is required, but nil was provided")
	}
	switch value.(type) {
	case cdktf.IResolvable:
		// ok
	case *[]*GoogleVertexAiEndpointWithModelGardenDeploymentModelConfigContainerSpecEnv:
		value := value.(*[]*GoogleVertexAiEndpointWithModelGardenDeploymentModelConfigContainerSpecEnv)
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	case []*GoogleVertexAiEndpointWithModelGardenDeploymentModelConfigContainerSpecEnv:
		value_ := value.([]*GoogleVertexAiEndpointWithModelGardenDeploymentModelConfigContainerSpecEnv)
		value := &value_
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	default:
		if !_jsii_.IsAnonymousProxy(value) {
			return fmt.Errorf("parameter value must be one of the allowed types: cdktf.IResolvable, *[]*GoogleVertexAiEndpointWithModelGardenDeploymentModelConfigContainerSpecEnv; received %#v (a %T)", value, value)
		}
	}

	return nil
}

func (g *jsiiProxy_GoogleVertexAiEndpointWithModelGardenDeploymentModelConfigContainerSpecOutputReference) validatePutGrpcPortsParameters(value interface{}) error {
	if value == nil {
		return fmt.Errorf("parameter value is required, but nil was provided")
	}
	switch value.(type) {
	case cdktf.IResolvable:
		// ok
	case *[]*GoogleVertexAiEndpointWithModelGardenDeploymentModelConfigContainerSpecGrpcPorts:
		value := value.(*[]*GoogleVertexAiEndpointWithModelGardenDeploymentModelConfigContainerSpecGrpcPorts)
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	case []*GoogleVertexAiEndpointWithModelGardenDeploymentModelConfigContainerSpecGrpcPorts:
		value_ := value.([]*GoogleVertexAiEndpointWithModelGardenDeploymentModelConfigContainerSpecGrpcPorts)
		value := &value_
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	default:
		if !_jsii_.IsAnonymousProxy(value) {
			return fmt.Errorf("parameter value must be one of the allowed types: cdktf.IResolvable, *[]*GoogleVertexAiEndpointWithModelGardenDeploymentModelConfigContainerSpecGrpcPorts; received %#v (a %T)", value, value)
		}
	}

	return nil
}

func (g *jsiiProxy_GoogleVertexAiEndpointWithModelGardenDeploymentModelConfigContainerSpecOutputReference) validatePutHealthProbeParameters(value *GoogleVertexAiEndpointWithModelGardenDeploymentModelConfigContainerSpecHealthProbe) error {
	if value == nil {
		return fmt.Errorf("parameter value is required, but nil was provided")
	}
	if err := _jsii_.ValidateStruct(value, func() string { return "parameter value" }); err != nil {
		return err
	}

	return nil
}

func (g *jsiiProxy_GoogleVertexAiEndpointWithModelGardenDeploymentModelConfigContainerSpecOutputReference) validatePutLivenessProbeParameters(value *GoogleVertexAiEndpointWithModelGardenDeploymentModelConfigContainerSpecLivenessProbe) error {
	if value == nil {
		return fmt.Errorf("parameter value is required, but nil was provided")
	}
	if err := _jsii_.ValidateStruct(value, func() string { return "parameter value" }); err != nil {
		return err
	}

	return nil
}

func (g *jsiiProxy_GoogleVertexAiEndpointWithModelGardenDeploymentModelConfigContainerSpecOutputReference) validatePutPortsParameters(value interface{}) error {
	if value == nil {
		return fmt.Errorf("parameter value is required, but nil was provided")
	}
	switch value.(type) {
	case cdktf.IResolvable:
		// ok
	case *[]*GoogleVertexAiEndpointWithModelGardenDeploymentModelConfigContainerSpecPorts:
		value := value.(*[]*GoogleVertexAiEndpointWithModelGardenDeploymentModelConfigContainerSpecPorts)
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	case []*GoogleVertexAiEndpointWithModelGardenDeploymentModelConfigContainerSpecPorts:
		value_ := value.([]*GoogleVertexAiEndpointWithModelGardenDeploymentModelConfigContainerSpecPorts)
		value := &value_
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	default:
		if !_jsii_.IsAnonymousProxy(value) {
			return fmt.Errorf("parameter value must be one of the allowed types: cdktf.IResolvable, *[]*GoogleVertexAiEndpointWithModelGardenDeploymentModelConfigContainerSpecPorts; received %#v (a %T)", value, value)
		}
	}

	return nil
}

func (g *jsiiProxy_GoogleVertexAiEndpointWithModelGardenDeploymentModelConfigContainerSpecOutputReference) validatePutStartupProbeParameters(value *GoogleVertexAiEndpointWithModelGardenDeploymentModelConfigContainerSpecStartupProbe) error {
	if value == nil {
		return fmt.Errorf("parameter value is required, but nil was provided")
	}
	if err := _jsii_.ValidateStruct(value, func() string { return "parameter value" }); err != nil {
		return err
	}

	return nil
}

func (g *jsiiProxy_GoogleVertexAiEndpointWithModelGardenDeploymentModelConfigContainerSpecOutputReference) validateResolveParameters(_context cdktf.IResolveContext) error {
	if _context == nil {
		return fmt.Errorf("parameter _context is required, but nil was provided")
	}

	return nil
}

func (j *jsiiProxy_GoogleVertexAiEndpointWithModelGardenDeploymentModelConfigContainerSpecOutputReference) validateSetArgsParameters(val *[]*string) error {
	if val == nil {
		return fmt.Errorf("parameter val is required, but nil was provided")
	}

	return nil
}

func (j *jsiiProxy_GoogleVertexAiEndpointWithModelGardenDeploymentModelConfigContainerSpecOutputReference) validateSetCommandParameters(val *[]*string) error {
	if val == nil {
		return fmt.Errorf("parameter val is required, but nil was provided")
	}

	return nil
}

func (j *jsiiProxy_GoogleVertexAiEndpointWithModelGardenDeploymentModelConfigContainerSpecOutputReference) validateSetComplexObjectIndexParameters(val interface{}) error {
	switch val.(type) {
	case *string:
		// ok
	case string:
		// ok
	case *float64:
		// ok
	case float64:
		// ok
	case *int:
		// ok
	case int:
		// ok
	case *uint:
		// ok
	case uint:
		// ok
	case *int8:
		// ok
	case int8:
		// ok
	case *int16:
		// ok
	case int16:
		// ok
	case *int32:
		// ok
	case int32:
		// ok
	case *int64:
		// ok
	case int64:
		// ok
	case *uint8:
		// ok
	case uint8:
		// ok
	case *uint16:
		// ok
	case uint16:
		// ok
	case *uint32:
		// ok
	case uint32:
		// ok
	case *uint64:
		// ok
	case uint64:
		// ok
	default:
		return fmt.Errorf("parameter val must be one of the allowed types: *string, *float64; received %#v (a %T)", val, val)
	}

	return nil
}

func (j *jsiiProxy_GoogleVertexAiEndpointWithModelGardenDeploymentModelConfigContainerSpecOutputReference) validateSetComplexObjectIsFromSetParameters(val *bool) error {
	if val == nil {
		return fmt.Errorf("parameter val is required, but nil was provided")
	}

	return nil
}

func (j *jsiiProxy_GoogleVertexAiEndpointWithModelGardenDeploymentModelConfigContainerSpecOutputReference) validateSetDeploymentTimeoutParameters(val *string) error {
	if val == nil {
		return fmt.Errorf("parameter val is required, but nil was provided")
	}

	return nil
}

func (j *jsiiProxy_GoogleVertexAiEndpointWithModelGardenDeploymentModelConfigContainerSpecOutputReference) validateSetHealthRouteParameters(val *string) error {
	if val == nil {
		return fmt.Errorf("parameter val is required, but nil was provided")
	}

	return nil
}

func (j *jsiiProxy_GoogleVertexAiEndpointWithModelGardenDeploymentModelConfigContainerSpecOutputReference) validateSetImageUriParameters(val *string) error {
	if val == nil {
		return fmt.Errorf("parameter val is required, but nil was provided")
	}

	return nil
}

func (j *jsiiProxy_GoogleVertexAiEndpointWithModelGardenDeploymentModelConfigContainerSpecOutputReference) validateSetInternalValueParameters(val *GoogleVertexAiEndpointWithModelGardenDeploymentModelConfigContainerSpec) error {
	if err := _jsii_.ValidateStruct(val, func() string { return "parameter val" }); err != nil {
		return err
	}

	return nil
}

func (j *jsiiProxy_GoogleVertexAiEndpointWithModelGardenDeploymentModelConfigContainerSpecOutputReference) validateSetPredictRouteParameters(val *string) error {
	if val == nil {
		return fmt.Errorf("parameter val is required, but nil was provided")
	}

	return nil
}

func (j *jsiiProxy_GoogleVertexAiEndpointWithModelGardenDeploymentModelConfigContainerSpecOutputReference) validateSetSharedMemorySizeMbParameters(val *string) error {
	if val == nil {
		return fmt.Errorf("parameter val is required, but nil was provided")
	}

	return nil
}

func (j *jsiiProxy_GoogleVertexAiEndpointWithModelGardenDeploymentModelConfigContainerSpecOutputReference) validateSetTerraformAttributeParameters(val *string) error {
	if val == nil {
		return fmt.Errorf("parameter val is required, but nil was provided")
	}

	return nil
}

func (j *jsiiProxy_GoogleVertexAiEndpointWithModelGardenDeploymentModelConfigContainerSpecOutputReference) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	if val == nil {
		return fmt.Errorf("parameter val is required, but nil was provided")
	}

	return nil
}

func validateNewGoogleVertexAiEndpointWithModelGardenDeploymentModelConfigContainerSpecOutputReferenceParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) error {
	if terraformResource == nil {
		return fmt.Errorf("parameter terraformResource is required, but nil was provided")
	}

	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

