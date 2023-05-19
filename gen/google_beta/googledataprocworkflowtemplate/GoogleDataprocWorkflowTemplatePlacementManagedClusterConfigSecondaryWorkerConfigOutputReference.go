package googledataprocworkflowtemplate

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/google_beta/jsii"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/sourcegraph/controller-cdktf/gen/google_beta/googledataprocworkflowtemplate/internal"
)

type GoogleDataprocWorkflowTemplatePlacementManagedClusterConfigSecondaryWorkerConfigOutputReference interface {
	cdktf.ComplexObject
	Accelerators() GoogleDataprocWorkflowTemplatePlacementManagedClusterConfigSecondaryWorkerConfigAcceleratorsList
	AcceleratorsInput() interface{}
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
	DiskConfig() GoogleDataprocWorkflowTemplatePlacementManagedClusterConfigSecondaryWorkerConfigDiskConfigOutputReference
	DiskConfigInput() *GoogleDataprocWorkflowTemplatePlacementManagedClusterConfigSecondaryWorkerConfigDiskConfig
	// Experimental.
	Fqn() *string
	Image() *string
	SetImage(val *string)
	ImageInput() *string
	InstanceNames() *[]*string
	InternalValue() *GoogleDataprocWorkflowTemplatePlacementManagedClusterConfigSecondaryWorkerConfig
	SetInternalValue(val *GoogleDataprocWorkflowTemplatePlacementManagedClusterConfigSecondaryWorkerConfig)
	IsPreemptible() cdktf.IResolvable
	MachineType() *string
	SetMachineType(val *string)
	MachineTypeInput() *string
	ManagedGroupConfig() GoogleDataprocWorkflowTemplatePlacementManagedClusterConfigSecondaryWorkerConfigManagedGroupConfigList
	MinCpuPlatform() *string
	SetMinCpuPlatform(val *string)
	MinCpuPlatformInput() *string
	NumInstances() *float64
	SetNumInstances(val *float64)
	NumInstancesInput() *float64
	Preemptibility() *string
	SetPreemptibility(val *string)
	PreemptibilityInput() *string
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
	PutAccelerators(value interface{})
	PutDiskConfig(value *GoogleDataprocWorkflowTemplatePlacementManagedClusterConfigSecondaryWorkerConfigDiskConfig)
	ResetAccelerators()
	ResetDiskConfig()
	ResetImage()
	ResetMachineType()
	ResetMinCpuPlatform()
	ResetNumInstances()
	ResetPreemptibility()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for GoogleDataprocWorkflowTemplatePlacementManagedClusterConfigSecondaryWorkerConfigOutputReference
type jsiiProxy_GoogleDataprocWorkflowTemplatePlacementManagedClusterConfigSecondaryWorkerConfigOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_GoogleDataprocWorkflowTemplatePlacementManagedClusterConfigSecondaryWorkerConfigOutputReference) Accelerators() GoogleDataprocWorkflowTemplatePlacementManagedClusterConfigSecondaryWorkerConfigAcceleratorsList {
	var returns GoogleDataprocWorkflowTemplatePlacementManagedClusterConfigSecondaryWorkerConfigAcceleratorsList
	_jsii_.Get(
		j,
		"accelerators",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataprocWorkflowTemplatePlacementManagedClusterConfigSecondaryWorkerConfigOutputReference) AcceleratorsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"acceleratorsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataprocWorkflowTemplatePlacementManagedClusterConfigSecondaryWorkerConfigOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataprocWorkflowTemplatePlacementManagedClusterConfigSecondaryWorkerConfigOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataprocWorkflowTemplatePlacementManagedClusterConfigSecondaryWorkerConfigOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataprocWorkflowTemplatePlacementManagedClusterConfigSecondaryWorkerConfigOutputReference) DiskConfig() GoogleDataprocWorkflowTemplatePlacementManagedClusterConfigSecondaryWorkerConfigDiskConfigOutputReference {
	var returns GoogleDataprocWorkflowTemplatePlacementManagedClusterConfigSecondaryWorkerConfigDiskConfigOutputReference
	_jsii_.Get(
		j,
		"diskConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataprocWorkflowTemplatePlacementManagedClusterConfigSecondaryWorkerConfigOutputReference) DiskConfigInput() *GoogleDataprocWorkflowTemplatePlacementManagedClusterConfigSecondaryWorkerConfigDiskConfig {
	var returns *GoogleDataprocWorkflowTemplatePlacementManagedClusterConfigSecondaryWorkerConfigDiskConfig
	_jsii_.Get(
		j,
		"diskConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataprocWorkflowTemplatePlacementManagedClusterConfigSecondaryWorkerConfigOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataprocWorkflowTemplatePlacementManagedClusterConfigSecondaryWorkerConfigOutputReference) Image() *string {
	var returns *string
	_jsii_.Get(
		j,
		"image",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataprocWorkflowTemplatePlacementManagedClusterConfigSecondaryWorkerConfigOutputReference) ImageInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"imageInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataprocWorkflowTemplatePlacementManagedClusterConfigSecondaryWorkerConfigOutputReference) InstanceNames() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"instanceNames",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataprocWorkflowTemplatePlacementManagedClusterConfigSecondaryWorkerConfigOutputReference) InternalValue() *GoogleDataprocWorkflowTemplatePlacementManagedClusterConfigSecondaryWorkerConfig {
	var returns *GoogleDataprocWorkflowTemplatePlacementManagedClusterConfigSecondaryWorkerConfig
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataprocWorkflowTemplatePlacementManagedClusterConfigSecondaryWorkerConfigOutputReference) IsPreemptible() cdktf.IResolvable {
	var returns cdktf.IResolvable
	_jsii_.Get(
		j,
		"isPreemptible",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataprocWorkflowTemplatePlacementManagedClusterConfigSecondaryWorkerConfigOutputReference) MachineType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"machineType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataprocWorkflowTemplatePlacementManagedClusterConfigSecondaryWorkerConfigOutputReference) MachineTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"machineTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataprocWorkflowTemplatePlacementManagedClusterConfigSecondaryWorkerConfigOutputReference) ManagedGroupConfig() GoogleDataprocWorkflowTemplatePlacementManagedClusterConfigSecondaryWorkerConfigManagedGroupConfigList {
	var returns GoogleDataprocWorkflowTemplatePlacementManagedClusterConfigSecondaryWorkerConfigManagedGroupConfigList
	_jsii_.Get(
		j,
		"managedGroupConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataprocWorkflowTemplatePlacementManagedClusterConfigSecondaryWorkerConfigOutputReference) MinCpuPlatform() *string {
	var returns *string
	_jsii_.Get(
		j,
		"minCpuPlatform",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataprocWorkflowTemplatePlacementManagedClusterConfigSecondaryWorkerConfigOutputReference) MinCpuPlatformInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"minCpuPlatformInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataprocWorkflowTemplatePlacementManagedClusterConfigSecondaryWorkerConfigOutputReference) NumInstances() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"numInstances",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataprocWorkflowTemplatePlacementManagedClusterConfigSecondaryWorkerConfigOutputReference) NumInstancesInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"numInstancesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataprocWorkflowTemplatePlacementManagedClusterConfigSecondaryWorkerConfigOutputReference) Preemptibility() *string {
	var returns *string
	_jsii_.Get(
		j,
		"preemptibility",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataprocWorkflowTemplatePlacementManagedClusterConfigSecondaryWorkerConfigOutputReference) PreemptibilityInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"preemptibilityInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataprocWorkflowTemplatePlacementManagedClusterConfigSecondaryWorkerConfigOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataprocWorkflowTemplatePlacementManagedClusterConfigSecondaryWorkerConfigOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewGoogleDataprocWorkflowTemplatePlacementManagedClusterConfigSecondaryWorkerConfigOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) GoogleDataprocWorkflowTemplatePlacementManagedClusterConfigSecondaryWorkerConfigOutputReference {
	_init_.Initialize()

	if err := validateNewGoogleDataprocWorkflowTemplatePlacementManagedClusterConfigSecondaryWorkerConfigOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_GoogleDataprocWorkflowTemplatePlacementManagedClusterConfigSecondaryWorkerConfigOutputReference{}

	_jsii_.Create(
		"google-beta.googleDataprocWorkflowTemplate.GoogleDataprocWorkflowTemplatePlacementManagedClusterConfigSecondaryWorkerConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewGoogleDataprocWorkflowTemplatePlacementManagedClusterConfigSecondaryWorkerConfigOutputReference_Override(g GoogleDataprocWorkflowTemplatePlacementManagedClusterConfigSecondaryWorkerConfigOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"google-beta.googleDataprocWorkflowTemplate.GoogleDataprocWorkflowTemplatePlacementManagedClusterConfigSecondaryWorkerConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		g,
	)
}

func (j *jsiiProxy_GoogleDataprocWorkflowTemplatePlacementManagedClusterConfigSecondaryWorkerConfigOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_GoogleDataprocWorkflowTemplatePlacementManagedClusterConfigSecondaryWorkerConfigOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_GoogleDataprocWorkflowTemplatePlacementManagedClusterConfigSecondaryWorkerConfigOutputReference)SetImage(val *string) {
	if err := j.validateSetImageParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"image",
		val,
	)
}

func (j *jsiiProxy_GoogleDataprocWorkflowTemplatePlacementManagedClusterConfigSecondaryWorkerConfigOutputReference)SetInternalValue(val *GoogleDataprocWorkflowTemplatePlacementManagedClusterConfigSecondaryWorkerConfig) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_GoogleDataprocWorkflowTemplatePlacementManagedClusterConfigSecondaryWorkerConfigOutputReference)SetMachineType(val *string) {
	if err := j.validateSetMachineTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"machineType",
		val,
	)
}

func (j *jsiiProxy_GoogleDataprocWorkflowTemplatePlacementManagedClusterConfigSecondaryWorkerConfigOutputReference)SetMinCpuPlatform(val *string) {
	if err := j.validateSetMinCpuPlatformParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"minCpuPlatform",
		val,
	)
}

func (j *jsiiProxy_GoogleDataprocWorkflowTemplatePlacementManagedClusterConfigSecondaryWorkerConfigOutputReference)SetNumInstances(val *float64) {
	if err := j.validateSetNumInstancesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"numInstances",
		val,
	)
}

func (j *jsiiProxy_GoogleDataprocWorkflowTemplatePlacementManagedClusterConfigSecondaryWorkerConfigOutputReference)SetPreemptibility(val *string) {
	if err := j.validateSetPreemptibilityParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"preemptibility",
		val,
	)
}

func (j *jsiiProxy_GoogleDataprocWorkflowTemplatePlacementManagedClusterConfigSecondaryWorkerConfigOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_GoogleDataprocWorkflowTemplatePlacementManagedClusterConfigSecondaryWorkerConfigOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (g *jsiiProxy_GoogleDataprocWorkflowTemplatePlacementManagedClusterConfigSecondaryWorkerConfigOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleDataprocWorkflowTemplatePlacementManagedClusterConfigSecondaryWorkerConfigOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (g *jsiiProxy_GoogleDataprocWorkflowTemplatePlacementManagedClusterConfigSecondaryWorkerConfigOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GoogleDataprocWorkflowTemplatePlacementManagedClusterConfigSecondaryWorkerConfigOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (g *jsiiProxy_GoogleDataprocWorkflowTemplatePlacementManagedClusterConfigSecondaryWorkerConfigOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (g *jsiiProxy_GoogleDataprocWorkflowTemplatePlacementManagedClusterConfigSecondaryWorkerConfigOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (g *jsiiProxy_GoogleDataprocWorkflowTemplatePlacementManagedClusterConfigSecondaryWorkerConfigOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (g *jsiiProxy_GoogleDataprocWorkflowTemplatePlacementManagedClusterConfigSecondaryWorkerConfigOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (g *jsiiProxy_GoogleDataprocWorkflowTemplatePlacementManagedClusterConfigSecondaryWorkerConfigOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (g *jsiiProxy_GoogleDataprocWorkflowTemplatePlacementManagedClusterConfigSecondaryWorkerConfigOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (g *jsiiProxy_GoogleDataprocWorkflowTemplatePlacementManagedClusterConfigSecondaryWorkerConfigOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		g,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleDataprocWorkflowTemplatePlacementManagedClusterConfigSecondaryWorkerConfigOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GoogleDataprocWorkflowTemplatePlacementManagedClusterConfigSecondaryWorkerConfigOutputReference) PutAccelerators(value interface{}) {
	if err := g.validatePutAcceleratorsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putAccelerators",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleDataprocWorkflowTemplatePlacementManagedClusterConfigSecondaryWorkerConfigOutputReference) PutDiskConfig(value *GoogleDataprocWorkflowTemplatePlacementManagedClusterConfigSecondaryWorkerConfigDiskConfig) {
	if err := g.validatePutDiskConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putDiskConfig",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleDataprocWorkflowTemplatePlacementManagedClusterConfigSecondaryWorkerConfigOutputReference) ResetAccelerators() {
	_jsii_.InvokeVoid(
		g,
		"resetAccelerators",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleDataprocWorkflowTemplatePlacementManagedClusterConfigSecondaryWorkerConfigOutputReference) ResetDiskConfig() {
	_jsii_.InvokeVoid(
		g,
		"resetDiskConfig",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleDataprocWorkflowTemplatePlacementManagedClusterConfigSecondaryWorkerConfigOutputReference) ResetImage() {
	_jsii_.InvokeVoid(
		g,
		"resetImage",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleDataprocWorkflowTemplatePlacementManagedClusterConfigSecondaryWorkerConfigOutputReference) ResetMachineType() {
	_jsii_.InvokeVoid(
		g,
		"resetMachineType",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleDataprocWorkflowTemplatePlacementManagedClusterConfigSecondaryWorkerConfigOutputReference) ResetMinCpuPlatform() {
	_jsii_.InvokeVoid(
		g,
		"resetMinCpuPlatform",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleDataprocWorkflowTemplatePlacementManagedClusterConfigSecondaryWorkerConfigOutputReference) ResetNumInstances() {
	_jsii_.InvokeVoid(
		g,
		"resetNumInstances",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleDataprocWorkflowTemplatePlacementManagedClusterConfigSecondaryWorkerConfigOutputReference) ResetPreemptibility() {
	_jsii_.InvokeVoid(
		g,
		"resetPreemptibility",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleDataprocWorkflowTemplatePlacementManagedClusterConfigSecondaryWorkerConfigOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (g *jsiiProxy_GoogleDataprocWorkflowTemplatePlacementManagedClusterConfigSecondaryWorkerConfigOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

