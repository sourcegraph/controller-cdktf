package googlecomputeresourcepolicy

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/google_beta/jsii"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/sourcegraph/controller-cdktf/gen/google_beta/googlecomputeresourcepolicy/internal"
)

type GoogleComputeResourcePolicyGroupPlacementPolicyOutputReference interface {
	cdktf.ComplexObject
	AvailabilityDomainCount() *float64
	SetAvailabilityDomainCount(val *float64)
	AvailabilityDomainCountInput() *float64
	Collocation() *string
	SetCollocation(val *string)
	CollocationInput() *string
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
	InternalValue() *GoogleComputeResourcePolicyGroupPlacementPolicy
	SetInternalValue(val *GoogleComputeResourcePolicyGroupPlacementPolicy)
	MaxDistance() *float64
	SetMaxDistance(val *float64)
	MaxDistanceInput() *float64
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	VmCount() *float64
	SetVmCount(val *float64)
	VmCountInput() *float64
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
	ResetAvailabilityDomainCount()
	ResetCollocation()
	ResetMaxDistance()
	ResetVmCount()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for GoogleComputeResourcePolicyGroupPlacementPolicyOutputReference
type jsiiProxy_GoogleComputeResourcePolicyGroupPlacementPolicyOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_GoogleComputeResourcePolicyGroupPlacementPolicyOutputReference) AvailabilityDomainCount() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"availabilityDomainCount",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeResourcePolicyGroupPlacementPolicyOutputReference) AvailabilityDomainCountInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"availabilityDomainCountInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeResourcePolicyGroupPlacementPolicyOutputReference) Collocation() *string {
	var returns *string
	_jsii_.Get(
		j,
		"collocation",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeResourcePolicyGroupPlacementPolicyOutputReference) CollocationInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"collocationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeResourcePolicyGroupPlacementPolicyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeResourcePolicyGroupPlacementPolicyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeResourcePolicyGroupPlacementPolicyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeResourcePolicyGroupPlacementPolicyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeResourcePolicyGroupPlacementPolicyOutputReference) InternalValue() *GoogleComputeResourcePolicyGroupPlacementPolicy {
	var returns *GoogleComputeResourcePolicyGroupPlacementPolicy
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeResourcePolicyGroupPlacementPolicyOutputReference) MaxDistance() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxDistance",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeResourcePolicyGroupPlacementPolicyOutputReference) MaxDistanceInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxDistanceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeResourcePolicyGroupPlacementPolicyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeResourcePolicyGroupPlacementPolicyOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeResourcePolicyGroupPlacementPolicyOutputReference) VmCount() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"vmCount",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeResourcePolicyGroupPlacementPolicyOutputReference) VmCountInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"vmCountInput",
		&returns,
	)
	return returns
}


func NewGoogleComputeResourcePolicyGroupPlacementPolicyOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) GoogleComputeResourcePolicyGroupPlacementPolicyOutputReference {
	_init_.Initialize()

	if err := validateNewGoogleComputeResourcePolicyGroupPlacementPolicyOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_GoogleComputeResourcePolicyGroupPlacementPolicyOutputReference{}

	_jsii_.Create(
		"google-beta.googleComputeResourcePolicy.GoogleComputeResourcePolicyGroupPlacementPolicyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewGoogleComputeResourcePolicyGroupPlacementPolicyOutputReference_Override(g GoogleComputeResourcePolicyGroupPlacementPolicyOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"google-beta.googleComputeResourcePolicy.GoogleComputeResourcePolicyGroupPlacementPolicyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		g,
	)
}

func (j *jsiiProxy_GoogleComputeResourcePolicyGroupPlacementPolicyOutputReference)SetAvailabilityDomainCount(val *float64) {
	if err := j.validateSetAvailabilityDomainCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"availabilityDomainCount",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeResourcePolicyGroupPlacementPolicyOutputReference)SetCollocation(val *string) {
	if err := j.validateSetCollocationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"collocation",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeResourcePolicyGroupPlacementPolicyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeResourcePolicyGroupPlacementPolicyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeResourcePolicyGroupPlacementPolicyOutputReference)SetInternalValue(val *GoogleComputeResourcePolicyGroupPlacementPolicy) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeResourcePolicyGroupPlacementPolicyOutputReference)SetMaxDistance(val *float64) {
	if err := j.validateSetMaxDistanceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"maxDistance",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeResourcePolicyGroupPlacementPolicyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeResourcePolicyGroupPlacementPolicyOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeResourcePolicyGroupPlacementPolicyOutputReference)SetVmCount(val *float64) {
	if err := j.validateSetVmCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"vmCount",
		val,
	)
}

func (g *jsiiProxy_GoogleComputeResourcePolicyGroupPlacementPolicyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleComputeResourcePolicyGroupPlacementPolicyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (g *jsiiProxy_GoogleComputeResourcePolicyGroupPlacementPolicyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GoogleComputeResourcePolicyGroupPlacementPolicyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (g *jsiiProxy_GoogleComputeResourcePolicyGroupPlacementPolicyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (g *jsiiProxy_GoogleComputeResourcePolicyGroupPlacementPolicyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (g *jsiiProxy_GoogleComputeResourcePolicyGroupPlacementPolicyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (g *jsiiProxy_GoogleComputeResourcePolicyGroupPlacementPolicyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (g *jsiiProxy_GoogleComputeResourcePolicyGroupPlacementPolicyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (g *jsiiProxy_GoogleComputeResourcePolicyGroupPlacementPolicyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (g *jsiiProxy_GoogleComputeResourcePolicyGroupPlacementPolicyOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		g,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleComputeResourcePolicyGroupPlacementPolicyOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GoogleComputeResourcePolicyGroupPlacementPolicyOutputReference) ResetAvailabilityDomainCount() {
	_jsii_.InvokeVoid(
		g,
		"resetAvailabilityDomainCount",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleComputeResourcePolicyGroupPlacementPolicyOutputReference) ResetCollocation() {
	_jsii_.InvokeVoid(
		g,
		"resetCollocation",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleComputeResourcePolicyGroupPlacementPolicyOutputReference) ResetMaxDistance() {
	_jsii_.InvokeVoid(
		g,
		"resetMaxDistance",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleComputeResourcePolicyGroupPlacementPolicyOutputReference) ResetVmCount() {
	_jsii_.InvokeVoid(
		g,
		"resetVmCount",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleComputeResourcePolicyGroupPlacementPolicyOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (g *jsiiProxy_GoogleComputeResourcePolicyGroupPlacementPolicyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

