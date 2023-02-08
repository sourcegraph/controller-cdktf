package googlecomputeinstancegroupmanager

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/google_beta/jsii"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/sourcegraph/controller-cdktf/gen/google_beta/googlecomputeinstancegroupmanager/internal"
)

type GoogleComputeInstanceGroupManagerUpdatePolicyOutputReference interface {
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
	InternalValue() *GoogleComputeInstanceGroupManagerUpdatePolicy
	SetInternalValue(val *GoogleComputeInstanceGroupManagerUpdatePolicy)
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

// The jsii proxy struct for GoogleComputeInstanceGroupManagerUpdatePolicyOutputReference
type jsiiProxy_GoogleComputeInstanceGroupManagerUpdatePolicyOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_GoogleComputeInstanceGroupManagerUpdatePolicyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeInstanceGroupManagerUpdatePolicyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeInstanceGroupManagerUpdatePolicyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeInstanceGroupManagerUpdatePolicyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeInstanceGroupManagerUpdatePolicyOutputReference) InternalValue() *GoogleComputeInstanceGroupManagerUpdatePolicy {
	var returns *GoogleComputeInstanceGroupManagerUpdatePolicy
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeInstanceGroupManagerUpdatePolicyOutputReference) MaxSurgeFixed() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxSurgeFixed",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeInstanceGroupManagerUpdatePolicyOutputReference) MaxSurgeFixedInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxSurgeFixedInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeInstanceGroupManagerUpdatePolicyOutputReference) MaxSurgePercent() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxSurgePercent",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeInstanceGroupManagerUpdatePolicyOutputReference) MaxSurgePercentInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxSurgePercentInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeInstanceGroupManagerUpdatePolicyOutputReference) MaxUnavailableFixed() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxUnavailableFixed",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeInstanceGroupManagerUpdatePolicyOutputReference) MaxUnavailableFixedInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxUnavailableFixedInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeInstanceGroupManagerUpdatePolicyOutputReference) MaxUnavailablePercent() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxUnavailablePercent",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeInstanceGroupManagerUpdatePolicyOutputReference) MaxUnavailablePercentInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxUnavailablePercentInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeInstanceGroupManagerUpdatePolicyOutputReference) MinimalAction() *string {
	var returns *string
	_jsii_.Get(
		j,
		"minimalAction",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeInstanceGroupManagerUpdatePolicyOutputReference) MinimalActionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"minimalActionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeInstanceGroupManagerUpdatePolicyOutputReference) MinReadySec() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"minReadySec",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeInstanceGroupManagerUpdatePolicyOutputReference) MinReadySecInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"minReadySecInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeInstanceGroupManagerUpdatePolicyOutputReference) MostDisruptiveAllowedAction() *string {
	var returns *string
	_jsii_.Get(
		j,
		"mostDisruptiveAllowedAction",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeInstanceGroupManagerUpdatePolicyOutputReference) MostDisruptiveAllowedActionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"mostDisruptiveAllowedActionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeInstanceGroupManagerUpdatePolicyOutputReference) ReplacementMethod() *string {
	var returns *string
	_jsii_.Get(
		j,
		"replacementMethod",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeInstanceGroupManagerUpdatePolicyOutputReference) ReplacementMethodInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"replacementMethodInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeInstanceGroupManagerUpdatePolicyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeInstanceGroupManagerUpdatePolicyOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeInstanceGroupManagerUpdatePolicyOutputReference) Type() *string {
	var returns *string
	_jsii_.Get(
		j,
		"type",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeInstanceGroupManagerUpdatePolicyOutputReference) TypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"typeInput",
		&returns,
	)
	return returns
}


func NewGoogleComputeInstanceGroupManagerUpdatePolicyOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) GoogleComputeInstanceGroupManagerUpdatePolicyOutputReference {
	_init_.Initialize()

	if err := validateNewGoogleComputeInstanceGroupManagerUpdatePolicyOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_GoogleComputeInstanceGroupManagerUpdatePolicyOutputReference{}

	_jsii_.Create(
		"google-beta.googleComputeInstanceGroupManager.GoogleComputeInstanceGroupManagerUpdatePolicyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewGoogleComputeInstanceGroupManagerUpdatePolicyOutputReference_Override(g GoogleComputeInstanceGroupManagerUpdatePolicyOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"google-beta.googleComputeInstanceGroupManager.GoogleComputeInstanceGroupManagerUpdatePolicyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		g,
	)
}

func (j *jsiiProxy_GoogleComputeInstanceGroupManagerUpdatePolicyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeInstanceGroupManagerUpdatePolicyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeInstanceGroupManagerUpdatePolicyOutputReference)SetInternalValue(val *GoogleComputeInstanceGroupManagerUpdatePolicy) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeInstanceGroupManagerUpdatePolicyOutputReference)SetMaxSurgeFixed(val *float64) {
	if err := j.validateSetMaxSurgeFixedParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"maxSurgeFixed",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeInstanceGroupManagerUpdatePolicyOutputReference)SetMaxSurgePercent(val *float64) {
	if err := j.validateSetMaxSurgePercentParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"maxSurgePercent",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeInstanceGroupManagerUpdatePolicyOutputReference)SetMaxUnavailableFixed(val *float64) {
	if err := j.validateSetMaxUnavailableFixedParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"maxUnavailableFixed",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeInstanceGroupManagerUpdatePolicyOutputReference)SetMaxUnavailablePercent(val *float64) {
	if err := j.validateSetMaxUnavailablePercentParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"maxUnavailablePercent",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeInstanceGroupManagerUpdatePolicyOutputReference)SetMinimalAction(val *string) {
	if err := j.validateSetMinimalActionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"minimalAction",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeInstanceGroupManagerUpdatePolicyOutputReference)SetMinReadySec(val *float64) {
	if err := j.validateSetMinReadySecParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"minReadySec",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeInstanceGroupManagerUpdatePolicyOutputReference)SetMostDisruptiveAllowedAction(val *string) {
	if err := j.validateSetMostDisruptiveAllowedActionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"mostDisruptiveAllowedAction",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeInstanceGroupManagerUpdatePolicyOutputReference)SetReplacementMethod(val *string) {
	if err := j.validateSetReplacementMethodParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"replacementMethod",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeInstanceGroupManagerUpdatePolicyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeInstanceGroupManagerUpdatePolicyOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeInstanceGroupManagerUpdatePolicyOutputReference)SetType(val *string) {
	if err := j.validateSetTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"type",
		val,
	)
}

func (g *jsiiProxy_GoogleComputeInstanceGroupManagerUpdatePolicyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleComputeInstanceGroupManagerUpdatePolicyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (g *jsiiProxy_GoogleComputeInstanceGroupManagerUpdatePolicyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GoogleComputeInstanceGroupManagerUpdatePolicyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (g *jsiiProxy_GoogleComputeInstanceGroupManagerUpdatePolicyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (g *jsiiProxy_GoogleComputeInstanceGroupManagerUpdatePolicyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (g *jsiiProxy_GoogleComputeInstanceGroupManagerUpdatePolicyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (g *jsiiProxy_GoogleComputeInstanceGroupManagerUpdatePolicyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (g *jsiiProxy_GoogleComputeInstanceGroupManagerUpdatePolicyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (g *jsiiProxy_GoogleComputeInstanceGroupManagerUpdatePolicyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (g *jsiiProxy_GoogleComputeInstanceGroupManagerUpdatePolicyOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		g,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleComputeInstanceGroupManagerUpdatePolicyOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GoogleComputeInstanceGroupManagerUpdatePolicyOutputReference) ResetMaxSurgeFixed() {
	_jsii_.InvokeVoid(
		g,
		"resetMaxSurgeFixed",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleComputeInstanceGroupManagerUpdatePolicyOutputReference) ResetMaxSurgePercent() {
	_jsii_.InvokeVoid(
		g,
		"resetMaxSurgePercent",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleComputeInstanceGroupManagerUpdatePolicyOutputReference) ResetMaxUnavailableFixed() {
	_jsii_.InvokeVoid(
		g,
		"resetMaxUnavailableFixed",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleComputeInstanceGroupManagerUpdatePolicyOutputReference) ResetMaxUnavailablePercent() {
	_jsii_.InvokeVoid(
		g,
		"resetMaxUnavailablePercent",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleComputeInstanceGroupManagerUpdatePolicyOutputReference) ResetMinReadySec() {
	_jsii_.InvokeVoid(
		g,
		"resetMinReadySec",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleComputeInstanceGroupManagerUpdatePolicyOutputReference) ResetMostDisruptiveAllowedAction() {
	_jsii_.InvokeVoid(
		g,
		"resetMostDisruptiveAllowedAction",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleComputeInstanceGroupManagerUpdatePolicyOutputReference) ResetReplacementMethod() {
	_jsii_.InvokeVoid(
		g,
		"resetReplacementMethod",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleComputeInstanceGroupManagerUpdatePolicyOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (g *jsiiProxy_GoogleComputeInstanceGroupManagerUpdatePolicyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

