package googledataproccluster

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/google_beta/jsii"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/sourcegraph/controller-cdktf/gen/google_beta/googledataproccluster/internal"
)

type GoogleDataprocClusterVirtualClusterConfigKubernetesClusterConfigOutputReference interface {
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
	// Experimental.
	Fqn() *string
	GkeClusterConfig() GoogleDataprocClusterVirtualClusterConfigKubernetesClusterConfigGkeClusterConfigOutputReference
	GkeClusterConfigInput() *GoogleDataprocClusterVirtualClusterConfigKubernetesClusterConfigGkeClusterConfig
	InternalValue() *GoogleDataprocClusterVirtualClusterConfigKubernetesClusterConfig
	SetInternalValue(val *GoogleDataprocClusterVirtualClusterConfigKubernetesClusterConfig)
	KubernetesNamespace() *string
	SetKubernetesNamespace(val *string)
	KubernetesNamespaceInput() *string
	KubernetesSoftwareConfig() GoogleDataprocClusterVirtualClusterConfigKubernetesClusterConfigKubernetesSoftwareConfigOutputReference
	KubernetesSoftwareConfigInput() *GoogleDataprocClusterVirtualClusterConfigKubernetesClusterConfigKubernetesSoftwareConfig
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
	PutGkeClusterConfig(value *GoogleDataprocClusterVirtualClusterConfigKubernetesClusterConfigGkeClusterConfig)
	PutKubernetesSoftwareConfig(value *GoogleDataprocClusterVirtualClusterConfigKubernetesClusterConfigKubernetesSoftwareConfig)
	ResetKubernetesNamespace()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for GoogleDataprocClusterVirtualClusterConfigKubernetesClusterConfigOutputReference
type jsiiProxy_GoogleDataprocClusterVirtualClusterConfigKubernetesClusterConfigOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_GoogleDataprocClusterVirtualClusterConfigKubernetesClusterConfigOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataprocClusterVirtualClusterConfigKubernetesClusterConfigOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataprocClusterVirtualClusterConfigKubernetesClusterConfigOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataprocClusterVirtualClusterConfigKubernetesClusterConfigOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataprocClusterVirtualClusterConfigKubernetesClusterConfigOutputReference) GkeClusterConfig() GoogleDataprocClusterVirtualClusterConfigKubernetesClusterConfigGkeClusterConfigOutputReference {
	var returns GoogleDataprocClusterVirtualClusterConfigKubernetesClusterConfigGkeClusterConfigOutputReference
	_jsii_.Get(
		j,
		"gkeClusterConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataprocClusterVirtualClusterConfigKubernetesClusterConfigOutputReference) GkeClusterConfigInput() *GoogleDataprocClusterVirtualClusterConfigKubernetesClusterConfigGkeClusterConfig {
	var returns *GoogleDataprocClusterVirtualClusterConfigKubernetesClusterConfigGkeClusterConfig
	_jsii_.Get(
		j,
		"gkeClusterConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataprocClusterVirtualClusterConfigKubernetesClusterConfigOutputReference) InternalValue() *GoogleDataprocClusterVirtualClusterConfigKubernetesClusterConfig {
	var returns *GoogleDataprocClusterVirtualClusterConfigKubernetesClusterConfig
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataprocClusterVirtualClusterConfigKubernetesClusterConfigOutputReference) KubernetesNamespace() *string {
	var returns *string
	_jsii_.Get(
		j,
		"kubernetesNamespace",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataprocClusterVirtualClusterConfigKubernetesClusterConfigOutputReference) KubernetesNamespaceInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"kubernetesNamespaceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataprocClusterVirtualClusterConfigKubernetesClusterConfigOutputReference) KubernetesSoftwareConfig() GoogleDataprocClusterVirtualClusterConfigKubernetesClusterConfigKubernetesSoftwareConfigOutputReference {
	var returns GoogleDataprocClusterVirtualClusterConfigKubernetesClusterConfigKubernetesSoftwareConfigOutputReference
	_jsii_.Get(
		j,
		"kubernetesSoftwareConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataprocClusterVirtualClusterConfigKubernetesClusterConfigOutputReference) KubernetesSoftwareConfigInput() *GoogleDataprocClusterVirtualClusterConfigKubernetesClusterConfigKubernetesSoftwareConfig {
	var returns *GoogleDataprocClusterVirtualClusterConfigKubernetesClusterConfigKubernetesSoftwareConfig
	_jsii_.Get(
		j,
		"kubernetesSoftwareConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataprocClusterVirtualClusterConfigKubernetesClusterConfigOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataprocClusterVirtualClusterConfigKubernetesClusterConfigOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewGoogleDataprocClusterVirtualClusterConfigKubernetesClusterConfigOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) GoogleDataprocClusterVirtualClusterConfigKubernetesClusterConfigOutputReference {
	_init_.Initialize()

	if err := validateNewGoogleDataprocClusterVirtualClusterConfigKubernetesClusterConfigOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_GoogleDataprocClusterVirtualClusterConfigKubernetesClusterConfigOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-google-beta.googleDataprocCluster.GoogleDataprocClusterVirtualClusterConfigKubernetesClusterConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewGoogleDataprocClusterVirtualClusterConfigKubernetesClusterConfigOutputReference_Override(g GoogleDataprocClusterVirtualClusterConfigKubernetesClusterConfigOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google-beta.googleDataprocCluster.GoogleDataprocClusterVirtualClusterConfigKubernetesClusterConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		g,
	)
}

func (j *jsiiProxy_GoogleDataprocClusterVirtualClusterConfigKubernetesClusterConfigOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_GoogleDataprocClusterVirtualClusterConfigKubernetesClusterConfigOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_GoogleDataprocClusterVirtualClusterConfigKubernetesClusterConfigOutputReference)SetInternalValue(val *GoogleDataprocClusterVirtualClusterConfigKubernetesClusterConfig) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_GoogleDataprocClusterVirtualClusterConfigKubernetesClusterConfigOutputReference)SetKubernetesNamespace(val *string) {
	if err := j.validateSetKubernetesNamespaceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"kubernetesNamespace",
		val,
	)
}

func (j *jsiiProxy_GoogleDataprocClusterVirtualClusterConfigKubernetesClusterConfigOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_GoogleDataprocClusterVirtualClusterConfigKubernetesClusterConfigOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (g *jsiiProxy_GoogleDataprocClusterVirtualClusterConfigKubernetesClusterConfigOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleDataprocClusterVirtualClusterConfigKubernetesClusterConfigOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (g *jsiiProxy_GoogleDataprocClusterVirtualClusterConfigKubernetesClusterConfigOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GoogleDataprocClusterVirtualClusterConfigKubernetesClusterConfigOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (g *jsiiProxy_GoogleDataprocClusterVirtualClusterConfigKubernetesClusterConfigOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (g *jsiiProxy_GoogleDataprocClusterVirtualClusterConfigKubernetesClusterConfigOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (g *jsiiProxy_GoogleDataprocClusterVirtualClusterConfigKubernetesClusterConfigOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (g *jsiiProxy_GoogleDataprocClusterVirtualClusterConfigKubernetesClusterConfigOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (g *jsiiProxy_GoogleDataprocClusterVirtualClusterConfigKubernetesClusterConfigOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (g *jsiiProxy_GoogleDataprocClusterVirtualClusterConfigKubernetesClusterConfigOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (g *jsiiProxy_GoogleDataprocClusterVirtualClusterConfigKubernetesClusterConfigOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		g,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleDataprocClusterVirtualClusterConfigKubernetesClusterConfigOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GoogleDataprocClusterVirtualClusterConfigKubernetesClusterConfigOutputReference) PutGkeClusterConfig(value *GoogleDataprocClusterVirtualClusterConfigKubernetesClusterConfigGkeClusterConfig) {
	if err := g.validatePutGkeClusterConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putGkeClusterConfig",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleDataprocClusterVirtualClusterConfigKubernetesClusterConfigOutputReference) PutKubernetesSoftwareConfig(value *GoogleDataprocClusterVirtualClusterConfigKubernetesClusterConfigKubernetesSoftwareConfig) {
	if err := g.validatePutKubernetesSoftwareConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putKubernetesSoftwareConfig",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleDataprocClusterVirtualClusterConfigKubernetesClusterConfigOutputReference) ResetKubernetesNamespace() {
	_jsii_.InvokeVoid(
		g,
		"resetKubernetesNamespace",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleDataprocClusterVirtualClusterConfigKubernetesClusterConfigOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (g *jsiiProxy_GoogleDataprocClusterVirtualClusterConfigKubernetesClusterConfigOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

