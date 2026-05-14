//go:build !no_runtime_type_checking

package backupdrrestoreworkload

import (
	"fmt"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

func (b *jsiiProxy_BackupDrRestoreWorkloadComputeInstanceRestorePropertiesDisksOutputReference) validateGetAnyMapAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (b *jsiiProxy_BackupDrRestoreWorkloadComputeInstanceRestorePropertiesDisksOutputReference) validateGetBooleanAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (b *jsiiProxy_BackupDrRestoreWorkloadComputeInstanceRestorePropertiesDisksOutputReference) validateGetBooleanMapAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (b *jsiiProxy_BackupDrRestoreWorkloadComputeInstanceRestorePropertiesDisksOutputReference) validateGetListAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (b *jsiiProxy_BackupDrRestoreWorkloadComputeInstanceRestorePropertiesDisksOutputReference) validateGetNumberAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (b *jsiiProxy_BackupDrRestoreWorkloadComputeInstanceRestorePropertiesDisksOutputReference) validateGetNumberListAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (b *jsiiProxy_BackupDrRestoreWorkloadComputeInstanceRestorePropertiesDisksOutputReference) validateGetNumberMapAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (b *jsiiProxy_BackupDrRestoreWorkloadComputeInstanceRestorePropertiesDisksOutputReference) validateGetStringAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (b *jsiiProxy_BackupDrRestoreWorkloadComputeInstanceRestorePropertiesDisksOutputReference) validateGetStringMapAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (b *jsiiProxy_BackupDrRestoreWorkloadComputeInstanceRestorePropertiesDisksOutputReference) validateInterpolationForAttributeParameters(property *string) error {
	if property == nil {
		return fmt.Errorf("parameter property is required, but nil was provided")
	}

	return nil
}

func (b *jsiiProxy_BackupDrRestoreWorkloadComputeInstanceRestorePropertiesDisksOutputReference) validatePutDiskEncryptionKeyParameters(value *BackupDrRestoreWorkloadComputeInstanceRestorePropertiesDisksDiskEncryptionKey) error {
	if value == nil {
		return fmt.Errorf("parameter value is required, but nil was provided")
	}
	if err := _jsii_.ValidateStruct(value, func() string { return "parameter value" }); err != nil {
		return err
	}

	return nil
}

func (b *jsiiProxy_BackupDrRestoreWorkloadComputeInstanceRestorePropertiesDisksOutputReference) validatePutGuestOsFeatureParameters(value interface{}) error {
	if value == nil {
		return fmt.Errorf("parameter value is required, but nil was provided")
	}
	switch value.(type) {
	case cdktf.IResolvable:
		// ok
	case *[]*BackupDrRestoreWorkloadComputeInstanceRestorePropertiesDisksGuestOsFeature:
		value := value.(*[]*BackupDrRestoreWorkloadComputeInstanceRestorePropertiesDisksGuestOsFeature)
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	case []*BackupDrRestoreWorkloadComputeInstanceRestorePropertiesDisksGuestOsFeature:
		value_ := value.([]*BackupDrRestoreWorkloadComputeInstanceRestorePropertiesDisksGuestOsFeature)
		value := &value_
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	default:
		if !_jsii_.IsAnonymousProxy(value) {
			return fmt.Errorf("parameter value must be one of the allowed types: cdktf.IResolvable, *[]*BackupDrRestoreWorkloadComputeInstanceRestorePropertiesDisksGuestOsFeature; received %#v (a %T)", value, value)
		}
	}

	return nil
}

func (b *jsiiProxy_BackupDrRestoreWorkloadComputeInstanceRestorePropertiesDisksOutputReference) validatePutInitializeParamsParameters(value *BackupDrRestoreWorkloadComputeInstanceRestorePropertiesDisksInitializeParams) error {
	if value == nil {
		return fmt.Errorf("parameter value is required, but nil was provided")
	}
	if err := _jsii_.ValidateStruct(value, func() string { return "parameter value" }); err != nil {
		return err
	}

	return nil
}

func (b *jsiiProxy_BackupDrRestoreWorkloadComputeInstanceRestorePropertiesDisksOutputReference) validateResolveParameters(_context cdktf.IResolveContext) error {
	if _context == nil {
		return fmt.Errorf("parameter _context is required, but nil was provided")
	}

	return nil
}

func (j *jsiiProxy_BackupDrRestoreWorkloadComputeInstanceRestorePropertiesDisksOutputReference) validateSetAutoDeleteParameters(val interface{}) error {
	if val == nil {
		return fmt.Errorf("parameter val is required, but nil was provided")
	}
	switch val.(type) {
	case *bool:
		// ok
	case bool:
		// ok
	case cdktf.IResolvable:
		// ok
	default:
		if !_jsii_.IsAnonymousProxy(val) {
			return fmt.Errorf("parameter val must be one of the allowed types: *bool, cdktf.IResolvable; received %#v (a %T)", val, val)
		}
	}

	return nil
}

func (j *jsiiProxy_BackupDrRestoreWorkloadComputeInstanceRestorePropertiesDisksOutputReference) validateSetBootParameters(val interface{}) error {
	if val == nil {
		return fmt.Errorf("parameter val is required, but nil was provided")
	}
	switch val.(type) {
	case *bool:
		// ok
	case bool:
		// ok
	case cdktf.IResolvable:
		// ok
	default:
		if !_jsii_.IsAnonymousProxy(val) {
			return fmt.Errorf("parameter val must be one of the allowed types: *bool, cdktf.IResolvable; received %#v (a %T)", val, val)
		}
	}

	return nil
}

func (j *jsiiProxy_BackupDrRestoreWorkloadComputeInstanceRestorePropertiesDisksOutputReference) validateSetComplexObjectIndexParameters(val interface{}) error {
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

func (j *jsiiProxy_BackupDrRestoreWorkloadComputeInstanceRestorePropertiesDisksOutputReference) validateSetComplexObjectIsFromSetParameters(val *bool) error {
	if val == nil {
		return fmt.Errorf("parameter val is required, but nil was provided")
	}

	return nil
}

func (j *jsiiProxy_BackupDrRestoreWorkloadComputeInstanceRestorePropertiesDisksOutputReference) validateSetDeviceNameParameters(val *string) error {
	if val == nil {
		return fmt.Errorf("parameter val is required, but nil was provided")
	}

	return nil
}

func (j *jsiiProxy_BackupDrRestoreWorkloadComputeInstanceRestorePropertiesDisksOutputReference) validateSetDiskInterfaceParameters(val *string) error {
	if val == nil {
		return fmt.Errorf("parameter val is required, but nil was provided")
	}

	return nil
}

func (j *jsiiProxy_BackupDrRestoreWorkloadComputeInstanceRestorePropertiesDisksOutputReference) validateSetDiskSizeGbParameters(val *float64) error {
	if val == nil {
		return fmt.Errorf("parameter val is required, but nil was provided")
	}

	return nil
}

func (j *jsiiProxy_BackupDrRestoreWorkloadComputeInstanceRestorePropertiesDisksOutputReference) validateSetDiskTypeParameters(val *string) error {
	if val == nil {
		return fmt.Errorf("parameter val is required, but nil was provided")
	}

	return nil
}

func (j *jsiiProxy_BackupDrRestoreWorkloadComputeInstanceRestorePropertiesDisksOutputReference) validateSetIndexParameters(val *float64) error {
	if val == nil {
		return fmt.Errorf("parameter val is required, but nil was provided")
	}

	return nil
}

func (j *jsiiProxy_BackupDrRestoreWorkloadComputeInstanceRestorePropertiesDisksOutputReference) validateSetInternalValueParameters(val interface{}) error {
	switch val.(type) {
	case cdktf.IResolvable:
		// ok
	case *BackupDrRestoreWorkloadComputeInstanceRestorePropertiesDisks:
		val := val.(*BackupDrRestoreWorkloadComputeInstanceRestorePropertiesDisks)
		if err := _jsii_.ValidateStruct(val, func() string { return "parameter val" }); err != nil {
			return err
		}
	case BackupDrRestoreWorkloadComputeInstanceRestorePropertiesDisks:
		val_ := val.(BackupDrRestoreWorkloadComputeInstanceRestorePropertiesDisks)
		val := &val_
		if err := _jsii_.ValidateStruct(val, func() string { return "parameter val" }); err != nil {
			return err
		}
	default:
		if !_jsii_.IsAnonymousProxy(val) {
			return fmt.Errorf("parameter val must be one of the allowed types: cdktf.IResolvable, *BackupDrRestoreWorkloadComputeInstanceRestorePropertiesDisks; received %#v (a %T)", val, val)
		}
	}

	return nil
}

func (j *jsiiProxy_BackupDrRestoreWorkloadComputeInstanceRestorePropertiesDisksOutputReference) validateSetKindParameters(val *string) error {
	if val == nil {
		return fmt.Errorf("parameter val is required, but nil was provided")
	}

	return nil
}

func (j *jsiiProxy_BackupDrRestoreWorkloadComputeInstanceRestorePropertiesDisksOutputReference) validateSetLicenseParameters(val *[]*string) error {
	if val == nil {
		return fmt.Errorf("parameter val is required, but nil was provided")
	}

	return nil
}

func (j *jsiiProxy_BackupDrRestoreWorkloadComputeInstanceRestorePropertiesDisksOutputReference) validateSetModeParameters(val *string) error {
	if val == nil {
		return fmt.Errorf("parameter val is required, but nil was provided")
	}

	return nil
}

func (j *jsiiProxy_BackupDrRestoreWorkloadComputeInstanceRestorePropertiesDisksOutputReference) validateSetSavedStateParameters(val *string) error {
	if val == nil {
		return fmt.Errorf("parameter val is required, but nil was provided")
	}

	return nil
}

func (j *jsiiProxy_BackupDrRestoreWorkloadComputeInstanceRestorePropertiesDisksOutputReference) validateSetSourceParameters(val *string) error {
	if val == nil {
		return fmt.Errorf("parameter val is required, but nil was provided")
	}

	return nil
}

func (j *jsiiProxy_BackupDrRestoreWorkloadComputeInstanceRestorePropertiesDisksOutputReference) validateSetTerraformAttributeParameters(val *string) error {
	if val == nil {
		return fmt.Errorf("parameter val is required, but nil was provided")
	}

	return nil
}

func (j *jsiiProxy_BackupDrRestoreWorkloadComputeInstanceRestorePropertiesDisksOutputReference) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	if val == nil {
		return fmt.Errorf("parameter val is required, but nil was provided")
	}

	return nil
}

func (j *jsiiProxy_BackupDrRestoreWorkloadComputeInstanceRestorePropertiesDisksOutputReference) validateSetTypeParameters(val *string) error {
	if val == nil {
		return fmt.Errorf("parameter val is required, but nil was provided")
	}

	return nil
}

func validateNewBackupDrRestoreWorkloadComputeInstanceRestorePropertiesDisksOutputReferenceParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) error {
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

