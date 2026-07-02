package medialivechannel

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/aws/jsii"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/sourcegraph/controller-cdktf/gen/aws/medialivechannel/internal"
)

type MedialiveChannelEncoderSettingsOutputReference interface {
	cdktf.ComplexObject
	AudioDescriptions() MedialiveChannelEncoderSettingsAudioDescriptionsList
	AudioDescriptionsInput() any
	AvailBlanking() MedialiveChannelEncoderSettingsAvailBlankingOutputReference
	AvailBlankingInput() *MedialiveChannelEncoderSettingsAvailBlanking
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
	// Experimental.
	Fqn() *string
	InternalValue() *MedialiveChannelEncoderSettings
	SetInternalValue(val *MedialiveChannelEncoderSettings)
	OutputGroups() MedialiveChannelEncoderSettingsOutputGroupsList
	OutputGroupsInput() any
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	TimecodeConfig() MedialiveChannelEncoderSettingsTimecodeConfigOutputReference
	TimecodeConfigInput() *MedialiveChannelEncoderSettingsTimecodeConfig
	VideoDescriptions() MedialiveChannelEncoderSettingsVideoDescriptionsList
	VideoDescriptionsInput() any
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
	PutAudioDescriptions(value any)
	PutAvailBlanking(value *MedialiveChannelEncoderSettingsAvailBlanking)
	PutOutputGroups(value any)
	PutTimecodeConfig(value *MedialiveChannelEncoderSettingsTimecodeConfig)
	PutVideoDescriptions(value any)
	ResetAudioDescriptions()
	ResetAvailBlanking()
	ResetVideoDescriptions()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) any
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for MedialiveChannelEncoderSettingsOutputReference
type jsiiProxy_MedialiveChannelEncoderSettingsOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_MedialiveChannelEncoderSettingsOutputReference) AudioDescriptions() MedialiveChannelEncoderSettingsAudioDescriptionsList {
	var returns MedialiveChannelEncoderSettingsAudioDescriptionsList
	_jsii_.Get(
		j,
		"audioDescriptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MedialiveChannelEncoderSettingsOutputReference) AudioDescriptionsInput() any {
	var returns any
	_jsii_.Get(
		j,
		"audioDescriptionsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MedialiveChannelEncoderSettingsOutputReference) AvailBlanking() MedialiveChannelEncoderSettingsAvailBlankingOutputReference {
	var returns MedialiveChannelEncoderSettingsAvailBlankingOutputReference
	_jsii_.Get(
		j,
		"availBlanking",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MedialiveChannelEncoderSettingsOutputReference) AvailBlankingInput() *MedialiveChannelEncoderSettingsAvailBlanking {
	var returns *MedialiveChannelEncoderSettingsAvailBlanking
	_jsii_.Get(
		j,
		"availBlankingInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MedialiveChannelEncoderSettingsOutputReference) ComplexObjectIndex() any {
	var returns any
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MedialiveChannelEncoderSettingsOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MedialiveChannelEncoderSettingsOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MedialiveChannelEncoderSettingsOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MedialiveChannelEncoderSettingsOutputReference) InternalValue() *MedialiveChannelEncoderSettings {
	var returns *MedialiveChannelEncoderSettings
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MedialiveChannelEncoderSettingsOutputReference) OutputGroups() MedialiveChannelEncoderSettingsOutputGroupsList {
	var returns MedialiveChannelEncoderSettingsOutputGroupsList
	_jsii_.Get(
		j,
		"outputGroups",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MedialiveChannelEncoderSettingsOutputReference) OutputGroupsInput() any {
	var returns any
	_jsii_.Get(
		j,
		"outputGroupsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MedialiveChannelEncoderSettingsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MedialiveChannelEncoderSettingsOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MedialiveChannelEncoderSettingsOutputReference) TimecodeConfig() MedialiveChannelEncoderSettingsTimecodeConfigOutputReference {
	var returns MedialiveChannelEncoderSettingsTimecodeConfigOutputReference
	_jsii_.Get(
		j,
		"timecodeConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MedialiveChannelEncoderSettingsOutputReference) TimecodeConfigInput() *MedialiveChannelEncoderSettingsTimecodeConfig {
	var returns *MedialiveChannelEncoderSettingsTimecodeConfig
	_jsii_.Get(
		j,
		"timecodeConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MedialiveChannelEncoderSettingsOutputReference) VideoDescriptions() MedialiveChannelEncoderSettingsVideoDescriptionsList {
	var returns MedialiveChannelEncoderSettingsVideoDescriptionsList
	_jsii_.Get(
		j,
		"videoDescriptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MedialiveChannelEncoderSettingsOutputReference) VideoDescriptionsInput() any {
	var returns any
	_jsii_.Get(
		j,
		"videoDescriptionsInput",
		&returns,
	)
	return returns
}

func NewMedialiveChannelEncoderSettingsOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) MedialiveChannelEncoderSettingsOutputReference {
	_init_.Initialize()

	if err := validateNewMedialiveChannelEncoderSettingsOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_MedialiveChannelEncoderSettingsOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-aws.medialiveChannel.MedialiveChannelEncoderSettingsOutputReference",
		[]any{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewMedialiveChannelEncoderSettingsOutputReference_Override(m MedialiveChannelEncoderSettingsOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-aws.medialiveChannel.MedialiveChannelEncoderSettingsOutputReference",
		[]any{terraformResource, terraformAttribute},
		m,
	)
}

func (j *jsiiProxy_MedialiveChannelEncoderSettingsOutputReference) SetComplexObjectIndex(val any) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_MedialiveChannelEncoderSettingsOutputReference) SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_MedialiveChannelEncoderSettingsOutputReference) SetInternalValue(val *MedialiveChannelEncoderSettings) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_MedialiveChannelEncoderSettingsOutputReference) SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_MedialiveChannelEncoderSettingsOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (m *jsiiProxy_MedialiveChannelEncoderSettingsOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		m,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MedialiveChannelEncoderSettingsOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]any {
	if err := m.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]any

	_jsii_.Invoke(
		m,
		"getAnyMapAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MedialiveChannelEncoderSettingsOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := m.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		m,
		"getBooleanAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MedialiveChannelEncoderSettingsOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := m.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		m,
		"getBooleanMapAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MedialiveChannelEncoderSettingsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := m.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		m,
		"getListAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MedialiveChannelEncoderSettingsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := m.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		m,
		"getNumberAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MedialiveChannelEncoderSettingsOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := m.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		m,
		"getNumberListAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MedialiveChannelEncoderSettingsOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := m.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		m,
		"getNumberMapAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MedialiveChannelEncoderSettingsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := m.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		m,
		"getStringAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MedialiveChannelEncoderSettingsOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := m.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		m,
		"getStringMapAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MedialiveChannelEncoderSettingsOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		m,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MedialiveChannelEncoderSettingsOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	if err := m.validateInterpolationForAttributeParameters(property); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		m,
		"interpolationForAttribute",
		[]any{property},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MedialiveChannelEncoderSettingsOutputReference) PutAudioDescriptions(value any) {
	if err := m.validatePutAudioDescriptionsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		m,
		"putAudioDescriptions",
		[]any{value},
	)
}

func (m *jsiiProxy_MedialiveChannelEncoderSettingsOutputReference) PutAvailBlanking(value *MedialiveChannelEncoderSettingsAvailBlanking) {
	if err := m.validatePutAvailBlankingParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		m,
		"putAvailBlanking",
		[]any{value},
	)
}

func (m *jsiiProxy_MedialiveChannelEncoderSettingsOutputReference) PutOutputGroups(value any) {
	if err := m.validatePutOutputGroupsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		m,
		"putOutputGroups",
		[]any{value},
	)
}

func (m *jsiiProxy_MedialiveChannelEncoderSettingsOutputReference) PutTimecodeConfig(value *MedialiveChannelEncoderSettingsTimecodeConfig) {
	if err := m.validatePutTimecodeConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		m,
		"putTimecodeConfig",
		[]any{value},
	)
}

func (m *jsiiProxy_MedialiveChannelEncoderSettingsOutputReference) PutVideoDescriptions(value any) {
	if err := m.validatePutVideoDescriptionsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		m,
		"putVideoDescriptions",
		[]any{value},
	)
}

func (m *jsiiProxy_MedialiveChannelEncoderSettingsOutputReference) ResetAudioDescriptions() {
	_jsii_.InvokeVoid(
		m,
		"resetAudioDescriptions",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MedialiveChannelEncoderSettingsOutputReference) ResetAvailBlanking() {
	_jsii_.InvokeVoid(
		m,
		"resetAvailBlanking",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MedialiveChannelEncoderSettingsOutputReference) ResetVideoDescriptions() {
	_jsii_.InvokeVoid(
		m,
		"resetVideoDescriptions",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MedialiveChannelEncoderSettingsOutputReference) Resolve(_context cdktf.IResolveContext) any {
	if err := m.validateResolveParameters(_context); err != nil {
		panic(err)
	}
	var returns any

	_jsii_.Invoke(
		m,
		"resolve",
		[]any{_context},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MedialiveChannelEncoderSettingsOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		m,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}
