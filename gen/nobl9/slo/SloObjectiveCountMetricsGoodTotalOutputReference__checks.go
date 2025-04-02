//go:build !no_runtime_type_checking

package slo

import (
	"fmt"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

func (s *jsiiProxy_SloObjectiveCountMetricsGoodTotalOutputReference) validateGetAnyMapAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (s *jsiiProxy_SloObjectiveCountMetricsGoodTotalOutputReference) validateGetBooleanAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (s *jsiiProxy_SloObjectiveCountMetricsGoodTotalOutputReference) validateGetBooleanMapAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (s *jsiiProxy_SloObjectiveCountMetricsGoodTotalOutputReference) validateGetListAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (s *jsiiProxy_SloObjectiveCountMetricsGoodTotalOutputReference) validateGetNumberAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (s *jsiiProxy_SloObjectiveCountMetricsGoodTotalOutputReference) validateGetNumberListAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (s *jsiiProxy_SloObjectiveCountMetricsGoodTotalOutputReference) validateGetNumberMapAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (s *jsiiProxy_SloObjectiveCountMetricsGoodTotalOutputReference) validateGetStringAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (s *jsiiProxy_SloObjectiveCountMetricsGoodTotalOutputReference) validateGetStringMapAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (s *jsiiProxy_SloObjectiveCountMetricsGoodTotalOutputReference) validateInterpolationForAttributeParameters(property *string) error {
	if property == nil {
		return fmt.Errorf("parameter property is required, but nil was provided")
	}

	return nil
}

func (s *jsiiProxy_SloObjectiveCountMetricsGoodTotalOutputReference) validatePutAmazonPrometheusParameters(value interface{}) error {
	if value == nil {
		return fmt.Errorf("parameter value is required, but nil was provided")
	}
	switch value.(type) {
	case cdktf.IResolvable:
		// ok
	case *[]*SloObjectiveCountMetricsGoodTotalAmazonPrometheus:
		value := value.(*[]*SloObjectiveCountMetricsGoodTotalAmazonPrometheus)
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	case []*SloObjectiveCountMetricsGoodTotalAmazonPrometheus:
		value_ := value.([]*SloObjectiveCountMetricsGoodTotalAmazonPrometheus)
		value := &value_
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	default:
		if !_jsii_.IsAnonymousProxy(value) {
			return fmt.Errorf("parameter value must be one of the allowed types: cdktf.IResolvable, *[]*SloObjectiveCountMetricsGoodTotalAmazonPrometheus; received %#v (a %T)", value, value)
		}
	}

	return nil
}

func (s *jsiiProxy_SloObjectiveCountMetricsGoodTotalOutputReference) validatePutAppdynamicsParameters(value interface{}) error {
	if value == nil {
		return fmt.Errorf("parameter value is required, but nil was provided")
	}
	switch value.(type) {
	case cdktf.IResolvable:
		// ok
	case *[]*SloObjectiveCountMetricsGoodTotalAppdynamics:
		value := value.(*[]*SloObjectiveCountMetricsGoodTotalAppdynamics)
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	case []*SloObjectiveCountMetricsGoodTotalAppdynamics:
		value_ := value.([]*SloObjectiveCountMetricsGoodTotalAppdynamics)
		value := &value_
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	default:
		if !_jsii_.IsAnonymousProxy(value) {
			return fmt.Errorf("parameter value must be one of the allowed types: cdktf.IResolvable, *[]*SloObjectiveCountMetricsGoodTotalAppdynamics; received %#v (a %T)", value, value)
		}
	}

	return nil
}

func (s *jsiiProxy_SloObjectiveCountMetricsGoodTotalOutputReference) validatePutAzureMonitorParameters(value interface{}) error {
	if value == nil {
		return fmt.Errorf("parameter value is required, but nil was provided")
	}
	switch value.(type) {
	case cdktf.IResolvable:
		// ok
	case *[]*SloObjectiveCountMetricsGoodTotalAzureMonitor:
		value := value.(*[]*SloObjectiveCountMetricsGoodTotalAzureMonitor)
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	case []*SloObjectiveCountMetricsGoodTotalAzureMonitor:
		value_ := value.([]*SloObjectiveCountMetricsGoodTotalAzureMonitor)
		value := &value_
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	default:
		if !_jsii_.IsAnonymousProxy(value) {
			return fmt.Errorf("parameter value must be one of the allowed types: cdktf.IResolvable, *[]*SloObjectiveCountMetricsGoodTotalAzureMonitor; received %#v (a %T)", value, value)
		}
	}

	return nil
}

func (s *jsiiProxy_SloObjectiveCountMetricsGoodTotalOutputReference) validatePutBigqueryParameters(value interface{}) error {
	if value == nil {
		return fmt.Errorf("parameter value is required, but nil was provided")
	}
	switch value.(type) {
	case cdktf.IResolvable:
		// ok
	case *[]*SloObjectiveCountMetricsGoodTotalBigquery:
		value := value.(*[]*SloObjectiveCountMetricsGoodTotalBigquery)
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	case []*SloObjectiveCountMetricsGoodTotalBigquery:
		value_ := value.([]*SloObjectiveCountMetricsGoodTotalBigquery)
		value := &value_
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	default:
		if !_jsii_.IsAnonymousProxy(value) {
			return fmt.Errorf("parameter value must be one of the allowed types: cdktf.IResolvable, *[]*SloObjectiveCountMetricsGoodTotalBigquery; received %#v (a %T)", value, value)
		}
	}

	return nil
}

func (s *jsiiProxy_SloObjectiveCountMetricsGoodTotalOutputReference) validatePutCloudwatchParameters(value interface{}) error {
	if value == nil {
		return fmt.Errorf("parameter value is required, but nil was provided")
	}
	switch value.(type) {
	case cdktf.IResolvable:
		// ok
	case *[]*SloObjectiveCountMetricsGoodTotalCloudwatch:
		value := value.(*[]*SloObjectiveCountMetricsGoodTotalCloudwatch)
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	case []*SloObjectiveCountMetricsGoodTotalCloudwatch:
		value_ := value.([]*SloObjectiveCountMetricsGoodTotalCloudwatch)
		value := &value_
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	default:
		if !_jsii_.IsAnonymousProxy(value) {
			return fmt.Errorf("parameter value must be one of the allowed types: cdktf.IResolvable, *[]*SloObjectiveCountMetricsGoodTotalCloudwatch; received %#v (a %T)", value, value)
		}
	}

	return nil
}

func (s *jsiiProxy_SloObjectiveCountMetricsGoodTotalOutputReference) validatePutDatadogParameters(value interface{}) error {
	if value == nil {
		return fmt.Errorf("parameter value is required, but nil was provided")
	}
	switch value.(type) {
	case cdktf.IResolvable:
		// ok
	case *[]*SloObjectiveCountMetricsGoodTotalDatadog:
		value := value.(*[]*SloObjectiveCountMetricsGoodTotalDatadog)
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	case []*SloObjectiveCountMetricsGoodTotalDatadog:
		value_ := value.([]*SloObjectiveCountMetricsGoodTotalDatadog)
		value := &value_
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	default:
		if !_jsii_.IsAnonymousProxy(value) {
			return fmt.Errorf("parameter value must be one of the allowed types: cdktf.IResolvable, *[]*SloObjectiveCountMetricsGoodTotalDatadog; received %#v (a %T)", value, value)
		}
	}

	return nil
}

func (s *jsiiProxy_SloObjectiveCountMetricsGoodTotalOutputReference) validatePutDynatraceParameters(value interface{}) error {
	if value == nil {
		return fmt.Errorf("parameter value is required, but nil was provided")
	}
	switch value.(type) {
	case cdktf.IResolvable:
		// ok
	case *[]*SloObjectiveCountMetricsGoodTotalDynatrace:
		value := value.(*[]*SloObjectiveCountMetricsGoodTotalDynatrace)
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	case []*SloObjectiveCountMetricsGoodTotalDynatrace:
		value_ := value.([]*SloObjectiveCountMetricsGoodTotalDynatrace)
		value := &value_
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	default:
		if !_jsii_.IsAnonymousProxy(value) {
			return fmt.Errorf("parameter value must be one of the allowed types: cdktf.IResolvable, *[]*SloObjectiveCountMetricsGoodTotalDynatrace; received %#v (a %T)", value, value)
		}
	}

	return nil
}

func (s *jsiiProxy_SloObjectiveCountMetricsGoodTotalOutputReference) validatePutElasticsearchParameters(value interface{}) error {
	if value == nil {
		return fmt.Errorf("parameter value is required, but nil was provided")
	}
	switch value.(type) {
	case cdktf.IResolvable:
		// ok
	case *[]*SloObjectiveCountMetricsGoodTotalElasticsearch:
		value := value.(*[]*SloObjectiveCountMetricsGoodTotalElasticsearch)
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	case []*SloObjectiveCountMetricsGoodTotalElasticsearch:
		value_ := value.([]*SloObjectiveCountMetricsGoodTotalElasticsearch)
		value := &value_
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	default:
		if !_jsii_.IsAnonymousProxy(value) {
			return fmt.Errorf("parameter value must be one of the allowed types: cdktf.IResolvable, *[]*SloObjectiveCountMetricsGoodTotalElasticsearch; received %#v (a %T)", value, value)
		}
	}

	return nil
}

func (s *jsiiProxy_SloObjectiveCountMetricsGoodTotalOutputReference) validatePutGcmParameters(value interface{}) error {
	if value == nil {
		return fmt.Errorf("parameter value is required, but nil was provided")
	}
	switch value.(type) {
	case cdktf.IResolvable:
		// ok
	case *[]*SloObjectiveCountMetricsGoodTotalGcm:
		value := value.(*[]*SloObjectiveCountMetricsGoodTotalGcm)
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	case []*SloObjectiveCountMetricsGoodTotalGcm:
		value_ := value.([]*SloObjectiveCountMetricsGoodTotalGcm)
		value := &value_
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	default:
		if !_jsii_.IsAnonymousProxy(value) {
			return fmt.Errorf("parameter value must be one of the allowed types: cdktf.IResolvable, *[]*SloObjectiveCountMetricsGoodTotalGcm; received %#v (a %T)", value, value)
		}
	}

	return nil
}

func (s *jsiiProxy_SloObjectiveCountMetricsGoodTotalOutputReference) validatePutGrafanaLokiParameters(value interface{}) error {
	if value == nil {
		return fmt.Errorf("parameter value is required, but nil was provided")
	}
	switch value.(type) {
	case cdktf.IResolvable:
		// ok
	case *[]*SloObjectiveCountMetricsGoodTotalGrafanaLoki:
		value := value.(*[]*SloObjectiveCountMetricsGoodTotalGrafanaLoki)
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	case []*SloObjectiveCountMetricsGoodTotalGrafanaLoki:
		value_ := value.([]*SloObjectiveCountMetricsGoodTotalGrafanaLoki)
		value := &value_
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	default:
		if !_jsii_.IsAnonymousProxy(value) {
			return fmt.Errorf("parameter value must be one of the allowed types: cdktf.IResolvable, *[]*SloObjectiveCountMetricsGoodTotalGrafanaLoki; received %#v (a %T)", value, value)
		}
	}

	return nil
}

func (s *jsiiProxy_SloObjectiveCountMetricsGoodTotalOutputReference) validatePutGraphiteParameters(value interface{}) error {
	if value == nil {
		return fmt.Errorf("parameter value is required, but nil was provided")
	}
	switch value.(type) {
	case cdktf.IResolvable:
		// ok
	case *[]*SloObjectiveCountMetricsGoodTotalGraphite:
		value := value.(*[]*SloObjectiveCountMetricsGoodTotalGraphite)
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	case []*SloObjectiveCountMetricsGoodTotalGraphite:
		value_ := value.([]*SloObjectiveCountMetricsGoodTotalGraphite)
		value := &value_
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	default:
		if !_jsii_.IsAnonymousProxy(value) {
			return fmt.Errorf("parameter value must be one of the allowed types: cdktf.IResolvable, *[]*SloObjectiveCountMetricsGoodTotalGraphite; received %#v (a %T)", value, value)
		}
	}

	return nil
}

func (s *jsiiProxy_SloObjectiveCountMetricsGoodTotalOutputReference) validatePutHoneycombParameters(value interface{}) error {
	if value == nil {
		return fmt.Errorf("parameter value is required, but nil was provided")
	}
	switch value.(type) {
	case cdktf.IResolvable:
		// ok
	case *[]*SloObjectiveCountMetricsGoodTotalHoneycomb:
		value := value.(*[]*SloObjectiveCountMetricsGoodTotalHoneycomb)
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	case []*SloObjectiveCountMetricsGoodTotalHoneycomb:
		value_ := value.([]*SloObjectiveCountMetricsGoodTotalHoneycomb)
		value := &value_
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	default:
		if !_jsii_.IsAnonymousProxy(value) {
			return fmt.Errorf("parameter value must be one of the allowed types: cdktf.IResolvable, *[]*SloObjectiveCountMetricsGoodTotalHoneycomb; received %#v (a %T)", value, value)
		}
	}

	return nil
}

func (s *jsiiProxy_SloObjectiveCountMetricsGoodTotalOutputReference) validatePutInfluxdbParameters(value interface{}) error {
	if value == nil {
		return fmt.Errorf("parameter value is required, but nil was provided")
	}
	switch value.(type) {
	case cdktf.IResolvable:
		// ok
	case *[]*SloObjectiveCountMetricsGoodTotalInfluxdb:
		value := value.(*[]*SloObjectiveCountMetricsGoodTotalInfluxdb)
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	case []*SloObjectiveCountMetricsGoodTotalInfluxdb:
		value_ := value.([]*SloObjectiveCountMetricsGoodTotalInfluxdb)
		value := &value_
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	default:
		if !_jsii_.IsAnonymousProxy(value) {
			return fmt.Errorf("parameter value must be one of the allowed types: cdktf.IResolvable, *[]*SloObjectiveCountMetricsGoodTotalInfluxdb; received %#v (a %T)", value, value)
		}
	}

	return nil
}

func (s *jsiiProxy_SloObjectiveCountMetricsGoodTotalOutputReference) validatePutInstanaParameters(value interface{}) error {
	if value == nil {
		return fmt.Errorf("parameter value is required, but nil was provided")
	}
	switch value.(type) {
	case cdktf.IResolvable:
		// ok
	case *[]*SloObjectiveCountMetricsGoodTotalInstana:
		value := value.(*[]*SloObjectiveCountMetricsGoodTotalInstana)
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	case []*SloObjectiveCountMetricsGoodTotalInstana:
		value_ := value.([]*SloObjectiveCountMetricsGoodTotalInstana)
		value := &value_
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	default:
		if !_jsii_.IsAnonymousProxy(value) {
			return fmt.Errorf("parameter value must be one of the allowed types: cdktf.IResolvable, *[]*SloObjectiveCountMetricsGoodTotalInstana; received %#v (a %T)", value, value)
		}
	}

	return nil
}

func (s *jsiiProxy_SloObjectiveCountMetricsGoodTotalOutputReference) validatePutLightstepParameters(value interface{}) error {
	if value == nil {
		return fmt.Errorf("parameter value is required, but nil was provided")
	}
	switch value.(type) {
	case cdktf.IResolvable:
		// ok
	case *[]*SloObjectiveCountMetricsGoodTotalLightstep:
		value := value.(*[]*SloObjectiveCountMetricsGoodTotalLightstep)
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	case []*SloObjectiveCountMetricsGoodTotalLightstep:
		value_ := value.([]*SloObjectiveCountMetricsGoodTotalLightstep)
		value := &value_
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	default:
		if !_jsii_.IsAnonymousProxy(value) {
			return fmt.Errorf("parameter value must be one of the allowed types: cdktf.IResolvable, *[]*SloObjectiveCountMetricsGoodTotalLightstep; received %#v (a %T)", value, value)
		}
	}

	return nil
}

func (s *jsiiProxy_SloObjectiveCountMetricsGoodTotalOutputReference) validatePutLogicMonitorParameters(value interface{}) error {
	if value == nil {
		return fmt.Errorf("parameter value is required, but nil was provided")
	}
	switch value.(type) {
	case cdktf.IResolvable:
		// ok
	case *[]*SloObjectiveCountMetricsGoodTotalLogicMonitor:
		value := value.(*[]*SloObjectiveCountMetricsGoodTotalLogicMonitor)
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	case []*SloObjectiveCountMetricsGoodTotalLogicMonitor:
		value_ := value.([]*SloObjectiveCountMetricsGoodTotalLogicMonitor)
		value := &value_
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	default:
		if !_jsii_.IsAnonymousProxy(value) {
			return fmt.Errorf("parameter value must be one of the allowed types: cdktf.IResolvable, *[]*SloObjectiveCountMetricsGoodTotalLogicMonitor; received %#v (a %T)", value, value)
		}
	}

	return nil
}

func (s *jsiiProxy_SloObjectiveCountMetricsGoodTotalOutputReference) validatePutNewrelicParameters(value interface{}) error {
	if value == nil {
		return fmt.Errorf("parameter value is required, but nil was provided")
	}
	switch value.(type) {
	case cdktf.IResolvable:
		// ok
	case *[]*SloObjectiveCountMetricsGoodTotalNewrelic:
		value := value.(*[]*SloObjectiveCountMetricsGoodTotalNewrelic)
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	case []*SloObjectiveCountMetricsGoodTotalNewrelic:
		value_ := value.([]*SloObjectiveCountMetricsGoodTotalNewrelic)
		value := &value_
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	default:
		if !_jsii_.IsAnonymousProxy(value) {
			return fmt.Errorf("parameter value must be one of the allowed types: cdktf.IResolvable, *[]*SloObjectiveCountMetricsGoodTotalNewrelic; received %#v (a %T)", value, value)
		}
	}

	return nil
}

func (s *jsiiProxy_SloObjectiveCountMetricsGoodTotalOutputReference) validatePutOpentsdbParameters(value interface{}) error {
	if value == nil {
		return fmt.Errorf("parameter value is required, but nil was provided")
	}
	switch value.(type) {
	case cdktf.IResolvable:
		// ok
	case *[]*SloObjectiveCountMetricsGoodTotalOpentsdb:
		value := value.(*[]*SloObjectiveCountMetricsGoodTotalOpentsdb)
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	case []*SloObjectiveCountMetricsGoodTotalOpentsdb:
		value_ := value.([]*SloObjectiveCountMetricsGoodTotalOpentsdb)
		value := &value_
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	default:
		if !_jsii_.IsAnonymousProxy(value) {
			return fmt.Errorf("parameter value must be one of the allowed types: cdktf.IResolvable, *[]*SloObjectiveCountMetricsGoodTotalOpentsdb; received %#v (a %T)", value, value)
		}
	}

	return nil
}

func (s *jsiiProxy_SloObjectiveCountMetricsGoodTotalOutputReference) validatePutPingdomParameters(value interface{}) error {
	if value == nil {
		return fmt.Errorf("parameter value is required, but nil was provided")
	}
	switch value.(type) {
	case cdktf.IResolvable:
		// ok
	case *[]*SloObjectiveCountMetricsGoodTotalPingdom:
		value := value.(*[]*SloObjectiveCountMetricsGoodTotalPingdom)
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	case []*SloObjectiveCountMetricsGoodTotalPingdom:
		value_ := value.([]*SloObjectiveCountMetricsGoodTotalPingdom)
		value := &value_
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	default:
		if !_jsii_.IsAnonymousProxy(value) {
			return fmt.Errorf("parameter value must be one of the allowed types: cdktf.IResolvable, *[]*SloObjectiveCountMetricsGoodTotalPingdom; received %#v (a %T)", value, value)
		}
	}

	return nil
}

func (s *jsiiProxy_SloObjectiveCountMetricsGoodTotalOutputReference) validatePutPrometheusParameters(value interface{}) error {
	if value == nil {
		return fmt.Errorf("parameter value is required, but nil was provided")
	}
	switch value.(type) {
	case cdktf.IResolvable:
		// ok
	case *[]*SloObjectiveCountMetricsGoodTotalPrometheus:
		value := value.(*[]*SloObjectiveCountMetricsGoodTotalPrometheus)
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	case []*SloObjectiveCountMetricsGoodTotalPrometheus:
		value_ := value.([]*SloObjectiveCountMetricsGoodTotalPrometheus)
		value := &value_
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	default:
		if !_jsii_.IsAnonymousProxy(value) {
			return fmt.Errorf("parameter value must be one of the allowed types: cdktf.IResolvable, *[]*SloObjectiveCountMetricsGoodTotalPrometheus; received %#v (a %T)", value, value)
		}
	}

	return nil
}

func (s *jsiiProxy_SloObjectiveCountMetricsGoodTotalOutputReference) validatePutRedshiftParameters(value interface{}) error {
	if value == nil {
		return fmt.Errorf("parameter value is required, but nil was provided")
	}
	switch value.(type) {
	case cdktf.IResolvable:
		// ok
	case *[]*SloObjectiveCountMetricsGoodTotalRedshift:
		value := value.(*[]*SloObjectiveCountMetricsGoodTotalRedshift)
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	case []*SloObjectiveCountMetricsGoodTotalRedshift:
		value_ := value.([]*SloObjectiveCountMetricsGoodTotalRedshift)
		value := &value_
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	default:
		if !_jsii_.IsAnonymousProxy(value) {
			return fmt.Errorf("parameter value must be one of the allowed types: cdktf.IResolvable, *[]*SloObjectiveCountMetricsGoodTotalRedshift; received %#v (a %T)", value, value)
		}
	}

	return nil
}

func (s *jsiiProxy_SloObjectiveCountMetricsGoodTotalOutputReference) validatePutSplunkParameters(value interface{}) error {
	if value == nil {
		return fmt.Errorf("parameter value is required, but nil was provided")
	}
	switch value.(type) {
	case cdktf.IResolvable:
		// ok
	case *[]*SloObjectiveCountMetricsGoodTotalSplunk:
		value := value.(*[]*SloObjectiveCountMetricsGoodTotalSplunk)
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	case []*SloObjectiveCountMetricsGoodTotalSplunk:
		value_ := value.([]*SloObjectiveCountMetricsGoodTotalSplunk)
		value := &value_
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	default:
		if !_jsii_.IsAnonymousProxy(value) {
			return fmt.Errorf("parameter value must be one of the allowed types: cdktf.IResolvable, *[]*SloObjectiveCountMetricsGoodTotalSplunk; received %#v (a %T)", value, value)
		}
	}

	return nil
}

func (s *jsiiProxy_SloObjectiveCountMetricsGoodTotalOutputReference) validatePutSplunkObservabilityParameters(value interface{}) error {
	if value == nil {
		return fmt.Errorf("parameter value is required, but nil was provided")
	}
	switch value.(type) {
	case cdktf.IResolvable:
		// ok
	case *[]*SloObjectiveCountMetricsGoodTotalSplunkObservability:
		value := value.(*[]*SloObjectiveCountMetricsGoodTotalSplunkObservability)
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	case []*SloObjectiveCountMetricsGoodTotalSplunkObservability:
		value_ := value.([]*SloObjectiveCountMetricsGoodTotalSplunkObservability)
		value := &value_
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	default:
		if !_jsii_.IsAnonymousProxy(value) {
			return fmt.Errorf("parameter value must be one of the allowed types: cdktf.IResolvable, *[]*SloObjectiveCountMetricsGoodTotalSplunkObservability; received %#v (a %T)", value, value)
		}
	}

	return nil
}

func (s *jsiiProxy_SloObjectiveCountMetricsGoodTotalOutputReference) validatePutSumologicParameters(value interface{}) error {
	if value == nil {
		return fmt.Errorf("parameter value is required, but nil was provided")
	}
	switch value.(type) {
	case cdktf.IResolvable:
		// ok
	case *[]*SloObjectiveCountMetricsGoodTotalSumologic:
		value := value.(*[]*SloObjectiveCountMetricsGoodTotalSumologic)
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	case []*SloObjectiveCountMetricsGoodTotalSumologic:
		value_ := value.([]*SloObjectiveCountMetricsGoodTotalSumologic)
		value := &value_
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	default:
		if !_jsii_.IsAnonymousProxy(value) {
			return fmt.Errorf("parameter value must be one of the allowed types: cdktf.IResolvable, *[]*SloObjectiveCountMetricsGoodTotalSumologic; received %#v (a %T)", value, value)
		}
	}

	return nil
}

func (s *jsiiProxy_SloObjectiveCountMetricsGoodTotalOutputReference) validatePutThousandeyesParameters(value interface{}) error {
	if value == nil {
		return fmt.Errorf("parameter value is required, but nil was provided")
	}
	switch value.(type) {
	case cdktf.IResolvable:
		// ok
	case *[]*SloObjectiveCountMetricsGoodTotalThousandeyes:
		value := value.(*[]*SloObjectiveCountMetricsGoodTotalThousandeyes)
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	case []*SloObjectiveCountMetricsGoodTotalThousandeyes:
		value_ := value.([]*SloObjectiveCountMetricsGoodTotalThousandeyes)
		value := &value_
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	default:
		if !_jsii_.IsAnonymousProxy(value) {
			return fmt.Errorf("parameter value must be one of the allowed types: cdktf.IResolvable, *[]*SloObjectiveCountMetricsGoodTotalThousandeyes; received %#v (a %T)", value, value)
		}
	}

	return nil
}

func (s *jsiiProxy_SloObjectiveCountMetricsGoodTotalOutputReference) validateResolveParameters(_context cdktf.IResolveContext) error {
	if _context == nil {
		return fmt.Errorf("parameter _context is required, but nil was provided")
	}

	return nil
}

func (j *jsiiProxy_SloObjectiveCountMetricsGoodTotalOutputReference) validateSetComplexObjectIndexParameters(val interface{}) error {
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

func (j *jsiiProxy_SloObjectiveCountMetricsGoodTotalOutputReference) validateSetComplexObjectIsFromSetParameters(val *bool) error {
	if val == nil {
		return fmt.Errorf("parameter val is required, but nil was provided")
	}

	return nil
}

func (j *jsiiProxy_SloObjectiveCountMetricsGoodTotalOutputReference) validateSetInternalValueParameters(val interface{}) error {
	switch val.(type) {
	case cdktf.IResolvable:
		// ok
	case *SloObjectiveCountMetricsGoodTotal:
		val := val.(*SloObjectiveCountMetricsGoodTotal)
		if err := _jsii_.ValidateStruct(val, func() string { return "parameter val" }); err != nil {
			return err
		}
	case SloObjectiveCountMetricsGoodTotal:
		val_ := val.(SloObjectiveCountMetricsGoodTotal)
		val := &val_
		if err := _jsii_.ValidateStruct(val, func() string { return "parameter val" }); err != nil {
			return err
		}
	default:
		if !_jsii_.IsAnonymousProxy(val) {
			return fmt.Errorf("parameter val must be one of the allowed types: cdktf.IResolvable, *SloObjectiveCountMetricsGoodTotal; received %#v (a %T)", val, val)
		}
	}

	return nil
}

func (j *jsiiProxy_SloObjectiveCountMetricsGoodTotalOutputReference) validateSetTerraformAttributeParameters(val *string) error {
	if val == nil {
		return fmt.Errorf("parameter val is required, but nil was provided")
	}

	return nil
}

func (j *jsiiProxy_SloObjectiveCountMetricsGoodTotalOutputReference) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	if val == nil {
		return fmt.Errorf("parameter val is required, but nil was provided")
	}

	return nil
}

func validateNewSloObjectiveCountMetricsGoodTotalOutputReferenceParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) error {
	if terraformResource == nil {
		return fmt.Errorf("parameter terraformResource is required, but nil was provided")
	}

	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	if complexObjectIndex == nil {
		return fmt.Errorf("parameter complexObjectIndex is required, but nil was provided")
	}

	if complexObjectIsFromSet == nil {
		return fmt.Errorf("parameter complexObjectIsFromSet is required, but nil was provided")
	}

	return nil
}

