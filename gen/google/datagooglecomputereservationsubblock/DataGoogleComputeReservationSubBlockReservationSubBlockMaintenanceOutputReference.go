package datagooglecomputereservationsubblock

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/google/jsii"

	"github.com/open-constructs/cdk-terrain-go/cdktn"
	"github.com/sourcegraph/controller-cdktf/gen/google/datagooglecomputereservationsubblock/internal"
)

type DataGoogleComputeReservationSubBlockReservationSubBlockMaintenanceOutputReference interface {
	cdktn.ComplexObject
	// the index of the complex object in a list.
	// Experimental.
	ComplexObjectIndex() interface{}
	// Experimental.
	SetComplexObjectIndex(val interface{})
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
	InternalValue() *DataGoogleComputeReservationSubBlockReservationSubBlockMaintenance
	SetInternalValue(val *DataGoogleComputeReservationSubBlockReservationSubBlockMaintenance)
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
	TerraformResource() cdktn.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktn.IInterpolatingParent)
	// Experimental.
	ComputeFqn() *string
	// Experimental.
	GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{}
	// Experimental.
	GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable
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
	InterpolationAsList() cdktn.IResolvable
	// Experimental.
	InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for DataGoogleComputeReservationSubBlockReservationSubBlockMaintenanceOutputReference
type jsiiProxy_DataGoogleComputeReservationSubBlockReservationSubBlockMaintenanceOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_DataGoogleComputeReservationSubBlockReservationSubBlockMaintenanceOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleComputeReservationSubBlockReservationSubBlockMaintenanceOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleComputeReservationSubBlockReservationSubBlockMaintenanceOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleComputeReservationSubBlockReservationSubBlockMaintenanceOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleComputeReservationSubBlockReservationSubBlockMaintenanceOutputReference) InstanceMaintenanceOngoingCount() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"instanceMaintenanceOngoingCount",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleComputeReservationSubBlockReservationSubBlockMaintenanceOutputReference) InstanceMaintenancePendingCount() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"instanceMaintenancePendingCount",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleComputeReservationSubBlockReservationSubBlockMaintenanceOutputReference) InternalValue() *DataGoogleComputeReservationSubBlockReservationSubBlockMaintenance {
	var returns *DataGoogleComputeReservationSubBlockReservationSubBlockMaintenance
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleComputeReservationSubBlockReservationSubBlockMaintenanceOutputReference) MaintenanceOngoingCount() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maintenanceOngoingCount",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleComputeReservationSubBlockReservationSubBlockMaintenanceOutputReference) MaintenancePendingCount() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maintenancePendingCount",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleComputeReservationSubBlockReservationSubBlockMaintenanceOutputReference) SchedulingType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"schedulingType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleComputeReservationSubBlockReservationSubBlockMaintenanceOutputReference) SubblockInfraMaintenanceOngoingCount() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"subblockInfraMaintenanceOngoingCount",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleComputeReservationSubBlockReservationSubBlockMaintenanceOutputReference) SubblockInfraMaintenancePendingCount() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"subblockInfraMaintenancePendingCount",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleComputeReservationSubBlockReservationSubBlockMaintenanceOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleComputeReservationSubBlockReservationSubBlockMaintenanceOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewDataGoogleComputeReservationSubBlockReservationSubBlockMaintenanceOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) DataGoogleComputeReservationSubBlockReservationSubBlockMaintenanceOutputReference {
	_init_.Initialize()

	if err := validateNewDataGoogleComputeReservationSubBlockReservationSubBlockMaintenanceOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_DataGoogleComputeReservationSubBlockReservationSubBlockMaintenanceOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-google.dataGoogleComputeReservationSubBlock.DataGoogleComputeReservationSubBlockReservationSubBlockMaintenanceOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewDataGoogleComputeReservationSubBlockReservationSubBlockMaintenanceOutputReference_Override(d DataGoogleComputeReservationSubBlockReservationSubBlockMaintenanceOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-google.dataGoogleComputeReservationSubBlock.DataGoogleComputeReservationSubBlockReservationSubBlockMaintenanceOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		d,
	)
}

func (j *jsiiProxy_DataGoogleComputeReservationSubBlockReservationSubBlockMaintenanceOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_DataGoogleComputeReservationSubBlockReservationSubBlockMaintenanceOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_DataGoogleComputeReservationSubBlockReservationSubBlockMaintenanceOutputReference)SetInternalValue(val *DataGoogleComputeReservationSubBlockReservationSubBlockMaintenance) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_DataGoogleComputeReservationSubBlockReservationSubBlockMaintenanceOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DataGoogleComputeReservationSubBlockReservationSubBlockMaintenanceOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (d *jsiiProxy_DataGoogleComputeReservationSubBlockReservationSubBlockMaintenanceOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataGoogleComputeReservationSubBlockReservationSubBlockMaintenanceOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := d.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		d,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataGoogleComputeReservationSubBlockReservationSubBlockMaintenanceOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := d.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		d,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataGoogleComputeReservationSubBlockReservationSubBlockMaintenanceOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := d.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		d,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataGoogleComputeReservationSubBlockReservationSubBlockMaintenanceOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := d.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		d,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataGoogleComputeReservationSubBlockReservationSubBlockMaintenanceOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := d.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		d,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataGoogleComputeReservationSubBlockReservationSubBlockMaintenanceOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := d.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		d,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataGoogleComputeReservationSubBlockReservationSubBlockMaintenanceOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := d.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		d,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataGoogleComputeReservationSubBlockReservationSubBlockMaintenanceOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := d.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		d,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataGoogleComputeReservationSubBlockReservationSubBlockMaintenanceOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := d.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		d,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataGoogleComputeReservationSubBlockReservationSubBlockMaintenanceOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataGoogleComputeReservationSubBlockReservationSubBlockMaintenanceOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := d.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataGoogleComputeReservationSubBlockReservationSubBlockMaintenanceOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
	if err := d.validateResolveParameters(context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		d,
		"resolve",
		[]interface{}{context},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataGoogleComputeReservationSubBlockReservationSubBlockMaintenanceOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

