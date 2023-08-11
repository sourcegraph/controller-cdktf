package googlegkeonpremvmwarenodepool

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/google_beta/jsii"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/sourcegraph/controller-cdktf/gen/google_beta/googlegkeonpremvmwarenodepool/internal"
)

type GoogleGkeonpremVmwareNodePoolConfigAOutputReference interface {
	cdktf.ComplexObject
	BootDiskSizeGb() *float64
	SetBootDiskSizeGb(val *float64)
	BootDiskSizeGbInput() *float64
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
	Cpus() *float64
	SetCpus(val *float64)
	CpusInput() *float64
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	EnableLoadBalancer() interface{}
	SetEnableLoadBalancer(val interface{})
	EnableLoadBalancerInput() interface{}
	// Experimental.
	Fqn() *string
	Image() *string
	SetImage(val *string)
	ImageInput() *string
	ImageType() *string
	SetImageType(val *string)
	ImageTypeInput() *string
	InternalValue() *GoogleGkeonpremVmwareNodePoolConfigA
	SetInternalValue(val *GoogleGkeonpremVmwareNodePoolConfigA)
	Labels() *map[string]*string
	SetLabels(val *map[string]*string)
	LabelsInput() *map[string]*string
	MemoryMb() *float64
	SetMemoryMb(val *float64)
	MemoryMbInput() *float64
	Replicas() *float64
	SetReplicas(val *float64)
	ReplicasInput() *float64
	Taints() GoogleGkeonpremVmwareNodePoolConfigTaintsList
	TaintsInput() interface{}
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	VsphereConfig() GoogleGkeonpremVmwareNodePoolConfigVsphereConfigList
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
	PutTaints(value interface{})
	ResetBootDiskSizeGb()
	ResetCpus()
	ResetEnableLoadBalancer()
	ResetImage()
	ResetLabels()
	ResetMemoryMb()
	ResetReplicas()
	ResetTaints()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for GoogleGkeonpremVmwareNodePoolConfigAOutputReference
type jsiiProxy_GoogleGkeonpremVmwareNodePoolConfigAOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_GoogleGkeonpremVmwareNodePoolConfigAOutputReference) BootDiskSizeGb() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"bootDiskSizeGb",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleGkeonpremVmwareNodePoolConfigAOutputReference) BootDiskSizeGbInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"bootDiskSizeGbInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleGkeonpremVmwareNodePoolConfigAOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleGkeonpremVmwareNodePoolConfigAOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleGkeonpremVmwareNodePoolConfigAOutputReference) Cpus() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"cpus",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleGkeonpremVmwareNodePoolConfigAOutputReference) CpusInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"cpusInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleGkeonpremVmwareNodePoolConfigAOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleGkeonpremVmwareNodePoolConfigAOutputReference) EnableLoadBalancer() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableLoadBalancer",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleGkeonpremVmwareNodePoolConfigAOutputReference) EnableLoadBalancerInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableLoadBalancerInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleGkeonpremVmwareNodePoolConfigAOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleGkeonpremVmwareNodePoolConfigAOutputReference) Image() *string {
	var returns *string
	_jsii_.Get(
		j,
		"image",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleGkeonpremVmwareNodePoolConfigAOutputReference) ImageInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"imageInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleGkeonpremVmwareNodePoolConfigAOutputReference) ImageType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"imageType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleGkeonpremVmwareNodePoolConfigAOutputReference) ImageTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"imageTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleGkeonpremVmwareNodePoolConfigAOutputReference) InternalValue() *GoogleGkeonpremVmwareNodePoolConfigA {
	var returns *GoogleGkeonpremVmwareNodePoolConfigA
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleGkeonpremVmwareNodePoolConfigAOutputReference) Labels() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"labels",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleGkeonpremVmwareNodePoolConfigAOutputReference) LabelsInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"labelsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleGkeonpremVmwareNodePoolConfigAOutputReference) MemoryMb() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"memoryMb",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleGkeonpremVmwareNodePoolConfigAOutputReference) MemoryMbInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"memoryMbInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleGkeonpremVmwareNodePoolConfigAOutputReference) Replicas() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"replicas",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleGkeonpremVmwareNodePoolConfigAOutputReference) ReplicasInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"replicasInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleGkeonpremVmwareNodePoolConfigAOutputReference) Taints() GoogleGkeonpremVmwareNodePoolConfigTaintsList {
	var returns GoogleGkeonpremVmwareNodePoolConfigTaintsList
	_jsii_.Get(
		j,
		"taints",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleGkeonpremVmwareNodePoolConfigAOutputReference) TaintsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"taintsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleGkeonpremVmwareNodePoolConfigAOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleGkeonpremVmwareNodePoolConfigAOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleGkeonpremVmwareNodePoolConfigAOutputReference) VsphereConfig() GoogleGkeonpremVmwareNodePoolConfigVsphereConfigList {
	var returns GoogleGkeonpremVmwareNodePoolConfigVsphereConfigList
	_jsii_.Get(
		j,
		"vsphereConfig",
		&returns,
	)
	return returns
}


func NewGoogleGkeonpremVmwareNodePoolConfigAOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) GoogleGkeonpremVmwareNodePoolConfigAOutputReference {
	_init_.Initialize()

	if err := validateNewGoogleGkeonpremVmwareNodePoolConfigAOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_GoogleGkeonpremVmwareNodePoolConfigAOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-google-beta.googleGkeonpremVmwareNodePool.GoogleGkeonpremVmwareNodePoolConfigAOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewGoogleGkeonpremVmwareNodePoolConfigAOutputReference_Override(g GoogleGkeonpremVmwareNodePoolConfigAOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google-beta.googleGkeonpremVmwareNodePool.GoogleGkeonpremVmwareNodePoolConfigAOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		g,
	)
}

func (j *jsiiProxy_GoogleGkeonpremVmwareNodePoolConfigAOutputReference)SetBootDiskSizeGb(val *float64) {
	if err := j.validateSetBootDiskSizeGbParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"bootDiskSizeGb",
		val,
	)
}

func (j *jsiiProxy_GoogleGkeonpremVmwareNodePoolConfigAOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_GoogleGkeonpremVmwareNodePoolConfigAOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_GoogleGkeonpremVmwareNodePoolConfigAOutputReference)SetCpus(val *float64) {
	if err := j.validateSetCpusParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"cpus",
		val,
	)
}

func (j *jsiiProxy_GoogleGkeonpremVmwareNodePoolConfigAOutputReference)SetEnableLoadBalancer(val interface{}) {
	if err := j.validateSetEnableLoadBalancerParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"enableLoadBalancer",
		val,
	)
}

func (j *jsiiProxy_GoogleGkeonpremVmwareNodePoolConfigAOutputReference)SetImage(val *string) {
	if err := j.validateSetImageParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"image",
		val,
	)
}

func (j *jsiiProxy_GoogleGkeonpremVmwareNodePoolConfigAOutputReference)SetImageType(val *string) {
	if err := j.validateSetImageTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"imageType",
		val,
	)
}

func (j *jsiiProxy_GoogleGkeonpremVmwareNodePoolConfigAOutputReference)SetInternalValue(val *GoogleGkeonpremVmwareNodePoolConfigA) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_GoogleGkeonpremVmwareNodePoolConfigAOutputReference)SetLabels(val *map[string]*string) {
	if err := j.validateSetLabelsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"labels",
		val,
	)
}

func (j *jsiiProxy_GoogleGkeonpremVmwareNodePoolConfigAOutputReference)SetMemoryMb(val *float64) {
	if err := j.validateSetMemoryMbParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"memoryMb",
		val,
	)
}

func (j *jsiiProxy_GoogleGkeonpremVmwareNodePoolConfigAOutputReference)SetReplicas(val *float64) {
	if err := j.validateSetReplicasParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"replicas",
		val,
	)
}

func (j *jsiiProxy_GoogleGkeonpremVmwareNodePoolConfigAOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_GoogleGkeonpremVmwareNodePoolConfigAOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (g *jsiiProxy_GoogleGkeonpremVmwareNodePoolConfigAOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleGkeonpremVmwareNodePoolConfigAOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (g *jsiiProxy_GoogleGkeonpremVmwareNodePoolConfigAOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GoogleGkeonpremVmwareNodePoolConfigAOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (g *jsiiProxy_GoogleGkeonpremVmwareNodePoolConfigAOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (g *jsiiProxy_GoogleGkeonpremVmwareNodePoolConfigAOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (g *jsiiProxy_GoogleGkeonpremVmwareNodePoolConfigAOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (g *jsiiProxy_GoogleGkeonpremVmwareNodePoolConfigAOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (g *jsiiProxy_GoogleGkeonpremVmwareNodePoolConfigAOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (g *jsiiProxy_GoogleGkeonpremVmwareNodePoolConfigAOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (g *jsiiProxy_GoogleGkeonpremVmwareNodePoolConfigAOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		g,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleGkeonpremVmwareNodePoolConfigAOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GoogleGkeonpremVmwareNodePoolConfigAOutputReference) PutTaints(value interface{}) {
	if err := g.validatePutTaintsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putTaints",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleGkeonpremVmwareNodePoolConfigAOutputReference) ResetBootDiskSizeGb() {
	_jsii_.InvokeVoid(
		g,
		"resetBootDiskSizeGb",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleGkeonpremVmwareNodePoolConfigAOutputReference) ResetCpus() {
	_jsii_.InvokeVoid(
		g,
		"resetCpus",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleGkeonpremVmwareNodePoolConfigAOutputReference) ResetEnableLoadBalancer() {
	_jsii_.InvokeVoid(
		g,
		"resetEnableLoadBalancer",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleGkeonpremVmwareNodePoolConfigAOutputReference) ResetImage() {
	_jsii_.InvokeVoid(
		g,
		"resetImage",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleGkeonpremVmwareNodePoolConfigAOutputReference) ResetLabels() {
	_jsii_.InvokeVoid(
		g,
		"resetLabels",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleGkeonpremVmwareNodePoolConfigAOutputReference) ResetMemoryMb() {
	_jsii_.InvokeVoid(
		g,
		"resetMemoryMb",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleGkeonpremVmwareNodePoolConfigAOutputReference) ResetReplicas() {
	_jsii_.InvokeVoid(
		g,
		"resetReplicas",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleGkeonpremVmwareNodePoolConfigAOutputReference) ResetTaints() {
	_jsii_.InvokeVoid(
		g,
		"resetTaints",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleGkeonpremVmwareNodePoolConfigAOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (g *jsiiProxy_GoogleGkeonpremVmwareNodePoolConfigAOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

