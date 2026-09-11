package googlecesapp

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/google_beta/jsii"

	"github.com/open-constructs/cdk-terrain-go/cdktn"
	"github.com/sourcegraph/controller-cdktf/gen/google_beta/googlecesapp/internal"
)

type GoogleCesAppAudioProcessingConfigAmbientSoundConfigOutputReference interface {
	cdktn.ComplexObject
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
	// Experimental.
	Fqn() *string
	GcsUri() *string
	SetGcsUri(val *string)
	GcsUriInput() *string
	InternalValue() *GoogleCesAppAudioProcessingConfigAmbientSoundConfig
	SetInternalValue(val *GoogleCesAppAudioProcessingConfigAmbientSoundConfig)
	PrebuiltAmbientSound() *string
	SetPrebuiltAmbientSound(val *string)
	PrebuiltAmbientSoundInput() *string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktn.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktn.IInterpolatingParent)
	VolumeGainDb() *float64
	SetVolumeGainDb(val *float64)
	VolumeGainDbInput() *float64
	// Experimental.
	ComputeFqn() *string
	// Experimental.
	GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{}
	// Experimental.
	GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable
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
	InterpolationAsList() cdktn.IResolvable
	// Experimental.
	InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable
	ResetGcsUri()
	ResetPrebuiltAmbientSound()
	ResetVolumeGainDb()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for GoogleCesAppAudioProcessingConfigAmbientSoundConfigOutputReference
type jsiiProxy_GoogleCesAppAudioProcessingConfigAmbientSoundConfigOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_GoogleCesAppAudioProcessingConfigAmbientSoundConfigOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCesAppAudioProcessingConfigAmbientSoundConfigOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCesAppAudioProcessingConfigAmbientSoundConfigOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCesAppAudioProcessingConfigAmbientSoundConfigOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCesAppAudioProcessingConfigAmbientSoundConfigOutputReference) GcsUri() *string {
	var returns *string
	_jsii_.Get(
		j,
		"gcsUri",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCesAppAudioProcessingConfigAmbientSoundConfigOutputReference) GcsUriInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"gcsUriInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCesAppAudioProcessingConfigAmbientSoundConfigOutputReference) InternalValue() *GoogleCesAppAudioProcessingConfigAmbientSoundConfig {
	var returns *GoogleCesAppAudioProcessingConfigAmbientSoundConfig
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCesAppAudioProcessingConfigAmbientSoundConfigOutputReference) PrebuiltAmbientSound() *string {
	var returns *string
	_jsii_.Get(
		j,
		"prebuiltAmbientSound",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCesAppAudioProcessingConfigAmbientSoundConfigOutputReference) PrebuiltAmbientSoundInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"prebuiltAmbientSoundInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCesAppAudioProcessingConfigAmbientSoundConfigOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCesAppAudioProcessingConfigAmbientSoundConfigOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCesAppAudioProcessingConfigAmbientSoundConfigOutputReference) VolumeGainDb() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"volumeGainDb",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCesAppAudioProcessingConfigAmbientSoundConfigOutputReference) VolumeGainDbInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"volumeGainDbInput",
		&returns,
	)
	return returns
}


func NewGoogleCesAppAudioProcessingConfigAmbientSoundConfigOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) GoogleCesAppAudioProcessingConfigAmbientSoundConfigOutputReference {
	_init_.Initialize()

	if err := validateNewGoogleCesAppAudioProcessingConfigAmbientSoundConfigOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_GoogleCesAppAudioProcessingConfigAmbientSoundConfigOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-google-beta.googleCesApp.GoogleCesAppAudioProcessingConfigAmbientSoundConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewGoogleCesAppAudioProcessingConfigAmbientSoundConfigOutputReference_Override(g GoogleCesAppAudioProcessingConfigAmbientSoundConfigOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-google-beta.googleCesApp.GoogleCesAppAudioProcessingConfigAmbientSoundConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		g,
	)
}

func (j *jsiiProxy_GoogleCesAppAudioProcessingConfigAmbientSoundConfigOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_GoogleCesAppAudioProcessingConfigAmbientSoundConfigOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_GoogleCesAppAudioProcessingConfigAmbientSoundConfigOutputReference)SetGcsUri(val *string) {
	if err := j.validateSetGcsUriParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"gcsUri",
		val,
	)
}

func (j *jsiiProxy_GoogleCesAppAudioProcessingConfigAmbientSoundConfigOutputReference)SetInternalValue(val *GoogleCesAppAudioProcessingConfigAmbientSoundConfig) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_GoogleCesAppAudioProcessingConfigAmbientSoundConfigOutputReference)SetPrebuiltAmbientSound(val *string) {
	if err := j.validateSetPrebuiltAmbientSoundParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"prebuiltAmbientSound",
		val,
	)
}

func (j *jsiiProxy_GoogleCesAppAudioProcessingConfigAmbientSoundConfigOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_GoogleCesAppAudioProcessingConfigAmbientSoundConfigOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_GoogleCesAppAudioProcessingConfigAmbientSoundConfigOutputReference)SetVolumeGainDb(val *float64) {
	if err := j.validateSetVolumeGainDbParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"volumeGainDb",
		val,
	)
}

func (g *jsiiProxy_GoogleCesAppAudioProcessingConfigAmbientSoundConfigOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleCesAppAudioProcessingConfigAmbientSoundConfigOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (g *jsiiProxy_GoogleCesAppAudioProcessingConfigAmbientSoundConfigOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := g.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		g,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleCesAppAudioProcessingConfigAmbientSoundConfigOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (g *jsiiProxy_GoogleCesAppAudioProcessingConfigAmbientSoundConfigOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (g *jsiiProxy_GoogleCesAppAudioProcessingConfigAmbientSoundConfigOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (g *jsiiProxy_GoogleCesAppAudioProcessingConfigAmbientSoundConfigOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (g *jsiiProxy_GoogleCesAppAudioProcessingConfigAmbientSoundConfigOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (g *jsiiProxy_GoogleCesAppAudioProcessingConfigAmbientSoundConfigOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (g *jsiiProxy_GoogleCesAppAudioProcessingConfigAmbientSoundConfigOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (g *jsiiProxy_GoogleCesAppAudioProcessingConfigAmbientSoundConfigOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		g,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleCesAppAudioProcessingConfigAmbientSoundConfigOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := g.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		g,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleCesAppAudioProcessingConfigAmbientSoundConfigOutputReference) ResetGcsUri() {
	_jsii_.InvokeVoid(
		g,
		"resetGcsUri",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleCesAppAudioProcessingConfigAmbientSoundConfigOutputReference) ResetPrebuiltAmbientSound() {
	_jsii_.InvokeVoid(
		g,
		"resetPrebuiltAmbientSound",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleCesAppAudioProcessingConfigAmbientSoundConfigOutputReference) ResetVolumeGainDb() {
	_jsii_.InvokeVoid(
		g,
		"resetVolumeGainDb",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleCesAppAudioProcessingConfigAmbientSoundConfigOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
	if err := g.validateResolveParameters(context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		g,
		"resolve",
		[]interface{}{context},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleCesAppAudioProcessingConfigAmbientSoundConfigOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

