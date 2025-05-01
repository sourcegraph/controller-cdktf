package googlepubsubtopic

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/google_beta/jsii"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/sourcegraph/controller-cdktf/gen/google_beta/googlepubsubtopic/internal"
)

type GooglePubsubTopicIngestionDataSourceSettingsAwsMskOutputReference interface {
	cdktf.ComplexObject
	AwsRoleArn() *string
	SetAwsRoleArn(val *string)
	AwsRoleArnInput() *string
	ClusterArn() *string
	SetClusterArn(val *string)
	ClusterArnInput() *string
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
	GcpServiceAccount() *string
	SetGcpServiceAccount(val *string)
	GcpServiceAccountInput() *string
	InternalValue() *GooglePubsubTopicIngestionDataSourceSettingsAwsMsk
	SetInternalValue(val *GooglePubsubTopicIngestionDataSourceSettingsAwsMsk)
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	Topic() *string
	SetTopic(val *string)
	TopicInput() *string
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
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for GooglePubsubTopicIngestionDataSourceSettingsAwsMskOutputReference
type jsiiProxy_GooglePubsubTopicIngestionDataSourceSettingsAwsMskOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_GooglePubsubTopicIngestionDataSourceSettingsAwsMskOutputReference) AwsRoleArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"awsRoleArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GooglePubsubTopicIngestionDataSourceSettingsAwsMskOutputReference) AwsRoleArnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"awsRoleArnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GooglePubsubTopicIngestionDataSourceSettingsAwsMskOutputReference) ClusterArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clusterArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GooglePubsubTopicIngestionDataSourceSettingsAwsMskOutputReference) ClusterArnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clusterArnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GooglePubsubTopicIngestionDataSourceSettingsAwsMskOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GooglePubsubTopicIngestionDataSourceSettingsAwsMskOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GooglePubsubTopicIngestionDataSourceSettingsAwsMskOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GooglePubsubTopicIngestionDataSourceSettingsAwsMskOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GooglePubsubTopicIngestionDataSourceSettingsAwsMskOutputReference) GcpServiceAccount() *string {
	var returns *string
	_jsii_.Get(
		j,
		"gcpServiceAccount",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GooglePubsubTopicIngestionDataSourceSettingsAwsMskOutputReference) GcpServiceAccountInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"gcpServiceAccountInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GooglePubsubTopicIngestionDataSourceSettingsAwsMskOutputReference) InternalValue() *GooglePubsubTopicIngestionDataSourceSettingsAwsMsk {
	var returns *GooglePubsubTopicIngestionDataSourceSettingsAwsMsk
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GooglePubsubTopicIngestionDataSourceSettingsAwsMskOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GooglePubsubTopicIngestionDataSourceSettingsAwsMskOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GooglePubsubTopicIngestionDataSourceSettingsAwsMskOutputReference) Topic() *string {
	var returns *string
	_jsii_.Get(
		j,
		"topic",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GooglePubsubTopicIngestionDataSourceSettingsAwsMskOutputReference) TopicInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"topicInput",
		&returns,
	)
	return returns
}


func NewGooglePubsubTopicIngestionDataSourceSettingsAwsMskOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) GooglePubsubTopicIngestionDataSourceSettingsAwsMskOutputReference {
	_init_.Initialize()

	if err := validateNewGooglePubsubTopicIngestionDataSourceSettingsAwsMskOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_GooglePubsubTopicIngestionDataSourceSettingsAwsMskOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-google-beta.googlePubsubTopic.GooglePubsubTopicIngestionDataSourceSettingsAwsMskOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewGooglePubsubTopicIngestionDataSourceSettingsAwsMskOutputReference_Override(g GooglePubsubTopicIngestionDataSourceSettingsAwsMskOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google-beta.googlePubsubTopic.GooglePubsubTopicIngestionDataSourceSettingsAwsMskOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		g,
	)
}

func (j *jsiiProxy_GooglePubsubTopicIngestionDataSourceSettingsAwsMskOutputReference)SetAwsRoleArn(val *string) {
	if err := j.validateSetAwsRoleArnParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"awsRoleArn",
		val,
	)
}

func (j *jsiiProxy_GooglePubsubTopicIngestionDataSourceSettingsAwsMskOutputReference)SetClusterArn(val *string) {
	if err := j.validateSetClusterArnParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"clusterArn",
		val,
	)
}

func (j *jsiiProxy_GooglePubsubTopicIngestionDataSourceSettingsAwsMskOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_GooglePubsubTopicIngestionDataSourceSettingsAwsMskOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_GooglePubsubTopicIngestionDataSourceSettingsAwsMskOutputReference)SetGcpServiceAccount(val *string) {
	if err := j.validateSetGcpServiceAccountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"gcpServiceAccount",
		val,
	)
}

func (j *jsiiProxy_GooglePubsubTopicIngestionDataSourceSettingsAwsMskOutputReference)SetInternalValue(val *GooglePubsubTopicIngestionDataSourceSettingsAwsMsk) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_GooglePubsubTopicIngestionDataSourceSettingsAwsMskOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_GooglePubsubTopicIngestionDataSourceSettingsAwsMskOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_GooglePubsubTopicIngestionDataSourceSettingsAwsMskOutputReference)SetTopic(val *string) {
	if err := j.validateSetTopicParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"topic",
		val,
	)
}

func (g *jsiiProxy_GooglePubsubTopicIngestionDataSourceSettingsAwsMskOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GooglePubsubTopicIngestionDataSourceSettingsAwsMskOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (g *jsiiProxy_GooglePubsubTopicIngestionDataSourceSettingsAwsMskOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GooglePubsubTopicIngestionDataSourceSettingsAwsMskOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (g *jsiiProxy_GooglePubsubTopicIngestionDataSourceSettingsAwsMskOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (g *jsiiProxy_GooglePubsubTopicIngestionDataSourceSettingsAwsMskOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (g *jsiiProxy_GooglePubsubTopicIngestionDataSourceSettingsAwsMskOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (g *jsiiProxy_GooglePubsubTopicIngestionDataSourceSettingsAwsMskOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (g *jsiiProxy_GooglePubsubTopicIngestionDataSourceSettingsAwsMskOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (g *jsiiProxy_GooglePubsubTopicIngestionDataSourceSettingsAwsMskOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (g *jsiiProxy_GooglePubsubTopicIngestionDataSourceSettingsAwsMskOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		g,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GooglePubsubTopicIngestionDataSourceSettingsAwsMskOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GooglePubsubTopicIngestionDataSourceSettingsAwsMskOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (g *jsiiProxy_GooglePubsubTopicIngestionDataSourceSettingsAwsMskOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

