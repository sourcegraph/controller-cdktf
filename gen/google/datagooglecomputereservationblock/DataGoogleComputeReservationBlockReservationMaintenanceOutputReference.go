package datagooglecomputereservationblock

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/google/jsii"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/sourcegraph/controller-cdktf/gen/google/datagooglecomputereservationblock/internal"
)

type DataGoogleComputeReservationBlockReservationMaintenanceOutputReference interface {
	cdktf.ComplexObject
	// the index of the complex object in a list.
	// Experimental.
	ComplexObjectIndex() any
	// Experimental.
	SetComplexObjectIndex(val any)
	// set to true if this item is from inside a set and needs tolist() for accessing it set to "0" for single list items.
	// Experimental.
	ComplexObjectIsFromSet() *bool
	// Experimental.
	SetComplexObjectIsFromSet(val *bool)
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	// Experimental.
	Fqn() *string
	InstanceMaintenanceOngoingCount() *float64
	InstanceMaintenancePendingCount() *float64
	InternalValue() *DataGoogleComputeReservationBlockReservationMaintenance
	SetInternalValue(val *DataGoogleComputeReservationBlockReservationMaintenance)
	MaintenanceOngoingCount() *float64
	MaintenancePendingCount() *float64
	SchedulingType() *string
	SubblockInfraMaintenanceOngoingCount() *float64
	SubblockInfraMaintenancePendingCount() *float64
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	// Experimental.
	ComputeFqn() *string
	// Experimental.
	GetAnyMapAttribute(terraformAttribute *string) *map[string]any
	// Experimental.
	GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable
	// Experimental.
	GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool
	// Experimental.
	GetListAttribute(terraformAttribute *string) *[]*string
	// Experimental.
	GetNumberAttribute(terraformAttribute *string) *float64
	// Experimental.
	GetNumberListAttribute(terraformAttribute *string) *[]*float64
	// Experimental.
	GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64
	// Experimental.
	GetStringAttribute(terraformAttribute *string) *string
	// Experimental.
	GetStringMapAttribute(terraformAttribute *string) *map[string]*string
	// Experimental.
	InterpolationAsList() cdktf.IResolvable
	// Experimental.
	InterpolationForAttribute(property *string) cdktf.IResolvable
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) any
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for DataGoogleComputeReservationBlockReservationMaintenanceOutputReference
type jsiiProxy_DataGoogleComputeReservationBlockReservationMaintenanceOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_DataGoogleComputeReservationBlockReservationMaintenanceOutputReference) ComplexObjectIndex() any {
	var returns any
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleComputeReservationBlockReservationMaintenanceOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleComputeReservationBlockReservationMaintenanceOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleComputeReservationBlockReservationMaintenanceOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleComputeReservationBlockReservationMaintenanceOutputReference) InstanceMaintenanceOngoingCount() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"instanceMaintenanceOngoingCount",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleComputeReservationBlockReservationMaintenanceOutputReference) InstanceMaintenancePendingCount() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"instanceMaintenancePendingCount",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleComputeReservationBlockReservationMaintenanceOutputReference) InternalValue() *DataGoogleComputeReservationBlockReservationMaintenance {
	var returns *DataGoogleComputeReservationBlockReservationMaintenance
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleComputeReservationBlockReservationMaintenanceOutputReference) MaintenanceOngoingCount() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maintenanceOngoingCount",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleComputeReservationBlockReservationMaintenanceOutputReference) MaintenancePendingCount() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maintenancePendingCount",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleComputeReservationBlockReservationMaintenanceOutputReference) SchedulingType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"schedulingType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleComputeReservationBlockReservationMaintenanceOutputReference) SubblockInfraMaintenanceOngoingCount() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"subblockInfraMaintenanceOngoingCount",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleComputeReservationBlockReservationMaintenanceOutputReference) SubblockInfraMaintenancePendingCount() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"subblockInfraMaintenancePendingCount",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleComputeReservationBlockReservationMaintenanceOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleComputeReservationBlockReservationMaintenanceOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func NewDataGoogleComputeReservationBlockReservationMaintenanceOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) DataGoogleComputeReservationBlockReservationMaintenanceOutputReference {
	_init_.Initialize()

	if err := validateNewDataGoogleComputeReservationBlockReservationMaintenanceOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_DataGoogleComputeReservationBlockReservationMaintenanceOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-google.dataGoogleComputeReservationBlock.DataGoogleComputeReservationBlockReservationMaintenanceOutputReference",
		[]any{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewDataGoogleComputeReservationBlockReservationMaintenanceOutputReference_Override(d DataGoogleComputeReservationBlockReservationMaintenanceOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google.dataGoogleComputeReservationBlock.DataGoogleComputeReservationBlockReservationMaintenanceOutputReference",
		[]any{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		d,
	)
}

func (j *jsiiProxy_DataGoogleComputeReservationBlockReservationMaintenanceOutputReference) SetComplexObjectIndex(val any) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_DataGoogleComputeReservationBlockReservationMaintenanceOutputReference) SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_DataGoogleComputeReservationBlockReservationMaintenanceOutputReference) SetInternalValue(val *DataGoogleComputeReservationBlockReservationMaintenance) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_DataGoogleComputeReservationBlockReservationMaintenanceOutputReference) SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DataGoogleComputeReservationBlockReservationMaintenanceOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (d *jsiiProxy_DataGoogleComputeReservationBlockReservationMaintenanceOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataGoogleComputeReservationBlockReservationMaintenanceOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]any {
	if err := d.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]any

	_jsii_.Invoke(
		d,
		"getAnyMapAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataGoogleComputeReservationBlockReservationMaintenanceOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := d.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"getBooleanAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataGoogleComputeReservationBlockReservationMaintenanceOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := d.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		d,
		"getBooleanMapAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataGoogleComputeReservationBlockReservationMaintenanceOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := d.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		d,
		"getListAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataGoogleComputeReservationBlockReservationMaintenanceOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := d.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		d,
		"getNumberAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataGoogleComputeReservationBlockReservationMaintenanceOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := d.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		d,
		"getNumberListAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataGoogleComputeReservationBlockReservationMaintenanceOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := d.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		d,
		"getNumberMapAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataGoogleComputeReservationBlockReservationMaintenanceOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := d.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		d,
		"getStringAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataGoogleComputeReservationBlockReservationMaintenanceOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := d.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		d,
		"getStringMapAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataGoogleComputeReservationBlockReservationMaintenanceOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataGoogleComputeReservationBlockReservationMaintenanceOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	if err := d.validateInterpolationForAttributeParameters(property); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationForAttribute",
		[]any{property},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataGoogleComputeReservationBlockReservationMaintenanceOutputReference) Resolve(_context cdktf.IResolveContext) any {
	if err := d.validateResolveParameters(_context); err != nil {
		panic(err)
	}
	var returns any

	_jsii_.Invoke(
		d,
		"resolve",
		[]any{_context},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataGoogleComputeReservationBlockReservationMaintenanceOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}
