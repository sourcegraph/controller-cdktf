package googlebigtableinstance

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/googlebeta/jsii"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/sourcegraph/controller-cdktf/gen/googlebeta/googlebigtableinstance/internal"
)

type GoogleBigtableInstanceClusterAutoscalingConfigOutputReference interface {
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
	CpuTarget() *float64
	SetCpuTarget(val *float64)
	CpuTargetInput() *float64
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	// Experimental.
	Fqn() *string
	InternalValue() *GoogleBigtableInstanceClusterAutoscalingConfig
	SetInternalValue(val *GoogleBigtableInstanceClusterAutoscalingConfig)
	MaxNodes() *float64
	SetMaxNodes(val *float64)
	MaxNodesInput() *float64
	MinNodes() *float64
	SetMinNodes(val *float64)
	MinNodesInput() *float64
	StorageTarget() *float64
	SetStorageTarget(val *float64)
	StorageTargetInput() *float64
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
	ResetStorageTarget()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for GoogleBigtableInstanceClusterAutoscalingConfigOutputReference
type jsiiProxy_GoogleBigtableInstanceClusterAutoscalingConfigOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_GoogleBigtableInstanceClusterAutoscalingConfigOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBigtableInstanceClusterAutoscalingConfigOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBigtableInstanceClusterAutoscalingConfigOutputReference) CpuTarget() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"cpuTarget",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBigtableInstanceClusterAutoscalingConfigOutputReference) CpuTargetInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"cpuTargetInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBigtableInstanceClusterAutoscalingConfigOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBigtableInstanceClusterAutoscalingConfigOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBigtableInstanceClusterAutoscalingConfigOutputReference) InternalValue() *GoogleBigtableInstanceClusterAutoscalingConfig {
	var returns *GoogleBigtableInstanceClusterAutoscalingConfig
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBigtableInstanceClusterAutoscalingConfigOutputReference) MaxNodes() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxNodes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBigtableInstanceClusterAutoscalingConfigOutputReference) MaxNodesInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxNodesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBigtableInstanceClusterAutoscalingConfigOutputReference) MinNodes() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"minNodes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBigtableInstanceClusterAutoscalingConfigOutputReference) MinNodesInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"minNodesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBigtableInstanceClusterAutoscalingConfigOutputReference) StorageTarget() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"storageTarget",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBigtableInstanceClusterAutoscalingConfigOutputReference) StorageTargetInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"storageTargetInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBigtableInstanceClusterAutoscalingConfigOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBigtableInstanceClusterAutoscalingConfigOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewGoogleBigtableInstanceClusterAutoscalingConfigOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) GoogleBigtableInstanceClusterAutoscalingConfigOutputReference {
	_init_.Initialize()

	if err := validateNewGoogleBigtableInstanceClusterAutoscalingConfigOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_GoogleBigtableInstanceClusterAutoscalingConfigOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-googlebeta.googleBigtableInstance.GoogleBigtableInstanceClusterAutoscalingConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewGoogleBigtableInstanceClusterAutoscalingConfigOutputReference_Override(g GoogleBigtableInstanceClusterAutoscalingConfigOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-googlebeta.googleBigtableInstance.GoogleBigtableInstanceClusterAutoscalingConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		g,
	)
}

func (j *jsiiProxy_GoogleBigtableInstanceClusterAutoscalingConfigOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_GoogleBigtableInstanceClusterAutoscalingConfigOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_GoogleBigtableInstanceClusterAutoscalingConfigOutputReference)SetCpuTarget(val *float64) {
	if err := j.validateSetCpuTargetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"cpuTarget",
		val,
	)
}

func (j *jsiiProxy_GoogleBigtableInstanceClusterAutoscalingConfigOutputReference)SetInternalValue(val *GoogleBigtableInstanceClusterAutoscalingConfig) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_GoogleBigtableInstanceClusterAutoscalingConfigOutputReference)SetMaxNodes(val *float64) {
	if err := j.validateSetMaxNodesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"maxNodes",
		val,
	)
}

func (j *jsiiProxy_GoogleBigtableInstanceClusterAutoscalingConfigOutputReference)SetMinNodes(val *float64) {
	if err := j.validateSetMinNodesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"minNodes",
		val,
	)
}

func (j *jsiiProxy_GoogleBigtableInstanceClusterAutoscalingConfigOutputReference)SetStorageTarget(val *float64) {
	if err := j.validateSetStorageTargetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"storageTarget",
		val,
	)
}

func (j *jsiiProxy_GoogleBigtableInstanceClusterAutoscalingConfigOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_GoogleBigtableInstanceClusterAutoscalingConfigOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (g *jsiiProxy_GoogleBigtableInstanceClusterAutoscalingConfigOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleBigtableInstanceClusterAutoscalingConfigOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (g *jsiiProxy_GoogleBigtableInstanceClusterAutoscalingConfigOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GoogleBigtableInstanceClusterAutoscalingConfigOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (g *jsiiProxy_GoogleBigtableInstanceClusterAutoscalingConfigOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (g *jsiiProxy_GoogleBigtableInstanceClusterAutoscalingConfigOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (g *jsiiProxy_GoogleBigtableInstanceClusterAutoscalingConfigOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (g *jsiiProxy_GoogleBigtableInstanceClusterAutoscalingConfigOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (g *jsiiProxy_GoogleBigtableInstanceClusterAutoscalingConfigOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (g *jsiiProxy_GoogleBigtableInstanceClusterAutoscalingConfigOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (g *jsiiProxy_GoogleBigtableInstanceClusterAutoscalingConfigOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		g,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleBigtableInstanceClusterAutoscalingConfigOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GoogleBigtableInstanceClusterAutoscalingConfigOutputReference) ResetStorageTarget() {
	_jsii_.InvokeVoid(
		g,
		"resetStorageTarget",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleBigtableInstanceClusterAutoscalingConfigOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (g *jsiiProxy_GoogleBigtableInstanceClusterAutoscalingConfigOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

