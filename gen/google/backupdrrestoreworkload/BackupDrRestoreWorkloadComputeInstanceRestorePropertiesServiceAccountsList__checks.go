//go:build !no_runtime_type_checking

package backupdrrestoreworkload

import (
	"fmt"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

func (b *jsiiProxy_BackupDrRestoreWorkloadComputeInstanceRestorePropertiesServiceAccountsList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	if mapKeyAttributeName == nil {
		return fmt.Errorf("parameter mapKeyAttributeName is required, but nil was provided")
	}

	return nil
}

func (b *jsiiProxy_BackupDrRestoreWorkloadComputeInstanceRestorePropertiesServiceAccountsList) validateGetParameters(index *float64) error {
	if index == nil {
		return fmt.Errorf("parameter index is required, but nil was provided")
	}

	return nil
}

func (b *jsiiProxy_BackupDrRestoreWorkloadComputeInstanceRestorePropertiesServiceAccountsList) validateResolveParameters(_context cdktf.IResolveContext) error {
	if _context == nil {
		return fmt.Errorf("parameter _context is required, but nil was provided")
	}

	return nil
}

func (j *jsiiProxy_BackupDrRestoreWorkloadComputeInstanceRestorePropertiesServiceAccountsList) validateSetInternalValueParameters(val interface{}) error {
	switch val.(type) {
	case cdktf.IResolvable:
		// ok
	case *[]*BackupDrRestoreWorkloadComputeInstanceRestorePropertiesServiceAccounts:
		val := val.(*[]*BackupDrRestoreWorkloadComputeInstanceRestorePropertiesServiceAccounts)
		for idx_97dfc6, v := range *val {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter val[%#v]", idx_97dfc6) }); err != nil {
				return err
			}
		}
	case []*BackupDrRestoreWorkloadComputeInstanceRestorePropertiesServiceAccounts:
		val_ := val.([]*BackupDrRestoreWorkloadComputeInstanceRestorePropertiesServiceAccounts)
		val := &val_
		for idx_97dfc6, v := range *val {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter val[%#v]", idx_97dfc6) }); err != nil {
				return err
			}
		}
	default:
		if !_jsii_.IsAnonymousProxy(val) {
			return fmt.Errorf("parameter val must be one of the allowed types: cdktf.IResolvable, *[]*BackupDrRestoreWorkloadComputeInstanceRestorePropertiesServiceAccounts; received %#v (a %T)", val, val)
		}
	}

	return nil
}

func (j *jsiiProxy_BackupDrRestoreWorkloadComputeInstanceRestorePropertiesServiceAccountsList) validateSetTerraformAttributeParameters(val *string) error {
	if val == nil {
		return fmt.Errorf("parameter val is required, but nil was provided")
	}

	return nil
}

func (j *jsiiProxy_BackupDrRestoreWorkloadComputeInstanceRestorePropertiesServiceAccountsList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	if val == nil {
		return fmt.Errorf("parameter val is required, but nil was provided")
	}

	return nil
}

func (j *jsiiProxy_BackupDrRestoreWorkloadComputeInstanceRestorePropertiesServiceAccountsList) validateSetWrapsSetParameters(val *bool) error {
	if val == nil {
		return fmt.Errorf("parameter val is required, but nil was provided")
	}

	return nil
}

func validateNewBackupDrRestoreWorkloadComputeInstanceRestorePropertiesServiceAccountsListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
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

