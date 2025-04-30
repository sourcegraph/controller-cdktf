package googletranscoderjob

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/google_beta/jsii"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/sourcegraph/controller-cdktf/gen/google_beta/googletranscoderjob/internal"
)

type GoogleTranscoderJobConfigEncryptionsOutputReference interface {
	cdktf.ComplexObject
	Aes128() GoogleTranscoderJobConfigEncryptionsAes128OutputReference
	Aes128Input() *GoogleTranscoderJobConfigEncryptionsAes128
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
	DrmSystems() GoogleTranscoderJobConfigEncryptionsDrmSystemsOutputReference
	DrmSystemsInput() *GoogleTranscoderJobConfigEncryptionsDrmSystems
	// Experimental.
	Fqn() *string
	Id() *string
	SetId(val *string)
	IdInput() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	MpegCenc() GoogleTranscoderJobConfigEncryptionsMpegCencOutputReference
	MpegCencInput() *GoogleTranscoderJobConfigEncryptionsMpegCenc
	SampleAes() GoogleTranscoderJobConfigEncryptionsSampleAesOutputReference
	SampleAesInput() *GoogleTranscoderJobConfigEncryptionsSampleAes
	SecretManagerKeySource() GoogleTranscoderJobConfigEncryptionsSecretManagerKeySourceOutputReference
	SecretManagerKeySourceInput() *GoogleTranscoderJobConfigEncryptionsSecretManagerKeySource
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
	PutAes128(value *GoogleTranscoderJobConfigEncryptionsAes128)
	PutDrmSystems(value *GoogleTranscoderJobConfigEncryptionsDrmSystems)
	PutMpegCenc(value *GoogleTranscoderJobConfigEncryptionsMpegCenc)
	PutSampleAes(value *GoogleTranscoderJobConfigEncryptionsSampleAes)
	PutSecretManagerKeySource(value *GoogleTranscoderJobConfigEncryptionsSecretManagerKeySource)
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

// The jsii proxy struct for GoogleTranscoderJobConfigEncryptionsOutputReference
type jsiiProxy_GoogleTranscoderJobConfigEncryptionsOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_GoogleTranscoderJobConfigEncryptionsOutputReference) Aes128() GoogleTranscoderJobConfigEncryptionsAes128OutputReference {
	var returns GoogleTranscoderJobConfigEncryptionsAes128OutputReference
	_jsii_.Get(
		j,
		"aes128",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleTranscoderJobConfigEncryptionsOutputReference) Aes128Input() *GoogleTranscoderJobConfigEncryptionsAes128 {
	var returns *GoogleTranscoderJobConfigEncryptionsAes128
	_jsii_.Get(
		j,
		"aes128Input",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleTranscoderJobConfigEncryptionsOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleTranscoderJobConfigEncryptionsOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleTranscoderJobConfigEncryptionsOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleTranscoderJobConfigEncryptionsOutputReference) DrmSystems() GoogleTranscoderJobConfigEncryptionsDrmSystemsOutputReference {
	var returns GoogleTranscoderJobConfigEncryptionsDrmSystemsOutputReference
	_jsii_.Get(
		j,
		"drmSystems",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleTranscoderJobConfigEncryptionsOutputReference) DrmSystemsInput() *GoogleTranscoderJobConfigEncryptionsDrmSystems {
	var returns *GoogleTranscoderJobConfigEncryptionsDrmSystems
	_jsii_.Get(
		j,
		"drmSystemsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleTranscoderJobConfigEncryptionsOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleTranscoderJobConfigEncryptionsOutputReference) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleTranscoderJobConfigEncryptionsOutputReference) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleTranscoderJobConfigEncryptionsOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleTranscoderJobConfigEncryptionsOutputReference) MpegCenc() GoogleTranscoderJobConfigEncryptionsMpegCencOutputReference {
	var returns GoogleTranscoderJobConfigEncryptionsMpegCencOutputReference
	_jsii_.Get(
		j,
		"mpegCenc",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleTranscoderJobConfigEncryptionsOutputReference) MpegCencInput() *GoogleTranscoderJobConfigEncryptionsMpegCenc {
	var returns *GoogleTranscoderJobConfigEncryptionsMpegCenc
	_jsii_.Get(
		j,
		"mpegCencInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleTranscoderJobConfigEncryptionsOutputReference) SampleAes() GoogleTranscoderJobConfigEncryptionsSampleAesOutputReference {
	var returns GoogleTranscoderJobConfigEncryptionsSampleAesOutputReference
	_jsii_.Get(
		j,
		"sampleAes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleTranscoderJobConfigEncryptionsOutputReference) SampleAesInput() *GoogleTranscoderJobConfigEncryptionsSampleAes {
	var returns *GoogleTranscoderJobConfigEncryptionsSampleAes
	_jsii_.Get(
		j,
		"sampleAesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleTranscoderJobConfigEncryptionsOutputReference) SecretManagerKeySource() GoogleTranscoderJobConfigEncryptionsSecretManagerKeySourceOutputReference {
	var returns GoogleTranscoderJobConfigEncryptionsSecretManagerKeySourceOutputReference
	_jsii_.Get(
		j,
		"secretManagerKeySource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleTranscoderJobConfigEncryptionsOutputReference) SecretManagerKeySourceInput() *GoogleTranscoderJobConfigEncryptionsSecretManagerKeySource {
	var returns *GoogleTranscoderJobConfigEncryptionsSecretManagerKeySource
	_jsii_.Get(
		j,
		"secretManagerKeySourceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleTranscoderJobConfigEncryptionsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleTranscoderJobConfigEncryptionsOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewGoogleTranscoderJobConfigEncryptionsOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) GoogleTranscoderJobConfigEncryptionsOutputReference {
	_init_.Initialize()

	if err := validateNewGoogleTranscoderJobConfigEncryptionsOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_GoogleTranscoderJobConfigEncryptionsOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-google-beta.googleTranscoderJob.GoogleTranscoderJobConfigEncryptionsOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewGoogleTranscoderJobConfigEncryptionsOutputReference_Override(g GoogleTranscoderJobConfigEncryptionsOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google-beta.googleTranscoderJob.GoogleTranscoderJobConfigEncryptionsOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		g,
	)
}

func (j *jsiiProxy_GoogleTranscoderJobConfigEncryptionsOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_GoogleTranscoderJobConfigEncryptionsOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_GoogleTranscoderJobConfigEncryptionsOutputReference)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_GoogleTranscoderJobConfigEncryptionsOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_GoogleTranscoderJobConfigEncryptionsOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_GoogleTranscoderJobConfigEncryptionsOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (g *jsiiProxy_GoogleTranscoderJobConfigEncryptionsOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleTranscoderJobConfigEncryptionsOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (g *jsiiProxy_GoogleTranscoderJobConfigEncryptionsOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GoogleTranscoderJobConfigEncryptionsOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (g *jsiiProxy_GoogleTranscoderJobConfigEncryptionsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (g *jsiiProxy_GoogleTranscoderJobConfigEncryptionsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (g *jsiiProxy_GoogleTranscoderJobConfigEncryptionsOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (g *jsiiProxy_GoogleTranscoderJobConfigEncryptionsOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (g *jsiiProxy_GoogleTranscoderJobConfigEncryptionsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (g *jsiiProxy_GoogleTranscoderJobConfigEncryptionsOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (g *jsiiProxy_GoogleTranscoderJobConfigEncryptionsOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		g,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleTranscoderJobConfigEncryptionsOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GoogleTranscoderJobConfigEncryptionsOutputReference) PutAes128(value *GoogleTranscoderJobConfigEncryptionsAes128) {
	if err := g.validatePutAes128Parameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putAes128",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleTranscoderJobConfigEncryptionsOutputReference) PutDrmSystems(value *GoogleTranscoderJobConfigEncryptionsDrmSystems) {
	if err := g.validatePutDrmSystemsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putDrmSystems",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleTranscoderJobConfigEncryptionsOutputReference) PutMpegCenc(value *GoogleTranscoderJobConfigEncryptionsMpegCenc) {
	if err := g.validatePutMpegCencParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putMpegCenc",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleTranscoderJobConfigEncryptionsOutputReference) PutSampleAes(value *GoogleTranscoderJobConfigEncryptionsSampleAes) {
	if err := g.validatePutSampleAesParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putSampleAes",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleTranscoderJobConfigEncryptionsOutputReference) PutSecretManagerKeySource(value *GoogleTranscoderJobConfigEncryptionsSecretManagerKeySource) {
	if err := g.validatePutSecretManagerKeySourceParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putSecretManagerKeySource",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleTranscoderJobConfigEncryptionsOutputReference) ResetAes128() {
	_jsii_.InvokeVoid(
		g,
		"resetAes128",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleTranscoderJobConfigEncryptionsOutputReference) ResetDrmSystems() {
	_jsii_.InvokeVoid(
		g,
		"resetDrmSystems",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleTranscoderJobConfigEncryptionsOutputReference) ResetMpegCenc() {
	_jsii_.InvokeVoid(
		g,
		"resetMpegCenc",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleTranscoderJobConfigEncryptionsOutputReference) ResetSampleAes() {
	_jsii_.InvokeVoid(
		g,
		"resetSampleAes",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleTranscoderJobConfigEncryptionsOutputReference) ResetSecretManagerKeySource() {
	_jsii_.InvokeVoid(
		g,
		"resetSecretManagerKeySource",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleTranscoderJobConfigEncryptionsOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (g *jsiiProxy_GoogleTranscoderJobConfigEncryptionsOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

