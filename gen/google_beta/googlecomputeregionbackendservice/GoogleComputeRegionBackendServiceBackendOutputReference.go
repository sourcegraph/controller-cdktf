package googlecomputeregionbackendservice

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/google_beta/jsii"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/sourcegraph/controller-cdktf/gen/google_beta/googlecomputeregionbackendservice/internal"
)

type GoogleComputeRegionBackendServiceBackendOutputReference interface {
	cdktf.ComplexObject
	BalancingMode() *string
	SetBalancingMode(val *string)
	BalancingModeInput() *string
	CapacityScaler() *float64
	SetCapacityScaler(val *float64)
	CapacityScalerInput() *float64
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
	Description() *string
	SetDescription(val *string)
	DescriptionInput() *string
	Failover() interface{}
	SetFailover(val interface{})
	FailoverInput() interface{}
	// Experimental.
	Fqn() *string
	Group() *string
	SetGroup(val *string)
	GroupInput() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	MaxConnections() *float64
	SetMaxConnections(val *float64)
	MaxConnectionsInput() *float64
	MaxConnectionsPerEndpoint() *float64
	SetMaxConnectionsPerEndpoint(val *float64)
	MaxConnectionsPerEndpointInput() *float64
	MaxConnectionsPerInstance() *float64
	SetMaxConnectionsPerInstance(val *float64)
	MaxConnectionsPerInstanceInput() *float64
	MaxRate() *float64
	SetMaxRate(val *float64)
	MaxRateInput() *float64
	MaxRatePerEndpoint() *float64
	SetMaxRatePerEndpoint(val *float64)
	MaxRatePerEndpointInput() *float64
	MaxRatePerInstance() *float64
	SetMaxRatePerInstance(val *float64)
	MaxRatePerInstanceInput() *float64
	MaxUtilization() *float64
	SetMaxUtilization(val *float64)
	MaxUtilizationInput() *float64
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
	ResetBalancingMode()
	ResetCapacityScaler()
	ResetDescription()
	ResetFailover()
	ResetMaxConnections()
	ResetMaxConnectionsPerEndpoint()
	ResetMaxConnectionsPerInstance()
	ResetMaxRate()
	ResetMaxRatePerEndpoint()
	ResetMaxRatePerInstance()
	ResetMaxUtilization()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for GoogleComputeRegionBackendServiceBackendOutputReference
type jsiiProxy_GoogleComputeRegionBackendServiceBackendOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_GoogleComputeRegionBackendServiceBackendOutputReference) BalancingMode() *string {
	var returns *string
	_jsii_.Get(
		j,
		"balancingMode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeRegionBackendServiceBackendOutputReference) BalancingModeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"balancingModeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeRegionBackendServiceBackendOutputReference) CapacityScaler() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"capacityScaler",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeRegionBackendServiceBackendOutputReference) CapacityScalerInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"capacityScalerInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeRegionBackendServiceBackendOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeRegionBackendServiceBackendOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeRegionBackendServiceBackendOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeRegionBackendServiceBackendOutputReference) Description() *string {
	var returns *string
	_jsii_.Get(
		j,
		"description",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeRegionBackendServiceBackendOutputReference) DescriptionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"descriptionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeRegionBackendServiceBackendOutputReference) Failover() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"failover",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeRegionBackendServiceBackendOutputReference) FailoverInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"failoverInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeRegionBackendServiceBackendOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeRegionBackendServiceBackendOutputReference) Group() *string {
	var returns *string
	_jsii_.Get(
		j,
		"group",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeRegionBackendServiceBackendOutputReference) GroupInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"groupInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeRegionBackendServiceBackendOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeRegionBackendServiceBackendOutputReference) MaxConnections() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxConnections",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeRegionBackendServiceBackendOutputReference) MaxConnectionsInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxConnectionsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeRegionBackendServiceBackendOutputReference) MaxConnectionsPerEndpoint() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxConnectionsPerEndpoint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeRegionBackendServiceBackendOutputReference) MaxConnectionsPerEndpointInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxConnectionsPerEndpointInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeRegionBackendServiceBackendOutputReference) MaxConnectionsPerInstance() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxConnectionsPerInstance",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeRegionBackendServiceBackendOutputReference) MaxConnectionsPerInstanceInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxConnectionsPerInstanceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeRegionBackendServiceBackendOutputReference) MaxRate() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxRate",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeRegionBackendServiceBackendOutputReference) MaxRateInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxRateInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeRegionBackendServiceBackendOutputReference) MaxRatePerEndpoint() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxRatePerEndpoint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeRegionBackendServiceBackendOutputReference) MaxRatePerEndpointInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxRatePerEndpointInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeRegionBackendServiceBackendOutputReference) MaxRatePerInstance() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxRatePerInstance",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeRegionBackendServiceBackendOutputReference) MaxRatePerInstanceInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxRatePerInstanceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeRegionBackendServiceBackendOutputReference) MaxUtilization() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxUtilization",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeRegionBackendServiceBackendOutputReference) MaxUtilizationInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxUtilizationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeRegionBackendServiceBackendOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeRegionBackendServiceBackendOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewGoogleComputeRegionBackendServiceBackendOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) GoogleComputeRegionBackendServiceBackendOutputReference {
	_init_.Initialize()

	if err := validateNewGoogleComputeRegionBackendServiceBackendOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_GoogleComputeRegionBackendServiceBackendOutputReference{}

	_jsii_.Create(
		"google-beta.googleComputeRegionBackendService.GoogleComputeRegionBackendServiceBackendOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewGoogleComputeRegionBackendServiceBackendOutputReference_Override(g GoogleComputeRegionBackendServiceBackendOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"google-beta.googleComputeRegionBackendService.GoogleComputeRegionBackendServiceBackendOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		g,
	)
}

func (j *jsiiProxy_GoogleComputeRegionBackendServiceBackendOutputReference)SetBalancingMode(val *string) {
	if err := j.validateSetBalancingModeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"balancingMode",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeRegionBackendServiceBackendOutputReference)SetCapacityScaler(val *float64) {
	if err := j.validateSetCapacityScalerParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"capacityScaler",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeRegionBackendServiceBackendOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeRegionBackendServiceBackendOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeRegionBackendServiceBackendOutputReference)SetDescription(val *string) {
	if err := j.validateSetDescriptionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"description",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeRegionBackendServiceBackendOutputReference)SetFailover(val interface{}) {
	if err := j.validateSetFailoverParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"failover",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeRegionBackendServiceBackendOutputReference)SetGroup(val *string) {
	if err := j.validateSetGroupParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"group",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeRegionBackendServiceBackendOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeRegionBackendServiceBackendOutputReference)SetMaxConnections(val *float64) {
	if err := j.validateSetMaxConnectionsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"maxConnections",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeRegionBackendServiceBackendOutputReference)SetMaxConnectionsPerEndpoint(val *float64) {
	if err := j.validateSetMaxConnectionsPerEndpointParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"maxConnectionsPerEndpoint",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeRegionBackendServiceBackendOutputReference)SetMaxConnectionsPerInstance(val *float64) {
	if err := j.validateSetMaxConnectionsPerInstanceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"maxConnectionsPerInstance",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeRegionBackendServiceBackendOutputReference)SetMaxRate(val *float64) {
	if err := j.validateSetMaxRateParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"maxRate",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeRegionBackendServiceBackendOutputReference)SetMaxRatePerEndpoint(val *float64) {
	if err := j.validateSetMaxRatePerEndpointParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"maxRatePerEndpoint",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeRegionBackendServiceBackendOutputReference)SetMaxRatePerInstance(val *float64) {
	if err := j.validateSetMaxRatePerInstanceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"maxRatePerInstance",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeRegionBackendServiceBackendOutputReference)SetMaxUtilization(val *float64) {
	if err := j.validateSetMaxUtilizationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"maxUtilization",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeRegionBackendServiceBackendOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeRegionBackendServiceBackendOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (g *jsiiProxy_GoogleComputeRegionBackendServiceBackendOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleComputeRegionBackendServiceBackendOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (g *jsiiProxy_GoogleComputeRegionBackendServiceBackendOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GoogleComputeRegionBackendServiceBackendOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (g *jsiiProxy_GoogleComputeRegionBackendServiceBackendOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (g *jsiiProxy_GoogleComputeRegionBackendServiceBackendOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (g *jsiiProxy_GoogleComputeRegionBackendServiceBackendOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (g *jsiiProxy_GoogleComputeRegionBackendServiceBackendOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (g *jsiiProxy_GoogleComputeRegionBackendServiceBackendOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (g *jsiiProxy_GoogleComputeRegionBackendServiceBackendOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (g *jsiiProxy_GoogleComputeRegionBackendServiceBackendOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		g,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleComputeRegionBackendServiceBackendOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GoogleComputeRegionBackendServiceBackendOutputReference) ResetBalancingMode() {
	_jsii_.InvokeVoid(
		g,
		"resetBalancingMode",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleComputeRegionBackendServiceBackendOutputReference) ResetCapacityScaler() {
	_jsii_.InvokeVoid(
		g,
		"resetCapacityScaler",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleComputeRegionBackendServiceBackendOutputReference) ResetDescription() {
	_jsii_.InvokeVoid(
		g,
		"resetDescription",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleComputeRegionBackendServiceBackendOutputReference) ResetFailover() {
	_jsii_.InvokeVoid(
		g,
		"resetFailover",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleComputeRegionBackendServiceBackendOutputReference) ResetMaxConnections() {
	_jsii_.InvokeVoid(
		g,
		"resetMaxConnections",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleComputeRegionBackendServiceBackendOutputReference) ResetMaxConnectionsPerEndpoint() {
	_jsii_.InvokeVoid(
		g,
		"resetMaxConnectionsPerEndpoint",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleComputeRegionBackendServiceBackendOutputReference) ResetMaxConnectionsPerInstance() {
	_jsii_.InvokeVoid(
		g,
		"resetMaxConnectionsPerInstance",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleComputeRegionBackendServiceBackendOutputReference) ResetMaxRate() {
	_jsii_.InvokeVoid(
		g,
		"resetMaxRate",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleComputeRegionBackendServiceBackendOutputReference) ResetMaxRatePerEndpoint() {
	_jsii_.InvokeVoid(
		g,
		"resetMaxRatePerEndpoint",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleComputeRegionBackendServiceBackendOutputReference) ResetMaxRatePerInstance() {
	_jsii_.InvokeVoid(
		g,
		"resetMaxRatePerInstance",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleComputeRegionBackendServiceBackendOutputReference) ResetMaxUtilization() {
	_jsii_.InvokeVoid(
		g,
		"resetMaxUtilization",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleComputeRegionBackendServiceBackendOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (g *jsiiProxy_GoogleComputeRegionBackendServiceBackendOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

