package googlecloudbuildtrigger

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/google_beta/jsii"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/sourcegraph/controller-cdktf/gen/google_beta/googlecloudbuildtrigger/internal"
)

type GoogleCloudbuildTriggerBuildStepOutputReference interface {
	cdktf.ComplexObject
	AllowExitCodes() *[]*float64
	SetAllowExitCodes(val *[]*float64)
	AllowExitCodesInput() *[]*float64
	AllowFailure() interface{}
	SetAllowFailure(val interface{})
	AllowFailureInput() interface{}
	Args() *[]*string
	SetArgs(val *[]*string)
	ArgsInput() *[]*string
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
	Dir() *string
	SetDir(val *string)
	DirInput() *string
	Entrypoint() *string
	SetEntrypoint(val *string)
	EntrypointInput() *string
	Env() *[]*string
	SetEnv(val *[]*string)
	EnvInput() *[]*string
	// Experimental.
	Fqn() *string
	Id() *string
	SetId(val *string)
	IdInput() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	Name() *string
	SetName(val *string)
	NameInput() *string
	Script() *string
	SetScript(val *string)
	ScriptInput() *string
	SecretEnv() *[]*string
	SetSecretEnv(val *[]*string)
	SecretEnvInput() *[]*string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	Timeout() *string
	SetTimeout(val *string)
	TimeoutInput() *string
	Timing() *string
	SetTiming(val *string)
	TimingInput() *string
	Volumes() GoogleCloudbuildTriggerBuildStepVolumesList
	VolumesInput() interface{}
	WaitFor() *[]*string
	SetWaitFor(val *[]*string)
	WaitForInput() *[]*string
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
	PutVolumes(value interface{})
	ResetAllowExitCodes()
	ResetAllowFailure()
	ResetArgs()
	ResetDir()
	ResetEntrypoint()
	ResetEnv()
	ResetId()
	ResetScript()
	ResetSecretEnv()
	ResetTimeout()
	ResetTiming()
	ResetVolumes()
	ResetWaitFor()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for GoogleCloudbuildTriggerBuildStepOutputReference
type jsiiProxy_GoogleCloudbuildTriggerBuildStepOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_GoogleCloudbuildTriggerBuildStepOutputReference) AllowExitCodes() *[]*float64 {
	var returns *[]*float64
	_jsii_.Get(
		j,
		"allowExitCodes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCloudbuildTriggerBuildStepOutputReference) AllowExitCodesInput() *[]*float64 {
	var returns *[]*float64
	_jsii_.Get(
		j,
		"allowExitCodesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCloudbuildTriggerBuildStepOutputReference) AllowFailure() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"allowFailure",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCloudbuildTriggerBuildStepOutputReference) AllowFailureInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"allowFailureInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCloudbuildTriggerBuildStepOutputReference) Args() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"args",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCloudbuildTriggerBuildStepOutputReference) ArgsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"argsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCloudbuildTriggerBuildStepOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCloudbuildTriggerBuildStepOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCloudbuildTriggerBuildStepOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCloudbuildTriggerBuildStepOutputReference) Dir() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dir",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCloudbuildTriggerBuildStepOutputReference) DirInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dirInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCloudbuildTriggerBuildStepOutputReference) Entrypoint() *string {
	var returns *string
	_jsii_.Get(
		j,
		"entrypoint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCloudbuildTriggerBuildStepOutputReference) EntrypointInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"entrypointInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCloudbuildTriggerBuildStepOutputReference) Env() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCloudbuildTriggerBuildStepOutputReference) EnvInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"envInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCloudbuildTriggerBuildStepOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCloudbuildTriggerBuildStepOutputReference) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCloudbuildTriggerBuildStepOutputReference) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCloudbuildTriggerBuildStepOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCloudbuildTriggerBuildStepOutputReference) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCloudbuildTriggerBuildStepOutputReference) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCloudbuildTriggerBuildStepOutputReference) Script() *string {
	var returns *string
	_jsii_.Get(
		j,
		"script",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCloudbuildTriggerBuildStepOutputReference) ScriptInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"scriptInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCloudbuildTriggerBuildStepOutputReference) SecretEnv() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"secretEnv",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCloudbuildTriggerBuildStepOutputReference) SecretEnvInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"secretEnvInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCloudbuildTriggerBuildStepOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCloudbuildTriggerBuildStepOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCloudbuildTriggerBuildStepOutputReference) Timeout() *string {
	var returns *string
	_jsii_.Get(
		j,
		"timeout",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCloudbuildTriggerBuildStepOutputReference) TimeoutInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"timeoutInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCloudbuildTriggerBuildStepOutputReference) Timing() *string {
	var returns *string
	_jsii_.Get(
		j,
		"timing",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCloudbuildTriggerBuildStepOutputReference) TimingInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"timingInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCloudbuildTriggerBuildStepOutputReference) Volumes() GoogleCloudbuildTriggerBuildStepVolumesList {
	var returns GoogleCloudbuildTriggerBuildStepVolumesList
	_jsii_.Get(
		j,
		"volumes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCloudbuildTriggerBuildStepOutputReference) VolumesInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"volumesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCloudbuildTriggerBuildStepOutputReference) WaitFor() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"waitFor",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCloudbuildTriggerBuildStepOutputReference) WaitForInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"waitForInput",
		&returns,
	)
	return returns
}


func NewGoogleCloudbuildTriggerBuildStepOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) GoogleCloudbuildTriggerBuildStepOutputReference {
	_init_.Initialize()

	if err := validateNewGoogleCloudbuildTriggerBuildStepOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_GoogleCloudbuildTriggerBuildStepOutputReference{}

	_jsii_.Create(
		"google-beta.googleCloudbuildTrigger.GoogleCloudbuildTriggerBuildStepOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewGoogleCloudbuildTriggerBuildStepOutputReference_Override(g GoogleCloudbuildTriggerBuildStepOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"google-beta.googleCloudbuildTrigger.GoogleCloudbuildTriggerBuildStepOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		g,
	)
}

func (j *jsiiProxy_GoogleCloudbuildTriggerBuildStepOutputReference)SetAllowExitCodes(val *[]*float64) {
	if err := j.validateSetAllowExitCodesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"allowExitCodes",
		val,
	)
}

func (j *jsiiProxy_GoogleCloudbuildTriggerBuildStepOutputReference)SetAllowFailure(val interface{}) {
	if err := j.validateSetAllowFailureParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"allowFailure",
		val,
	)
}

func (j *jsiiProxy_GoogleCloudbuildTriggerBuildStepOutputReference)SetArgs(val *[]*string) {
	if err := j.validateSetArgsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"args",
		val,
	)
}

func (j *jsiiProxy_GoogleCloudbuildTriggerBuildStepOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_GoogleCloudbuildTriggerBuildStepOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_GoogleCloudbuildTriggerBuildStepOutputReference)SetDir(val *string) {
	if err := j.validateSetDirParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"dir",
		val,
	)
}

func (j *jsiiProxy_GoogleCloudbuildTriggerBuildStepOutputReference)SetEntrypoint(val *string) {
	if err := j.validateSetEntrypointParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"entrypoint",
		val,
	)
}

func (j *jsiiProxy_GoogleCloudbuildTriggerBuildStepOutputReference)SetEnv(val *[]*string) {
	if err := j.validateSetEnvParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"env",
		val,
	)
}

func (j *jsiiProxy_GoogleCloudbuildTriggerBuildStepOutputReference)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_GoogleCloudbuildTriggerBuildStepOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_GoogleCloudbuildTriggerBuildStepOutputReference)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_GoogleCloudbuildTriggerBuildStepOutputReference)SetScript(val *string) {
	if err := j.validateSetScriptParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"script",
		val,
	)
}

func (j *jsiiProxy_GoogleCloudbuildTriggerBuildStepOutputReference)SetSecretEnv(val *[]*string) {
	if err := j.validateSetSecretEnvParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"secretEnv",
		val,
	)
}

func (j *jsiiProxy_GoogleCloudbuildTriggerBuildStepOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_GoogleCloudbuildTriggerBuildStepOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_GoogleCloudbuildTriggerBuildStepOutputReference)SetTimeout(val *string) {
	if err := j.validateSetTimeoutParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"timeout",
		val,
	)
}

func (j *jsiiProxy_GoogleCloudbuildTriggerBuildStepOutputReference)SetTiming(val *string) {
	if err := j.validateSetTimingParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"timing",
		val,
	)
}

func (j *jsiiProxy_GoogleCloudbuildTriggerBuildStepOutputReference)SetWaitFor(val *[]*string) {
	if err := j.validateSetWaitForParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"waitFor",
		val,
	)
}

func (g *jsiiProxy_GoogleCloudbuildTriggerBuildStepOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleCloudbuildTriggerBuildStepOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (g *jsiiProxy_GoogleCloudbuildTriggerBuildStepOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GoogleCloudbuildTriggerBuildStepOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (g *jsiiProxy_GoogleCloudbuildTriggerBuildStepOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (g *jsiiProxy_GoogleCloudbuildTriggerBuildStepOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (g *jsiiProxy_GoogleCloudbuildTriggerBuildStepOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (g *jsiiProxy_GoogleCloudbuildTriggerBuildStepOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (g *jsiiProxy_GoogleCloudbuildTriggerBuildStepOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (g *jsiiProxy_GoogleCloudbuildTriggerBuildStepOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (g *jsiiProxy_GoogleCloudbuildTriggerBuildStepOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		g,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleCloudbuildTriggerBuildStepOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GoogleCloudbuildTriggerBuildStepOutputReference) PutVolumes(value interface{}) {
	if err := g.validatePutVolumesParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putVolumes",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleCloudbuildTriggerBuildStepOutputReference) ResetAllowExitCodes() {
	_jsii_.InvokeVoid(
		g,
		"resetAllowExitCodes",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleCloudbuildTriggerBuildStepOutputReference) ResetAllowFailure() {
	_jsii_.InvokeVoid(
		g,
		"resetAllowFailure",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleCloudbuildTriggerBuildStepOutputReference) ResetArgs() {
	_jsii_.InvokeVoid(
		g,
		"resetArgs",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleCloudbuildTriggerBuildStepOutputReference) ResetDir() {
	_jsii_.InvokeVoid(
		g,
		"resetDir",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleCloudbuildTriggerBuildStepOutputReference) ResetEntrypoint() {
	_jsii_.InvokeVoid(
		g,
		"resetEntrypoint",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleCloudbuildTriggerBuildStepOutputReference) ResetEnv() {
	_jsii_.InvokeVoid(
		g,
		"resetEnv",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleCloudbuildTriggerBuildStepOutputReference) ResetId() {
	_jsii_.InvokeVoid(
		g,
		"resetId",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleCloudbuildTriggerBuildStepOutputReference) ResetScript() {
	_jsii_.InvokeVoid(
		g,
		"resetScript",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleCloudbuildTriggerBuildStepOutputReference) ResetSecretEnv() {
	_jsii_.InvokeVoid(
		g,
		"resetSecretEnv",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleCloudbuildTriggerBuildStepOutputReference) ResetTimeout() {
	_jsii_.InvokeVoid(
		g,
		"resetTimeout",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleCloudbuildTriggerBuildStepOutputReference) ResetTiming() {
	_jsii_.InvokeVoid(
		g,
		"resetTiming",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleCloudbuildTriggerBuildStepOutputReference) ResetVolumes() {
	_jsii_.InvokeVoid(
		g,
		"resetVolumes",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleCloudbuildTriggerBuildStepOutputReference) ResetWaitFor() {
	_jsii_.InvokeVoid(
		g,
		"resetWaitFor",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleCloudbuildTriggerBuildStepOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (g *jsiiProxy_GoogleCloudbuildTriggerBuildStepOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

