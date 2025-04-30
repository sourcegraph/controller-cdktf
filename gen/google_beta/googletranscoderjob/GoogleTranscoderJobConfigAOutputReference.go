package googletranscoderjob

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/google_beta/jsii"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/sourcegraph/controller-cdktf/gen/google_beta/googletranscoderjob/internal"
)

type GoogleTranscoderJobConfigAOutputReference interface {
	cdktf.ComplexObject
	AdBreaks() GoogleTranscoderJobConfigAdBreaksList
	AdBreaksInput() interface{}
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
	EditList() GoogleTranscoderJobConfigEditListStructList
	EditListInput() interface{}
	ElementaryStreams() GoogleTranscoderJobConfigElementaryStreamsList
	ElementaryStreamsInput() interface{}
	Encryptions() GoogleTranscoderJobConfigEncryptionsList
	EncryptionsInput() interface{}
	// Experimental.
	Fqn() *string
	Inputs() GoogleTranscoderJobConfigInputsList
	InputsInput() interface{}
	InternalValue() *GoogleTranscoderJobConfigA
	SetInternalValue(val *GoogleTranscoderJobConfigA)
	Manifests() GoogleTranscoderJobConfigManifestsList
	ManifestsInput() interface{}
	MuxStreams() GoogleTranscoderJobConfigMuxStreamsList
	MuxStreamsInput() interface{}
	Output() GoogleTranscoderJobConfigOutputOutputReference
	OutputInput() *GoogleTranscoderJobConfigOutput
	Overlays() GoogleTranscoderJobConfigOverlaysList
	OverlaysInput() interface{}
	PubsubDestination() GoogleTranscoderJobConfigPubsubDestinationOutputReference
	PubsubDestinationInput() *GoogleTranscoderJobConfigPubsubDestination
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
	PutAdBreaks(value interface{})
	PutEditList(value interface{})
	PutElementaryStreams(value interface{})
	PutEncryptions(value interface{})
	PutInputs(value interface{})
	PutManifests(value interface{})
	PutMuxStreams(value interface{})
	PutOutput(value *GoogleTranscoderJobConfigOutput)
	PutOverlays(value interface{})
	PutPubsubDestination(value *GoogleTranscoderJobConfigPubsubDestination)
	ResetAdBreaks()
	ResetEditList()
	ResetElementaryStreams()
	ResetEncryptions()
	ResetInputs()
	ResetManifests()
	ResetMuxStreams()
	ResetOutput()
	ResetOverlays()
	ResetPubsubDestination()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for GoogleTranscoderJobConfigAOutputReference
type jsiiProxy_GoogleTranscoderJobConfigAOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_GoogleTranscoderJobConfigAOutputReference) AdBreaks() GoogleTranscoderJobConfigAdBreaksList {
	var returns GoogleTranscoderJobConfigAdBreaksList
	_jsii_.Get(
		j,
		"adBreaks",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleTranscoderJobConfigAOutputReference) AdBreaksInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"adBreaksInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleTranscoderJobConfigAOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleTranscoderJobConfigAOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleTranscoderJobConfigAOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleTranscoderJobConfigAOutputReference) EditList() GoogleTranscoderJobConfigEditListStructList {
	var returns GoogleTranscoderJobConfigEditListStructList
	_jsii_.Get(
		j,
		"editList",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleTranscoderJobConfigAOutputReference) EditListInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"editListInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleTranscoderJobConfigAOutputReference) ElementaryStreams() GoogleTranscoderJobConfigElementaryStreamsList {
	var returns GoogleTranscoderJobConfigElementaryStreamsList
	_jsii_.Get(
		j,
		"elementaryStreams",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleTranscoderJobConfigAOutputReference) ElementaryStreamsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"elementaryStreamsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleTranscoderJobConfigAOutputReference) Encryptions() GoogleTranscoderJobConfigEncryptionsList {
	var returns GoogleTranscoderJobConfigEncryptionsList
	_jsii_.Get(
		j,
		"encryptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleTranscoderJobConfigAOutputReference) EncryptionsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"encryptionsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleTranscoderJobConfigAOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleTranscoderJobConfigAOutputReference) Inputs() GoogleTranscoderJobConfigInputsList {
	var returns GoogleTranscoderJobConfigInputsList
	_jsii_.Get(
		j,
		"inputs",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleTranscoderJobConfigAOutputReference) InputsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"inputsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleTranscoderJobConfigAOutputReference) InternalValue() *GoogleTranscoderJobConfigA {
	var returns *GoogleTranscoderJobConfigA
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleTranscoderJobConfigAOutputReference) Manifests() GoogleTranscoderJobConfigManifestsList {
	var returns GoogleTranscoderJobConfigManifestsList
	_jsii_.Get(
		j,
		"manifests",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleTranscoderJobConfigAOutputReference) ManifestsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"manifestsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleTranscoderJobConfigAOutputReference) MuxStreams() GoogleTranscoderJobConfigMuxStreamsList {
	var returns GoogleTranscoderJobConfigMuxStreamsList
	_jsii_.Get(
		j,
		"muxStreams",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleTranscoderJobConfigAOutputReference) MuxStreamsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"muxStreamsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleTranscoderJobConfigAOutputReference) Output() GoogleTranscoderJobConfigOutputOutputReference {
	var returns GoogleTranscoderJobConfigOutputOutputReference
	_jsii_.Get(
		j,
		"output",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleTranscoderJobConfigAOutputReference) OutputInput() *GoogleTranscoderJobConfigOutput {
	var returns *GoogleTranscoderJobConfigOutput
	_jsii_.Get(
		j,
		"outputInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleTranscoderJobConfigAOutputReference) Overlays() GoogleTranscoderJobConfigOverlaysList {
	var returns GoogleTranscoderJobConfigOverlaysList
	_jsii_.Get(
		j,
		"overlays",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleTranscoderJobConfigAOutputReference) OverlaysInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"overlaysInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleTranscoderJobConfigAOutputReference) PubsubDestination() GoogleTranscoderJobConfigPubsubDestinationOutputReference {
	var returns GoogleTranscoderJobConfigPubsubDestinationOutputReference
	_jsii_.Get(
		j,
		"pubsubDestination",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleTranscoderJobConfigAOutputReference) PubsubDestinationInput() *GoogleTranscoderJobConfigPubsubDestination {
	var returns *GoogleTranscoderJobConfigPubsubDestination
	_jsii_.Get(
		j,
		"pubsubDestinationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleTranscoderJobConfigAOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleTranscoderJobConfigAOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewGoogleTranscoderJobConfigAOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) GoogleTranscoderJobConfigAOutputReference {
	_init_.Initialize()

	if err := validateNewGoogleTranscoderJobConfigAOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_GoogleTranscoderJobConfigAOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-google-beta.googleTranscoderJob.GoogleTranscoderJobConfigAOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewGoogleTranscoderJobConfigAOutputReference_Override(g GoogleTranscoderJobConfigAOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google-beta.googleTranscoderJob.GoogleTranscoderJobConfigAOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		g,
	)
}

func (j *jsiiProxy_GoogleTranscoderJobConfigAOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_GoogleTranscoderJobConfigAOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_GoogleTranscoderJobConfigAOutputReference)SetInternalValue(val *GoogleTranscoderJobConfigA) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_GoogleTranscoderJobConfigAOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_GoogleTranscoderJobConfigAOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (g *jsiiProxy_GoogleTranscoderJobConfigAOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleTranscoderJobConfigAOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (g *jsiiProxy_GoogleTranscoderJobConfigAOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GoogleTranscoderJobConfigAOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (g *jsiiProxy_GoogleTranscoderJobConfigAOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (g *jsiiProxy_GoogleTranscoderJobConfigAOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (g *jsiiProxy_GoogleTranscoderJobConfigAOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (g *jsiiProxy_GoogleTranscoderJobConfigAOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (g *jsiiProxy_GoogleTranscoderJobConfigAOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (g *jsiiProxy_GoogleTranscoderJobConfigAOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (g *jsiiProxy_GoogleTranscoderJobConfigAOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		g,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleTranscoderJobConfigAOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GoogleTranscoderJobConfigAOutputReference) PutAdBreaks(value interface{}) {
	if err := g.validatePutAdBreaksParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putAdBreaks",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleTranscoderJobConfigAOutputReference) PutEditList(value interface{}) {
	if err := g.validatePutEditListParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putEditList",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleTranscoderJobConfigAOutputReference) PutElementaryStreams(value interface{}) {
	if err := g.validatePutElementaryStreamsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putElementaryStreams",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleTranscoderJobConfigAOutputReference) PutEncryptions(value interface{}) {
	if err := g.validatePutEncryptionsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putEncryptions",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleTranscoderJobConfigAOutputReference) PutInputs(value interface{}) {
	if err := g.validatePutInputsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putInputs",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleTranscoderJobConfigAOutputReference) PutManifests(value interface{}) {
	if err := g.validatePutManifestsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putManifests",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleTranscoderJobConfigAOutputReference) PutMuxStreams(value interface{}) {
	if err := g.validatePutMuxStreamsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putMuxStreams",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleTranscoderJobConfigAOutputReference) PutOutput(value *GoogleTranscoderJobConfigOutput) {
	if err := g.validatePutOutputParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putOutput",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleTranscoderJobConfigAOutputReference) PutOverlays(value interface{}) {
	if err := g.validatePutOverlaysParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putOverlays",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleTranscoderJobConfigAOutputReference) PutPubsubDestination(value *GoogleTranscoderJobConfigPubsubDestination) {
	if err := g.validatePutPubsubDestinationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putPubsubDestination",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleTranscoderJobConfigAOutputReference) ResetAdBreaks() {
	_jsii_.InvokeVoid(
		g,
		"resetAdBreaks",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleTranscoderJobConfigAOutputReference) ResetEditList() {
	_jsii_.InvokeVoid(
		g,
		"resetEditList",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleTranscoderJobConfigAOutputReference) ResetElementaryStreams() {
	_jsii_.InvokeVoid(
		g,
		"resetElementaryStreams",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleTranscoderJobConfigAOutputReference) ResetEncryptions() {
	_jsii_.InvokeVoid(
		g,
		"resetEncryptions",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleTranscoderJobConfigAOutputReference) ResetInputs() {
	_jsii_.InvokeVoid(
		g,
		"resetInputs",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleTranscoderJobConfigAOutputReference) ResetManifests() {
	_jsii_.InvokeVoid(
		g,
		"resetManifests",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleTranscoderJobConfigAOutputReference) ResetMuxStreams() {
	_jsii_.InvokeVoid(
		g,
		"resetMuxStreams",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleTranscoderJobConfigAOutputReference) ResetOutput() {
	_jsii_.InvokeVoid(
		g,
		"resetOutput",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleTranscoderJobConfigAOutputReference) ResetOverlays() {
	_jsii_.InvokeVoid(
		g,
		"resetOverlays",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleTranscoderJobConfigAOutputReference) ResetPubsubDestination() {
	_jsii_.InvokeVoid(
		g,
		"resetPubsubDestination",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleTranscoderJobConfigAOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (g *jsiiProxy_GoogleTranscoderJobConfigAOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

