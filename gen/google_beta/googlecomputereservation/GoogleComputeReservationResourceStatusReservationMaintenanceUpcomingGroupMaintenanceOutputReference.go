package googlecomputereservation

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/google_beta/jsii"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/sourcegraph/controller-cdktf/gen/google_beta/googlecomputereservation/internal"
)

type GoogleComputeReservationResourceStatusReservationMaintenanceUpcomingGroupMaintenanceOutputReference interface {
	cdktf.ComplexObject
	CanReschedule() cdktf.IResolvable
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
	InternalValue() *GoogleComputeReservationResourceStatusReservationMaintenanceUpcomingGroupMaintenance
	SetInternalValue(val *GoogleComputeReservationResourceStatusReservationMaintenanceUpcomingGroupMaintenance)
	LatestWindowStartTime() *string
	MaintenanceOnShutdown() cdktf.IResolvable
	MaintenanceReasons() *[]*string
	MaintenanceStatus() *string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	Type() *string
	WindowEndTime() *string
	WindowStartTime() *string
	// Experimental.
	ComputeFqn() *string
	// Experimental.
	GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{}
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
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for GoogleComputeReservationResourceStatusReservationMaintenanceUpcomingGroupMaintenanceOutputReference
type jsiiProxy_GoogleComputeReservationResourceStatusReservationMaintenanceUpcomingGroupMaintenanceOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_GoogleComputeReservationResourceStatusReservationMaintenanceUpcomingGroupMaintenanceOutputReference) CanReschedule() cdktf.IResolvable {
	var returns cdktf.IResolvable
	_jsii_.Get(
		j,
		"canReschedule",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeReservationResourceStatusReservationMaintenanceUpcomingGroupMaintenanceOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeReservationResourceStatusReservationMaintenanceUpcomingGroupMaintenanceOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeReservationResourceStatusReservationMaintenanceUpcomingGroupMaintenanceOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeReservationResourceStatusReservationMaintenanceUpcomingGroupMaintenanceOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeReservationResourceStatusReservationMaintenanceUpcomingGroupMaintenanceOutputReference) InternalValue() *GoogleComputeReservationResourceStatusReservationMaintenanceUpcomingGroupMaintenance {
	var returns *GoogleComputeReservationResourceStatusReservationMaintenanceUpcomingGroupMaintenance
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeReservationResourceStatusReservationMaintenanceUpcomingGroupMaintenanceOutputReference) LatestWindowStartTime() *string {
	var returns *string
	_jsii_.Get(
		j,
		"latestWindowStartTime",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeReservationResourceStatusReservationMaintenanceUpcomingGroupMaintenanceOutputReference) MaintenanceOnShutdown() cdktf.IResolvable {
	var returns cdktf.IResolvable
	_jsii_.Get(
		j,
		"maintenanceOnShutdown",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeReservationResourceStatusReservationMaintenanceUpcomingGroupMaintenanceOutputReference) MaintenanceReasons() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"maintenanceReasons",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeReservationResourceStatusReservationMaintenanceUpcomingGroupMaintenanceOutputReference) MaintenanceStatus() *string {
	var returns *string
	_jsii_.Get(
		j,
		"maintenanceStatus",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeReservationResourceStatusReservationMaintenanceUpcomingGroupMaintenanceOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeReservationResourceStatusReservationMaintenanceUpcomingGroupMaintenanceOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeReservationResourceStatusReservationMaintenanceUpcomingGroupMaintenanceOutputReference) Type() *string {
	var returns *string
	_jsii_.Get(
		j,
		"type",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeReservationResourceStatusReservationMaintenanceUpcomingGroupMaintenanceOutputReference) WindowEndTime() *string {
	var returns *string
	_jsii_.Get(
		j,
		"windowEndTime",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeReservationResourceStatusReservationMaintenanceUpcomingGroupMaintenanceOutputReference) WindowStartTime() *string {
	var returns *string
	_jsii_.Get(
		j,
		"windowStartTime",
		&returns,
	)
	return returns
}


func NewGoogleComputeReservationResourceStatusReservationMaintenanceUpcomingGroupMaintenanceOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) GoogleComputeReservationResourceStatusReservationMaintenanceUpcomingGroupMaintenanceOutputReference {
	_init_.Initialize()

	if err := validateNewGoogleComputeReservationResourceStatusReservationMaintenanceUpcomingGroupMaintenanceOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_GoogleComputeReservationResourceStatusReservationMaintenanceUpcomingGroupMaintenanceOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-google-beta.googleComputeReservation.GoogleComputeReservationResourceStatusReservationMaintenanceUpcomingGroupMaintenanceOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewGoogleComputeReservationResourceStatusReservationMaintenanceUpcomingGroupMaintenanceOutputReference_Override(g GoogleComputeReservationResourceStatusReservationMaintenanceUpcomingGroupMaintenanceOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google-beta.googleComputeReservation.GoogleComputeReservationResourceStatusReservationMaintenanceUpcomingGroupMaintenanceOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		g,
	)
}

func (j *jsiiProxy_GoogleComputeReservationResourceStatusReservationMaintenanceUpcomingGroupMaintenanceOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeReservationResourceStatusReservationMaintenanceUpcomingGroupMaintenanceOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeReservationResourceStatusReservationMaintenanceUpcomingGroupMaintenanceOutputReference)SetInternalValue(val *GoogleComputeReservationResourceStatusReservationMaintenanceUpcomingGroupMaintenance) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeReservationResourceStatusReservationMaintenanceUpcomingGroupMaintenanceOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeReservationResourceStatusReservationMaintenanceUpcomingGroupMaintenanceOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (g *jsiiProxy_GoogleComputeReservationResourceStatusReservationMaintenanceUpcomingGroupMaintenanceOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleComputeReservationResourceStatusReservationMaintenanceUpcomingGroupMaintenanceOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := g.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		g,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleComputeReservationResourceStatusReservationMaintenanceUpcomingGroupMaintenanceOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := g.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		g,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleComputeReservationResourceStatusReservationMaintenanceUpcomingGroupMaintenanceOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := g.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		g,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleComputeReservationResourceStatusReservationMaintenanceUpcomingGroupMaintenanceOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := g.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		g,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleComputeReservationResourceStatusReservationMaintenanceUpcomingGroupMaintenanceOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := g.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		g,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleComputeReservationResourceStatusReservationMaintenanceUpcomingGroupMaintenanceOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := g.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		g,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleComputeReservationResourceStatusReservationMaintenanceUpcomingGroupMaintenanceOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := g.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		g,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleComputeReservationResourceStatusReservationMaintenanceUpcomingGroupMaintenanceOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := g.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		g,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleComputeReservationResourceStatusReservationMaintenanceUpcomingGroupMaintenanceOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := g.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		g,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleComputeReservationResourceStatusReservationMaintenanceUpcomingGroupMaintenanceOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		g,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleComputeReservationResourceStatusReservationMaintenanceUpcomingGroupMaintenanceOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	if err := g.validateInterpolationForAttributeParameters(property); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		g,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleComputeReservationResourceStatusReservationMaintenanceUpcomingGroupMaintenanceOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	if err := g.validateResolveParameters(_context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		g,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleComputeReservationResourceStatusReservationMaintenanceUpcomingGroupMaintenanceOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

