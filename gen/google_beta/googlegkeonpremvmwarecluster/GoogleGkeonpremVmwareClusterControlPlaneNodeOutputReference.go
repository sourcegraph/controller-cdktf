package googlegkeonpremvmwarecluster

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/google_beta/jsii"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/sourcegraph/controller-cdktf/gen/google_beta/googlegkeonpremvmwarecluster/internal"
)

type GoogleGkeonpremVmwareClusterControlPlaneNodeOutputReference interface {
	cdktf.ComplexObject
	AutoResizeConfig() GoogleGkeonpremVmwareClusterControlPlaneNodeAutoResizeConfigOutputReference
	AutoResizeConfigInput() *GoogleGkeonpremVmwareClusterControlPlaneNodeAutoResizeConfig
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
	// Experimental.
	Fqn() *string
	InternalValue() *GoogleGkeonpremVmwareClusterControlPlaneNode
	SetInternalValue(val *GoogleGkeonpremVmwareClusterControlPlaneNode)
	Memory() *float64
	SetMemory(val *float64)
	MemoryInput() *float64
	Replicas() *float64
	SetReplicas(val *float64)
	ReplicasInput() *float64
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	VsphereConfig() GoogleGkeonpremVmwareClusterControlPlaneNodeVsphereConfigList
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
	PutAutoResizeConfig(value *GoogleGkeonpremVmwareClusterControlPlaneNodeAutoResizeConfig)
	ResetAutoResizeConfig()
	ResetCpus()
	ResetMemory()
	ResetReplicas()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for GoogleGkeonpremVmwareClusterControlPlaneNodeOutputReference
type jsiiProxy_GoogleGkeonpremVmwareClusterControlPlaneNodeOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_GoogleGkeonpremVmwareClusterControlPlaneNodeOutputReference) AutoResizeConfig() GoogleGkeonpremVmwareClusterControlPlaneNodeAutoResizeConfigOutputReference {
	var returns GoogleGkeonpremVmwareClusterControlPlaneNodeAutoResizeConfigOutputReference
	_jsii_.Get(
		j,
		"autoResizeConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleGkeonpremVmwareClusterControlPlaneNodeOutputReference) AutoResizeConfigInput() *GoogleGkeonpremVmwareClusterControlPlaneNodeAutoResizeConfig {
	var returns *GoogleGkeonpremVmwareClusterControlPlaneNodeAutoResizeConfig
	_jsii_.Get(
		j,
		"autoResizeConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleGkeonpremVmwareClusterControlPlaneNodeOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleGkeonpremVmwareClusterControlPlaneNodeOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleGkeonpremVmwareClusterControlPlaneNodeOutputReference) Cpus() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"cpus",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleGkeonpremVmwareClusterControlPlaneNodeOutputReference) CpusInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"cpusInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleGkeonpremVmwareClusterControlPlaneNodeOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleGkeonpremVmwareClusterControlPlaneNodeOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleGkeonpremVmwareClusterControlPlaneNodeOutputReference) InternalValue() *GoogleGkeonpremVmwareClusterControlPlaneNode {
	var returns *GoogleGkeonpremVmwareClusterControlPlaneNode
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleGkeonpremVmwareClusterControlPlaneNodeOutputReference) Memory() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"memory",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleGkeonpremVmwareClusterControlPlaneNodeOutputReference) MemoryInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"memoryInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleGkeonpremVmwareClusterControlPlaneNodeOutputReference) Replicas() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"replicas",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleGkeonpremVmwareClusterControlPlaneNodeOutputReference) ReplicasInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"replicasInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleGkeonpremVmwareClusterControlPlaneNodeOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleGkeonpremVmwareClusterControlPlaneNodeOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleGkeonpremVmwareClusterControlPlaneNodeOutputReference) VsphereConfig() GoogleGkeonpremVmwareClusterControlPlaneNodeVsphereConfigList {
	var returns GoogleGkeonpremVmwareClusterControlPlaneNodeVsphereConfigList
	_jsii_.Get(
		j,
		"vsphereConfig",
		&returns,
	)
	return returns
}


func NewGoogleGkeonpremVmwareClusterControlPlaneNodeOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) GoogleGkeonpremVmwareClusterControlPlaneNodeOutputReference {
	_init_.Initialize()

	if err := validateNewGoogleGkeonpremVmwareClusterControlPlaneNodeOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_GoogleGkeonpremVmwareClusterControlPlaneNodeOutputReference{}

	_jsii_.Create(
		"google-beta.googleGkeonpremVmwareCluster.GoogleGkeonpremVmwareClusterControlPlaneNodeOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewGoogleGkeonpremVmwareClusterControlPlaneNodeOutputReference_Override(g GoogleGkeonpremVmwareClusterControlPlaneNodeOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"google-beta.googleGkeonpremVmwareCluster.GoogleGkeonpremVmwareClusterControlPlaneNodeOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		g,
	)
}

func (j *jsiiProxy_GoogleGkeonpremVmwareClusterControlPlaneNodeOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_GoogleGkeonpremVmwareClusterControlPlaneNodeOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_GoogleGkeonpremVmwareClusterControlPlaneNodeOutputReference)SetCpus(val *float64) {
	if err := j.validateSetCpusParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"cpus",
		val,
	)
}

func (j *jsiiProxy_GoogleGkeonpremVmwareClusterControlPlaneNodeOutputReference)SetInternalValue(val *GoogleGkeonpremVmwareClusterControlPlaneNode) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_GoogleGkeonpremVmwareClusterControlPlaneNodeOutputReference)SetMemory(val *float64) {
	if err := j.validateSetMemoryParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"memory",
		val,
	)
}

func (j *jsiiProxy_GoogleGkeonpremVmwareClusterControlPlaneNodeOutputReference)SetReplicas(val *float64) {
	if err := j.validateSetReplicasParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"replicas",
		val,
	)
}

func (j *jsiiProxy_GoogleGkeonpremVmwareClusterControlPlaneNodeOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_GoogleGkeonpremVmwareClusterControlPlaneNodeOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (g *jsiiProxy_GoogleGkeonpremVmwareClusterControlPlaneNodeOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleGkeonpremVmwareClusterControlPlaneNodeOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (g *jsiiProxy_GoogleGkeonpremVmwareClusterControlPlaneNodeOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GoogleGkeonpremVmwareClusterControlPlaneNodeOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (g *jsiiProxy_GoogleGkeonpremVmwareClusterControlPlaneNodeOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (g *jsiiProxy_GoogleGkeonpremVmwareClusterControlPlaneNodeOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (g *jsiiProxy_GoogleGkeonpremVmwareClusterControlPlaneNodeOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (g *jsiiProxy_GoogleGkeonpremVmwareClusterControlPlaneNodeOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (g *jsiiProxy_GoogleGkeonpremVmwareClusterControlPlaneNodeOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (g *jsiiProxy_GoogleGkeonpremVmwareClusterControlPlaneNodeOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (g *jsiiProxy_GoogleGkeonpremVmwareClusterControlPlaneNodeOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		g,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleGkeonpremVmwareClusterControlPlaneNodeOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GoogleGkeonpremVmwareClusterControlPlaneNodeOutputReference) PutAutoResizeConfig(value *GoogleGkeonpremVmwareClusterControlPlaneNodeAutoResizeConfig) {
	if err := g.validatePutAutoResizeConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putAutoResizeConfig",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleGkeonpremVmwareClusterControlPlaneNodeOutputReference) ResetAutoResizeConfig() {
	_jsii_.InvokeVoid(
		g,
		"resetAutoResizeConfig",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleGkeonpremVmwareClusterControlPlaneNodeOutputReference) ResetCpus() {
	_jsii_.InvokeVoid(
		g,
		"resetCpus",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleGkeonpremVmwareClusterControlPlaneNodeOutputReference) ResetMemory() {
	_jsii_.InvokeVoid(
		g,
		"resetMemory",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleGkeonpremVmwareClusterControlPlaneNodeOutputReference) ResetReplicas() {
	_jsii_.InvokeVoid(
		g,
		"resetReplicas",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleGkeonpremVmwareClusterControlPlaneNodeOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (g *jsiiProxy_GoogleGkeonpremVmwareClusterControlPlaneNodeOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

