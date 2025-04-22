package googletranscoderjob

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/google_beta/jsii"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/sourcegraph/controller-cdktf/gen/google_beta/googletranscoderjob/internal"
)

type GoogleTranscoderJobConfigElementaryStreamsVideoStreamH264OutputReference interface {
	cdktf.ComplexObject
	BitrateBps() *float64
	SetBitrateBps(val *float64)
	BitrateBpsInput() *float64
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
	CrfLevel() *float64
	SetCrfLevel(val *float64)
	CrfLevelInput() *float64
	EntropyCoder() *string
	SetEntropyCoder(val *string)
	EntropyCoderInput() *string
	// Experimental.
	Fqn() *string
	FrameRate() *float64
	SetFrameRate(val *float64)
	FrameRateInput() *float64
	GopDuration() *string
	SetGopDuration(val *string)
	GopDurationInput() *string
	HeightPixels() *float64
	SetHeightPixels(val *float64)
	HeightPixelsInput() *float64
	Hlg() GoogleTranscoderJobConfigElementaryStreamsVideoStreamH264HlgOutputReference
	HlgInput() *GoogleTranscoderJobConfigElementaryStreamsVideoStreamH264Hlg
	InternalValue() *GoogleTranscoderJobConfigElementaryStreamsVideoStreamH264
	SetInternalValue(val *GoogleTranscoderJobConfigElementaryStreamsVideoStreamH264)
	PixelFormat() *string
	SetPixelFormat(val *string)
	PixelFormatInput() *string
	Preset() *string
	SetPreset(val *string)
	PresetInput() *string
	Profile() *string
	SetProfile(val *string)
	ProfileInput() *string
	RateControlMode() *string
	SetRateControlMode(val *string)
	RateControlModeInput() *string
	Sdr() GoogleTranscoderJobConfigElementaryStreamsVideoStreamH264SdrOutputReference
	SdrInput() *GoogleTranscoderJobConfigElementaryStreamsVideoStreamH264Sdr
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	VbvFullnessBits() *float64
	SetVbvFullnessBits(val *float64)
	VbvFullnessBitsInput() *float64
	VbvSizeBits() *float64
	SetVbvSizeBits(val *float64)
	VbvSizeBitsInput() *float64
	WidthPixels() *float64
	SetWidthPixels(val *float64)
	WidthPixelsInput() *float64
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
	PutHlg(value *GoogleTranscoderJobConfigElementaryStreamsVideoStreamH264Hlg)
	PutSdr(value *GoogleTranscoderJobConfigElementaryStreamsVideoStreamH264Sdr)
	ResetCrfLevel()
	ResetEntropyCoder()
	ResetGopDuration()
	ResetHeightPixels()
	ResetHlg()
	ResetPixelFormat()
	ResetPreset()
	ResetProfile()
	ResetRateControlMode()
	ResetSdr()
	ResetVbvFullnessBits()
	ResetVbvSizeBits()
	ResetWidthPixels()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for GoogleTranscoderJobConfigElementaryStreamsVideoStreamH264OutputReference
type jsiiProxy_GoogleTranscoderJobConfigElementaryStreamsVideoStreamH264OutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_GoogleTranscoderJobConfigElementaryStreamsVideoStreamH264OutputReference) BitrateBps() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"bitrateBps",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleTranscoderJobConfigElementaryStreamsVideoStreamH264OutputReference) BitrateBpsInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"bitrateBpsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleTranscoderJobConfigElementaryStreamsVideoStreamH264OutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleTranscoderJobConfigElementaryStreamsVideoStreamH264OutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleTranscoderJobConfigElementaryStreamsVideoStreamH264OutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleTranscoderJobConfigElementaryStreamsVideoStreamH264OutputReference) CrfLevel() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"crfLevel",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleTranscoderJobConfigElementaryStreamsVideoStreamH264OutputReference) CrfLevelInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"crfLevelInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleTranscoderJobConfigElementaryStreamsVideoStreamH264OutputReference) EntropyCoder() *string {
	var returns *string
	_jsii_.Get(
		j,
		"entropyCoder",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleTranscoderJobConfigElementaryStreamsVideoStreamH264OutputReference) EntropyCoderInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"entropyCoderInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleTranscoderJobConfigElementaryStreamsVideoStreamH264OutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleTranscoderJobConfigElementaryStreamsVideoStreamH264OutputReference) FrameRate() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"frameRate",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleTranscoderJobConfigElementaryStreamsVideoStreamH264OutputReference) FrameRateInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"frameRateInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleTranscoderJobConfigElementaryStreamsVideoStreamH264OutputReference) GopDuration() *string {
	var returns *string
	_jsii_.Get(
		j,
		"gopDuration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleTranscoderJobConfigElementaryStreamsVideoStreamH264OutputReference) GopDurationInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"gopDurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleTranscoderJobConfigElementaryStreamsVideoStreamH264OutputReference) HeightPixels() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"heightPixels",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleTranscoderJobConfigElementaryStreamsVideoStreamH264OutputReference) HeightPixelsInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"heightPixelsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleTranscoderJobConfigElementaryStreamsVideoStreamH264OutputReference) Hlg() GoogleTranscoderJobConfigElementaryStreamsVideoStreamH264HlgOutputReference {
	var returns GoogleTranscoderJobConfigElementaryStreamsVideoStreamH264HlgOutputReference
	_jsii_.Get(
		j,
		"hlg",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleTranscoderJobConfigElementaryStreamsVideoStreamH264OutputReference) HlgInput() *GoogleTranscoderJobConfigElementaryStreamsVideoStreamH264Hlg {
	var returns *GoogleTranscoderJobConfigElementaryStreamsVideoStreamH264Hlg
	_jsii_.Get(
		j,
		"hlgInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleTranscoderJobConfigElementaryStreamsVideoStreamH264OutputReference) InternalValue() *GoogleTranscoderJobConfigElementaryStreamsVideoStreamH264 {
	var returns *GoogleTranscoderJobConfigElementaryStreamsVideoStreamH264
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleTranscoderJobConfigElementaryStreamsVideoStreamH264OutputReference) PixelFormat() *string {
	var returns *string
	_jsii_.Get(
		j,
		"pixelFormat",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleTranscoderJobConfigElementaryStreamsVideoStreamH264OutputReference) PixelFormatInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"pixelFormatInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleTranscoderJobConfigElementaryStreamsVideoStreamH264OutputReference) Preset() *string {
	var returns *string
	_jsii_.Get(
		j,
		"preset",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleTranscoderJobConfigElementaryStreamsVideoStreamH264OutputReference) PresetInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"presetInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleTranscoderJobConfigElementaryStreamsVideoStreamH264OutputReference) Profile() *string {
	var returns *string
	_jsii_.Get(
		j,
		"profile",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleTranscoderJobConfigElementaryStreamsVideoStreamH264OutputReference) ProfileInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"profileInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleTranscoderJobConfigElementaryStreamsVideoStreamH264OutputReference) RateControlMode() *string {
	var returns *string
	_jsii_.Get(
		j,
		"rateControlMode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleTranscoderJobConfigElementaryStreamsVideoStreamH264OutputReference) RateControlModeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"rateControlModeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleTranscoderJobConfigElementaryStreamsVideoStreamH264OutputReference) Sdr() GoogleTranscoderJobConfigElementaryStreamsVideoStreamH264SdrOutputReference {
	var returns GoogleTranscoderJobConfigElementaryStreamsVideoStreamH264SdrOutputReference
	_jsii_.Get(
		j,
		"sdr",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleTranscoderJobConfigElementaryStreamsVideoStreamH264OutputReference) SdrInput() *GoogleTranscoderJobConfigElementaryStreamsVideoStreamH264Sdr {
	var returns *GoogleTranscoderJobConfigElementaryStreamsVideoStreamH264Sdr
	_jsii_.Get(
		j,
		"sdrInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleTranscoderJobConfigElementaryStreamsVideoStreamH264OutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleTranscoderJobConfigElementaryStreamsVideoStreamH264OutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleTranscoderJobConfigElementaryStreamsVideoStreamH264OutputReference) VbvFullnessBits() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"vbvFullnessBits",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleTranscoderJobConfigElementaryStreamsVideoStreamH264OutputReference) VbvFullnessBitsInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"vbvFullnessBitsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleTranscoderJobConfigElementaryStreamsVideoStreamH264OutputReference) VbvSizeBits() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"vbvSizeBits",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleTranscoderJobConfigElementaryStreamsVideoStreamH264OutputReference) VbvSizeBitsInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"vbvSizeBitsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleTranscoderJobConfigElementaryStreamsVideoStreamH264OutputReference) WidthPixels() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"widthPixels",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleTranscoderJobConfigElementaryStreamsVideoStreamH264OutputReference) WidthPixelsInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"widthPixelsInput",
		&returns,
	)
	return returns
}


func NewGoogleTranscoderJobConfigElementaryStreamsVideoStreamH264OutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) GoogleTranscoderJobConfigElementaryStreamsVideoStreamH264OutputReference {
	_init_.Initialize()

	if err := validateNewGoogleTranscoderJobConfigElementaryStreamsVideoStreamH264OutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_GoogleTranscoderJobConfigElementaryStreamsVideoStreamH264OutputReference{}

	_jsii_.Create(
		"@cdktf/provider-google-beta.googleTranscoderJob.GoogleTranscoderJobConfigElementaryStreamsVideoStreamH264OutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewGoogleTranscoderJobConfigElementaryStreamsVideoStreamH264OutputReference_Override(g GoogleTranscoderJobConfigElementaryStreamsVideoStreamH264OutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google-beta.googleTranscoderJob.GoogleTranscoderJobConfigElementaryStreamsVideoStreamH264OutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		g,
	)
}

func (j *jsiiProxy_GoogleTranscoderJobConfigElementaryStreamsVideoStreamH264OutputReference)SetBitrateBps(val *float64) {
	if err := j.validateSetBitrateBpsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"bitrateBps",
		val,
	)
}

func (j *jsiiProxy_GoogleTranscoderJobConfigElementaryStreamsVideoStreamH264OutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_GoogleTranscoderJobConfigElementaryStreamsVideoStreamH264OutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_GoogleTranscoderJobConfigElementaryStreamsVideoStreamH264OutputReference)SetCrfLevel(val *float64) {
	if err := j.validateSetCrfLevelParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"crfLevel",
		val,
	)
}

func (j *jsiiProxy_GoogleTranscoderJobConfigElementaryStreamsVideoStreamH264OutputReference)SetEntropyCoder(val *string) {
	if err := j.validateSetEntropyCoderParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"entropyCoder",
		val,
	)
}

func (j *jsiiProxy_GoogleTranscoderJobConfigElementaryStreamsVideoStreamH264OutputReference)SetFrameRate(val *float64) {
	if err := j.validateSetFrameRateParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"frameRate",
		val,
	)
}

func (j *jsiiProxy_GoogleTranscoderJobConfigElementaryStreamsVideoStreamH264OutputReference)SetGopDuration(val *string) {
	if err := j.validateSetGopDurationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"gopDuration",
		val,
	)
}

func (j *jsiiProxy_GoogleTranscoderJobConfigElementaryStreamsVideoStreamH264OutputReference)SetHeightPixels(val *float64) {
	if err := j.validateSetHeightPixelsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"heightPixels",
		val,
	)
}

func (j *jsiiProxy_GoogleTranscoderJobConfigElementaryStreamsVideoStreamH264OutputReference)SetInternalValue(val *GoogleTranscoderJobConfigElementaryStreamsVideoStreamH264) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_GoogleTranscoderJobConfigElementaryStreamsVideoStreamH264OutputReference)SetPixelFormat(val *string) {
	if err := j.validateSetPixelFormatParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"pixelFormat",
		val,
	)
}

func (j *jsiiProxy_GoogleTranscoderJobConfigElementaryStreamsVideoStreamH264OutputReference)SetPreset(val *string) {
	if err := j.validateSetPresetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"preset",
		val,
	)
}

func (j *jsiiProxy_GoogleTranscoderJobConfigElementaryStreamsVideoStreamH264OutputReference)SetProfile(val *string) {
	if err := j.validateSetProfileParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"profile",
		val,
	)
}

func (j *jsiiProxy_GoogleTranscoderJobConfigElementaryStreamsVideoStreamH264OutputReference)SetRateControlMode(val *string) {
	if err := j.validateSetRateControlModeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"rateControlMode",
		val,
	)
}

func (j *jsiiProxy_GoogleTranscoderJobConfigElementaryStreamsVideoStreamH264OutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_GoogleTranscoderJobConfigElementaryStreamsVideoStreamH264OutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_GoogleTranscoderJobConfigElementaryStreamsVideoStreamH264OutputReference)SetVbvFullnessBits(val *float64) {
	if err := j.validateSetVbvFullnessBitsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"vbvFullnessBits",
		val,
	)
}

func (j *jsiiProxy_GoogleTranscoderJobConfigElementaryStreamsVideoStreamH264OutputReference)SetVbvSizeBits(val *float64) {
	if err := j.validateSetVbvSizeBitsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"vbvSizeBits",
		val,
	)
}

func (j *jsiiProxy_GoogleTranscoderJobConfigElementaryStreamsVideoStreamH264OutputReference)SetWidthPixels(val *float64) {
	if err := j.validateSetWidthPixelsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"widthPixels",
		val,
	)
}

func (g *jsiiProxy_GoogleTranscoderJobConfigElementaryStreamsVideoStreamH264OutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleTranscoderJobConfigElementaryStreamsVideoStreamH264OutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (g *jsiiProxy_GoogleTranscoderJobConfigElementaryStreamsVideoStreamH264OutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GoogleTranscoderJobConfigElementaryStreamsVideoStreamH264OutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (g *jsiiProxy_GoogleTranscoderJobConfigElementaryStreamsVideoStreamH264OutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (g *jsiiProxy_GoogleTranscoderJobConfigElementaryStreamsVideoStreamH264OutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (g *jsiiProxy_GoogleTranscoderJobConfigElementaryStreamsVideoStreamH264OutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (g *jsiiProxy_GoogleTranscoderJobConfigElementaryStreamsVideoStreamH264OutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (g *jsiiProxy_GoogleTranscoderJobConfigElementaryStreamsVideoStreamH264OutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (g *jsiiProxy_GoogleTranscoderJobConfigElementaryStreamsVideoStreamH264OutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (g *jsiiProxy_GoogleTranscoderJobConfigElementaryStreamsVideoStreamH264OutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		g,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleTranscoderJobConfigElementaryStreamsVideoStreamH264OutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GoogleTranscoderJobConfigElementaryStreamsVideoStreamH264OutputReference) PutHlg(value *GoogleTranscoderJobConfigElementaryStreamsVideoStreamH264Hlg) {
	if err := g.validatePutHlgParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putHlg",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleTranscoderJobConfigElementaryStreamsVideoStreamH264OutputReference) PutSdr(value *GoogleTranscoderJobConfigElementaryStreamsVideoStreamH264Sdr) {
	if err := g.validatePutSdrParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putSdr",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleTranscoderJobConfigElementaryStreamsVideoStreamH264OutputReference) ResetCrfLevel() {
	_jsii_.InvokeVoid(
		g,
		"resetCrfLevel",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleTranscoderJobConfigElementaryStreamsVideoStreamH264OutputReference) ResetEntropyCoder() {
	_jsii_.InvokeVoid(
		g,
		"resetEntropyCoder",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleTranscoderJobConfigElementaryStreamsVideoStreamH264OutputReference) ResetGopDuration() {
	_jsii_.InvokeVoid(
		g,
		"resetGopDuration",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleTranscoderJobConfigElementaryStreamsVideoStreamH264OutputReference) ResetHeightPixels() {
	_jsii_.InvokeVoid(
		g,
		"resetHeightPixels",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleTranscoderJobConfigElementaryStreamsVideoStreamH264OutputReference) ResetHlg() {
	_jsii_.InvokeVoid(
		g,
		"resetHlg",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleTranscoderJobConfigElementaryStreamsVideoStreamH264OutputReference) ResetPixelFormat() {
	_jsii_.InvokeVoid(
		g,
		"resetPixelFormat",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleTranscoderJobConfigElementaryStreamsVideoStreamH264OutputReference) ResetPreset() {
	_jsii_.InvokeVoid(
		g,
		"resetPreset",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleTranscoderJobConfigElementaryStreamsVideoStreamH264OutputReference) ResetProfile() {
	_jsii_.InvokeVoid(
		g,
		"resetProfile",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleTranscoderJobConfigElementaryStreamsVideoStreamH264OutputReference) ResetRateControlMode() {
	_jsii_.InvokeVoid(
		g,
		"resetRateControlMode",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleTranscoderJobConfigElementaryStreamsVideoStreamH264OutputReference) ResetSdr() {
	_jsii_.InvokeVoid(
		g,
		"resetSdr",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleTranscoderJobConfigElementaryStreamsVideoStreamH264OutputReference) ResetVbvFullnessBits() {
	_jsii_.InvokeVoid(
		g,
		"resetVbvFullnessBits",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleTranscoderJobConfigElementaryStreamsVideoStreamH264OutputReference) ResetVbvSizeBits() {
	_jsii_.InvokeVoid(
		g,
		"resetVbvSizeBits",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleTranscoderJobConfigElementaryStreamsVideoStreamH264OutputReference) ResetWidthPixels() {
	_jsii_.InvokeVoid(
		g,
		"resetWidthPixels",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleTranscoderJobConfigElementaryStreamsVideoStreamH264OutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (g *jsiiProxy_GoogleTranscoderJobConfigElementaryStreamsVideoStreamH264OutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

