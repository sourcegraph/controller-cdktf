package transcoderjobtemplate

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/google/jsii"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/sourcegraph/controller-cdktf/gen/google/transcoderjobtemplate/internal"
)

type TranscoderJobTemplateConfigAOutputReference interface {
	cdktf.ComplexObject
	AdBreaks() TranscoderJobTemplateConfigAdBreaksList
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
	EditList() TranscoderJobTemplateConfigEditListStructList
	EditListInput() any
	ElementaryStreams() TranscoderJobTemplateConfigElementaryStreamsList
	ElementaryStreamsInput() any
	Encryptions() TranscoderJobTemplateConfigEncryptionsList
	EncryptionsInput() any
	// Experimental.
	Fqn() *string
	Inputs() TranscoderJobTemplateConfigInputsList
	InputsInput() any
	InternalValue() *TranscoderJobTemplateConfigA
	SetInternalValue(val *TranscoderJobTemplateConfigA)
	Manifests() TranscoderJobTemplateConfigManifestsList
	ManifestsInput() any
	MuxStreams() TranscoderJobTemplateConfigMuxStreamsList
	MuxStreamsInput() any
	Output() TranscoderJobTemplateConfigOutputOutputReference
	OutputInput() *TranscoderJobTemplateConfigOutput
	Overlays() TranscoderJobTemplateConfigOverlaysList
	OverlaysInput() any
	PubsubDestination() TranscoderJobTemplateConfigPubsubDestinationOutputReference
	PubsubDestinationInput() *TranscoderJobTemplateConfigPubsubDestination
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
	PutOutput(value *TranscoderJobTemplateConfigOutput)
	PutOverlays(value any)
	PutPubsubDestination(value *TranscoderJobTemplateConfigPubsubDestination)
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

// The jsii proxy struct for TranscoderJobTemplateConfigAOutputReference
type jsiiProxy_TranscoderJobTemplateConfigAOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_TranscoderJobTemplateConfigAOutputReference) AdBreaks() TranscoderJobTemplateConfigAdBreaksList {
	var returns TranscoderJobTemplateConfigAdBreaksList
	_jsii_.Get(
		j,
		"adBreaks",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TranscoderJobTemplateConfigAOutputReference) AdBreaksInput() any {
	var returns any
	_jsii_.Get(
		j,
		"adBreaksInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TranscoderJobTemplateConfigAOutputReference) ComplexObjectIndex() any {
	var returns any
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TranscoderJobTemplateConfigAOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TranscoderJobTemplateConfigAOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TranscoderJobTemplateConfigAOutputReference) EditList() TranscoderJobTemplateConfigEditListStructList {
	var returns TranscoderJobTemplateConfigEditListStructList
	_jsii_.Get(
		j,
		"editList",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TranscoderJobTemplateConfigAOutputReference) EditListInput() any {
	var returns any
	_jsii_.Get(
		j,
		"editListInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TranscoderJobTemplateConfigAOutputReference) ElementaryStreams() TranscoderJobTemplateConfigElementaryStreamsList {
	var returns TranscoderJobTemplateConfigElementaryStreamsList
	_jsii_.Get(
		j,
		"elementaryStreams",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TranscoderJobTemplateConfigAOutputReference) ElementaryStreamsInput() any {
	var returns any
	_jsii_.Get(
		j,
		"elementaryStreamsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TranscoderJobTemplateConfigAOutputReference) Encryptions() TranscoderJobTemplateConfigEncryptionsList {
	var returns TranscoderJobTemplateConfigEncryptionsList
	_jsii_.Get(
		j,
		"encryptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TranscoderJobTemplateConfigAOutputReference) EncryptionsInput() any {
	var returns any
	_jsii_.Get(
		j,
		"encryptionsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TranscoderJobTemplateConfigAOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TranscoderJobTemplateConfigAOutputReference) Inputs() TranscoderJobTemplateConfigInputsList {
	var returns TranscoderJobTemplateConfigInputsList
	_jsii_.Get(
		j,
		"inputs",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TranscoderJobTemplateConfigAOutputReference) InputsInput() any {
	var returns any
	_jsii_.Get(
		j,
		"inputsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TranscoderJobTemplateConfigAOutputReference) InternalValue() *TranscoderJobTemplateConfigA {
	var returns *TranscoderJobTemplateConfigA
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TranscoderJobTemplateConfigAOutputReference) Manifests() TranscoderJobTemplateConfigManifestsList {
	var returns TranscoderJobTemplateConfigManifestsList
	_jsii_.Get(
		j,
		"manifests",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TranscoderJobTemplateConfigAOutputReference) ManifestsInput() any {
	var returns any
	_jsii_.Get(
		j,
		"manifestsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TranscoderJobTemplateConfigAOutputReference) MuxStreams() TranscoderJobTemplateConfigMuxStreamsList {
	var returns TranscoderJobTemplateConfigMuxStreamsList
	_jsii_.Get(
		j,
		"muxStreams",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TranscoderJobTemplateConfigAOutputReference) MuxStreamsInput() any {
	var returns any
	_jsii_.Get(
		j,
		"muxStreamsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TranscoderJobTemplateConfigAOutputReference) Output() TranscoderJobTemplateConfigOutputOutputReference {
	var returns TranscoderJobTemplateConfigOutputOutputReference
	_jsii_.Get(
		j,
		"output",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TranscoderJobTemplateConfigAOutputReference) OutputInput() *TranscoderJobTemplateConfigOutput {
	var returns *TranscoderJobTemplateConfigOutput
	_jsii_.Get(
		j,
		"outputInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TranscoderJobTemplateConfigAOutputReference) Overlays() TranscoderJobTemplateConfigOverlaysList {
	var returns TranscoderJobTemplateConfigOverlaysList
	_jsii_.Get(
		j,
		"overlays",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TranscoderJobTemplateConfigAOutputReference) OverlaysInput() any {
	var returns any
	_jsii_.Get(
		j,
		"overlaysInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TranscoderJobTemplateConfigAOutputReference) PubsubDestination() TranscoderJobTemplateConfigPubsubDestinationOutputReference {
	var returns TranscoderJobTemplateConfigPubsubDestinationOutputReference
	_jsii_.Get(
		j,
		"pubsubDestination",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TranscoderJobTemplateConfigAOutputReference) PubsubDestinationInput() *TranscoderJobTemplateConfigPubsubDestination {
	var returns *TranscoderJobTemplateConfigPubsubDestination
	_jsii_.Get(
		j,
		"pubsubDestinationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TranscoderJobTemplateConfigAOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TranscoderJobTemplateConfigAOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func NewTranscoderJobTemplateConfigAOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) TranscoderJobTemplateConfigAOutputReference {
	_init_.Initialize()

	if err := validateNewTranscoderJobTemplateConfigAOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_TranscoderJobTemplateConfigAOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-google.transcoderJobTemplate.TranscoderJobTemplateConfigAOutputReference",
		[]any{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewTranscoderJobTemplateConfigAOutputReference_Override(t TranscoderJobTemplateConfigAOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google.transcoderJobTemplate.TranscoderJobTemplateConfigAOutputReference",
		[]any{terraformResource, terraformAttribute},
		t,
	)
}

func (j *jsiiProxy_TranscoderJobTemplateConfigAOutputReference) SetComplexObjectIndex(val any) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_TranscoderJobTemplateConfigAOutputReference) SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_TranscoderJobTemplateConfigAOutputReference) SetInternalValue(val *TranscoderJobTemplateConfigA) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_TranscoderJobTemplateConfigAOutputReference) SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_TranscoderJobTemplateConfigAOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (t *jsiiProxy_TranscoderJobTemplateConfigAOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TranscoderJobTemplateConfigAOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]any {
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

func (t *jsiiProxy_TranscoderJobTemplateConfigAOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (t *jsiiProxy_TranscoderJobTemplateConfigAOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (t *jsiiProxy_TranscoderJobTemplateConfigAOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (t *jsiiProxy_TranscoderJobTemplateConfigAOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (t *jsiiProxy_TranscoderJobTemplateConfigAOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (t *jsiiProxy_TranscoderJobTemplateConfigAOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (t *jsiiProxy_TranscoderJobTemplateConfigAOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (t *jsiiProxy_TranscoderJobTemplateConfigAOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (t *jsiiProxy_TranscoderJobTemplateConfigAOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		t,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TranscoderJobTemplateConfigAOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (t *jsiiProxy_TranscoderJobTemplateConfigAOutputReference) PutAdBreaks(value any) {
	if err := t.validatePutAdBreaksParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putAdBreaks",
		[]any{value},
	)
}

func (t *jsiiProxy_TranscoderJobTemplateConfigAOutputReference) PutEditList(value any) {
	if err := t.validatePutEditListParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putEditList",
		[]any{value},
	)
}

func (t *jsiiProxy_TranscoderJobTemplateConfigAOutputReference) PutElementaryStreams(value any) {
	if err := t.validatePutElementaryStreamsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putElementaryStreams",
		[]any{value},
	)
}

func (t *jsiiProxy_TranscoderJobTemplateConfigAOutputReference) PutEncryptions(value any) {
	if err := t.validatePutEncryptionsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putEncryptions",
		[]any{value},
	)
}

func (t *jsiiProxy_TranscoderJobTemplateConfigAOutputReference) PutInputs(value any) {
	if err := t.validatePutInputsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putInputs",
		[]any{value},
	)
}

func (t *jsiiProxy_TranscoderJobTemplateConfigAOutputReference) PutManifests(value any) {
	if err := t.validatePutManifestsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putManifests",
		[]any{value},
	)
}

func (t *jsiiProxy_TranscoderJobTemplateConfigAOutputReference) PutMuxStreams(value any) {
	if err := t.validatePutMuxStreamsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putMuxStreams",
		[]any{value},
	)
}

func (t *jsiiProxy_TranscoderJobTemplateConfigAOutputReference) PutOutput(value *TranscoderJobTemplateConfigOutput) {
	if err := t.validatePutOutputParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putOutput",
		[]any{value},
	)
}

func (t *jsiiProxy_TranscoderJobTemplateConfigAOutputReference) PutOverlays(value any) {
	if err := t.validatePutOverlaysParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putOverlays",
		[]any{value},
	)
}

func (t *jsiiProxy_TranscoderJobTemplateConfigAOutputReference) PutPubsubDestination(value *TranscoderJobTemplateConfigPubsubDestination) {
	if err := t.validatePutPubsubDestinationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putPubsubDestination",
		[]any{value},
	)
}

func (t *jsiiProxy_TranscoderJobTemplateConfigAOutputReference) ResetAdBreaks() {
	_jsii_.InvokeVoid(
		t,
		"resetAdBreaks",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TranscoderJobTemplateConfigAOutputReference) ResetEditList() {
	_jsii_.InvokeVoid(
		t,
		"resetEditList",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TranscoderJobTemplateConfigAOutputReference) ResetElementaryStreams() {
	_jsii_.InvokeVoid(
		t,
		"resetElementaryStreams",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TranscoderJobTemplateConfigAOutputReference) ResetEncryptions() {
	_jsii_.InvokeVoid(
		t,
		"resetEncryptions",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TranscoderJobTemplateConfigAOutputReference) ResetInputs() {
	_jsii_.InvokeVoid(
		t,
		"resetInputs",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TranscoderJobTemplateConfigAOutputReference) ResetManifests() {
	_jsii_.InvokeVoid(
		t,
		"resetManifests",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TranscoderJobTemplateConfigAOutputReference) ResetMuxStreams() {
	_jsii_.InvokeVoid(
		t,
		"resetMuxStreams",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TranscoderJobTemplateConfigAOutputReference) ResetOutput() {
	_jsii_.InvokeVoid(
		t,
		"resetOutput",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TranscoderJobTemplateConfigAOutputReference) ResetOverlays() {
	_jsii_.InvokeVoid(
		t,
		"resetOverlays",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TranscoderJobTemplateConfigAOutputReference) ResetPubsubDestination() {
	_jsii_.InvokeVoid(
		t,
		"resetPubsubDestination",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TranscoderJobTemplateConfigAOutputReference) Resolve(_context cdktf.IResolveContext) any {
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

func (t *jsiiProxy_TranscoderJobTemplateConfigAOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}
