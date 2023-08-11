package googledataproccluster

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/google_beta/jsii"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/sourcegraph/controller-cdktf/gen/google_beta/googledataproccluster/internal"
)

type GoogleDataprocClusterVirtualClusterConfigOutputReference interface {
	cdktf.ComplexObject
	AuxiliaryServicesConfig() GoogleDataprocClusterVirtualClusterConfigAuxiliaryServicesConfigOutputReference
	AuxiliaryServicesConfigInput() *GoogleDataprocClusterVirtualClusterConfigAuxiliaryServicesConfig
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
	InternalValue() *GoogleDataprocClusterVirtualClusterConfig
	SetInternalValue(val *GoogleDataprocClusterVirtualClusterConfig)
	KubernetesClusterConfig() GoogleDataprocClusterVirtualClusterConfigKubernetesClusterConfigOutputReference
	KubernetesClusterConfigInput() *GoogleDataprocClusterVirtualClusterConfigKubernetesClusterConfig
	StagingBucket() *string
	SetStagingBucket(val *string)
	StagingBucketInput() *string
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
	PutAuxiliaryServicesConfig(value *GoogleDataprocClusterVirtualClusterConfigAuxiliaryServicesConfig)
	PutKubernetesClusterConfig(value *GoogleDataprocClusterVirtualClusterConfigKubernetesClusterConfig)
	ResetAuxiliaryServicesConfig()
	ResetKubernetesClusterConfig()
	ResetStagingBucket()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for GoogleDataprocClusterVirtualClusterConfigOutputReference
type jsiiProxy_GoogleDataprocClusterVirtualClusterConfigOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_GoogleDataprocClusterVirtualClusterConfigOutputReference) AuxiliaryServicesConfig() GoogleDataprocClusterVirtualClusterConfigAuxiliaryServicesConfigOutputReference {
	var returns GoogleDataprocClusterVirtualClusterConfigAuxiliaryServicesConfigOutputReference
	_jsii_.Get(
		j,
		"auxiliaryServicesConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataprocClusterVirtualClusterConfigOutputReference) AuxiliaryServicesConfigInput() *GoogleDataprocClusterVirtualClusterConfigAuxiliaryServicesConfig {
	var returns *GoogleDataprocClusterVirtualClusterConfigAuxiliaryServicesConfig
	_jsii_.Get(
		j,
		"auxiliaryServicesConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataprocClusterVirtualClusterConfigOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataprocClusterVirtualClusterConfigOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataprocClusterVirtualClusterConfigOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataprocClusterVirtualClusterConfigOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataprocClusterVirtualClusterConfigOutputReference) InternalValue() *GoogleDataprocClusterVirtualClusterConfig {
	var returns *GoogleDataprocClusterVirtualClusterConfig
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataprocClusterVirtualClusterConfigOutputReference) KubernetesClusterConfig() GoogleDataprocClusterVirtualClusterConfigKubernetesClusterConfigOutputReference {
	var returns GoogleDataprocClusterVirtualClusterConfigKubernetesClusterConfigOutputReference
	_jsii_.Get(
		j,
		"kubernetesClusterConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataprocClusterVirtualClusterConfigOutputReference) KubernetesClusterConfigInput() *GoogleDataprocClusterVirtualClusterConfigKubernetesClusterConfig {
	var returns *GoogleDataprocClusterVirtualClusterConfigKubernetesClusterConfig
	_jsii_.Get(
		j,
		"kubernetesClusterConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataprocClusterVirtualClusterConfigOutputReference) StagingBucket() *string {
	var returns *string
	_jsii_.Get(
		j,
		"stagingBucket",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataprocClusterVirtualClusterConfigOutputReference) StagingBucketInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"stagingBucketInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataprocClusterVirtualClusterConfigOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataprocClusterVirtualClusterConfigOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewGoogleDataprocClusterVirtualClusterConfigOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) GoogleDataprocClusterVirtualClusterConfigOutputReference {
	_init_.Initialize()

	if err := validateNewGoogleDataprocClusterVirtualClusterConfigOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_GoogleDataprocClusterVirtualClusterConfigOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-google-beta.googleDataprocCluster.GoogleDataprocClusterVirtualClusterConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewGoogleDataprocClusterVirtualClusterConfigOutputReference_Override(g GoogleDataprocClusterVirtualClusterConfigOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google-beta.googleDataprocCluster.GoogleDataprocClusterVirtualClusterConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		g,
	)
}

func (j *jsiiProxy_GoogleDataprocClusterVirtualClusterConfigOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_GoogleDataprocClusterVirtualClusterConfigOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_GoogleDataprocClusterVirtualClusterConfigOutputReference)SetInternalValue(val *GoogleDataprocClusterVirtualClusterConfig) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_GoogleDataprocClusterVirtualClusterConfigOutputReference)SetStagingBucket(val *string) {
	if err := j.validateSetStagingBucketParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"stagingBucket",
		val,
	)
}

func (j *jsiiProxy_GoogleDataprocClusterVirtualClusterConfigOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_GoogleDataprocClusterVirtualClusterConfigOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (g *jsiiProxy_GoogleDataprocClusterVirtualClusterConfigOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleDataprocClusterVirtualClusterConfigOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (g *jsiiProxy_GoogleDataprocClusterVirtualClusterConfigOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GoogleDataprocClusterVirtualClusterConfigOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (g *jsiiProxy_GoogleDataprocClusterVirtualClusterConfigOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (g *jsiiProxy_GoogleDataprocClusterVirtualClusterConfigOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (g *jsiiProxy_GoogleDataprocClusterVirtualClusterConfigOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (g *jsiiProxy_GoogleDataprocClusterVirtualClusterConfigOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (g *jsiiProxy_GoogleDataprocClusterVirtualClusterConfigOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (g *jsiiProxy_GoogleDataprocClusterVirtualClusterConfigOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (g *jsiiProxy_GoogleDataprocClusterVirtualClusterConfigOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		g,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleDataprocClusterVirtualClusterConfigOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GoogleDataprocClusterVirtualClusterConfigOutputReference) PutAuxiliaryServicesConfig(value *GoogleDataprocClusterVirtualClusterConfigAuxiliaryServicesConfig) {
	if err := g.validatePutAuxiliaryServicesConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putAuxiliaryServicesConfig",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleDataprocClusterVirtualClusterConfigOutputReference) PutKubernetesClusterConfig(value *GoogleDataprocClusterVirtualClusterConfigKubernetesClusterConfig) {
	if err := g.validatePutKubernetesClusterConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putKubernetesClusterConfig",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleDataprocClusterVirtualClusterConfigOutputReference) ResetAuxiliaryServicesConfig() {
	_jsii_.InvokeVoid(
		g,
		"resetAuxiliaryServicesConfig",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleDataprocClusterVirtualClusterConfigOutputReference) ResetKubernetesClusterConfig() {
	_jsii_.InvokeVoid(
		g,
		"resetKubernetesClusterConfig",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleDataprocClusterVirtualClusterConfigOutputReference) ResetStagingBucket() {
	_jsii_.InvokeVoid(
		g,
		"resetStagingBucket",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleDataprocClusterVirtualClusterConfigOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (g *jsiiProxy_GoogleDataprocClusterVirtualClusterConfigOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

