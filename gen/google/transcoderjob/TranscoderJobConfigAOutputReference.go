package transcoderjob

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/google/jsii"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/sourcegraph/controller-cdktf/gen/google/transcoderjob/internal"
)

type TranscoderJobConfigAOutputReference interface {
	cdktf.ComplexObject
	AdBreaks() TranscoderJobConfigAdBreaksList
	AdBreaksInput() any
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
	EditList() TranscoderJobConfigEditListStructList
	EditListInput() any
	ElementaryStreams() TranscoderJobConfigElementaryStreamsList
	ElementaryStreamsInput() any
	Encryptions() TranscoderJobConfigEncryptionsList
	EncryptionsInput() any
	// Experimental.
	Fqn() *string
	Inputs() TranscoderJobConfigInputsList
	InputsInput() any
	InternalValue() *TranscoderJobConfigA
	SetInternalValue(val *TranscoderJobConfigA)
	Manifests() TranscoderJobConfigManifestsList
	ManifestsInput() any
	MuxStreams() TranscoderJobConfigMuxStreamsList
	MuxStreamsInput() any
	Output() TranscoderJobConfigOutputOutputReference
	OutputInput() *TranscoderJobConfigOutput
	Overlays() TranscoderJobConfigOverlaysList
	OverlaysInput() any
	PubsubDestination() TranscoderJobConfigPubsubDestinationOutputReference
	PubsubDestinationInput() *TranscoderJobConfigPubsubDestination
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
	PutAdBreaks(value any)
	PutEditList(value any)
	PutElementaryStreams(value any)
	PutEncryptions(value any)
	PutInputs(value any)
	PutManifests(value any)
	PutMuxStreams(value any)
	PutOutput(value *TranscoderJobConfigOutput)
	PutOverlays(value any)
	PutPubsubDestination(value *TranscoderJobConfigPubsubDestination)
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
	Resolve(_context cdktf.IResolveContext) any
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for TranscoderJobConfigAOutputReference
type jsiiProxy_TranscoderJobConfigAOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_TranscoderJobConfigAOutputReference) AdBreaks() TranscoderJobConfigAdBreaksList {
	var returns TranscoderJobConfigAdBreaksList
	_jsii_.Get(
		j,
		"adBreaks",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TranscoderJobConfigAOutputReference) AdBreaksInput() any {
	var returns any
	_jsii_.Get(
		j,
		"adBreaksInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TranscoderJobConfigAOutputReference) ComplexObjectIndex() any {
	var returns any
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TranscoderJobConfigAOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TranscoderJobConfigAOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TranscoderJobConfigAOutputReference) EditList() TranscoderJobConfigEditListStructList {
	var returns TranscoderJobConfigEditListStructList
	_jsii_.Get(
		j,
		"editList",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TranscoderJobConfigAOutputReference) EditListInput() any {
	var returns any
	_jsii_.Get(
		j,
		"editListInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TranscoderJobConfigAOutputReference) ElementaryStreams() TranscoderJobConfigElementaryStreamsList {
	var returns TranscoderJobConfigElementaryStreamsList
	_jsii_.Get(
		j,
		"elementaryStreams",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TranscoderJobConfigAOutputReference) ElementaryStreamsInput() any {
	var returns any
	_jsii_.Get(
		j,
		"elementaryStreamsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TranscoderJobConfigAOutputReference) Encryptions() TranscoderJobConfigEncryptionsList {
	var returns TranscoderJobConfigEncryptionsList
	_jsii_.Get(
		j,
		"encryptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TranscoderJobConfigAOutputReference) EncryptionsInput() any {
	var returns any
	_jsii_.Get(
		j,
		"encryptionsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TranscoderJobConfigAOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TranscoderJobConfigAOutputReference) Inputs() TranscoderJobConfigInputsList {
	var returns TranscoderJobConfigInputsList
	_jsii_.Get(
		j,
		"inputs",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TranscoderJobConfigAOutputReference) InputsInput() any {
	var returns any
	_jsii_.Get(
		j,
		"inputsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TranscoderJobConfigAOutputReference) InternalValue() *TranscoderJobConfigA {
	var returns *TranscoderJobConfigA
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TranscoderJobConfigAOutputReference) Manifests() TranscoderJobConfigManifestsList {
	var returns TranscoderJobConfigManifestsList
	_jsii_.Get(
		j,
		"manifests",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TranscoderJobConfigAOutputReference) ManifestsInput() any {
	var returns any
	_jsii_.Get(
		j,
		"manifestsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TranscoderJobConfigAOutputReference) MuxStreams() TranscoderJobConfigMuxStreamsList {
	var returns TranscoderJobConfigMuxStreamsList
	_jsii_.Get(
		j,
		"muxStreams",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TranscoderJobConfigAOutputReference) MuxStreamsInput() any {
	var returns any
	_jsii_.Get(
		j,
		"muxStreamsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TranscoderJobConfigAOutputReference) Output() TranscoderJobConfigOutputOutputReference {
	var returns TranscoderJobConfigOutputOutputReference
	_jsii_.Get(
		j,
		"output",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TranscoderJobConfigAOutputReference) OutputInput() *TranscoderJobConfigOutput {
	var returns *TranscoderJobConfigOutput
	_jsii_.Get(
		j,
		"outputInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TranscoderJobConfigAOutputReference) Overlays() TranscoderJobConfigOverlaysList {
	var returns TranscoderJobConfigOverlaysList
	_jsii_.Get(
		j,
		"overlays",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TranscoderJobConfigAOutputReference) OverlaysInput() any {
	var returns any
	_jsii_.Get(
		j,
		"overlaysInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TranscoderJobConfigAOutputReference) PubsubDestination() TranscoderJobConfigPubsubDestinationOutputReference {
	var returns TranscoderJobConfigPubsubDestinationOutputReference
	_jsii_.Get(
		j,
		"pubsubDestination",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TranscoderJobConfigAOutputReference) PubsubDestinationInput() *TranscoderJobConfigPubsubDestination {
	var returns *TranscoderJobConfigPubsubDestination
	_jsii_.Get(
		j,
		"pubsubDestinationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TranscoderJobConfigAOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TranscoderJobConfigAOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func NewTranscoderJobConfigAOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) TranscoderJobConfigAOutputReference {
	_init_.Initialize()

	if err := validateNewTranscoderJobConfigAOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_TranscoderJobConfigAOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-google.transcoderJob.TranscoderJobConfigAOutputReference",
		[]any{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewTranscoderJobConfigAOutputReference_Override(t TranscoderJobConfigAOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google.transcoderJob.TranscoderJobConfigAOutputReference",
		[]any{terraformResource, terraformAttribute},
		t,
	)
}

func (j *jsiiProxy_TranscoderJobConfigAOutputReference) SetComplexObjectIndex(val any) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_TranscoderJobConfigAOutputReference) SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_TranscoderJobConfigAOutputReference) SetInternalValue(val *TranscoderJobConfigA) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_TranscoderJobConfigAOutputReference) SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_TranscoderJobConfigAOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (t *jsiiProxy_TranscoderJobConfigAOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TranscoderJobConfigAOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]any {
	if err := t.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]any

	_jsii_.Invoke(
		t,
		"getAnyMapAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TranscoderJobConfigAOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := t.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		t,
		"getBooleanAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TranscoderJobConfigAOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := t.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		t,
		"getBooleanMapAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TranscoderJobConfigAOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := t.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		t,
		"getListAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TranscoderJobConfigAOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := t.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		t,
		"getNumberAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TranscoderJobConfigAOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := t.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		t,
		"getNumberListAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TranscoderJobConfigAOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := t.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		t,
		"getNumberMapAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TranscoderJobConfigAOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := t.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		t,
		"getStringAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TranscoderJobConfigAOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := t.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		t,
		"getStringMapAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TranscoderJobConfigAOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		t,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TranscoderJobConfigAOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	if err := t.validateInterpolationForAttributeParameters(property); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		t,
		"interpolationForAttribute",
		[]any{property},
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TranscoderJobConfigAOutputReference) PutAdBreaks(value any) {
	if err := t.validatePutAdBreaksParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putAdBreaks",
		[]any{value},
	)
}

func (t *jsiiProxy_TranscoderJobConfigAOutputReference) PutEditList(value any) {
	if err := t.validatePutEditListParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putEditList",
		[]any{value},
	)
}

func (t *jsiiProxy_TranscoderJobConfigAOutputReference) PutElementaryStreams(value any) {
	if err := t.validatePutElementaryStreamsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putElementaryStreams",
		[]any{value},
	)
}

func (t *jsiiProxy_TranscoderJobConfigAOutputReference) PutEncryptions(value any) {
	if err := t.validatePutEncryptionsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putEncryptions",
		[]any{value},
	)
}

func (t *jsiiProxy_TranscoderJobConfigAOutputReference) PutInputs(value any) {
	if err := t.validatePutInputsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putInputs",
		[]any{value},
	)
}

func (t *jsiiProxy_TranscoderJobConfigAOutputReference) PutManifests(value any) {
	if err := t.validatePutManifestsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putManifests",
		[]any{value},
	)
}

func (t *jsiiProxy_TranscoderJobConfigAOutputReference) PutMuxStreams(value any) {
	if err := t.validatePutMuxStreamsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putMuxStreams",
		[]any{value},
	)
}

func (t *jsiiProxy_TranscoderJobConfigAOutputReference) PutOutput(value *TranscoderJobConfigOutput) {
	if err := t.validatePutOutputParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putOutput",
		[]any{value},
	)
}

func (t *jsiiProxy_TranscoderJobConfigAOutputReference) PutOverlays(value any) {
	if err := t.validatePutOverlaysParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putOverlays",
		[]any{value},
	)
}

func (t *jsiiProxy_TranscoderJobConfigAOutputReference) PutPubsubDestination(value *TranscoderJobConfigPubsubDestination) {
	if err := t.validatePutPubsubDestinationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putPubsubDestination",
		[]any{value},
	)
}

func (t *jsiiProxy_TranscoderJobConfigAOutputReference) ResetAdBreaks() {
	_jsii_.InvokeVoid(
		t,
		"resetAdBreaks",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TranscoderJobConfigAOutputReference) ResetEditList() {
	_jsii_.InvokeVoid(
		t,
		"resetEditList",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TranscoderJobConfigAOutputReference) ResetElementaryStreams() {
	_jsii_.InvokeVoid(
		t,
		"resetElementaryStreams",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TranscoderJobConfigAOutputReference) ResetEncryptions() {
	_jsii_.InvokeVoid(
		t,
		"resetEncryptions",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TranscoderJobConfigAOutputReference) ResetInputs() {
	_jsii_.InvokeVoid(
		t,
		"resetInputs",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TranscoderJobConfigAOutputReference) ResetManifests() {
	_jsii_.InvokeVoid(
		t,
		"resetManifests",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TranscoderJobConfigAOutputReference) ResetMuxStreams() {
	_jsii_.InvokeVoid(
		t,
		"resetMuxStreams",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TranscoderJobConfigAOutputReference) ResetOutput() {
	_jsii_.InvokeVoid(
		t,
		"resetOutput",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TranscoderJobConfigAOutputReference) ResetOverlays() {
	_jsii_.InvokeVoid(
		t,
		"resetOverlays",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TranscoderJobConfigAOutputReference) ResetPubsubDestination() {
	_jsii_.InvokeVoid(
		t,
		"resetPubsubDestination",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TranscoderJobConfigAOutputReference) Resolve(_context cdktf.IResolveContext) any {
	if err := t.validateResolveParameters(_context); err != nil {
		panic(err)
	}
	var returns any

	_jsii_.Invoke(
		t,
		"resolve",
		[]any{_context},
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TranscoderJobConfigAOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}
