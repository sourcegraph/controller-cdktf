package googlepubsubtopic

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/google_beta/jsii"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/sourcegraph/controller-cdktf/gen/google_beta/googlepubsubtopic/internal"
)

type GooglePubsubTopicIngestionDataSourceSettingsAwsKinesisOutputReference interface {
	cdktf.ComplexObject
	AwsRoleArn() *string
	SetAwsRoleArn(val *string)
	AwsRoleArnInput() *string
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
	ConsumerArn() *string
	SetConsumerArn(val *string)
	ConsumerArnInput() *string
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
	InternalValue() *GooglePubsubTopicIngestionDataSourceSettingsAwsKinesis
	SetInternalValue(val *GooglePubsubTopicIngestionDataSourceSettingsAwsKinesis)
	StreamArn() *string
	SetStreamArn(val *string)
	StreamArnInput() *string
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
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for GooglePubsubTopicIngestionDataSourceSettingsAwsKinesisOutputReference
type jsiiProxy_GooglePubsubTopicIngestionDataSourceSettingsAwsKinesisOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_GooglePubsubTopicIngestionDataSourceSettingsAwsKinesisOutputReference) AwsRoleArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"awsRoleArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GooglePubsubTopicIngestionDataSourceSettingsAwsKinesisOutputReference) AwsRoleArnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"awsRoleArnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GooglePubsubTopicIngestionDataSourceSettingsAwsKinesisOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GooglePubsubTopicIngestionDataSourceSettingsAwsKinesisOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GooglePubsubTopicIngestionDataSourceSettingsAwsKinesisOutputReference) ConsumerArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"consumerArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GooglePubsubTopicIngestionDataSourceSettingsAwsKinesisOutputReference) ConsumerArnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"consumerArnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GooglePubsubTopicIngestionDataSourceSettingsAwsKinesisOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GooglePubsubTopicIngestionDataSourceSettingsAwsKinesisOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GooglePubsubTopicIngestionDataSourceSettingsAwsKinesisOutputReference) GcpServiceAccount() *string {
	var returns *string
	_jsii_.Get(
		j,
		"gcpServiceAccount",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GooglePubsubTopicIngestionDataSourceSettingsAwsKinesisOutputReference) GcpServiceAccountInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"gcpServiceAccountInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GooglePubsubTopicIngestionDataSourceSettingsAwsKinesisOutputReference) InternalValue() *GooglePubsubTopicIngestionDataSourceSettingsAwsKinesis {
	var returns *GooglePubsubTopicIngestionDataSourceSettingsAwsKinesis
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GooglePubsubTopicIngestionDataSourceSettingsAwsKinesisOutputReference) StreamArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"streamArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GooglePubsubTopicIngestionDataSourceSettingsAwsKinesisOutputReference) StreamArnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"streamArnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GooglePubsubTopicIngestionDataSourceSettingsAwsKinesisOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GooglePubsubTopicIngestionDataSourceSettingsAwsKinesisOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewGooglePubsubTopicIngestionDataSourceSettingsAwsKinesisOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) GooglePubsubTopicIngestionDataSourceSettingsAwsKinesisOutputReference {
	_init_.Initialize()

	if err := validateNewGooglePubsubTopicIngestionDataSourceSettingsAwsKinesisOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_GooglePubsubTopicIngestionDataSourceSettingsAwsKinesisOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-google-beta.googlePubsubTopic.GooglePubsubTopicIngestionDataSourceSettingsAwsKinesisOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewGooglePubsubTopicIngestionDataSourceSettingsAwsKinesisOutputReference_Override(g GooglePubsubTopicIngestionDataSourceSettingsAwsKinesisOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google-beta.googlePubsubTopic.GooglePubsubTopicIngestionDataSourceSettingsAwsKinesisOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		g,
	)
}

func (j *jsiiProxy_GooglePubsubTopicIngestionDataSourceSettingsAwsKinesisOutputReference)SetAwsRoleArn(val *string) {
	if err := j.validateSetAwsRoleArnParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"awsRoleArn",
		val,
	)
}

func (j *jsiiProxy_GooglePubsubTopicIngestionDataSourceSettingsAwsKinesisOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_GooglePubsubTopicIngestionDataSourceSettingsAwsKinesisOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_GooglePubsubTopicIngestionDataSourceSettingsAwsKinesisOutputReference)SetConsumerArn(val *string) {
	if err := j.validateSetConsumerArnParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"consumerArn",
		val,
	)
}

func (j *jsiiProxy_GooglePubsubTopicIngestionDataSourceSettingsAwsKinesisOutputReference)SetGcpServiceAccount(val *string) {
	if err := j.validateSetGcpServiceAccountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"gcpServiceAccount",
		val,
	)
}

func (j *jsiiProxy_GooglePubsubTopicIngestionDataSourceSettingsAwsKinesisOutputReference)SetInternalValue(val *GooglePubsubTopicIngestionDataSourceSettingsAwsKinesis) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_GooglePubsubTopicIngestionDataSourceSettingsAwsKinesisOutputReference)SetStreamArn(val *string) {
	if err := j.validateSetStreamArnParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"streamArn",
		val,
	)
}

func (j *jsiiProxy_GooglePubsubTopicIngestionDataSourceSettingsAwsKinesisOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_GooglePubsubTopicIngestionDataSourceSettingsAwsKinesisOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (g *jsiiProxy_GooglePubsubTopicIngestionDataSourceSettingsAwsKinesisOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GooglePubsubTopicIngestionDataSourceSettingsAwsKinesisOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (g *jsiiProxy_GooglePubsubTopicIngestionDataSourceSettingsAwsKinesisOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GooglePubsubTopicIngestionDataSourceSettingsAwsKinesisOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (g *jsiiProxy_GooglePubsubTopicIngestionDataSourceSettingsAwsKinesisOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (g *jsiiProxy_GooglePubsubTopicIngestionDataSourceSettingsAwsKinesisOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (g *jsiiProxy_GooglePubsubTopicIngestionDataSourceSettingsAwsKinesisOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (g *jsiiProxy_GooglePubsubTopicIngestionDataSourceSettingsAwsKinesisOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (g *jsiiProxy_GooglePubsubTopicIngestionDataSourceSettingsAwsKinesisOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (g *jsiiProxy_GooglePubsubTopicIngestionDataSourceSettingsAwsKinesisOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (g *jsiiProxy_GooglePubsubTopicIngestionDataSourceSettingsAwsKinesisOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		g,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GooglePubsubTopicIngestionDataSourceSettingsAwsKinesisOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GooglePubsubTopicIngestionDataSourceSettingsAwsKinesisOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (g *jsiiProxy_GooglePubsubTopicIngestionDataSourceSettingsAwsKinesisOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

