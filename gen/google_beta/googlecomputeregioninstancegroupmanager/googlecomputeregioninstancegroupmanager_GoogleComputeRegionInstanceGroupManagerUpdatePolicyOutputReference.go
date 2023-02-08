package googlecomputeregioninstancegroupmanager

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/google_beta/jsii"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/sourcegraph/controller-cdktf/gen/google_beta/googlecomputeregioninstancegroupmanager/internal"
)

type GoogleComputeRegionInstanceGroupManagerUpdatePolicyOutputReference interface {
	cdktf.ComplexObject
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
	InstanceRedistributionType() *string
	SetInstanceRedistributionType(val *string)
	InstanceRedistributionTypeInput() *string
	InternalValue() *GoogleComputeRegionInstanceGroupManagerUpdatePolicy
	SetInternalValue(val *GoogleComputeRegionInstanceGroupManagerUpdatePolicy)
	MaxSurgeFixed() *float64
	SetMaxSurgeFixed(val *float64)
	MaxSurgeFixedInput() *float64
	MaxSurgePercent() *float64
	SetMaxSurgePercent(val *float64)
	MaxSurgePercentInput() *float64
	MaxUnavailableFixed() *float64
	SetMaxUnavailableFixed(val *float64)
	MaxUnavailableFixedInput() *float64
	MaxUnavailablePercent() *float64
	SetMaxUnavailablePercent(val *float64)
	MaxUnavailablePercentInput() *float64
	MinimalAction() *string
	SetMinimalAction(val *string)
	MinimalActionInput() *string
	MinReadySec() *float64
	SetMinReadySec(val *float64)
	MinReadySecInput() *float64
	MostDisruptiveAllowedAction() *string
	SetMostDisruptiveAllowedAction(val *string)
	MostDisruptiveAllowedActionInput() *string
	ReplacementMethod() *string
	SetReplacementMethod(val *string)
	ReplacementMethodInput() *string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	Type() *string
	SetType(val *string)
	TypeInput() *string
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
	ResetInstanceRedistributionType()
	ResetMaxSurgeFixed()
	ResetMaxSurgePercent()
	ResetMaxUnavailableFixed()
	ResetMaxUnavailablePercent()
	ResetMinReadySec()
	ResetMostDisruptiveAllowedAction()
	ResetReplacementMethod()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for GoogleComputeRegionInstanceGroupManagerUpdatePolicyOutputReference
type jsiiProxy_GoogleComputeRegionInstanceGroupManagerUpdatePolicyOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_GoogleComputeRegionInstanceGroupManagerUpdatePolicyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeRegionInstanceGroupManagerUpdatePolicyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeRegionInstanceGroupManagerUpdatePolicyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeRegionInstanceGroupManagerUpdatePolicyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeRegionInstanceGroupManagerUpdatePolicyOutputReference) InstanceRedistributionType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"instanceRedistributionType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeRegionInstanceGroupManagerUpdatePolicyOutputReference) InstanceRedistributionTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"instanceRedistributionTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeRegionInstanceGroupManagerUpdatePolicyOutputReference) InternalValue() *GoogleComputeRegionInstanceGroupManagerUpdatePolicy {
	var returns *GoogleComputeRegionInstanceGroupManagerUpdatePolicy
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeRegionInstanceGroupManagerUpdatePolicyOutputReference) MaxSurgeFixed() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxSurgeFixed",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeRegionInstanceGroupManagerUpdatePolicyOutputReference) MaxSurgeFixedInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxSurgeFixedInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeRegionInstanceGroupManagerUpdatePolicyOutputReference) MaxSurgePercent() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxSurgePercent",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeRegionInstanceGroupManagerUpdatePolicyOutputReference) MaxSurgePercentInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxSurgePercentInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeRegionInstanceGroupManagerUpdatePolicyOutputReference) MaxUnavailableFixed() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxUnavailableFixed",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeRegionInstanceGroupManagerUpdatePolicyOutputReference) MaxUnavailableFixedInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxUnavailableFixedInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeRegionInstanceGroupManagerUpdatePolicyOutputReference) MaxUnavailablePercent() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxUnavailablePercent",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeRegionInstanceGroupManagerUpdatePolicyOutputReference) MaxUnavailablePercentInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxUnavailablePercentInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeRegionInstanceGroupManagerUpdatePolicyOutputReference) MinimalAction() *string {
	var returns *string
	_jsii_.Get(
		j,
		"minimalAction",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeRegionInstanceGroupManagerUpdatePolicyOutputReference) MinimalActionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"minimalActionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeRegionInstanceGroupManagerUpdatePolicyOutputReference) MinReadySec() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"minReadySec",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeRegionInstanceGroupManagerUpdatePolicyOutputReference) MinReadySecInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"minReadySecInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeRegionInstanceGroupManagerUpdatePolicyOutputReference) MostDisruptiveAllowedAction() *string {
	var returns *string
	_jsii_.Get(
		j,
		"mostDisruptiveAllowedAction",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeRegionInstanceGroupManagerUpdatePolicyOutputReference) MostDisruptiveAllowedActionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"mostDisruptiveAllowedActionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeRegionInstanceGroupManagerUpdatePolicyOutputReference) ReplacementMethod() *string {
	var returns *string
	_jsii_.Get(
		j,
		"replacementMethod",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeRegionInstanceGroupManagerUpdatePolicyOutputReference) ReplacementMethodInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"replacementMethodInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeRegionInstanceGroupManagerUpdatePolicyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeRegionInstanceGroupManagerUpdatePolicyOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeRegionInstanceGroupManagerUpdatePolicyOutputReference) Type() *string {
	var returns *string
	_jsii_.Get(
		j,
		"type",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeRegionInstanceGroupManagerUpdatePolicyOutputReference) TypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"typeInput",
		&returns,
	)
	return returns
}


func NewGoogleComputeRegionInstanceGroupManagerUpdatePolicyOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) GoogleComputeRegionInstanceGroupManagerUpdatePolicyOutputReference {
	_init_.Initialize()

	if err := validateNewGoogleComputeRegionInstanceGroupManagerUpdatePolicyOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_GoogleComputeRegionInstanceGroupManagerUpdatePolicyOutputReference{}

	_jsii_.Create(
		"google-beta.googleComputeRegionInstanceGroupManager.GoogleComputeRegionInstanceGroupManagerUpdatePolicyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewGoogleComputeRegionInstanceGroupManagerUpdatePolicyOutputReference_Override(g GoogleComputeRegionInstanceGroupManagerUpdatePolicyOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"google-beta.googleComputeRegionInstanceGroupManager.GoogleComputeRegionInstanceGroupManagerUpdatePolicyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		g,
	)
}

func (j *jsiiProxy_GoogleComputeRegionInstanceGroupManagerUpdatePolicyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeRegionInstanceGroupManagerUpdatePolicyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeRegionInstanceGroupManagerUpdatePolicyOutputReference)SetInstanceRedistributionType(val *string) {
	if err := j.validateSetInstanceRedistributionTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"instanceRedistributionType",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeRegionInstanceGroupManagerUpdatePolicyOutputReference)SetInternalValue(val *GoogleComputeRegionInstanceGroupManagerUpdatePolicy) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeRegionInstanceGroupManagerUpdatePolicyOutputReference)SetMaxSurgeFixed(val *float64) {
	if err := j.validateSetMaxSurgeFixedParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"maxSurgeFixed",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeRegionInstanceGroupManagerUpdatePolicyOutputReference)SetMaxSurgePercent(val *float64) {
	if err := j.validateSetMaxSurgePercentParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"maxSurgePercent",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeRegionInstanceGroupManagerUpdatePolicyOutputReference)SetMaxUnavailableFixed(val *float64) {
	if err := j.validateSetMaxUnavailableFixedParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"maxUnavailableFixed",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeRegionInstanceGroupManagerUpdatePolicyOutputReference)SetMaxUnavailablePercent(val *float64) {
	if err := j.validateSetMaxUnavailablePercentParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"maxUnavailablePercent",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeRegionInstanceGroupManagerUpdatePolicyOutputReference)SetMinimalAction(val *string) {
	if err := j.validateSetMinimalActionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"minimalAction",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeRegionInstanceGroupManagerUpdatePolicyOutputReference)SetMinReadySec(val *float64) {
	if err := j.validateSetMinReadySecParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"minReadySec",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeRegionInstanceGroupManagerUpdatePolicyOutputReference)SetMostDisruptiveAllowedAction(val *string) {
	if err := j.validateSetMostDisruptiveAllowedActionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"mostDisruptiveAllowedAction",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeRegionInstanceGroupManagerUpdatePolicyOutputReference)SetReplacementMethod(val *string) {
	if err := j.validateSetReplacementMethodParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"replacementMethod",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeRegionInstanceGroupManagerUpdatePolicyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeRegionInstanceGroupManagerUpdatePolicyOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeRegionInstanceGroupManagerUpdatePolicyOutputReference)SetType(val *string) {
	if err := j.validateSetTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"type",
		val,
	)
}

func (g *jsiiProxy_GoogleComputeRegionInstanceGroupManagerUpdatePolicyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleComputeRegionInstanceGroupManagerUpdatePolicyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (g *jsiiProxy_GoogleComputeRegionInstanceGroupManagerUpdatePolicyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GoogleComputeRegionInstanceGroupManagerUpdatePolicyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (g *jsiiProxy_GoogleComputeRegionInstanceGroupManagerUpdatePolicyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (g *jsiiProxy_GoogleComputeRegionInstanceGroupManagerUpdatePolicyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (g *jsiiProxy_GoogleComputeRegionInstanceGroupManagerUpdatePolicyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (g *jsiiProxy_GoogleComputeRegionInstanceGroupManagerUpdatePolicyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (g *jsiiProxy_GoogleComputeRegionInstanceGroupManagerUpdatePolicyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (g *jsiiProxy_GoogleComputeRegionInstanceGroupManagerUpdatePolicyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (g *jsiiProxy_GoogleComputeRegionInstanceGroupManagerUpdatePolicyOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		g,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleComputeRegionInstanceGroupManagerUpdatePolicyOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GoogleComputeRegionInstanceGroupManagerUpdatePolicyOutputReference) ResetInstanceRedistributionType() {
	_jsii_.InvokeVoid(
		g,
		"resetInstanceRedistributionType",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleComputeRegionInstanceGroupManagerUpdatePolicyOutputReference) ResetMaxSurgeFixed() {
	_jsii_.InvokeVoid(
		g,
		"resetMaxSurgeFixed",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleComputeRegionInstanceGroupManagerUpdatePolicyOutputReference) ResetMaxSurgePercent() {
	_jsii_.InvokeVoid(
		g,
		"resetMaxSurgePercent",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleComputeRegionInstanceGroupManagerUpdatePolicyOutputReference) ResetMaxUnavailableFixed() {
	_jsii_.InvokeVoid(
		g,
		"resetMaxUnavailableFixed",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleComputeRegionInstanceGroupManagerUpdatePolicyOutputReference) ResetMaxUnavailablePercent() {
	_jsii_.InvokeVoid(
		g,
		"resetMaxUnavailablePercent",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleComputeRegionInstanceGroupManagerUpdatePolicyOutputReference) ResetMinReadySec() {
	_jsii_.InvokeVoid(
		g,
		"resetMinReadySec",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleComputeRegionInstanceGroupManagerUpdatePolicyOutputReference) ResetMostDisruptiveAllowedAction() {
	_jsii_.InvokeVoid(
		g,
		"resetMostDisruptiveAllowedAction",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleComputeRegionInstanceGroupManagerUpdatePolicyOutputReference) ResetReplacementMethod() {
	_jsii_.InvokeVoid(
		g,
		"resetReplacementMethod",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleComputeRegionInstanceGroupManagerUpdatePolicyOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (g *jsiiProxy_GoogleComputeRegionInstanceGroupManagerUpdatePolicyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

