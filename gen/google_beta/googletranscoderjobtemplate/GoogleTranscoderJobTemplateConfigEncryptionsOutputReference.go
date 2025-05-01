package googletranscoderjobtemplate

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/google_beta/jsii"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/sourcegraph/controller-cdktf/gen/google_beta/googletranscoderjobtemplate/internal"
)

type GoogleTranscoderJobTemplateConfigEncryptionsOutputReference interface {
	cdktf.ComplexObject
	Aes128() GoogleTranscoderJobTemplateConfigEncryptionsAes128OutputReference
	Aes128Input() *GoogleTranscoderJobTemplateConfigEncryptionsAes128
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
	DrmSystems() GoogleTranscoderJobTemplateConfigEncryptionsDrmSystemsOutputReference
	DrmSystemsInput() *GoogleTranscoderJobTemplateConfigEncryptionsDrmSystems
	// Experimental.
	Fqn() *string
	Id() *string
	SetId(val *string)
	IdInput() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	MpegCenc() GoogleTranscoderJobTemplateConfigEncryptionsMpegCencOutputReference
	MpegCencInput() *GoogleTranscoderJobTemplateConfigEncryptionsMpegCenc
	SampleAes() GoogleTranscoderJobTemplateConfigEncryptionsSampleAesOutputReference
	SampleAesInput() *GoogleTranscoderJobTemplateConfigEncryptionsSampleAes
	SecretManagerKeySource() GoogleTranscoderJobTemplateConfigEncryptionsSecretManagerKeySourceOutputReference
	SecretManagerKeySourceInput() *GoogleTranscoderJobTemplateConfigEncryptionsSecretManagerKeySource
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
	PutAes128(value *GoogleTranscoderJobTemplateConfigEncryptionsAes128)
	PutDrmSystems(value *GoogleTranscoderJobTemplateConfigEncryptionsDrmSystems)
	PutMpegCenc(value *GoogleTranscoderJobTemplateConfigEncryptionsMpegCenc)
	PutSampleAes(value *GoogleTranscoderJobTemplateConfigEncryptionsSampleAes)
	PutSecretManagerKeySource(value *GoogleTranscoderJobTemplateConfigEncryptionsSecretManagerKeySource)
	ResetAes128()
	ResetDrmSystems()
	ResetMpegCenc()
	ResetSampleAes()
	ResetSecretManagerKeySource()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for GoogleTranscoderJobTemplateConfigEncryptionsOutputReference
type jsiiProxy_GoogleTranscoderJobTemplateConfigEncryptionsOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_GoogleTranscoderJobTemplateConfigEncryptionsOutputReference) Aes128() GoogleTranscoderJobTemplateConfigEncryptionsAes128OutputReference {
	var returns GoogleTranscoderJobTemplateConfigEncryptionsAes128OutputReference
	_jsii_.Get(
		j,
		"aes128",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleTranscoderJobTemplateConfigEncryptionsOutputReference) Aes128Input() *GoogleTranscoderJobTemplateConfigEncryptionsAes128 {
	var returns *GoogleTranscoderJobTemplateConfigEncryptionsAes128
	_jsii_.Get(
		j,
		"aes128Input",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleTranscoderJobTemplateConfigEncryptionsOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleTranscoderJobTemplateConfigEncryptionsOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleTranscoderJobTemplateConfigEncryptionsOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleTranscoderJobTemplateConfigEncryptionsOutputReference) DrmSystems() GoogleTranscoderJobTemplateConfigEncryptionsDrmSystemsOutputReference {
	var returns GoogleTranscoderJobTemplateConfigEncryptionsDrmSystemsOutputReference
	_jsii_.Get(
		j,
		"drmSystems",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleTranscoderJobTemplateConfigEncryptionsOutputReference) DrmSystemsInput() *GoogleTranscoderJobTemplateConfigEncryptionsDrmSystems {
	var returns *GoogleTranscoderJobTemplateConfigEncryptionsDrmSystems
	_jsii_.Get(
		j,
		"drmSystemsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleTranscoderJobTemplateConfigEncryptionsOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleTranscoderJobTemplateConfigEncryptionsOutputReference) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleTranscoderJobTemplateConfigEncryptionsOutputReference) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleTranscoderJobTemplateConfigEncryptionsOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleTranscoderJobTemplateConfigEncryptionsOutputReference) MpegCenc() GoogleTranscoderJobTemplateConfigEncryptionsMpegCencOutputReference {
	var returns GoogleTranscoderJobTemplateConfigEncryptionsMpegCencOutputReference
	_jsii_.Get(
		j,
		"mpegCenc",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleTranscoderJobTemplateConfigEncryptionsOutputReference) MpegCencInput() *GoogleTranscoderJobTemplateConfigEncryptionsMpegCenc {
	var returns *GoogleTranscoderJobTemplateConfigEncryptionsMpegCenc
	_jsii_.Get(
		j,
		"mpegCencInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleTranscoderJobTemplateConfigEncryptionsOutputReference) SampleAes() GoogleTranscoderJobTemplateConfigEncryptionsSampleAesOutputReference {
	var returns GoogleTranscoderJobTemplateConfigEncryptionsSampleAesOutputReference
	_jsii_.Get(
		j,
		"sampleAes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleTranscoderJobTemplateConfigEncryptionsOutputReference) SampleAesInput() *GoogleTranscoderJobTemplateConfigEncryptionsSampleAes {
	var returns *GoogleTranscoderJobTemplateConfigEncryptionsSampleAes
	_jsii_.Get(
		j,
		"sampleAesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleTranscoderJobTemplateConfigEncryptionsOutputReference) SecretManagerKeySource() GoogleTranscoderJobTemplateConfigEncryptionsSecretManagerKeySourceOutputReference {
	var returns GoogleTranscoderJobTemplateConfigEncryptionsSecretManagerKeySourceOutputReference
	_jsii_.Get(
		j,
		"secretManagerKeySource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleTranscoderJobTemplateConfigEncryptionsOutputReference) SecretManagerKeySourceInput() *GoogleTranscoderJobTemplateConfigEncryptionsSecretManagerKeySource {
	var returns *GoogleTranscoderJobTemplateConfigEncryptionsSecretManagerKeySource
	_jsii_.Get(
		j,
		"secretManagerKeySourceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleTranscoderJobTemplateConfigEncryptionsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleTranscoderJobTemplateConfigEncryptionsOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewGoogleTranscoderJobTemplateConfigEncryptionsOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) GoogleTranscoderJobTemplateConfigEncryptionsOutputReference {
	_init_.Initialize()

	if err := validateNewGoogleTranscoderJobTemplateConfigEncryptionsOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_GoogleTranscoderJobTemplateConfigEncryptionsOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-google-beta.googleTranscoderJobTemplate.GoogleTranscoderJobTemplateConfigEncryptionsOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewGoogleTranscoderJobTemplateConfigEncryptionsOutputReference_Override(g GoogleTranscoderJobTemplateConfigEncryptionsOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google-beta.googleTranscoderJobTemplate.GoogleTranscoderJobTemplateConfigEncryptionsOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		g,
	)
}

func (j *jsiiProxy_GoogleTranscoderJobTemplateConfigEncryptionsOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_GoogleTranscoderJobTemplateConfigEncryptionsOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_GoogleTranscoderJobTemplateConfigEncryptionsOutputReference)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_GoogleTranscoderJobTemplateConfigEncryptionsOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_GoogleTranscoderJobTemplateConfigEncryptionsOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_GoogleTranscoderJobTemplateConfigEncryptionsOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (g *jsiiProxy_GoogleTranscoderJobTemplateConfigEncryptionsOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleTranscoderJobTemplateConfigEncryptionsOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (g *jsiiProxy_GoogleTranscoderJobTemplateConfigEncryptionsOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GoogleTranscoderJobTemplateConfigEncryptionsOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (g *jsiiProxy_GoogleTranscoderJobTemplateConfigEncryptionsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (g *jsiiProxy_GoogleTranscoderJobTemplateConfigEncryptionsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (g *jsiiProxy_GoogleTranscoderJobTemplateConfigEncryptionsOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (g *jsiiProxy_GoogleTranscoderJobTemplateConfigEncryptionsOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (g *jsiiProxy_GoogleTranscoderJobTemplateConfigEncryptionsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (g *jsiiProxy_GoogleTranscoderJobTemplateConfigEncryptionsOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (g *jsiiProxy_GoogleTranscoderJobTemplateConfigEncryptionsOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		g,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleTranscoderJobTemplateConfigEncryptionsOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GoogleTranscoderJobTemplateConfigEncryptionsOutputReference) PutAes128(value *GoogleTranscoderJobTemplateConfigEncryptionsAes128) {
	if err := g.validatePutAes128Parameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putAes128",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleTranscoderJobTemplateConfigEncryptionsOutputReference) PutDrmSystems(value *GoogleTranscoderJobTemplateConfigEncryptionsDrmSystems) {
	if err := g.validatePutDrmSystemsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putDrmSystems",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleTranscoderJobTemplateConfigEncryptionsOutputReference) PutMpegCenc(value *GoogleTranscoderJobTemplateConfigEncryptionsMpegCenc) {
	if err := g.validatePutMpegCencParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putMpegCenc",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleTranscoderJobTemplateConfigEncryptionsOutputReference) PutSampleAes(value *GoogleTranscoderJobTemplateConfigEncryptionsSampleAes) {
	if err := g.validatePutSampleAesParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putSampleAes",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleTranscoderJobTemplateConfigEncryptionsOutputReference) PutSecretManagerKeySource(value *GoogleTranscoderJobTemplateConfigEncryptionsSecretManagerKeySource) {
	if err := g.validatePutSecretManagerKeySourceParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putSecretManagerKeySource",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleTranscoderJobTemplateConfigEncryptionsOutputReference) ResetAes128() {
	_jsii_.InvokeVoid(
		g,
		"resetAes128",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleTranscoderJobTemplateConfigEncryptionsOutputReference) ResetDrmSystems() {
	_jsii_.InvokeVoid(
		g,
		"resetDrmSystems",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleTranscoderJobTemplateConfigEncryptionsOutputReference) ResetMpegCenc() {
	_jsii_.InvokeVoid(
		g,
		"resetMpegCenc",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleTranscoderJobTemplateConfigEncryptionsOutputReference) ResetSampleAes() {
	_jsii_.InvokeVoid(
		g,
		"resetSampleAes",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleTranscoderJobTemplateConfigEncryptionsOutputReference) ResetSecretManagerKeySource() {
	_jsii_.InvokeVoid(
		g,
		"resetSecretManagerKeySource",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleTranscoderJobTemplateConfigEncryptionsOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (g *jsiiProxy_GoogleTranscoderJobTemplateConfigEncryptionsOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

