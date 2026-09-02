package computereservation

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/google/jsii"

	"github.com/open-constructs/cdk-terrain-go/cdktn"
	"github.com/sourcegraph/controller-cdktf/gen/google/computereservation/internal"
)

type ComputeReservationResourceStatusReservationMaintenanceOutputReference interface {
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
	InternalValue() *ComputeReservationResourceStatusReservationMaintenance
	SetInternalValue(val *ComputeReservationResourceStatusReservationMaintenance)
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
	UpcomingGroupMaintenance() ComputeReservationResourceStatusReservationMaintenanceUpcomingGroupMaintenanceList
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

// The jsii proxy struct for ComputeReservationResourceStatusReservationMaintenanceOutputReference
type jsiiProxy_ComputeReservationResourceStatusReservationMaintenanceOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_ComputeReservationResourceStatusReservationMaintenanceOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeReservationResourceStatusReservationMaintenanceOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeReservationResourceStatusReservationMaintenanceOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeReservationResourceStatusReservationMaintenanceOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeReservationResourceStatusReservationMaintenanceOutputReference) InstanceMaintenanceOngoingCount() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"instanceMaintenanceOngoingCount",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeReservationResourceStatusReservationMaintenanceOutputReference) InstanceMaintenancePendingCount() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"instanceMaintenancePendingCount",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeReservationResourceStatusReservationMaintenanceOutputReference) InternalValue() *ComputeReservationResourceStatusReservationMaintenance {
	var returns *ComputeReservationResourceStatusReservationMaintenance
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeReservationResourceStatusReservationMaintenanceOutputReference) MaintenanceOngoingCount() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maintenanceOngoingCount",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeReservationResourceStatusReservationMaintenanceOutputReference) MaintenancePendingCount() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maintenancePendingCount",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeReservationResourceStatusReservationMaintenanceOutputReference) SchedulingType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"schedulingType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeReservationResourceStatusReservationMaintenanceOutputReference) SubblockInfraMaintenanceOngoingCount() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"subblockInfraMaintenanceOngoingCount",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeReservationResourceStatusReservationMaintenanceOutputReference) SubblockInfraMaintenancePendingCount() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"subblockInfraMaintenancePendingCount",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeReservationResourceStatusReservationMaintenanceOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeReservationResourceStatusReservationMaintenanceOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeReservationResourceStatusReservationMaintenanceOutputReference) UpcomingGroupMaintenance() ComputeReservationResourceStatusReservationMaintenanceUpcomingGroupMaintenanceList {
	var returns ComputeReservationResourceStatusReservationMaintenanceUpcomingGroupMaintenanceList
	_jsii_.Get(
		j,
		"upcomingGroupMaintenance",
		&returns,
	)
	return returns
}


func NewComputeReservationResourceStatusReservationMaintenanceOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) ComputeReservationResourceStatusReservationMaintenanceOutputReference {
	_init_.Initialize()

	if err := validateNewComputeReservationResourceStatusReservationMaintenanceOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_ComputeReservationResourceStatusReservationMaintenanceOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-google.computeReservation.ComputeReservationResourceStatusReservationMaintenanceOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewComputeReservationResourceStatusReservationMaintenanceOutputReference_Override(c ComputeReservationResourceStatusReservationMaintenanceOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-google.computeReservation.ComputeReservationResourceStatusReservationMaintenanceOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		c,
	)
}

func (j *jsiiProxy_ComputeReservationResourceStatusReservationMaintenanceOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_ComputeReservationResourceStatusReservationMaintenanceOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_ComputeReservationResourceStatusReservationMaintenanceOutputReference)SetInternalValue(val *ComputeReservationResourceStatusReservationMaintenance) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_ComputeReservationResourceStatusReservationMaintenanceOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_ComputeReservationResourceStatusReservationMaintenanceOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (c *jsiiProxy_ComputeReservationResourceStatusReservationMaintenanceOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ComputeReservationResourceStatusReservationMaintenanceOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := c.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		c,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ComputeReservationResourceStatusReservationMaintenanceOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := c.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		c,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ComputeReservationResourceStatusReservationMaintenanceOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := c.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		c,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ComputeReservationResourceStatusReservationMaintenanceOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := c.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		c,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ComputeReservationResourceStatusReservationMaintenanceOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := c.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		c,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ComputeReservationResourceStatusReservationMaintenanceOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := c.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		c,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ComputeReservationResourceStatusReservationMaintenanceOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := c.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		c,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ComputeReservationResourceStatusReservationMaintenanceOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := c.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		c,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ComputeReservationResourceStatusReservationMaintenanceOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := c.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		c,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ComputeReservationResourceStatusReservationMaintenanceOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		c,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ComputeReservationResourceStatusReservationMaintenanceOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := c.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		c,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ComputeReservationResourceStatusReservationMaintenanceOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
	if err := c.validateResolveParameters(context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		c,
		"resolve",
		[]interface{}{context},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ComputeReservationResourceStatusReservationMaintenanceOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

