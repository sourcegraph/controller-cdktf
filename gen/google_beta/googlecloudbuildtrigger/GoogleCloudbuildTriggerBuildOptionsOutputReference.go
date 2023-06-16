package googlecloudbuildtrigger

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/google_beta/jsii"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/sourcegraph/controller-cdktf/gen/google_beta/googlecloudbuildtrigger/internal"
)

type GoogleCloudbuildTriggerBuildOptionsOutputReference interface {
	cdktf.ComplexObject
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
	DiskSizeGb() *float64
	SetDiskSizeGb(val *float64)
	DiskSizeGbInput() *float64
	DynamicSubstitutions() interface{}
	SetDynamicSubstitutions(val interface{})
	DynamicSubstitutionsInput() interface{}
	Env() *[]*string
	SetEnv(val *[]*string)
	EnvInput() *[]*string
	// Experimental.
	Fqn() *string
	InternalValue() *GoogleCloudbuildTriggerBuildOptions
	SetInternalValue(val *GoogleCloudbuildTriggerBuildOptions)
	Logging() *string
	SetLogging(val *string)
	LoggingInput() *string
	LogStreamingOption() *string
	SetLogStreamingOption(val *string)
	LogStreamingOptionInput() *string
	MachineType() *string
	SetMachineType(val *string)
	MachineTypeInput() *string
	RequestedVerifyOption() *string
	SetRequestedVerifyOption(val *string)
	RequestedVerifyOptionInput() *string
	SecretEnv() *[]*string
	SetSecretEnv(val *[]*string)
	SecretEnvInput() *[]*string
	SourceProvenanceHash() *[]*string
	SetSourceProvenanceHash(val *[]*string)
	SourceProvenanceHashInput() *[]*string
	SubstitutionOption() *string
	SetSubstitutionOption(val *string)
	SubstitutionOptionInput() *string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	Volumes() GoogleCloudbuildTriggerBuildOptionsVolumesList
	VolumesInput() interface{}
	WorkerPool() *string
	SetWorkerPool(val *string)
	WorkerPoolInput() *string
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
	ResetDiskSizeGb()
	ResetDynamicSubstitutions()
	ResetEnv()
	ResetLogging()
	ResetLogStreamingOption()
	ResetMachineType()
	ResetRequestedVerifyOption()
	ResetSecretEnv()
	ResetSourceProvenanceHash()
	ResetSubstitutionOption()
	ResetVolumes()
	ResetWorkerPool()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for GoogleCloudbuildTriggerBuildOptionsOutputReference
type jsiiProxy_GoogleCloudbuildTriggerBuildOptionsOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_GoogleCloudbuildTriggerBuildOptionsOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCloudbuildTriggerBuildOptionsOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCloudbuildTriggerBuildOptionsOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCloudbuildTriggerBuildOptionsOutputReference) DiskSizeGb() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"diskSizeGb",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCloudbuildTriggerBuildOptionsOutputReference) DiskSizeGbInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"diskSizeGbInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCloudbuildTriggerBuildOptionsOutputReference) DynamicSubstitutions() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"dynamicSubstitutions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCloudbuildTriggerBuildOptionsOutputReference) DynamicSubstitutionsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"dynamicSubstitutionsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCloudbuildTriggerBuildOptionsOutputReference) Env() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCloudbuildTriggerBuildOptionsOutputReference) EnvInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"envInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCloudbuildTriggerBuildOptionsOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCloudbuildTriggerBuildOptionsOutputReference) InternalValue() *GoogleCloudbuildTriggerBuildOptions {
	var returns *GoogleCloudbuildTriggerBuildOptions
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCloudbuildTriggerBuildOptionsOutputReference) Logging() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logging",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCloudbuildTriggerBuildOptionsOutputReference) LoggingInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"loggingInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCloudbuildTriggerBuildOptionsOutputReference) LogStreamingOption() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logStreamingOption",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCloudbuildTriggerBuildOptionsOutputReference) LogStreamingOptionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logStreamingOptionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCloudbuildTriggerBuildOptionsOutputReference) MachineType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"machineType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCloudbuildTriggerBuildOptionsOutputReference) MachineTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"machineTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCloudbuildTriggerBuildOptionsOutputReference) RequestedVerifyOption() *string {
	var returns *string
	_jsii_.Get(
		j,
		"requestedVerifyOption",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCloudbuildTriggerBuildOptionsOutputReference) RequestedVerifyOptionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"requestedVerifyOptionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCloudbuildTriggerBuildOptionsOutputReference) SecretEnv() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"secretEnv",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCloudbuildTriggerBuildOptionsOutputReference) SecretEnvInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"secretEnvInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCloudbuildTriggerBuildOptionsOutputReference) SourceProvenanceHash() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"sourceProvenanceHash",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCloudbuildTriggerBuildOptionsOutputReference) SourceProvenanceHashInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"sourceProvenanceHashInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCloudbuildTriggerBuildOptionsOutputReference) SubstitutionOption() *string {
	var returns *string
	_jsii_.Get(
		j,
		"substitutionOption",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCloudbuildTriggerBuildOptionsOutputReference) SubstitutionOptionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"substitutionOptionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCloudbuildTriggerBuildOptionsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCloudbuildTriggerBuildOptionsOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCloudbuildTriggerBuildOptionsOutputReference) Volumes() GoogleCloudbuildTriggerBuildOptionsVolumesList {
	var returns GoogleCloudbuildTriggerBuildOptionsVolumesList
	_jsii_.Get(
		j,
		"volumes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCloudbuildTriggerBuildOptionsOutputReference) VolumesInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"volumesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCloudbuildTriggerBuildOptionsOutputReference) WorkerPool() *string {
	var returns *string
	_jsii_.Get(
		j,
		"workerPool",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCloudbuildTriggerBuildOptionsOutputReference) WorkerPoolInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"workerPoolInput",
		&returns,
	)
	return returns
}


func NewGoogleCloudbuildTriggerBuildOptionsOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) GoogleCloudbuildTriggerBuildOptionsOutputReference {
	_init_.Initialize()

	if err := validateNewGoogleCloudbuildTriggerBuildOptionsOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_GoogleCloudbuildTriggerBuildOptionsOutputReference{}

	_jsii_.Create(
		"google-beta.googleCloudbuildTrigger.GoogleCloudbuildTriggerBuildOptionsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewGoogleCloudbuildTriggerBuildOptionsOutputReference_Override(g GoogleCloudbuildTriggerBuildOptionsOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"google-beta.googleCloudbuildTrigger.GoogleCloudbuildTriggerBuildOptionsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		g,
	)
}

func (j *jsiiProxy_GoogleCloudbuildTriggerBuildOptionsOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_GoogleCloudbuildTriggerBuildOptionsOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_GoogleCloudbuildTriggerBuildOptionsOutputReference)SetDiskSizeGb(val *float64) {
	if err := j.validateSetDiskSizeGbParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"diskSizeGb",
		val,
	)
}

func (j *jsiiProxy_GoogleCloudbuildTriggerBuildOptionsOutputReference)SetDynamicSubstitutions(val interface{}) {
	if err := j.validateSetDynamicSubstitutionsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"dynamicSubstitutions",
		val,
	)
}

func (j *jsiiProxy_GoogleCloudbuildTriggerBuildOptionsOutputReference)SetEnv(val *[]*string) {
	if err := j.validateSetEnvParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"env",
		val,
	)
}

func (j *jsiiProxy_GoogleCloudbuildTriggerBuildOptionsOutputReference)SetInternalValue(val *GoogleCloudbuildTriggerBuildOptions) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_GoogleCloudbuildTriggerBuildOptionsOutputReference)SetLogging(val *string) {
	if err := j.validateSetLoggingParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"logging",
		val,
	)
}

func (j *jsiiProxy_GoogleCloudbuildTriggerBuildOptionsOutputReference)SetLogStreamingOption(val *string) {
	if err := j.validateSetLogStreamingOptionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"logStreamingOption",
		val,
	)
}

func (j *jsiiProxy_GoogleCloudbuildTriggerBuildOptionsOutputReference)SetMachineType(val *string) {
	if err := j.validateSetMachineTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"machineType",
		val,
	)
}

func (j *jsiiProxy_GoogleCloudbuildTriggerBuildOptionsOutputReference)SetRequestedVerifyOption(val *string) {
	if err := j.validateSetRequestedVerifyOptionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"requestedVerifyOption",
		val,
	)
}

func (j *jsiiProxy_GoogleCloudbuildTriggerBuildOptionsOutputReference)SetSecretEnv(val *[]*string) {
	if err := j.validateSetSecretEnvParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"secretEnv",
		val,
	)
}

func (j *jsiiProxy_GoogleCloudbuildTriggerBuildOptionsOutputReference)SetSourceProvenanceHash(val *[]*string) {
	if err := j.validateSetSourceProvenanceHashParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"sourceProvenanceHash",
		val,
	)
}

func (j *jsiiProxy_GoogleCloudbuildTriggerBuildOptionsOutputReference)SetSubstitutionOption(val *string) {
	if err := j.validateSetSubstitutionOptionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"substitutionOption",
		val,
	)
}

func (j *jsiiProxy_GoogleCloudbuildTriggerBuildOptionsOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_GoogleCloudbuildTriggerBuildOptionsOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_GoogleCloudbuildTriggerBuildOptionsOutputReference)SetWorkerPool(val *string) {
	if err := j.validateSetWorkerPoolParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"workerPool",
		val,
	)
}

func (g *jsiiProxy_GoogleCloudbuildTriggerBuildOptionsOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleCloudbuildTriggerBuildOptionsOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (g *jsiiProxy_GoogleCloudbuildTriggerBuildOptionsOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GoogleCloudbuildTriggerBuildOptionsOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (g *jsiiProxy_GoogleCloudbuildTriggerBuildOptionsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (g *jsiiProxy_GoogleCloudbuildTriggerBuildOptionsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (g *jsiiProxy_GoogleCloudbuildTriggerBuildOptionsOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (g *jsiiProxy_GoogleCloudbuildTriggerBuildOptionsOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (g *jsiiProxy_GoogleCloudbuildTriggerBuildOptionsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (g *jsiiProxy_GoogleCloudbuildTriggerBuildOptionsOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (g *jsiiProxy_GoogleCloudbuildTriggerBuildOptionsOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		g,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleCloudbuildTriggerBuildOptionsOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GoogleCloudbuildTriggerBuildOptionsOutputReference) PutVolumes(value interface{}) {
	if err := g.validatePutVolumesParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putVolumes",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleCloudbuildTriggerBuildOptionsOutputReference) ResetDiskSizeGb() {
	_jsii_.InvokeVoid(
		g,
		"resetDiskSizeGb",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleCloudbuildTriggerBuildOptionsOutputReference) ResetDynamicSubstitutions() {
	_jsii_.InvokeVoid(
		g,
		"resetDynamicSubstitutions",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleCloudbuildTriggerBuildOptionsOutputReference) ResetEnv() {
	_jsii_.InvokeVoid(
		g,
		"resetEnv",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleCloudbuildTriggerBuildOptionsOutputReference) ResetLogging() {
	_jsii_.InvokeVoid(
		g,
		"resetLogging",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleCloudbuildTriggerBuildOptionsOutputReference) ResetLogStreamingOption() {
	_jsii_.InvokeVoid(
		g,
		"resetLogStreamingOption",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleCloudbuildTriggerBuildOptionsOutputReference) ResetMachineType() {
	_jsii_.InvokeVoid(
		g,
		"resetMachineType",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleCloudbuildTriggerBuildOptionsOutputReference) ResetRequestedVerifyOption() {
	_jsii_.InvokeVoid(
		g,
		"resetRequestedVerifyOption",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleCloudbuildTriggerBuildOptionsOutputReference) ResetSecretEnv() {
	_jsii_.InvokeVoid(
		g,
		"resetSecretEnv",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleCloudbuildTriggerBuildOptionsOutputReference) ResetSourceProvenanceHash() {
	_jsii_.InvokeVoid(
		g,
		"resetSourceProvenanceHash",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleCloudbuildTriggerBuildOptionsOutputReference) ResetSubstitutionOption() {
	_jsii_.InvokeVoid(
		g,
		"resetSubstitutionOption",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleCloudbuildTriggerBuildOptionsOutputReference) ResetVolumes() {
	_jsii_.InvokeVoid(
		g,
		"resetVolumes",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleCloudbuildTriggerBuildOptionsOutputReference) ResetWorkerPool() {
	_jsii_.InvokeVoid(
		g,
		"resetWorkerPool",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleCloudbuildTriggerBuildOptionsOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (g *jsiiProxy_GoogleCloudbuildTriggerBuildOptionsOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

