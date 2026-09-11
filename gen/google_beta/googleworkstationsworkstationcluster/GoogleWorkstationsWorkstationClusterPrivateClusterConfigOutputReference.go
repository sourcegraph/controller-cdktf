package googleworkstationsworkstationcluster

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/google_beta/jsii"

	"github.com/open-constructs/cdk-terrain-go/cdktn"
	"github.com/sourcegraph/controller-cdktf/gen/google_beta/googleworkstationsworkstationcluster/internal"
)

type GoogleWorkstationsWorkstationClusterPrivateClusterConfigOutputReference interface {
	cdktn.ComplexObject
	AllowedProjects() *[]*string
	SetAllowedProjects(val *[]*string)
	AllowedProjectsInput() *[]*string
	ClusterHostname() *string
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
	EnablePrivateEndpoint() interface{}
	SetEnablePrivateEndpoint(val interface{})
	EnablePrivateEndpointInput() interface{}
	// Experimental.
	Fqn() *string
	InternalValue() *GoogleWorkstationsWorkstationClusterPrivateClusterConfig
	SetInternalValue(val *GoogleWorkstationsWorkstationClusterPrivateClusterConfig)
	ServiceAttachmentUri() *string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktn.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktn.IInterpolatingParent)
	// Experimental.
	ComputeFqn() *string
	// Experimental.
	GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{}
	// Experimental.
	GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable
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
	InterpolationAsList() cdktn.IResolvable
	// Experimental.
	InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable
	ResetAllowedProjects()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for GoogleWorkstationsWorkstationClusterPrivateClusterConfigOutputReference
type jsiiProxy_GoogleWorkstationsWorkstationClusterPrivateClusterConfigOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_GoogleWorkstationsWorkstationClusterPrivateClusterConfigOutputReference) AllowedProjects() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"allowedProjects",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleWorkstationsWorkstationClusterPrivateClusterConfigOutputReference) AllowedProjectsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"allowedProjectsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleWorkstationsWorkstationClusterPrivateClusterConfigOutputReference) ClusterHostname() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clusterHostname",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleWorkstationsWorkstationClusterPrivateClusterConfigOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleWorkstationsWorkstationClusterPrivateClusterConfigOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleWorkstationsWorkstationClusterPrivateClusterConfigOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleWorkstationsWorkstationClusterPrivateClusterConfigOutputReference) EnablePrivateEndpoint() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enablePrivateEndpoint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleWorkstationsWorkstationClusterPrivateClusterConfigOutputReference) EnablePrivateEndpointInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enablePrivateEndpointInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleWorkstationsWorkstationClusterPrivateClusterConfigOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleWorkstationsWorkstationClusterPrivateClusterConfigOutputReference) InternalValue() *GoogleWorkstationsWorkstationClusterPrivateClusterConfig {
	var returns *GoogleWorkstationsWorkstationClusterPrivateClusterConfig
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleWorkstationsWorkstationClusterPrivateClusterConfigOutputReference) ServiceAttachmentUri() *string {
	var returns *string
	_jsii_.Get(
		j,
		"serviceAttachmentUri",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleWorkstationsWorkstationClusterPrivateClusterConfigOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleWorkstationsWorkstationClusterPrivateClusterConfigOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewGoogleWorkstationsWorkstationClusterPrivateClusterConfigOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) GoogleWorkstationsWorkstationClusterPrivateClusterConfigOutputReference {
	_init_.Initialize()

	if err := validateNewGoogleWorkstationsWorkstationClusterPrivateClusterConfigOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_GoogleWorkstationsWorkstationClusterPrivateClusterConfigOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-google-beta.googleWorkstationsWorkstationCluster.GoogleWorkstationsWorkstationClusterPrivateClusterConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewGoogleWorkstationsWorkstationClusterPrivateClusterConfigOutputReference_Override(g GoogleWorkstationsWorkstationClusterPrivateClusterConfigOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-google-beta.googleWorkstationsWorkstationCluster.GoogleWorkstationsWorkstationClusterPrivateClusterConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		g,
	)
}

func (j *jsiiProxy_GoogleWorkstationsWorkstationClusterPrivateClusterConfigOutputReference)SetAllowedProjects(val *[]*string) {
	if err := j.validateSetAllowedProjectsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"allowedProjects",
		val,
	)
}

func (j *jsiiProxy_GoogleWorkstationsWorkstationClusterPrivateClusterConfigOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_GoogleWorkstationsWorkstationClusterPrivateClusterConfigOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_GoogleWorkstationsWorkstationClusterPrivateClusterConfigOutputReference)SetEnablePrivateEndpoint(val interface{}) {
	if err := j.validateSetEnablePrivateEndpointParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"enablePrivateEndpoint",
		val,
	)
}

func (j *jsiiProxy_GoogleWorkstationsWorkstationClusterPrivateClusterConfigOutputReference)SetInternalValue(val *GoogleWorkstationsWorkstationClusterPrivateClusterConfig) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_GoogleWorkstationsWorkstationClusterPrivateClusterConfigOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_GoogleWorkstationsWorkstationClusterPrivateClusterConfigOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (g *jsiiProxy_GoogleWorkstationsWorkstationClusterPrivateClusterConfigOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleWorkstationsWorkstationClusterPrivateClusterConfigOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (g *jsiiProxy_GoogleWorkstationsWorkstationClusterPrivateClusterConfigOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := g.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		g,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleWorkstationsWorkstationClusterPrivateClusterConfigOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (g *jsiiProxy_GoogleWorkstationsWorkstationClusterPrivateClusterConfigOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (g *jsiiProxy_GoogleWorkstationsWorkstationClusterPrivateClusterConfigOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (g *jsiiProxy_GoogleWorkstationsWorkstationClusterPrivateClusterConfigOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (g *jsiiProxy_GoogleWorkstationsWorkstationClusterPrivateClusterConfigOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (g *jsiiProxy_GoogleWorkstationsWorkstationClusterPrivateClusterConfigOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (g *jsiiProxy_GoogleWorkstationsWorkstationClusterPrivateClusterConfigOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (g *jsiiProxy_GoogleWorkstationsWorkstationClusterPrivateClusterConfigOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		g,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleWorkstationsWorkstationClusterPrivateClusterConfigOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := g.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		g,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleWorkstationsWorkstationClusterPrivateClusterConfigOutputReference) ResetAllowedProjects() {
	_jsii_.InvokeVoid(
		g,
		"resetAllowedProjects",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleWorkstationsWorkstationClusterPrivateClusterConfigOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
	if err := g.validateResolveParameters(context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		g,
		"resolve",
		[]interface{}{context},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleWorkstationsWorkstationClusterPrivateClusterConfigOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

