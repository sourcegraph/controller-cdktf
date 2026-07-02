//go:build !no_runtime_type_checking

package chronicledashboardchart

import (
	"fmt"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

func (c *jsiiProxy_ChronicleDashboardChartDashboardChartDrillDownConfigRightDrillDownsCustomSettingsFilterDashboardFiltersList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	if mapKeyAttributeName == nil {
		return fmt.Errorf("parameter mapKeyAttributeName is required, but nil was provided")
	}

	return nil
}

func (c *jsiiProxy_ChronicleDashboardChartDashboardChartDrillDownConfigRightDrillDownsCustomSettingsFilterDashboardFiltersList) validateGetParameters(index *float64) error {
	if index == nil {
		return fmt.Errorf("parameter index is required, but nil was provided")
	}

	return nil
}

func (c *jsiiProxy_ChronicleDashboardChartDashboardChartDrillDownConfigRightDrillDownsCustomSettingsFilterDashboardFiltersList) validateResolveParameters(_context cdktf.IResolveContext) error {
	if _context == nil {
		return fmt.Errorf("parameter _context is required, but nil was provided")
	}

	return nil
}

func (j *jsiiProxy_ChronicleDashboardChartDashboardChartDrillDownConfigRightDrillDownsCustomSettingsFilterDashboardFiltersList) validateSetInternalValueParameters(val any) error {
	switch val.(type) {
	case cdktf.IResolvable:
		// ok
	case *[]*ChronicleDashboardChartDashboardChartDrillDownConfigRightDrillDownsCustomSettingsFilterDashboardFilters:
		val := val.(*[]*ChronicleDashboardChartDashboardChartDrillDownConfigRightDrillDownsCustomSettingsFilterDashboardFilters)
		for idx_97dfc6, v := range *val {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter val[%#v]", idx_97dfc6) }); err != nil {
				return err
			}
		}
	case []*ChronicleDashboardChartDashboardChartDrillDownConfigRightDrillDownsCustomSettingsFilterDashboardFilters:
		val_ := val.([]*ChronicleDashboardChartDashboardChartDrillDownConfigRightDrillDownsCustomSettingsFilterDashboardFilters)
		val := &val_
		for idx_97dfc6, v := range *val {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter val[%#v]", idx_97dfc6) }); err != nil {
				return err
			}
		}
	default:
		if !_jsii_.IsAnonymousProxy(val) {
			return fmt.Errorf("parameter val must be one of the allowed types: cdktf.IResolvable, *[]*ChronicleDashboardChartDashboardChartDrillDownConfigRightDrillDownsCustomSettingsFilterDashboardFilters; received %#v (a %T)", val, val)
		}
	}

	return nil
}

func (j *jsiiProxy_ChronicleDashboardChartDashboardChartDrillDownConfigRightDrillDownsCustomSettingsFilterDashboardFiltersList) validateSetTerraformAttributeParameters(val *string) error {
	if val == nil {
		return fmt.Errorf("parameter val is required, but nil was provided")
	}

	return nil
}

func (j *jsiiProxy_ChronicleDashboardChartDashboardChartDrillDownConfigRightDrillDownsCustomSettingsFilterDashboardFiltersList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	if val == nil {
		return fmt.Errorf("parameter val is required, but nil was provided")
	}

	return nil
}

func (j *jsiiProxy_ChronicleDashboardChartDashboardChartDrillDownConfigRightDrillDownsCustomSettingsFilterDashboardFiltersList) validateSetWrapsSetParameters(val *bool) error {
	if val == nil {
		return fmt.Errorf("parameter val is required, but nil was provided")
	}

	return nil
}

func validateNewChronicleDashboardChartDashboardChartDrillDownConfigRightDrillDownsCustomSettingsFilterDashboardFiltersListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	if terraformResource == nil {
		return fmt.Errorf("parameter terraformResource is required, but nil was provided")
	}

	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	if wrapsSet == nil {
		return fmt.Errorf("parameter wrapsSet is required, but nil was provided")
	}

	return nil
}
