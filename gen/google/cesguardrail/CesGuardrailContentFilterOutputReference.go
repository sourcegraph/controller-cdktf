package cesguardrail

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/google/jsii"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/sourcegraph/controller-cdktf/gen/google/cesguardrail/internal"
)

type CesGuardrailContentFilterOutputReference interface {
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
	DisregardDiacritics() any
	SetDisregardDiacritics(val any)
	DisregardDiacriticsInput() any
	// Experimental.
	Fqn() *string
	InternalValue() *CesGuardrailContentFilter
	SetInternalValue(val *CesGuardrailContentFilter)
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
	ResetBannedContents()
	ResetBannedContentsInAgentResponse()
	ResetBannedContentsInUserInput()
	ResetDisregardDiacritics()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) any
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for CesGuardrailContentFilterOutputReference
type jsiiProxy_CesGuardrailContentFilterOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_CesGuardrailContentFilterOutputReference) BannedContents() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"bannedContents",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CesGuardrailContentFilterOutputReference) BannedContentsInAgentResponse() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"bannedContentsInAgentResponse",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CesGuardrailContentFilterOutputReference) BannedContentsInAgentResponseInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"bannedContentsInAgentResponseInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CesGuardrailContentFilterOutputReference) BannedContentsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"bannedContentsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CesGuardrailContentFilterOutputReference) BannedContentsInUserInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"bannedContentsInUserInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CesGuardrailContentFilterOutputReference) BannedContentsInUserInputInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"bannedContentsInUserInputInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CesGuardrailContentFilterOutputReference) ComplexObjectIndex() any {
	var returns any
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CesGuardrailContentFilterOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CesGuardrailContentFilterOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CesGuardrailContentFilterOutputReference) DisregardDiacritics() any {
	var returns any
	_jsii_.Get(
		j,
		"disregardDiacritics",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CesGuardrailContentFilterOutputReference) DisregardDiacriticsInput() any {
	var returns any
	_jsii_.Get(
		j,
		"disregardDiacriticsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CesGuardrailContentFilterOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CesGuardrailContentFilterOutputReference) InternalValue() *CesGuardrailContentFilter {
	var returns *CesGuardrailContentFilter
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CesGuardrailContentFilterOutputReference) MatchType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"matchType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CesGuardrailContentFilterOutputReference) MatchTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"matchTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CesGuardrailContentFilterOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CesGuardrailContentFilterOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func NewCesGuardrailContentFilterOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) CesGuardrailContentFilterOutputReference {
	_init_.Initialize()

	if err := validateNewCesGuardrailContentFilterOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_CesGuardrailContentFilterOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-google.cesGuardrail.CesGuardrailContentFilterOutputReference",
		[]any{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewCesGuardrailContentFilterOutputReference_Override(c CesGuardrailContentFilterOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google.cesGuardrail.CesGuardrailContentFilterOutputReference",
		[]any{terraformResource, terraformAttribute},
		c,
	)
}

func (j *jsiiProxy_CesGuardrailContentFilterOutputReference) SetBannedContents(val *[]*string) {
	if err := j.validateSetBannedContentsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"bannedContents",
		val,
	)
}

func (j *jsiiProxy_CesGuardrailContentFilterOutputReference) SetBannedContentsInAgentResponse(val *[]*string) {
	if err := j.validateSetBannedContentsInAgentResponseParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"bannedContentsInAgentResponse",
		val,
	)
}

func (j *jsiiProxy_CesGuardrailContentFilterOutputReference) SetBannedContentsInUserInput(val *[]*string) {
	if err := j.validateSetBannedContentsInUserInputParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"bannedContentsInUserInput",
		val,
	)
}

func (j *jsiiProxy_CesGuardrailContentFilterOutputReference) SetComplexObjectIndex(val any) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_CesGuardrailContentFilterOutputReference) SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_CesGuardrailContentFilterOutputReference) SetDisregardDiacritics(val any) {
	if err := j.validateSetDisregardDiacriticsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"disregardDiacritics",
		val,
	)
}

func (j *jsiiProxy_CesGuardrailContentFilterOutputReference) SetInternalValue(val *CesGuardrailContentFilter) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_CesGuardrailContentFilterOutputReference) SetMatchType(val *string) {
	if err := j.validateSetMatchTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"matchType",
		val,
	)
}

func (j *jsiiProxy_CesGuardrailContentFilterOutputReference) SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_CesGuardrailContentFilterOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (c *jsiiProxy_CesGuardrailContentFilterOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CesGuardrailContentFilterOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]any {
	if err := c.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]any

	_jsii_.Invoke(
		c,
		"getAnyMapAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CesGuardrailContentFilterOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := c.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		c,
		"getBooleanAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CesGuardrailContentFilterOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := c.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		c,
		"getBooleanMapAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CesGuardrailContentFilterOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := c.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		c,
		"getListAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CesGuardrailContentFilterOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := c.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		c,
		"getNumberAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CesGuardrailContentFilterOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := c.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		c,
		"getNumberListAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CesGuardrailContentFilterOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := c.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		c,
		"getNumberMapAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CesGuardrailContentFilterOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := c.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		c,
		"getStringAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CesGuardrailContentFilterOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := c.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		c,
		"getStringMapAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CesGuardrailContentFilterOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		c,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CesGuardrailContentFilterOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	if err := c.validateInterpolationForAttributeParameters(property); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		c,
		"interpolationForAttribute",
		[]any{property},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CesGuardrailContentFilterOutputReference) ResetBannedContents() {
	_jsii_.InvokeVoid(
		c,
		"resetBannedContents",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CesGuardrailContentFilterOutputReference) ResetBannedContentsInAgentResponse() {
	_jsii_.InvokeVoid(
		c,
		"resetBannedContentsInAgentResponse",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CesGuardrailContentFilterOutputReference) ResetBannedContentsInUserInput() {
	_jsii_.InvokeVoid(
		c,
		"resetBannedContentsInUserInput",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CesGuardrailContentFilterOutputReference) ResetDisregardDiacritics() {
	_jsii_.InvokeVoid(
		c,
		"resetDisregardDiacritics",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CesGuardrailContentFilterOutputReference) Resolve(_context cdktf.IResolveContext) any {
	if err := c.validateResolveParameters(_context); err != nil {
		panic(err)
	}
	var returns any

	_jsii_.Invoke(
		c,
		"resolve",
		[]any{_context},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CesGuardrailContentFilterOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}
