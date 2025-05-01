package googlerediscluster

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/google_beta/jsii"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/sourcegraph/controller-cdktf/gen/google_beta/googlerediscluster/internal"
)

type GoogleRedisClusterCrossClusterReplicationConfigOutputReference interface {
	cdktf.ComplexObject
	ClusterRole() *string
	SetClusterRole(val *string)
	ClusterRoleInput() *string
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
	InternalValue() *GoogleRedisClusterCrossClusterReplicationConfig
	SetInternalValue(val *GoogleRedisClusterCrossClusterReplicationConfig)
	Membership() GoogleRedisClusterCrossClusterReplicationConfigMembershipList
	PrimaryCluster() GoogleRedisClusterCrossClusterReplicationConfigPrimaryClusterOutputReference
	PrimaryClusterInput() *GoogleRedisClusterCrossClusterReplicationConfigPrimaryCluster
	SecondaryClusters() GoogleRedisClusterCrossClusterReplicationConfigSecondaryClustersList
	SecondaryClustersInput() interface{}
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	UpdateTime() *string
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
	PutPrimaryCluster(value *GoogleRedisClusterCrossClusterReplicationConfigPrimaryCluster)
	PutSecondaryClusters(value interface{})
	ResetClusterRole()
	ResetPrimaryCluster()
	ResetSecondaryClusters()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for GoogleRedisClusterCrossClusterReplicationConfigOutputReference
type jsiiProxy_GoogleRedisClusterCrossClusterReplicationConfigOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_GoogleRedisClusterCrossClusterReplicationConfigOutputReference) ClusterRole() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clusterRole",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleRedisClusterCrossClusterReplicationConfigOutputReference) ClusterRoleInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clusterRoleInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleRedisClusterCrossClusterReplicationConfigOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleRedisClusterCrossClusterReplicationConfigOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleRedisClusterCrossClusterReplicationConfigOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleRedisClusterCrossClusterReplicationConfigOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleRedisClusterCrossClusterReplicationConfigOutputReference) InternalValue() *GoogleRedisClusterCrossClusterReplicationConfig {
	var returns *GoogleRedisClusterCrossClusterReplicationConfig
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleRedisClusterCrossClusterReplicationConfigOutputReference) Membership() GoogleRedisClusterCrossClusterReplicationConfigMembershipList {
	var returns GoogleRedisClusterCrossClusterReplicationConfigMembershipList
	_jsii_.Get(
		j,
		"membership",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleRedisClusterCrossClusterReplicationConfigOutputReference) PrimaryCluster() GoogleRedisClusterCrossClusterReplicationConfigPrimaryClusterOutputReference {
	var returns GoogleRedisClusterCrossClusterReplicationConfigPrimaryClusterOutputReference
	_jsii_.Get(
		j,
		"primaryCluster",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleRedisClusterCrossClusterReplicationConfigOutputReference) PrimaryClusterInput() *GoogleRedisClusterCrossClusterReplicationConfigPrimaryCluster {
	var returns *GoogleRedisClusterCrossClusterReplicationConfigPrimaryCluster
	_jsii_.Get(
		j,
		"primaryClusterInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleRedisClusterCrossClusterReplicationConfigOutputReference) SecondaryClusters() GoogleRedisClusterCrossClusterReplicationConfigSecondaryClustersList {
	var returns GoogleRedisClusterCrossClusterReplicationConfigSecondaryClustersList
	_jsii_.Get(
		j,
		"secondaryClusters",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleRedisClusterCrossClusterReplicationConfigOutputReference) SecondaryClustersInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"secondaryClustersInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleRedisClusterCrossClusterReplicationConfigOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleRedisClusterCrossClusterReplicationConfigOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleRedisClusterCrossClusterReplicationConfigOutputReference) UpdateTime() *string {
	var returns *string
	_jsii_.Get(
		j,
		"updateTime",
		&returns,
	)
	return returns
}


func NewGoogleRedisClusterCrossClusterReplicationConfigOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) GoogleRedisClusterCrossClusterReplicationConfigOutputReference {
	_init_.Initialize()

	if err := validateNewGoogleRedisClusterCrossClusterReplicationConfigOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_GoogleRedisClusterCrossClusterReplicationConfigOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-google-beta.googleRedisCluster.GoogleRedisClusterCrossClusterReplicationConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewGoogleRedisClusterCrossClusterReplicationConfigOutputReference_Override(g GoogleRedisClusterCrossClusterReplicationConfigOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google-beta.googleRedisCluster.GoogleRedisClusterCrossClusterReplicationConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		g,
	)
}

func (j *jsiiProxy_GoogleRedisClusterCrossClusterReplicationConfigOutputReference)SetClusterRole(val *string) {
	if err := j.validateSetClusterRoleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"clusterRole",
		val,
	)
}

func (j *jsiiProxy_GoogleRedisClusterCrossClusterReplicationConfigOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_GoogleRedisClusterCrossClusterReplicationConfigOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_GoogleRedisClusterCrossClusterReplicationConfigOutputReference)SetInternalValue(val *GoogleRedisClusterCrossClusterReplicationConfig) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_GoogleRedisClusterCrossClusterReplicationConfigOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_GoogleRedisClusterCrossClusterReplicationConfigOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (g *jsiiProxy_GoogleRedisClusterCrossClusterReplicationConfigOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleRedisClusterCrossClusterReplicationConfigOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (g *jsiiProxy_GoogleRedisClusterCrossClusterReplicationConfigOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GoogleRedisClusterCrossClusterReplicationConfigOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (g *jsiiProxy_GoogleRedisClusterCrossClusterReplicationConfigOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (g *jsiiProxy_GoogleRedisClusterCrossClusterReplicationConfigOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (g *jsiiProxy_GoogleRedisClusterCrossClusterReplicationConfigOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (g *jsiiProxy_GoogleRedisClusterCrossClusterReplicationConfigOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (g *jsiiProxy_GoogleRedisClusterCrossClusterReplicationConfigOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (g *jsiiProxy_GoogleRedisClusterCrossClusterReplicationConfigOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (g *jsiiProxy_GoogleRedisClusterCrossClusterReplicationConfigOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		g,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleRedisClusterCrossClusterReplicationConfigOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GoogleRedisClusterCrossClusterReplicationConfigOutputReference) PutPrimaryCluster(value *GoogleRedisClusterCrossClusterReplicationConfigPrimaryCluster) {
	if err := g.validatePutPrimaryClusterParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putPrimaryCluster",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleRedisClusterCrossClusterReplicationConfigOutputReference) PutSecondaryClusters(value interface{}) {
	if err := g.validatePutSecondaryClustersParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putSecondaryClusters",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleRedisClusterCrossClusterReplicationConfigOutputReference) ResetClusterRole() {
	_jsii_.InvokeVoid(
		g,
		"resetClusterRole",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleRedisClusterCrossClusterReplicationConfigOutputReference) ResetPrimaryCluster() {
	_jsii_.InvokeVoid(
		g,
		"resetPrimaryCluster",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleRedisClusterCrossClusterReplicationConfigOutputReference) ResetSecondaryClusters() {
	_jsii_.InvokeVoid(
		g,
		"resetSecondaryClusters",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleRedisClusterCrossClusterReplicationConfigOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (g *jsiiProxy_GoogleRedisClusterCrossClusterReplicationConfigOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

