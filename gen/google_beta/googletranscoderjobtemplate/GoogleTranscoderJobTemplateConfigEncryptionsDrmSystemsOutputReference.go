package googletranscoderjobtemplate

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/google_beta/jsii"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/sourcegraph/controller-cdktf/gen/google_beta/googletranscoderjobtemplate/internal"
)

type GoogleTranscoderJobTemplateConfigEncryptionsDrmSystemsOutputReference interface {
	cdktf.ComplexObject
	Clearkey() GoogleTranscoderJobTemplateConfigEncryptionsDrmSystemsClearkeyOutputReference
	ClearkeyInput() *GoogleTranscoderJobTemplateConfigEncryptionsDrmSystemsClearkey
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
	Fairplay() GoogleTranscoderJobTemplateConfigEncryptionsDrmSystemsFairplayOutputReference
	FairplayInput() *GoogleTranscoderJobTemplateConfigEncryptionsDrmSystemsFairplay
	// Experimental.
	Fqn() *string
	InternalValue() *GoogleTranscoderJobTemplateConfigEncryptionsDrmSystems
	SetInternalValue(val *GoogleTranscoderJobTemplateConfigEncryptionsDrmSystems)
	Playready() GoogleTranscoderJobTemplateConfigEncryptionsDrmSystemsPlayreadyOutputReference
	PlayreadyInput() *GoogleTranscoderJobTemplateConfigEncryptionsDrmSystemsPlayready
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	Widevine() GoogleTranscoderJobTemplateConfigEncryptionsDrmSystemsWidevineOutputReference
	WidevineInput() *GoogleTranscoderJobTemplateConfigEncryptionsDrmSystemsWidevine
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
	PutClearkey(value *GoogleTranscoderJobTemplateConfigEncryptionsDrmSystemsClearkey)
	PutFairplay(value *GoogleTranscoderJobTemplateConfigEncryptionsDrmSystemsFairplay)
	PutPlayready(value *GoogleTranscoderJobTemplateConfigEncryptionsDrmSystemsPlayready)
	PutWidevine(value *GoogleTranscoderJobTemplateConfigEncryptionsDrmSystemsWidevine)
	ResetClearkey()
	ResetFairplay()
	ResetPlayready()
	ResetWidevine()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for GoogleTranscoderJobTemplateConfigEncryptionsDrmSystemsOutputReference
type jsiiProxy_GoogleTranscoderJobTemplateConfigEncryptionsDrmSystemsOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_GoogleTranscoderJobTemplateConfigEncryptionsDrmSystemsOutputReference) Clearkey() GoogleTranscoderJobTemplateConfigEncryptionsDrmSystemsClearkeyOutputReference {
	var returns GoogleTranscoderJobTemplateConfigEncryptionsDrmSystemsClearkeyOutputReference
	_jsii_.Get(
		j,
		"clearkey",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleTranscoderJobTemplateConfigEncryptionsDrmSystemsOutputReference) ClearkeyInput() *GoogleTranscoderJobTemplateConfigEncryptionsDrmSystemsClearkey {
	var returns *GoogleTranscoderJobTemplateConfigEncryptionsDrmSystemsClearkey
	_jsii_.Get(
		j,
		"clearkeyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleTranscoderJobTemplateConfigEncryptionsDrmSystemsOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleTranscoderJobTemplateConfigEncryptionsDrmSystemsOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleTranscoderJobTemplateConfigEncryptionsDrmSystemsOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleTranscoderJobTemplateConfigEncryptionsDrmSystemsOutputReference) Fairplay() GoogleTranscoderJobTemplateConfigEncryptionsDrmSystemsFairplayOutputReference {
	var returns GoogleTranscoderJobTemplateConfigEncryptionsDrmSystemsFairplayOutputReference
	_jsii_.Get(
		j,
		"fairplay",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleTranscoderJobTemplateConfigEncryptionsDrmSystemsOutputReference) FairplayInput() *GoogleTranscoderJobTemplateConfigEncryptionsDrmSystemsFairplay {
	var returns *GoogleTranscoderJobTemplateConfigEncryptionsDrmSystemsFairplay
	_jsii_.Get(
		j,
		"fairplayInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleTranscoderJobTemplateConfigEncryptionsDrmSystemsOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleTranscoderJobTemplateConfigEncryptionsDrmSystemsOutputReference) InternalValue() *GoogleTranscoderJobTemplateConfigEncryptionsDrmSystems {
	var returns *GoogleTranscoderJobTemplateConfigEncryptionsDrmSystems
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleTranscoderJobTemplateConfigEncryptionsDrmSystemsOutputReference) Playready() GoogleTranscoderJobTemplateConfigEncryptionsDrmSystemsPlayreadyOutputReference {
	var returns GoogleTranscoderJobTemplateConfigEncryptionsDrmSystemsPlayreadyOutputReference
	_jsii_.Get(
		j,
		"playready",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleTranscoderJobTemplateConfigEncryptionsDrmSystemsOutputReference) PlayreadyInput() *GoogleTranscoderJobTemplateConfigEncryptionsDrmSystemsPlayready {
	var returns *GoogleTranscoderJobTemplateConfigEncryptionsDrmSystemsPlayready
	_jsii_.Get(
		j,
		"playreadyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleTranscoderJobTemplateConfigEncryptionsDrmSystemsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleTranscoderJobTemplateConfigEncryptionsDrmSystemsOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleTranscoderJobTemplateConfigEncryptionsDrmSystemsOutputReference) Widevine() GoogleTranscoderJobTemplateConfigEncryptionsDrmSystemsWidevineOutputReference {
	var returns GoogleTranscoderJobTemplateConfigEncryptionsDrmSystemsWidevineOutputReference
	_jsii_.Get(
		j,
		"widevine",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleTranscoderJobTemplateConfigEncryptionsDrmSystemsOutputReference) WidevineInput() *GoogleTranscoderJobTemplateConfigEncryptionsDrmSystemsWidevine {
	var returns *GoogleTranscoderJobTemplateConfigEncryptionsDrmSystemsWidevine
	_jsii_.Get(
		j,
		"widevineInput",
		&returns,
	)
	return returns
}


func NewGoogleTranscoderJobTemplateConfigEncryptionsDrmSystemsOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) GoogleTranscoderJobTemplateConfigEncryptionsDrmSystemsOutputReference {
	_init_.Initialize()

	if err := validateNewGoogleTranscoderJobTemplateConfigEncryptionsDrmSystemsOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_GoogleTranscoderJobTemplateConfigEncryptionsDrmSystemsOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-google-beta.googleTranscoderJobTemplate.GoogleTranscoderJobTemplateConfigEncryptionsDrmSystemsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewGoogleTranscoderJobTemplateConfigEncryptionsDrmSystemsOutputReference_Override(g GoogleTranscoderJobTemplateConfigEncryptionsDrmSystemsOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google-beta.googleTranscoderJobTemplate.GoogleTranscoderJobTemplateConfigEncryptionsDrmSystemsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		g,
	)
}

func (j *jsiiProxy_GoogleTranscoderJobTemplateConfigEncryptionsDrmSystemsOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_GoogleTranscoderJobTemplateConfigEncryptionsDrmSystemsOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_GoogleTranscoderJobTemplateConfigEncryptionsDrmSystemsOutputReference)SetInternalValue(val *GoogleTranscoderJobTemplateConfigEncryptionsDrmSystems) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_GoogleTranscoderJobTemplateConfigEncryptionsDrmSystemsOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_GoogleTranscoderJobTemplateConfigEncryptionsDrmSystemsOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (g *jsiiProxy_GoogleTranscoderJobTemplateConfigEncryptionsDrmSystemsOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleTranscoderJobTemplateConfigEncryptionsDrmSystemsOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (g *jsiiProxy_GoogleTranscoderJobTemplateConfigEncryptionsDrmSystemsOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GoogleTranscoderJobTemplateConfigEncryptionsDrmSystemsOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (g *jsiiProxy_GoogleTranscoderJobTemplateConfigEncryptionsDrmSystemsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (g *jsiiProxy_GoogleTranscoderJobTemplateConfigEncryptionsDrmSystemsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (g *jsiiProxy_GoogleTranscoderJobTemplateConfigEncryptionsDrmSystemsOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (g *jsiiProxy_GoogleTranscoderJobTemplateConfigEncryptionsDrmSystemsOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (g *jsiiProxy_GoogleTranscoderJobTemplateConfigEncryptionsDrmSystemsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (g *jsiiProxy_GoogleTranscoderJobTemplateConfigEncryptionsDrmSystemsOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (g *jsiiProxy_GoogleTranscoderJobTemplateConfigEncryptionsDrmSystemsOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		g,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleTranscoderJobTemplateConfigEncryptionsDrmSystemsOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GoogleTranscoderJobTemplateConfigEncryptionsDrmSystemsOutputReference) PutClearkey(value *GoogleTranscoderJobTemplateConfigEncryptionsDrmSystemsClearkey) {
	if err := g.validatePutClearkeyParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putClearkey",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleTranscoderJobTemplateConfigEncryptionsDrmSystemsOutputReference) PutFairplay(value *GoogleTranscoderJobTemplateConfigEncryptionsDrmSystemsFairplay) {
	if err := g.validatePutFairplayParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putFairplay",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleTranscoderJobTemplateConfigEncryptionsDrmSystemsOutputReference) PutPlayready(value *GoogleTranscoderJobTemplateConfigEncryptionsDrmSystemsPlayready) {
	if err := g.validatePutPlayreadyParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putPlayready",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleTranscoderJobTemplateConfigEncryptionsDrmSystemsOutputReference) PutWidevine(value *GoogleTranscoderJobTemplateConfigEncryptionsDrmSystemsWidevine) {
	if err := g.validatePutWidevineParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putWidevine",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleTranscoderJobTemplateConfigEncryptionsDrmSystemsOutputReference) ResetClearkey() {
	_jsii_.InvokeVoid(
		g,
		"resetClearkey",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleTranscoderJobTemplateConfigEncryptionsDrmSystemsOutputReference) ResetFairplay() {
	_jsii_.InvokeVoid(
		g,
		"resetFairplay",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleTranscoderJobTemplateConfigEncryptionsDrmSystemsOutputReference) ResetPlayready() {
	_jsii_.InvokeVoid(
		g,
		"resetPlayready",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleTranscoderJobTemplateConfigEncryptionsDrmSystemsOutputReference) ResetWidevine() {
	_jsii_.InvokeVoid(
		g,
		"resetWidevine",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleTranscoderJobTemplateConfigEncryptionsDrmSystemsOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (g *jsiiProxy_GoogleTranscoderJobTemplateConfigEncryptionsDrmSystemsOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

