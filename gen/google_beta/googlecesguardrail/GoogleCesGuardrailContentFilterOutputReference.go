package googlecesguardrail

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/google_beta/jsii"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/sourcegraph/controller-cdktf/gen/google_beta/googlecesguardrail/internal"
)

type GoogleCesGuardrailContentFilterOutputReference interface {
	cdktf.ComplexObject
	BannedContents() *[]*string
	SetBannedContents(val *[]*string)
	BannedContentsInAgentResponse() *[]*string
	SetBannedContentsInAgentResponse(val *[]*string)
	BannedContentsInAgentResponseInput() *[]*string
	BannedContentsInput() *[]*string
	BannedContentsInUserInput() *[]*string
	SetBannedContentsInUserInput(val *[]*string)
	BannedContentsInUserInputInput() *[]*string
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
	DisregardDiacritics() interface{}
	SetDisregardDiacritics(val interface{})
	DisregardDiacriticsInput() interface{}
	// Experimental.
	Fqn() *string
	InternalValue() *GoogleCesGuardrailContentFilter
	SetInternalValue(val *GoogleCesGuardrailContentFilter)
	MatchType() *string
	SetMatchType(val *string)
	MatchTypeInput() *string
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
	ResetBannedContents()
	ResetBannedContentsInAgentResponse()
	ResetBannedContentsInUserInput()
	ResetDisregardDiacritics()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for GoogleCesGuardrailContentFilterOutputReference
type jsiiProxy_GoogleCesGuardrailContentFilterOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_GoogleCesGuardrailContentFilterOutputReference) BannedContents() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"bannedContents",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCesGuardrailContentFilterOutputReference) BannedContentsInAgentResponse() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"bannedContentsInAgentResponse",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCesGuardrailContentFilterOutputReference) BannedContentsInAgentResponseInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"bannedContentsInAgentResponseInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCesGuardrailContentFilterOutputReference) BannedContentsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"bannedContentsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCesGuardrailContentFilterOutputReference) BannedContentsInUserInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"bannedContentsInUserInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCesGuardrailContentFilterOutputReference) BannedContentsInUserInputInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"bannedContentsInUserInputInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCesGuardrailContentFilterOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCesGuardrailContentFilterOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCesGuardrailContentFilterOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCesGuardrailContentFilterOutputReference) DisregardDiacritics() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"disregardDiacritics",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCesGuardrailContentFilterOutputReference) DisregardDiacriticsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"disregardDiacriticsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCesGuardrailContentFilterOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCesGuardrailContentFilterOutputReference) InternalValue() *GoogleCesGuardrailContentFilter {
	var returns *GoogleCesGuardrailContentFilter
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCesGuardrailContentFilterOutputReference) MatchType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"matchType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCesGuardrailContentFilterOutputReference) MatchTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"matchTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCesGuardrailContentFilterOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCesGuardrailContentFilterOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewGoogleCesGuardrailContentFilterOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) GoogleCesGuardrailContentFilterOutputReference {
	_init_.Initialize()

	if err := validateNewGoogleCesGuardrailContentFilterOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_GoogleCesGuardrailContentFilterOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-google-beta.googleCesGuardrail.GoogleCesGuardrailContentFilterOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewGoogleCesGuardrailContentFilterOutputReference_Override(g GoogleCesGuardrailContentFilterOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google-beta.googleCesGuardrail.GoogleCesGuardrailContentFilterOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		g,
	)
}

func (j *jsiiProxy_GoogleCesGuardrailContentFilterOutputReference)SetBannedContents(val *[]*string) {
	if err := j.validateSetBannedContentsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"bannedContents",
		val,
	)
}

func (j *jsiiProxy_GoogleCesGuardrailContentFilterOutputReference)SetBannedContentsInAgentResponse(val *[]*string) {
	if err := j.validateSetBannedContentsInAgentResponseParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"bannedContentsInAgentResponse",
		val,
	)
}

func (j *jsiiProxy_GoogleCesGuardrailContentFilterOutputReference)SetBannedContentsInUserInput(val *[]*string) {
	if err := j.validateSetBannedContentsInUserInputParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"bannedContentsInUserInput",
		val,
	)
}

func (j *jsiiProxy_GoogleCesGuardrailContentFilterOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_GoogleCesGuardrailContentFilterOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_GoogleCesGuardrailContentFilterOutputReference)SetDisregardDiacritics(val interface{}) {
	if err := j.validateSetDisregardDiacriticsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"disregardDiacritics",
		val,
	)
}

func (j *jsiiProxy_GoogleCesGuardrailContentFilterOutputReference)SetInternalValue(val *GoogleCesGuardrailContentFilter) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_GoogleCesGuardrailContentFilterOutputReference)SetMatchType(val *string) {
	if err := j.validateSetMatchTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"matchType",
		val,
	)
}

func (j *jsiiProxy_GoogleCesGuardrailContentFilterOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_GoogleCesGuardrailContentFilterOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (g *jsiiProxy_GoogleCesGuardrailContentFilterOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleCesGuardrailContentFilterOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (g *jsiiProxy_GoogleCesGuardrailContentFilterOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GoogleCesGuardrailContentFilterOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (g *jsiiProxy_GoogleCesGuardrailContentFilterOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (g *jsiiProxy_GoogleCesGuardrailContentFilterOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (g *jsiiProxy_GoogleCesGuardrailContentFilterOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (g *jsiiProxy_GoogleCesGuardrailContentFilterOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (g *jsiiProxy_GoogleCesGuardrailContentFilterOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (g *jsiiProxy_GoogleCesGuardrailContentFilterOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (g *jsiiProxy_GoogleCesGuardrailContentFilterOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		g,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleCesGuardrailContentFilterOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GoogleCesGuardrailContentFilterOutputReference) ResetBannedContents() {
	_jsii_.InvokeVoid(
		g,
		"resetBannedContents",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleCesGuardrailContentFilterOutputReference) ResetBannedContentsInAgentResponse() {
	_jsii_.InvokeVoid(
		g,
		"resetBannedContentsInAgentResponse",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleCesGuardrailContentFilterOutputReference) ResetBannedContentsInUserInput() {
	_jsii_.InvokeVoid(
		g,
		"resetBannedContentsInUserInput",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleCesGuardrailContentFilterOutputReference) ResetDisregardDiacritics() {
	_jsii_.InvokeVoid(
		g,
		"resetDisregardDiacritics",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleCesGuardrailContentFilterOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (g *jsiiProxy_GoogleCesGuardrailContentFilterOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

