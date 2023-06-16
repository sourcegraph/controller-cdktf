package googledataproccluster

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/google_beta/jsii"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/sourcegraph/controller-cdktf/gen/google_beta/googledataproccluster/internal"
)

type GoogleDataprocClusterClusterConfigPreemptibleWorkerConfigDiskConfigOutputReference interface {
	cdktf.ComplexObject
	BootDiskSizeGb() *float64
	SetBootDiskSizeGb(val *float64)
	BootDiskSizeGbInput() *float64
	BootDiskType() *string
	SetBootDiskType(val *string)
	BootDiskTypeInput() *string
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
	InternalValue() *GoogleDataprocClusterClusterConfigPreemptibleWorkerConfigDiskConfig
	SetInternalValue(val *GoogleDataprocClusterClusterConfigPreemptibleWorkerConfigDiskConfig)
	NumLocalSsds() *float64
	SetNumLocalSsds(val *float64)
	NumLocalSsdsInput() *float64
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
	ResetBootDiskSizeGb()
	ResetBootDiskType()
	ResetNumLocalSsds()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for GoogleDataprocClusterClusterConfigPreemptibleWorkerConfigDiskConfigOutputReference
type jsiiProxy_GoogleDataprocClusterClusterConfigPreemptibleWorkerConfigDiskConfigOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_GoogleDataprocClusterClusterConfigPreemptibleWorkerConfigDiskConfigOutputReference) BootDiskSizeGb() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"bootDiskSizeGb",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataprocClusterClusterConfigPreemptibleWorkerConfigDiskConfigOutputReference) BootDiskSizeGbInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"bootDiskSizeGbInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataprocClusterClusterConfigPreemptibleWorkerConfigDiskConfigOutputReference) BootDiskType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"bootDiskType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataprocClusterClusterConfigPreemptibleWorkerConfigDiskConfigOutputReference) BootDiskTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"bootDiskTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataprocClusterClusterConfigPreemptibleWorkerConfigDiskConfigOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataprocClusterClusterConfigPreemptibleWorkerConfigDiskConfigOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataprocClusterClusterConfigPreemptibleWorkerConfigDiskConfigOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataprocClusterClusterConfigPreemptibleWorkerConfigDiskConfigOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataprocClusterClusterConfigPreemptibleWorkerConfigDiskConfigOutputReference) InternalValue() *GoogleDataprocClusterClusterConfigPreemptibleWorkerConfigDiskConfig {
	var returns *GoogleDataprocClusterClusterConfigPreemptibleWorkerConfigDiskConfig
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataprocClusterClusterConfigPreemptibleWorkerConfigDiskConfigOutputReference) NumLocalSsds() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"numLocalSsds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataprocClusterClusterConfigPreemptibleWorkerConfigDiskConfigOutputReference) NumLocalSsdsInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"numLocalSsdsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataprocClusterClusterConfigPreemptibleWorkerConfigDiskConfigOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataprocClusterClusterConfigPreemptibleWorkerConfigDiskConfigOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewGoogleDataprocClusterClusterConfigPreemptibleWorkerConfigDiskConfigOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) GoogleDataprocClusterClusterConfigPreemptibleWorkerConfigDiskConfigOutputReference {
	_init_.Initialize()

	if err := validateNewGoogleDataprocClusterClusterConfigPreemptibleWorkerConfigDiskConfigOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_GoogleDataprocClusterClusterConfigPreemptibleWorkerConfigDiskConfigOutputReference{}

	_jsii_.Create(
		"google-beta.googleDataprocCluster.GoogleDataprocClusterClusterConfigPreemptibleWorkerConfigDiskConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewGoogleDataprocClusterClusterConfigPreemptibleWorkerConfigDiskConfigOutputReference_Override(g GoogleDataprocClusterClusterConfigPreemptibleWorkerConfigDiskConfigOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"google-beta.googleDataprocCluster.GoogleDataprocClusterClusterConfigPreemptibleWorkerConfigDiskConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		g,
	)
}

func (j *jsiiProxy_GoogleDataprocClusterClusterConfigPreemptibleWorkerConfigDiskConfigOutputReference)SetBootDiskSizeGb(val *float64) {
	if err := j.validateSetBootDiskSizeGbParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"bootDiskSizeGb",
		val,
	)
}

func (j *jsiiProxy_GoogleDataprocClusterClusterConfigPreemptibleWorkerConfigDiskConfigOutputReference)SetBootDiskType(val *string) {
	if err := j.validateSetBootDiskTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"bootDiskType",
		val,
	)
}

func (j *jsiiProxy_GoogleDataprocClusterClusterConfigPreemptibleWorkerConfigDiskConfigOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_GoogleDataprocClusterClusterConfigPreemptibleWorkerConfigDiskConfigOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_GoogleDataprocClusterClusterConfigPreemptibleWorkerConfigDiskConfigOutputReference)SetInternalValue(val *GoogleDataprocClusterClusterConfigPreemptibleWorkerConfigDiskConfig) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_GoogleDataprocClusterClusterConfigPreemptibleWorkerConfigDiskConfigOutputReference)SetNumLocalSsds(val *float64) {
	if err := j.validateSetNumLocalSsdsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"numLocalSsds",
		val,
	)
}

func (j *jsiiProxy_GoogleDataprocClusterClusterConfigPreemptibleWorkerConfigDiskConfigOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_GoogleDataprocClusterClusterConfigPreemptibleWorkerConfigDiskConfigOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (g *jsiiProxy_GoogleDataprocClusterClusterConfigPreemptibleWorkerConfigDiskConfigOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleDataprocClusterClusterConfigPreemptibleWorkerConfigDiskConfigOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (g *jsiiProxy_GoogleDataprocClusterClusterConfigPreemptibleWorkerConfigDiskConfigOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GoogleDataprocClusterClusterConfigPreemptibleWorkerConfigDiskConfigOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (g *jsiiProxy_GoogleDataprocClusterClusterConfigPreemptibleWorkerConfigDiskConfigOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (g *jsiiProxy_GoogleDataprocClusterClusterConfigPreemptibleWorkerConfigDiskConfigOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (g *jsiiProxy_GoogleDataprocClusterClusterConfigPreemptibleWorkerConfigDiskConfigOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (g *jsiiProxy_GoogleDataprocClusterClusterConfigPreemptibleWorkerConfigDiskConfigOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (g *jsiiProxy_GoogleDataprocClusterClusterConfigPreemptibleWorkerConfigDiskConfigOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (g *jsiiProxy_GoogleDataprocClusterClusterConfigPreemptibleWorkerConfigDiskConfigOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (g *jsiiProxy_GoogleDataprocClusterClusterConfigPreemptibleWorkerConfigDiskConfigOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		g,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleDataprocClusterClusterConfigPreemptibleWorkerConfigDiskConfigOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GoogleDataprocClusterClusterConfigPreemptibleWorkerConfigDiskConfigOutputReference) ResetBootDiskSizeGb() {
	_jsii_.InvokeVoid(
		g,
		"resetBootDiskSizeGb",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleDataprocClusterClusterConfigPreemptibleWorkerConfigDiskConfigOutputReference) ResetBootDiskType() {
	_jsii_.InvokeVoid(
		g,
		"resetBootDiskType",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleDataprocClusterClusterConfigPreemptibleWorkerConfigDiskConfigOutputReference) ResetNumLocalSsds() {
	_jsii_.InvokeVoid(
		g,
		"resetNumLocalSsds",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleDataprocClusterClusterConfigPreemptibleWorkerConfigDiskConfigOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (g *jsiiProxy_GoogleDataprocClusterClusterConfigPreemptibleWorkerConfigDiskConfigOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

