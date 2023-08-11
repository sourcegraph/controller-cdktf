package googlecloudrunv2job

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/google_beta/jsii"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/sourcegraph/controller-cdktf/gen/google_beta/googlecloudrunv2job/internal"
)

type GoogleCloudRunV2JobTemplateTemplateContainersOutputReference interface {
	cdktf.ComplexObject
	Args() *[]*string
	SetArgs(val *[]*string)
	ArgsInput() *[]*string
	Command() *[]*string
	SetCommand(val *[]*string)
	CommandInput() *[]*string
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
	Env() GoogleCloudRunV2JobTemplateTemplateContainersEnvList
	EnvInput() interface{}
	// Experimental.
	Fqn() *string
	Image() *string
	SetImage(val *string)
	ImageInput() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	LivenessProbe() GoogleCloudRunV2JobTemplateTemplateContainersLivenessProbeOutputReference
	LivenessProbeInput() *GoogleCloudRunV2JobTemplateTemplateContainersLivenessProbe
	Name() *string
	SetName(val *string)
	NameInput() *string
	Ports() GoogleCloudRunV2JobTemplateTemplateContainersPortsList
	PortsInput() interface{}
	Resources() GoogleCloudRunV2JobTemplateTemplateContainersResourcesOutputReference
	ResourcesInput() *GoogleCloudRunV2JobTemplateTemplateContainersResources
	StartupProbe() GoogleCloudRunV2JobTemplateTemplateContainersStartupProbeOutputReference
	StartupProbeInput() *GoogleCloudRunV2JobTemplateTemplateContainersStartupProbe
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	VolumeMounts() GoogleCloudRunV2JobTemplateTemplateContainersVolumeMountsList
	VolumeMountsInput() interface{}
	WorkingDir() *string
	SetWorkingDir(val *string)
	WorkingDirInput() *string
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
	PutEnv(value interface{})
	PutLivenessProbe(value *GoogleCloudRunV2JobTemplateTemplateContainersLivenessProbe)
	PutPorts(value interface{})
	PutResources(value *GoogleCloudRunV2JobTemplateTemplateContainersResources)
	PutStartupProbe(value *GoogleCloudRunV2JobTemplateTemplateContainersStartupProbe)
	PutVolumeMounts(value interface{})
	ResetArgs()
	ResetCommand()
	ResetEnv()
	ResetLivenessProbe()
	ResetName()
	ResetPorts()
	ResetResources()
	ResetStartupProbe()
	ResetVolumeMounts()
	ResetWorkingDir()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for GoogleCloudRunV2JobTemplateTemplateContainersOutputReference
type jsiiProxy_GoogleCloudRunV2JobTemplateTemplateContainersOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_GoogleCloudRunV2JobTemplateTemplateContainersOutputReference) Args() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"args",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCloudRunV2JobTemplateTemplateContainersOutputReference) ArgsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"argsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCloudRunV2JobTemplateTemplateContainersOutputReference) Command() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"command",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCloudRunV2JobTemplateTemplateContainersOutputReference) CommandInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"commandInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCloudRunV2JobTemplateTemplateContainersOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCloudRunV2JobTemplateTemplateContainersOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCloudRunV2JobTemplateTemplateContainersOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCloudRunV2JobTemplateTemplateContainersOutputReference) Env() GoogleCloudRunV2JobTemplateTemplateContainersEnvList {
	var returns GoogleCloudRunV2JobTemplateTemplateContainersEnvList
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCloudRunV2JobTemplateTemplateContainersOutputReference) EnvInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"envInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCloudRunV2JobTemplateTemplateContainersOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCloudRunV2JobTemplateTemplateContainersOutputReference) Image() *string {
	var returns *string
	_jsii_.Get(
		j,
		"image",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCloudRunV2JobTemplateTemplateContainersOutputReference) ImageInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"imageInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCloudRunV2JobTemplateTemplateContainersOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCloudRunV2JobTemplateTemplateContainersOutputReference) LivenessProbe() GoogleCloudRunV2JobTemplateTemplateContainersLivenessProbeOutputReference {
	var returns GoogleCloudRunV2JobTemplateTemplateContainersLivenessProbeOutputReference
	_jsii_.Get(
		j,
		"livenessProbe",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCloudRunV2JobTemplateTemplateContainersOutputReference) LivenessProbeInput() *GoogleCloudRunV2JobTemplateTemplateContainersLivenessProbe {
	var returns *GoogleCloudRunV2JobTemplateTemplateContainersLivenessProbe
	_jsii_.Get(
		j,
		"livenessProbeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCloudRunV2JobTemplateTemplateContainersOutputReference) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCloudRunV2JobTemplateTemplateContainersOutputReference) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCloudRunV2JobTemplateTemplateContainersOutputReference) Ports() GoogleCloudRunV2JobTemplateTemplateContainersPortsList {
	var returns GoogleCloudRunV2JobTemplateTemplateContainersPortsList
	_jsii_.Get(
		j,
		"ports",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCloudRunV2JobTemplateTemplateContainersOutputReference) PortsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"portsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCloudRunV2JobTemplateTemplateContainersOutputReference) Resources() GoogleCloudRunV2JobTemplateTemplateContainersResourcesOutputReference {
	var returns GoogleCloudRunV2JobTemplateTemplateContainersResourcesOutputReference
	_jsii_.Get(
		j,
		"resources",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCloudRunV2JobTemplateTemplateContainersOutputReference) ResourcesInput() *GoogleCloudRunV2JobTemplateTemplateContainersResources {
	var returns *GoogleCloudRunV2JobTemplateTemplateContainersResources
	_jsii_.Get(
		j,
		"resourcesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCloudRunV2JobTemplateTemplateContainersOutputReference) StartupProbe() GoogleCloudRunV2JobTemplateTemplateContainersStartupProbeOutputReference {
	var returns GoogleCloudRunV2JobTemplateTemplateContainersStartupProbeOutputReference
	_jsii_.Get(
		j,
		"startupProbe",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCloudRunV2JobTemplateTemplateContainersOutputReference) StartupProbeInput() *GoogleCloudRunV2JobTemplateTemplateContainersStartupProbe {
	var returns *GoogleCloudRunV2JobTemplateTemplateContainersStartupProbe
	_jsii_.Get(
		j,
		"startupProbeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCloudRunV2JobTemplateTemplateContainersOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCloudRunV2JobTemplateTemplateContainersOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCloudRunV2JobTemplateTemplateContainersOutputReference) VolumeMounts() GoogleCloudRunV2JobTemplateTemplateContainersVolumeMountsList {
	var returns GoogleCloudRunV2JobTemplateTemplateContainersVolumeMountsList
	_jsii_.Get(
		j,
		"volumeMounts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCloudRunV2JobTemplateTemplateContainersOutputReference) VolumeMountsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"volumeMountsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCloudRunV2JobTemplateTemplateContainersOutputReference) WorkingDir() *string {
	var returns *string
	_jsii_.Get(
		j,
		"workingDir",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCloudRunV2JobTemplateTemplateContainersOutputReference) WorkingDirInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"workingDirInput",
		&returns,
	)
	return returns
}


func NewGoogleCloudRunV2JobTemplateTemplateContainersOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) GoogleCloudRunV2JobTemplateTemplateContainersOutputReference {
	_init_.Initialize()

	if err := validateNewGoogleCloudRunV2JobTemplateTemplateContainersOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_GoogleCloudRunV2JobTemplateTemplateContainersOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-google-beta.googleCloudRunV2Job.GoogleCloudRunV2JobTemplateTemplateContainersOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewGoogleCloudRunV2JobTemplateTemplateContainersOutputReference_Override(g GoogleCloudRunV2JobTemplateTemplateContainersOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google-beta.googleCloudRunV2Job.GoogleCloudRunV2JobTemplateTemplateContainersOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		g,
	)
}

func (j *jsiiProxy_GoogleCloudRunV2JobTemplateTemplateContainersOutputReference)SetArgs(val *[]*string) {
	if err := j.validateSetArgsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"args",
		val,
	)
}

func (j *jsiiProxy_GoogleCloudRunV2JobTemplateTemplateContainersOutputReference)SetCommand(val *[]*string) {
	if err := j.validateSetCommandParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"command",
		val,
	)
}

func (j *jsiiProxy_GoogleCloudRunV2JobTemplateTemplateContainersOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_GoogleCloudRunV2JobTemplateTemplateContainersOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_GoogleCloudRunV2JobTemplateTemplateContainersOutputReference)SetImage(val *string) {
	if err := j.validateSetImageParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"image",
		val,
	)
}

func (j *jsiiProxy_GoogleCloudRunV2JobTemplateTemplateContainersOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_GoogleCloudRunV2JobTemplateTemplateContainersOutputReference)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_GoogleCloudRunV2JobTemplateTemplateContainersOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_GoogleCloudRunV2JobTemplateTemplateContainersOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_GoogleCloudRunV2JobTemplateTemplateContainersOutputReference)SetWorkingDir(val *string) {
	if err := j.validateSetWorkingDirParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"workingDir",
		val,
	)
}

func (g *jsiiProxy_GoogleCloudRunV2JobTemplateTemplateContainersOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleCloudRunV2JobTemplateTemplateContainersOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (g *jsiiProxy_GoogleCloudRunV2JobTemplateTemplateContainersOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GoogleCloudRunV2JobTemplateTemplateContainersOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (g *jsiiProxy_GoogleCloudRunV2JobTemplateTemplateContainersOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (g *jsiiProxy_GoogleCloudRunV2JobTemplateTemplateContainersOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (g *jsiiProxy_GoogleCloudRunV2JobTemplateTemplateContainersOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (g *jsiiProxy_GoogleCloudRunV2JobTemplateTemplateContainersOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (g *jsiiProxy_GoogleCloudRunV2JobTemplateTemplateContainersOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (g *jsiiProxy_GoogleCloudRunV2JobTemplateTemplateContainersOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (g *jsiiProxy_GoogleCloudRunV2JobTemplateTemplateContainersOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		g,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleCloudRunV2JobTemplateTemplateContainersOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GoogleCloudRunV2JobTemplateTemplateContainersOutputReference) PutEnv(value interface{}) {
	if err := g.validatePutEnvParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putEnv",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleCloudRunV2JobTemplateTemplateContainersOutputReference) PutLivenessProbe(value *GoogleCloudRunV2JobTemplateTemplateContainersLivenessProbe) {
	if err := g.validatePutLivenessProbeParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putLivenessProbe",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleCloudRunV2JobTemplateTemplateContainersOutputReference) PutPorts(value interface{}) {
	if err := g.validatePutPortsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putPorts",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleCloudRunV2JobTemplateTemplateContainersOutputReference) PutResources(value *GoogleCloudRunV2JobTemplateTemplateContainersResources) {
	if err := g.validatePutResourcesParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putResources",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleCloudRunV2JobTemplateTemplateContainersOutputReference) PutStartupProbe(value *GoogleCloudRunV2JobTemplateTemplateContainersStartupProbe) {
	if err := g.validatePutStartupProbeParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putStartupProbe",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleCloudRunV2JobTemplateTemplateContainersOutputReference) PutVolumeMounts(value interface{}) {
	if err := g.validatePutVolumeMountsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putVolumeMounts",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleCloudRunV2JobTemplateTemplateContainersOutputReference) ResetArgs() {
	_jsii_.InvokeVoid(
		g,
		"resetArgs",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleCloudRunV2JobTemplateTemplateContainersOutputReference) ResetCommand() {
	_jsii_.InvokeVoid(
		g,
		"resetCommand",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleCloudRunV2JobTemplateTemplateContainersOutputReference) ResetEnv() {
	_jsii_.InvokeVoid(
		g,
		"resetEnv",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleCloudRunV2JobTemplateTemplateContainersOutputReference) ResetLivenessProbe() {
	_jsii_.InvokeVoid(
		g,
		"resetLivenessProbe",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleCloudRunV2JobTemplateTemplateContainersOutputReference) ResetName() {
	_jsii_.InvokeVoid(
		g,
		"resetName",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleCloudRunV2JobTemplateTemplateContainersOutputReference) ResetPorts() {
	_jsii_.InvokeVoid(
		g,
		"resetPorts",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleCloudRunV2JobTemplateTemplateContainersOutputReference) ResetResources() {
	_jsii_.InvokeVoid(
		g,
		"resetResources",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleCloudRunV2JobTemplateTemplateContainersOutputReference) ResetStartupProbe() {
	_jsii_.InvokeVoid(
		g,
		"resetStartupProbe",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleCloudRunV2JobTemplateTemplateContainersOutputReference) ResetVolumeMounts() {
	_jsii_.InvokeVoid(
		g,
		"resetVolumeMounts",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleCloudRunV2JobTemplateTemplateContainersOutputReference) ResetWorkingDir() {
	_jsii_.InvokeVoid(
		g,
		"resetWorkingDir",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleCloudRunV2JobTemplateTemplateContainersOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (g *jsiiProxy_GoogleCloudRunV2JobTemplateTemplateContainersOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

