package googlevertexaiendpointwithmodelgardendeployment

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/google_beta/jsii"

	"github.com/open-constructs/cdk-terrain-go/cdktn"
	"github.com/sourcegraph/controller-cdktf/gen/google_beta/googlevertexaiendpointwithmodelgardendeployment/internal"
)

type GoogleVertexAiEndpointWithModelGardenDeploymentDeployConfigDedicatedResourcesMachineSpecOutputReference interface {
	cdktn.ComplexObject
	AcceleratorCount() *float64
	SetAcceleratorCount(val *float64)
	AcceleratorCountInput() *float64
	AcceleratorType() *string
	SetAcceleratorType(val *string)
	AcceleratorTypeInput() *string
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
	InternalValue() *GoogleVertexAiEndpointWithModelGardenDeploymentDeployConfigDedicatedResourcesMachineSpec
	SetInternalValue(val *GoogleVertexAiEndpointWithModelGardenDeploymentDeployConfigDedicatedResourcesMachineSpec)
	MachineType() *string
	SetMachineType(val *string)
	MachineTypeInput() *string
	MultihostGpuNodeCount() *float64
	SetMultihostGpuNodeCount(val *float64)
	MultihostGpuNodeCountInput() *float64
	ReservationAffinity() GoogleVertexAiEndpointWithModelGardenDeploymentDeployConfigDedicatedResourcesMachineSpecReservationAffinityOutputReference
	ReservationAffinityInput() *GoogleVertexAiEndpointWithModelGardenDeploymentDeployConfigDedicatedResourcesMachineSpecReservationAffinity
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktn.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktn.IInterpolatingParent)
	TpuTopology() *string
	SetTpuTopology(val *string)
	TpuTopologyInput() *string
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
	PutReservationAffinity(value *GoogleVertexAiEndpointWithModelGardenDeploymentDeployConfigDedicatedResourcesMachineSpecReservationAffinity)
	ResetAcceleratorCount()
	ResetAcceleratorType()
	ResetMachineType()
	ResetMultihostGpuNodeCount()
	ResetReservationAffinity()
	ResetTpuTopology()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for GoogleVertexAiEndpointWithModelGardenDeploymentDeployConfigDedicatedResourcesMachineSpecOutputReference
type jsiiProxy_GoogleVertexAiEndpointWithModelGardenDeploymentDeployConfigDedicatedResourcesMachineSpecOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_GoogleVertexAiEndpointWithModelGardenDeploymentDeployConfigDedicatedResourcesMachineSpecOutputReference) AcceleratorCount() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"acceleratorCount",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleVertexAiEndpointWithModelGardenDeploymentDeployConfigDedicatedResourcesMachineSpecOutputReference) AcceleratorCountInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"acceleratorCountInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleVertexAiEndpointWithModelGardenDeploymentDeployConfigDedicatedResourcesMachineSpecOutputReference) AcceleratorType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"acceleratorType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleVertexAiEndpointWithModelGardenDeploymentDeployConfigDedicatedResourcesMachineSpecOutputReference) AcceleratorTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"acceleratorTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleVertexAiEndpointWithModelGardenDeploymentDeployConfigDedicatedResourcesMachineSpecOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleVertexAiEndpointWithModelGardenDeploymentDeployConfigDedicatedResourcesMachineSpecOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleVertexAiEndpointWithModelGardenDeploymentDeployConfigDedicatedResourcesMachineSpecOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleVertexAiEndpointWithModelGardenDeploymentDeployConfigDedicatedResourcesMachineSpecOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleVertexAiEndpointWithModelGardenDeploymentDeployConfigDedicatedResourcesMachineSpecOutputReference) InternalValue() *GoogleVertexAiEndpointWithModelGardenDeploymentDeployConfigDedicatedResourcesMachineSpec {
	var returns *GoogleVertexAiEndpointWithModelGardenDeploymentDeployConfigDedicatedResourcesMachineSpec
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleVertexAiEndpointWithModelGardenDeploymentDeployConfigDedicatedResourcesMachineSpecOutputReference) MachineType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"machineType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleVertexAiEndpointWithModelGardenDeploymentDeployConfigDedicatedResourcesMachineSpecOutputReference) MachineTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"machineTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleVertexAiEndpointWithModelGardenDeploymentDeployConfigDedicatedResourcesMachineSpecOutputReference) MultihostGpuNodeCount() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"multihostGpuNodeCount",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleVertexAiEndpointWithModelGardenDeploymentDeployConfigDedicatedResourcesMachineSpecOutputReference) MultihostGpuNodeCountInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"multihostGpuNodeCountInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleVertexAiEndpointWithModelGardenDeploymentDeployConfigDedicatedResourcesMachineSpecOutputReference) ReservationAffinity() GoogleVertexAiEndpointWithModelGardenDeploymentDeployConfigDedicatedResourcesMachineSpecReservationAffinityOutputReference {
	var returns GoogleVertexAiEndpointWithModelGardenDeploymentDeployConfigDedicatedResourcesMachineSpecReservationAffinityOutputReference
	_jsii_.Get(
		j,
		"reservationAffinity",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleVertexAiEndpointWithModelGardenDeploymentDeployConfigDedicatedResourcesMachineSpecOutputReference) ReservationAffinityInput() *GoogleVertexAiEndpointWithModelGardenDeploymentDeployConfigDedicatedResourcesMachineSpecReservationAffinity {
	var returns *GoogleVertexAiEndpointWithModelGardenDeploymentDeployConfigDedicatedResourcesMachineSpecReservationAffinity
	_jsii_.Get(
		j,
		"reservationAffinityInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleVertexAiEndpointWithModelGardenDeploymentDeployConfigDedicatedResourcesMachineSpecOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleVertexAiEndpointWithModelGardenDeploymentDeployConfigDedicatedResourcesMachineSpecOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleVertexAiEndpointWithModelGardenDeploymentDeployConfigDedicatedResourcesMachineSpecOutputReference) TpuTopology() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tpuTopology",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleVertexAiEndpointWithModelGardenDeploymentDeployConfigDedicatedResourcesMachineSpecOutputReference) TpuTopologyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tpuTopologyInput",
		&returns,
	)
	return returns
}


func NewGoogleVertexAiEndpointWithModelGardenDeploymentDeployConfigDedicatedResourcesMachineSpecOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) GoogleVertexAiEndpointWithModelGardenDeploymentDeployConfigDedicatedResourcesMachineSpecOutputReference {
	_init_.Initialize()

	if err := validateNewGoogleVertexAiEndpointWithModelGardenDeploymentDeployConfigDedicatedResourcesMachineSpecOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_GoogleVertexAiEndpointWithModelGardenDeploymentDeployConfigDedicatedResourcesMachineSpecOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-google-beta.googleVertexAiEndpointWithModelGardenDeployment.GoogleVertexAiEndpointWithModelGardenDeploymentDeployConfigDedicatedResourcesMachineSpecOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewGoogleVertexAiEndpointWithModelGardenDeploymentDeployConfigDedicatedResourcesMachineSpecOutputReference_Override(g GoogleVertexAiEndpointWithModelGardenDeploymentDeployConfigDedicatedResourcesMachineSpecOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-google-beta.googleVertexAiEndpointWithModelGardenDeployment.GoogleVertexAiEndpointWithModelGardenDeploymentDeployConfigDedicatedResourcesMachineSpecOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		g,
	)
}

func (j *jsiiProxy_GoogleVertexAiEndpointWithModelGardenDeploymentDeployConfigDedicatedResourcesMachineSpecOutputReference)SetAcceleratorCount(val *float64) {
	if err := j.validateSetAcceleratorCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"acceleratorCount",
		val,
	)
}

func (j *jsiiProxy_GoogleVertexAiEndpointWithModelGardenDeploymentDeployConfigDedicatedResourcesMachineSpecOutputReference)SetAcceleratorType(val *string) {
	if err := j.validateSetAcceleratorTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"acceleratorType",
		val,
	)
}

func (j *jsiiProxy_GoogleVertexAiEndpointWithModelGardenDeploymentDeployConfigDedicatedResourcesMachineSpecOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_GoogleVertexAiEndpointWithModelGardenDeploymentDeployConfigDedicatedResourcesMachineSpecOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_GoogleVertexAiEndpointWithModelGardenDeploymentDeployConfigDedicatedResourcesMachineSpecOutputReference)SetInternalValue(val *GoogleVertexAiEndpointWithModelGardenDeploymentDeployConfigDedicatedResourcesMachineSpec) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_GoogleVertexAiEndpointWithModelGardenDeploymentDeployConfigDedicatedResourcesMachineSpecOutputReference)SetMachineType(val *string) {
	if err := j.validateSetMachineTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"machineType",
		val,
	)
}

func (j *jsiiProxy_GoogleVertexAiEndpointWithModelGardenDeploymentDeployConfigDedicatedResourcesMachineSpecOutputReference)SetMultihostGpuNodeCount(val *float64) {
	if err := j.validateSetMultihostGpuNodeCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"multihostGpuNodeCount",
		val,
	)
}

func (j *jsiiProxy_GoogleVertexAiEndpointWithModelGardenDeploymentDeployConfigDedicatedResourcesMachineSpecOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_GoogleVertexAiEndpointWithModelGardenDeploymentDeployConfigDedicatedResourcesMachineSpecOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_GoogleVertexAiEndpointWithModelGardenDeploymentDeployConfigDedicatedResourcesMachineSpecOutputReference)SetTpuTopology(val *string) {
	if err := j.validateSetTpuTopologyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tpuTopology",
		val,
	)
}

func (g *jsiiProxy_GoogleVertexAiEndpointWithModelGardenDeploymentDeployConfigDedicatedResourcesMachineSpecOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleVertexAiEndpointWithModelGardenDeploymentDeployConfigDedicatedResourcesMachineSpecOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (g *jsiiProxy_GoogleVertexAiEndpointWithModelGardenDeploymentDeployConfigDedicatedResourcesMachineSpecOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := g.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		g,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleVertexAiEndpointWithModelGardenDeploymentDeployConfigDedicatedResourcesMachineSpecOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (g *jsiiProxy_GoogleVertexAiEndpointWithModelGardenDeploymentDeployConfigDedicatedResourcesMachineSpecOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (g *jsiiProxy_GoogleVertexAiEndpointWithModelGardenDeploymentDeployConfigDedicatedResourcesMachineSpecOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (g *jsiiProxy_GoogleVertexAiEndpointWithModelGardenDeploymentDeployConfigDedicatedResourcesMachineSpecOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (g *jsiiProxy_GoogleVertexAiEndpointWithModelGardenDeploymentDeployConfigDedicatedResourcesMachineSpecOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (g *jsiiProxy_GoogleVertexAiEndpointWithModelGardenDeploymentDeployConfigDedicatedResourcesMachineSpecOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (g *jsiiProxy_GoogleVertexAiEndpointWithModelGardenDeploymentDeployConfigDedicatedResourcesMachineSpecOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (g *jsiiProxy_GoogleVertexAiEndpointWithModelGardenDeploymentDeployConfigDedicatedResourcesMachineSpecOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		g,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleVertexAiEndpointWithModelGardenDeploymentDeployConfigDedicatedResourcesMachineSpecOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := g.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		g,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleVertexAiEndpointWithModelGardenDeploymentDeployConfigDedicatedResourcesMachineSpecOutputReference) PutReservationAffinity(value *GoogleVertexAiEndpointWithModelGardenDeploymentDeployConfigDedicatedResourcesMachineSpecReservationAffinity) {
	if err := g.validatePutReservationAffinityParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putReservationAffinity",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleVertexAiEndpointWithModelGardenDeploymentDeployConfigDedicatedResourcesMachineSpecOutputReference) ResetAcceleratorCount() {
	_jsii_.InvokeVoid(
		g,
		"resetAcceleratorCount",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleVertexAiEndpointWithModelGardenDeploymentDeployConfigDedicatedResourcesMachineSpecOutputReference) ResetAcceleratorType() {
	_jsii_.InvokeVoid(
		g,
		"resetAcceleratorType",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleVertexAiEndpointWithModelGardenDeploymentDeployConfigDedicatedResourcesMachineSpecOutputReference) ResetMachineType() {
	_jsii_.InvokeVoid(
		g,
		"resetMachineType",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleVertexAiEndpointWithModelGardenDeploymentDeployConfigDedicatedResourcesMachineSpecOutputReference) ResetMultihostGpuNodeCount() {
	_jsii_.InvokeVoid(
		g,
		"resetMultihostGpuNodeCount",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleVertexAiEndpointWithModelGardenDeploymentDeployConfigDedicatedResourcesMachineSpecOutputReference) ResetReservationAffinity() {
	_jsii_.InvokeVoid(
		g,
		"resetReservationAffinity",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleVertexAiEndpointWithModelGardenDeploymentDeployConfigDedicatedResourcesMachineSpecOutputReference) ResetTpuTopology() {
	_jsii_.InvokeVoid(
		g,
		"resetTpuTopology",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleVertexAiEndpointWithModelGardenDeploymentDeployConfigDedicatedResourcesMachineSpecOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
	if err := g.validateResolveParameters(context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		g,
		"resolve",
		[]interface{}{context},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleVertexAiEndpointWithModelGardenDeploymentDeployConfigDedicatedResourcesMachineSpecOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

