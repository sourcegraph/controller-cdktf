package cloudrunv2workerpool

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/google/jsii"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/sourcegraph/controller-cdktf/gen/google/cloudrunv2workerpool/internal"
)

type CloudRunV2WorkerPoolTemplateContainersOutputReference interface {
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
	DependsOn() *[]*string
	SetDependsOn(val *[]*string)
	DependsOnInput() *[]*string
	Env() CloudRunV2WorkerPoolTemplateContainersEnvList
	EnvInput() any
	// Experimental.
	Fqn() *string
	Image() *string
	SetImage(val *string)
	ImageInput() *string
	InternalValue() any
	SetInternalValue(val any)
	LivenessProbe() CloudRunV2WorkerPoolTemplateContainersLivenessProbeOutputReference
	LivenessProbeInput() *CloudRunV2WorkerPoolTemplateContainersLivenessProbe
	Name() *string
	SetName(val *string)
	NameInput() *string
	Resources() CloudRunV2WorkerPoolTemplateContainersResourcesOutputReference
	ResourcesInput() *CloudRunV2WorkerPoolTemplateContainersResources
	StartupProbe() CloudRunV2WorkerPoolTemplateContainersStartupProbeOutputReference
	StartupProbeInput() *CloudRunV2WorkerPoolTemplateContainersStartupProbe
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	VolumeMounts() CloudRunV2WorkerPoolTemplateContainersVolumeMountsList
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
	PutLivenessProbe(value *CloudRunV2WorkerPoolTemplateContainersLivenessProbe)
	PutResources(value *CloudRunV2WorkerPoolTemplateContainersResources)
	PutStartupProbe(value *CloudRunV2WorkerPoolTemplateContainersStartupProbe)
	PutVolumeMounts(value any)
	ResetArgs()
	ResetCommand()
	ResetDependsOn()
	ResetEnv()
	ResetLivenessProbe()
	ResetName()
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

// The jsii proxy struct for CloudRunV2WorkerPoolTemplateContainersOutputReference
type jsiiProxy_CloudRunV2WorkerPoolTemplateContainersOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_CloudRunV2WorkerPoolTemplateContainersOutputReference) Args() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"args",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudRunV2WorkerPoolTemplateContainersOutputReference) ArgsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"argsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudRunV2WorkerPoolTemplateContainersOutputReference) Command() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"command",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudRunV2WorkerPoolTemplateContainersOutputReference) CommandInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"commandInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudRunV2WorkerPoolTemplateContainersOutputReference) ComplexObjectIndex() any {
	var returns any
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudRunV2WorkerPoolTemplateContainersOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudRunV2WorkerPoolTemplateContainersOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudRunV2WorkerPoolTemplateContainersOutputReference) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudRunV2WorkerPoolTemplateContainersOutputReference) DependsOnInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudRunV2WorkerPoolTemplateContainersOutputReference) Env() CloudRunV2WorkerPoolTemplateContainersEnvList {
	var returns CloudRunV2WorkerPoolTemplateContainersEnvList
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudRunV2WorkerPoolTemplateContainersOutputReference) EnvInput() any {
	var returns any
	_jsii_.Get(
		j,
		"envInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudRunV2WorkerPoolTemplateContainersOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudRunV2WorkerPoolTemplateContainersOutputReference) Image() *string {
	var returns *string
	_jsii_.Get(
		j,
		"image",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudRunV2WorkerPoolTemplateContainersOutputReference) ImageInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"imageInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudRunV2WorkerPoolTemplateContainersOutputReference) InternalValue() any {
	var returns any
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudRunV2WorkerPoolTemplateContainersOutputReference) LivenessProbe() CloudRunV2WorkerPoolTemplateContainersLivenessProbeOutputReference {
	var returns CloudRunV2WorkerPoolTemplateContainersLivenessProbeOutputReference
	_jsii_.Get(
		j,
		"livenessProbe",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudRunV2WorkerPoolTemplateContainersOutputReference) LivenessProbeInput() *CloudRunV2WorkerPoolTemplateContainersLivenessProbe {
	var returns *CloudRunV2WorkerPoolTemplateContainersLivenessProbe
	_jsii_.Get(
		j,
		"livenessProbeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudRunV2WorkerPoolTemplateContainersOutputReference) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudRunV2WorkerPoolTemplateContainersOutputReference) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudRunV2WorkerPoolTemplateContainersOutputReference) Resources() CloudRunV2WorkerPoolTemplateContainersResourcesOutputReference {
	var returns CloudRunV2WorkerPoolTemplateContainersResourcesOutputReference
	_jsii_.Get(
		j,
		"resources",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudRunV2WorkerPoolTemplateContainersOutputReference) ResourcesInput() *CloudRunV2WorkerPoolTemplateContainersResources {
	var returns *CloudRunV2WorkerPoolTemplateContainersResources
	_jsii_.Get(
		j,
		"resourcesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudRunV2WorkerPoolTemplateContainersOutputReference) StartupProbe() CloudRunV2WorkerPoolTemplateContainersStartupProbeOutputReference {
	var returns CloudRunV2WorkerPoolTemplateContainersStartupProbeOutputReference
	_jsii_.Get(
		j,
		"startupProbe",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudRunV2WorkerPoolTemplateContainersOutputReference) StartupProbeInput() *CloudRunV2WorkerPoolTemplateContainersStartupProbe {
	var returns *CloudRunV2WorkerPoolTemplateContainersStartupProbe
	_jsii_.Get(
		j,
		"startupProbeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudRunV2WorkerPoolTemplateContainersOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudRunV2WorkerPoolTemplateContainersOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudRunV2WorkerPoolTemplateContainersOutputReference) VolumeMounts() CloudRunV2WorkerPoolTemplateContainersVolumeMountsList {
	var returns CloudRunV2WorkerPoolTemplateContainersVolumeMountsList
	_jsii_.Get(
		j,
		"volumeMounts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudRunV2WorkerPoolTemplateContainersOutputReference) VolumeMountsInput() any {
	var returns any
	_jsii_.Get(
		j,
		"volumeMountsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudRunV2WorkerPoolTemplateContainersOutputReference) WorkingDir() *string {
	var returns *string
	_jsii_.Get(
		j,
		"workingDir",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudRunV2WorkerPoolTemplateContainersOutputReference) WorkingDirInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"workingDirInput",
		&returns,
	)
	return returns
}

func NewCloudRunV2WorkerPoolTemplateContainersOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) CloudRunV2WorkerPoolTemplateContainersOutputReference {
	_init_.Initialize()

	if err := validateNewCloudRunV2WorkerPoolTemplateContainersOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_CloudRunV2WorkerPoolTemplateContainersOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-google.cloudRunV2WorkerPool.CloudRunV2WorkerPoolTemplateContainersOutputReference",
		[]any{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewCloudRunV2WorkerPoolTemplateContainersOutputReference_Override(c CloudRunV2WorkerPoolTemplateContainersOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google.cloudRunV2WorkerPool.CloudRunV2WorkerPoolTemplateContainersOutputReference",
		[]any{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		c,
	)
}

func (j *jsiiProxy_CloudRunV2WorkerPoolTemplateContainersOutputReference) SetArgs(val *[]*string) {
	if err := j.validateSetArgsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"args",
		val,
	)
}

func (j *jsiiProxy_CloudRunV2WorkerPoolTemplateContainersOutputReference) SetCommand(val *[]*string) {
	if err := j.validateSetCommandParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"command",
		val,
	)
}

func (j *jsiiProxy_CloudRunV2WorkerPoolTemplateContainersOutputReference) SetComplexObjectIndex(val any) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_CloudRunV2WorkerPoolTemplateContainersOutputReference) SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_CloudRunV2WorkerPoolTemplateContainersOutputReference) SetDependsOn(val *[]*string) {
	if err := j.validateSetDependsOnParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_CloudRunV2WorkerPoolTemplateContainersOutputReference) SetImage(val *string) {
	if err := j.validateSetImageParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"image",
		val,
	)
}

func (j *jsiiProxy_CloudRunV2WorkerPoolTemplateContainersOutputReference) SetInternalValue(val any) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_CloudRunV2WorkerPoolTemplateContainersOutputReference) SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_CloudRunV2WorkerPoolTemplateContainersOutputReference) SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_CloudRunV2WorkerPoolTemplateContainersOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_CloudRunV2WorkerPoolTemplateContainersOutputReference) SetWorkingDir(val *string) {
	if err := j.validateSetWorkingDirParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"workingDir",
		val,
	)
}

func (c *jsiiProxy_CloudRunV2WorkerPoolTemplateContainersOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CloudRunV2WorkerPoolTemplateContainersOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]any {
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

func (c *jsiiProxy_CloudRunV2WorkerPoolTemplateContainersOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (c *jsiiProxy_CloudRunV2WorkerPoolTemplateContainersOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (c *jsiiProxy_CloudRunV2WorkerPoolTemplateContainersOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (c *jsiiProxy_CloudRunV2WorkerPoolTemplateContainersOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (c *jsiiProxy_CloudRunV2WorkerPoolTemplateContainersOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (c *jsiiProxy_CloudRunV2WorkerPoolTemplateContainersOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (c *jsiiProxy_CloudRunV2WorkerPoolTemplateContainersOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (c *jsiiProxy_CloudRunV2WorkerPoolTemplateContainersOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (c *jsiiProxy_CloudRunV2WorkerPoolTemplateContainersOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		c,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CloudRunV2WorkerPoolTemplateContainersOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (c *jsiiProxy_CloudRunV2WorkerPoolTemplateContainersOutputReference) PutEnv(value any) {
	if err := c.validatePutEnvParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putEnv",
		[]any{value},
	)
}

func (c *jsiiProxy_CloudRunV2WorkerPoolTemplateContainersOutputReference) PutLivenessProbe(value *CloudRunV2WorkerPoolTemplateContainersLivenessProbe) {
	if err := c.validatePutLivenessProbeParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putLivenessProbe",
		[]any{value},
	)
}

func (c *jsiiProxy_CloudRunV2WorkerPoolTemplateContainersOutputReference) PutResources(value *CloudRunV2WorkerPoolTemplateContainersResources) {
	if err := c.validatePutResourcesParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putResources",
		[]any{value},
	)
}

func (c *jsiiProxy_CloudRunV2WorkerPoolTemplateContainersOutputReference) PutStartupProbe(value *CloudRunV2WorkerPoolTemplateContainersStartupProbe) {
	if err := c.validatePutStartupProbeParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putStartupProbe",
		[]any{value},
	)
}

func (c *jsiiProxy_CloudRunV2WorkerPoolTemplateContainersOutputReference) PutVolumeMounts(value any) {
	if err := c.validatePutVolumeMountsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putVolumeMounts",
		[]any{value},
	)
}

func (c *jsiiProxy_CloudRunV2WorkerPoolTemplateContainersOutputReference) ResetArgs() {
	_jsii_.InvokeVoid(
		c,
		"resetArgs",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CloudRunV2WorkerPoolTemplateContainersOutputReference) ResetCommand() {
	_jsii_.InvokeVoid(
		c,
		"resetCommand",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CloudRunV2WorkerPoolTemplateContainersOutputReference) ResetDependsOn() {
	_jsii_.InvokeVoid(
		c,
		"resetDependsOn",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CloudRunV2WorkerPoolTemplateContainersOutputReference) ResetEnv() {
	_jsii_.InvokeVoid(
		c,
		"resetEnv",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CloudRunV2WorkerPoolTemplateContainersOutputReference) ResetLivenessProbe() {
	_jsii_.InvokeVoid(
		c,
		"resetLivenessProbe",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CloudRunV2WorkerPoolTemplateContainersOutputReference) ResetName() {
	_jsii_.InvokeVoid(
		c,
		"resetName",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CloudRunV2WorkerPoolTemplateContainersOutputReference) ResetResources() {
	_jsii_.InvokeVoid(
		c,
		"resetResources",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CloudRunV2WorkerPoolTemplateContainersOutputReference) ResetStartupProbe() {
	_jsii_.InvokeVoid(
		c,
		"resetStartupProbe",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CloudRunV2WorkerPoolTemplateContainersOutputReference) ResetVolumeMounts() {
	_jsii_.InvokeVoid(
		c,
		"resetVolumeMounts",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CloudRunV2WorkerPoolTemplateContainersOutputReference) ResetWorkingDir() {
	_jsii_.InvokeVoid(
		c,
		"resetWorkingDir",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CloudRunV2WorkerPoolTemplateContainersOutputReference) Resolve(_context cdktf.IResolveContext) any {
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

func (c *jsiiProxy_CloudRunV2WorkerPoolTemplateContainersOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}
