package cloudrunservice

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/google/jsii"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/sourcegraph/controller-cdktf/gen/google/cloudrunservice/internal"
)

type CloudRunServiceTemplateSpecContainersOutputReference interface {
	cdktf.ComplexObject
	Args() *[]*string
	SetArgs(val *[]*string)
	ArgsInput() *[]*string
	Command() *[]*string
	SetCommand(val *[]*string)
	CommandInput() *[]*string
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
	Env() CloudRunServiceTemplateSpecContainersEnvList
	EnvFrom() CloudRunServiceTemplateSpecContainersEnvFromList
	EnvFromInput() any
	EnvInput() any
	// Experimental.
	Fqn() *string
	Image() *string
	SetImage(val *string)
	ImageInput() *string
	InternalValue() any
	SetInternalValue(val any)
	LivenessProbe() CloudRunServiceTemplateSpecContainersLivenessProbeOutputReference
	LivenessProbeInput() *CloudRunServiceTemplateSpecContainersLivenessProbe
	Name() *string
	SetName(val *string)
	NameInput() *string
	Ports() CloudRunServiceTemplateSpecContainersPortsList
	PortsInput() any
	ReadinessProbe() CloudRunServiceTemplateSpecContainersReadinessProbeOutputReference
	ReadinessProbeInput() *CloudRunServiceTemplateSpecContainersReadinessProbe
	Resources() CloudRunServiceTemplateSpecContainersResourcesOutputReference
	ResourcesInput() *CloudRunServiceTemplateSpecContainersResources
	StartupProbe() CloudRunServiceTemplateSpecContainersStartupProbeOutputReference
	StartupProbeInput() *CloudRunServiceTemplateSpecContainersStartupProbe
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	VolumeMounts() CloudRunServiceTemplateSpecContainersVolumeMountsList
	VolumeMountsInput() any
	WorkingDir() *string
	SetWorkingDir(val *string)
	WorkingDirInput() *string
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
	PutEnv(value any)
	PutEnvFrom(value any)
	PutLivenessProbe(value *CloudRunServiceTemplateSpecContainersLivenessProbe)
	PutPorts(value any)
	PutReadinessProbe(value *CloudRunServiceTemplateSpecContainersReadinessProbe)
	PutResources(value *CloudRunServiceTemplateSpecContainersResources)
	PutStartupProbe(value *CloudRunServiceTemplateSpecContainersStartupProbe)
	PutVolumeMounts(value any)
	ResetArgs()
	ResetCommand()
	ResetEnv()
	ResetEnvFrom()
	ResetLivenessProbe()
	ResetName()
	ResetPorts()
	ResetReadinessProbe()
	ResetResources()
	ResetStartupProbe()
	ResetVolumeMounts()
	ResetWorkingDir()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) any
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for CloudRunServiceTemplateSpecContainersOutputReference
type jsiiProxy_CloudRunServiceTemplateSpecContainersOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_CloudRunServiceTemplateSpecContainersOutputReference) Args() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"args",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudRunServiceTemplateSpecContainersOutputReference) ArgsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"argsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudRunServiceTemplateSpecContainersOutputReference) Command() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"command",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudRunServiceTemplateSpecContainersOutputReference) CommandInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"commandInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudRunServiceTemplateSpecContainersOutputReference) ComplexObjectIndex() any {
	var returns any
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudRunServiceTemplateSpecContainersOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudRunServiceTemplateSpecContainersOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudRunServiceTemplateSpecContainersOutputReference) Env() CloudRunServiceTemplateSpecContainersEnvList {
	var returns CloudRunServiceTemplateSpecContainersEnvList
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudRunServiceTemplateSpecContainersOutputReference) EnvFrom() CloudRunServiceTemplateSpecContainersEnvFromList {
	var returns CloudRunServiceTemplateSpecContainersEnvFromList
	_jsii_.Get(
		j,
		"envFrom",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudRunServiceTemplateSpecContainersOutputReference) EnvFromInput() any {
	var returns any
	_jsii_.Get(
		j,
		"envFromInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudRunServiceTemplateSpecContainersOutputReference) EnvInput() any {
	var returns any
	_jsii_.Get(
		j,
		"envInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudRunServiceTemplateSpecContainersOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudRunServiceTemplateSpecContainersOutputReference) Image() *string {
	var returns *string
	_jsii_.Get(
		j,
		"image",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudRunServiceTemplateSpecContainersOutputReference) ImageInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"imageInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudRunServiceTemplateSpecContainersOutputReference) InternalValue() any {
	var returns any
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudRunServiceTemplateSpecContainersOutputReference) LivenessProbe() CloudRunServiceTemplateSpecContainersLivenessProbeOutputReference {
	var returns CloudRunServiceTemplateSpecContainersLivenessProbeOutputReference
	_jsii_.Get(
		j,
		"livenessProbe",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudRunServiceTemplateSpecContainersOutputReference) LivenessProbeInput() *CloudRunServiceTemplateSpecContainersLivenessProbe {
	var returns *CloudRunServiceTemplateSpecContainersLivenessProbe
	_jsii_.Get(
		j,
		"livenessProbeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudRunServiceTemplateSpecContainersOutputReference) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudRunServiceTemplateSpecContainersOutputReference) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudRunServiceTemplateSpecContainersOutputReference) Ports() CloudRunServiceTemplateSpecContainersPortsList {
	var returns CloudRunServiceTemplateSpecContainersPortsList
	_jsii_.Get(
		j,
		"ports",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudRunServiceTemplateSpecContainersOutputReference) PortsInput() any {
	var returns any
	_jsii_.Get(
		j,
		"portsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudRunServiceTemplateSpecContainersOutputReference) ReadinessProbe() CloudRunServiceTemplateSpecContainersReadinessProbeOutputReference {
	var returns CloudRunServiceTemplateSpecContainersReadinessProbeOutputReference
	_jsii_.Get(
		j,
		"readinessProbe",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudRunServiceTemplateSpecContainersOutputReference) ReadinessProbeInput() *CloudRunServiceTemplateSpecContainersReadinessProbe {
	var returns *CloudRunServiceTemplateSpecContainersReadinessProbe
	_jsii_.Get(
		j,
		"readinessProbeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudRunServiceTemplateSpecContainersOutputReference) Resources() CloudRunServiceTemplateSpecContainersResourcesOutputReference {
	var returns CloudRunServiceTemplateSpecContainersResourcesOutputReference
	_jsii_.Get(
		j,
		"resources",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudRunServiceTemplateSpecContainersOutputReference) ResourcesInput() *CloudRunServiceTemplateSpecContainersResources {
	var returns *CloudRunServiceTemplateSpecContainersResources
	_jsii_.Get(
		j,
		"resourcesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudRunServiceTemplateSpecContainersOutputReference) StartupProbe() CloudRunServiceTemplateSpecContainersStartupProbeOutputReference {
	var returns CloudRunServiceTemplateSpecContainersStartupProbeOutputReference
	_jsii_.Get(
		j,
		"startupProbe",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudRunServiceTemplateSpecContainersOutputReference) StartupProbeInput() *CloudRunServiceTemplateSpecContainersStartupProbe {
	var returns *CloudRunServiceTemplateSpecContainersStartupProbe
	_jsii_.Get(
		j,
		"startupProbeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudRunServiceTemplateSpecContainersOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudRunServiceTemplateSpecContainersOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudRunServiceTemplateSpecContainersOutputReference) VolumeMounts() CloudRunServiceTemplateSpecContainersVolumeMountsList {
	var returns CloudRunServiceTemplateSpecContainersVolumeMountsList
	_jsii_.Get(
		j,
		"volumeMounts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudRunServiceTemplateSpecContainersOutputReference) VolumeMountsInput() any {
	var returns any
	_jsii_.Get(
		j,
		"volumeMountsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudRunServiceTemplateSpecContainersOutputReference) WorkingDir() *string {
	var returns *string
	_jsii_.Get(
		j,
		"workingDir",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudRunServiceTemplateSpecContainersOutputReference) WorkingDirInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"workingDirInput",
		&returns,
	)
	return returns
}

func NewCloudRunServiceTemplateSpecContainersOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) CloudRunServiceTemplateSpecContainersOutputReference {
	_init_.Initialize()

	if err := validateNewCloudRunServiceTemplateSpecContainersOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_CloudRunServiceTemplateSpecContainersOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-google.cloudRunService.CloudRunServiceTemplateSpecContainersOutputReference",
		[]any{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewCloudRunServiceTemplateSpecContainersOutputReference_Override(c CloudRunServiceTemplateSpecContainersOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google.cloudRunService.CloudRunServiceTemplateSpecContainersOutputReference",
		[]any{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		c,
	)
}

func (j *jsiiProxy_CloudRunServiceTemplateSpecContainersOutputReference) SetArgs(val *[]*string) {
	if err := j.validateSetArgsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"args",
		val,
	)
}

func (j *jsiiProxy_CloudRunServiceTemplateSpecContainersOutputReference) SetCommand(val *[]*string) {
	if err := j.validateSetCommandParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"command",
		val,
	)
}

func (j *jsiiProxy_CloudRunServiceTemplateSpecContainersOutputReference) SetComplexObjectIndex(val any) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_CloudRunServiceTemplateSpecContainersOutputReference) SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_CloudRunServiceTemplateSpecContainersOutputReference) SetImage(val *string) {
	if err := j.validateSetImageParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"image",
		val,
	)
}

func (j *jsiiProxy_CloudRunServiceTemplateSpecContainersOutputReference) SetInternalValue(val any) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_CloudRunServiceTemplateSpecContainersOutputReference) SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_CloudRunServiceTemplateSpecContainersOutputReference) SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_CloudRunServiceTemplateSpecContainersOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_CloudRunServiceTemplateSpecContainersOutputReference) SetWorkingDir(val *string) {
	if err := j.validateSetWorkingDirParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"workingDir",
		val,
	)
}

func (c *jsiiProxy_CloudRunServiceTemplateSpecContainersOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CloudRunServiceTemplateSpecContainersOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]any {
	if err := c.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]any

	_jsii_.Invoke(
		c,
		"getAnyMapAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CloudRunServiceTemplateSpecContainersOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := c.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		c,
		"getBooleanAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CloudRunServiceTemplateSpecContainersOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := c.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		c,
		"getBooleanMapAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CloudRunServiceTemplateSpecContainersOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := c.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		c,
		"getListAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CloudRunServiceTemplateSpecContainersOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := c.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		c,
		"getNumberAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CloudRunServiceTemplateSpecContainersOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := c.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		c,
		"getNumberListAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CloudRunServiceTemplateSpecContainersOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := c.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		c,
		"getNumberMapAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CloudRunServiceTemplateSpecContainersOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := c.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		c,
		"getStringAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CloudRunServiceTemplateSpecContainersOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := c.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		c,
		"getStringMapAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CloudRunServiceTemplateSpecContainersOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		c,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CloudRunServiceTemplateSpecContainersOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	if err := c.validateInterpolationForAttributeParameters(property); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		c,
		"interpolationForAttribute",
		[]any{property},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CloudRunServiceTemplateSpecContainersOutputReference) PutEnv(value any) {
	if err := c.validatePutEnvParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putEnv",
		[]any{value},
	)
}

func (c *jsiiProxy_CloudRunServiceTemplateSpecContainersOutputReference) PutEnvFrom(value any) {
	if err := c.validatePutEnvFromParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putEnvFrom",
		[]any{value},
	)
}

func (c *jsiiProxy_CloudRunServiceTemplateSpecContainersOutputReference) PutLivenessProbe(value *CloudRunServiceTemplateSpecContainersLivenessProbe) {
	if err := c.validatePutLivenessProbeParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putLivenessProbe",
		[]any{value},
	)
}

func (c *jsiiProxy_CloudRunServiceTemplateSpecContainersOutputReference) PutPorts(value any) {
	if err := c.validatePutPortsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putPorts",
		[]any{value},
	)
}

func (c *jsiiProxy_CloudRunServiceTemplateSpecContainersOutputReference) PutReadinessProbe(value *CloudRunServiceTemplateSpecContainersReadinessProbe) {
	if err := c.validatePutReadinessProbeParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putReadinessProbe",
		[]any{value},
	)
}

func (c *jsiiProxy_CloudRunServiceTemplateSpecContainersOutputReference) PutResources(value *CloudRunServiceTemplateSpecContainersResources) {
	if err := c.validatePutResourcesParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putResources",
		[]any{value},
	)
}

func (c *jsiiProxy_CloudRunServiceTemplateSpecContainersOutputReference) PutStartupProbe(value *CloudRunServiceTemplateSpecContainersStartupProbe) {
	if err := c.validatePutStartupProbeParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putStartupProbe",
		[]any{value},
	)
}

func (c *jsiiProxy_CloudRunServiceTemplateSpecContainersOutputReference) PutVolumeMounts(value any) {
	if err := c.validatePutVolumeMountsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putVolumeMounts",
		[]any{value},
	)
}

func (c *jsiiProxy_CloudRunServiceTemplateSpecContainersOutputReference) ResetArgs() {
	_jsii_.InvokeVoid(
		c,
		"resetArgs",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CloudRunServiceTemplateSpecContainersOutputReference) ResetCommand() {
	_jsii_.InvokeVoid(
		c,
		"resetCommand",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CloudRunServiceTemplateSpecContainersOutputReference) ResetEnv() {
	_jsii_.InvokeVoid(
		c,
		"resetEnv",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CloudRunServiceTemplateSpecContainersOutputReference) ResetEnvFrom() {
	_jsii_.InvokeVoid(
		c,
		"resetEnvFrom",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CloudRunServiceTemplateSpecContainersOutputReference) ResetLivenessProbe() {
	_jsii_.InvokeVoid(
		c,
		"resetLivenessProbe",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CloudRunServiceTemplateSpecContainersOutputReference) ResetName() {
	_jsii_.InvokeVoid(
		c,
		"resetName",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CloudRunServiceTemplateSpecContainersOutputReference) ResetPorts() {
	_jsii_.InvokeVoid(
		c,
		"resetPorts",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CloudRunServiceTemplateSpecContainersOutputReference) ResetReadinessProbe() {
	_jsii_.InvokeVoid(
		c,
		"resetReadinessProbe",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CloudRunServiceTemplateSpecContainersOutputReference) ResetResources() {
	_jsii_.InvokeVoid(
		c,
		"resetResources",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CloudRunServiceTemplateSpecContainersOutputReference) ResetStartupProbe() {
	_jsii_.InvokeVoid(
		c,
		"resetStartupProbe",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CloudRunServiceTemplateSpecContainersOutputReference) ResetVolumeMounts() {
	_jsii_.InvokeVoid(
		c,
		"resetVolumeMounts",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CloudRunServiceTemplateSpecContainersOutputReference) ResetWorkingDir() {
	_jsii_.InvokeVoid(
		c,
		"resetWorkingDir",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CloudRunServiceTemplateSpecContainersOutputReference) Resolve(_context cdktf.IResolveContext) any {
	if err := c.validateResolveParameters(_context); err != nil {
		panic(err)
	}
	var returns any

	_jsii_.Invoke(
		c,
		"resolve",
		[]any{_context},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CloudRunServiceTemplateSpecContainersOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}
