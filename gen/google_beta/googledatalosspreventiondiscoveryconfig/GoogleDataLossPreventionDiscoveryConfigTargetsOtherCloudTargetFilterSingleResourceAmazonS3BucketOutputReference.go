package googledatalosspreventiondiscoveryconfig

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/google_beta/jsii"

	"github.com/open-constructs/cdk-terrain-go/cdktn"
	"github.com/sourcegraph/controller-cdktf/gen/google_beta/googledatalosspreventiondiscoveryconfig/internal"
)

type GoogleDataLossPreventionDiscoveryConfigTargetsOtherCloudTargetFilterSingleResourceAmazonS3BucketOutputReference interface {
	cdktn.ComplexObject
	AwsAccount() GoogleDataLossPreventionDiscoveryConfigTargetsOtherCloudTargetFilterSingleResourceAmazonS3BucketAwsAccountOutputReference
	AwsAccountInput() *GoogleDataLossPreventionDiscoveryConfigTargetsOtherCloudTargetFilterSingleResourceAmazonS3BucketAwsAccount
	BucketName() *string
	SetBucketName(val *string)
	BucketNameInput() *string
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
	InternalValue() *GoogleDataLossPreventionDiscoveryConfigTargetsOtherCloudTargetFilterSingleResourceAmazonS3Bucket
	SetInternalValue(val *GoogleDataLossPreventionDiscoveryConfigTargetsOtherCloudTargetFilterSingleResourceAmazonS3Bucket)
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
	PutAwsAccount(value *GoogleDataLossPreventionDiscoveryConfigTargetsOtherCloudTargetFilterSingleResourceAmazonS3BucketAwsAccount)
	ResetAwsAccount()
	ResetBucketName()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for GoogleDataLossPreventionDiscoveryConfigTargetsOtherCloudTargetFilterSingleResourceAmazonS3BucketOutputReference
type jsiiProxy_GoogleDataLossPreventionDiscoveryConfigTargetsOtherCloudTargetFilterSingleResourceAmazonS3BucketOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_GoogleDataLossPreventionDiscoveryConfigTargetsOtherCloudTargetFilterSingleResourceAmazonS3BucketOutputReference) AwsAccount() GoogleDataLossPreventionDiscoveryConfigTargetsOtherCloudTargetFilterSingleResourceAmazonS3BucketAwsAccountOutputReference {
	var returns GoogleDataLossPreventionDiscoveryConfigTargetsOtherCloudTargetFilterSingleResourceAmazonS3BucketAwsAccountOutputReference
	_jsii_.Get(
		j,
		"awsAccount",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataLossPreventionDiscoveryConfigTargetsOtherCloudTargetFilterSingleResourceAmazonS3BucketOutputReference) AwsAccountInput() *GoogleDataLossPreventionDiscoveryConfigTargetsOtherCloudTargetFilterSingleResourceAmazonS3BucketAwsAccount {
	var returns *GoogleDataLossPreventionDiscoveryConfigTargetsOtherCloudTargetFilterSingleResourceAmazonS3BucketAwsAccount
	_jsii_.Get(
		j,
		"awsAccountInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataLossPreventionDiscoveryConfigTargetsOtherCloudTargetFilterSingleResourceAmazonS3BucketOutputReference) BucketName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"bucketName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataLossPreventionDiscoveryConfigTargetsOtherCloudTargetFilterSingleResourceAmazonS3BucketOutputReference) BucketNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"bucketNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataLossPreventionDiscoveryConfigTargetsOtherCloudTargetFilterSingleResourceAmazonS3BucketOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataLossPreventionDiscoveryConfigTargetsOtherCloudTargetFilterSingleResourceAmazonS3BucketOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataLossPreventionDiscoveryConfigTargetsOtherCloudTargetFilterSingleResourceAmazonS3BucketOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataLossPreventionDiscoveryConfigTargetsOtherCloudTargetFilterSingleResourceAmazonS3BucketOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataLossPreventionDiscoveryConfigTargetsOtherCloudTargetFilterSingleResourceAmazonS3BucketOutputReference) InternalValue() *GoogleDataLossPreventionDiscoveryConfigTargetsOtherCloudTargetFilterSingleResourceAmazonS3Bucket {
	var returns *GoogleDataLossPreventionDiscoveryConfigTargetsOtherCloudTargetFilterSingleResourceAmazonS3Bucket
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataLossPreventionDiscoveryConfigTargetsOtherCloudTargetFilterSingleResourceAmazonS3BucketOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataLossPreventionDiscoveryConfigTargetsOtherCloudTargetFilterSingleResourceAmazonS3BucketOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewGoogleDataLossPreventionDiscoveryConfigTargetsOtherCloudTargetFilterSingleResourceAmazonS3BucketOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) GoogleDataLossPreventionDiscoveryConfigTargetsOtherCloudTargetFilterSingleResourceAmazonS3BucketOutputReference {
	_init_.Initialize()

	if err := validateNewGoogleDataLossPreventionDiscoveryConfigTargetsOtherCloudTargetFilterSingleResourceAmazonS3BucketOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_GoogleDataLossPreventionDiscoveryConfigTargetsOtherCloudTargetFilterSingleResourceAmazonS3BucketOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-google-beta.googleDataLossPreventionDiscoveryConfig.GoogleDataLossPreventionDiscoveryConfigTargetsOtherCloudTargetFilterSingleResourceAmazonS3BucketOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewGoogleDataLossPreventionDiscoveryConfigTargetsOtherCloudTargetFilterSingleResourceAmazonS3BucketOutputReference_Override(g GoogleDataLossPreventionDiscoveryConfigTargetsOtherCloudTargetFilterSingleResourceAmazonS3BucketOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-google-beta.googleDataLossPreventionDiscoveryConfig.GoogleDataLossPreventionDiscoveryConfigTargetsOtherCloudTargetFilterSingleResourceAmazonS3BucketOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		g,
	)
}

func (j *jsiiProxy_GoogleDataLossPreventionDiscoveryConfigTargetsOtherCloudTargetFilterSingleResourceAmazonS3BucketOutputReference)SetBucketName(val *string) {
	if err := j.validateSetBucketNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"bucketName",
		val,
	)
}

func (j *jsiiProxy_GoogleDataLossPreventionDiscoveryConfigTargetsOtherCloudTargetFilterSingleResourceAmazonS3BucketOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_GoogleDataLossPreventionDiscoveryConfigTargetsOtherCloudTargetFilterSingleResourceAmazonS3BucketOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_GoogleDataLossPreventionDiscoveryConfigTargetsOtherCloudTargetFilterSingleResourceAmazonS3BucketOutputReference)SetInternalValue(val *GoogleDataLossPreventionDiscoveryConfigTargetsOtherCloudTargetFilterSingleResourceAmazonS3Bucket) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_GoogleDataLossPreventionDiscoveryConfigTargetsOtherCloudTargetFilterSingleResourceAmazonS3BucketOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_GoogleDataLossPreventionDiscoveryConfigTargetsOtherCloudTargetFilterSingleResourceAmazonS3BucketOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (g *jsiiProxy_GoogleDataLossPreventionDiscoveryConfigTargetsOtherCloudTargetFilterSingleResourceAmazonS3BucketOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleDataLossPreventionDiscoveryConfigTargetsOtherCloudTargetFilterSingleResourceAmazonS3BucketOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (g *jsiiProxy_GoogleDataLossPreventionDiscoveryConfigTargetsOtherCloudTargetFilterSingleResourceAmazonS3BucketOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (g *jsiiProxy_GoogleDataLossPreventionDiscoveryConfigTargetsOtherCloudTargetFilterSingleResourceAmazonS3BucketOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (g *jsiiProxy_GoogleDataLossPreventionDiscoveryConfigTargetsOtherCloudTargetFilterSingleResourceAmazonS3BucketOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (g *jsiiProxy_GoogleDataLossPreventionDiscoveryConfigTargetsOtherCloudTargetFilterSingleResourceAmazonS3BucketOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (g *jsiiProxy_GoogleDataLossPreventionDiscoveryConfigTargetsOtherCloudTargetFilterSingleResourceAmazonS3BucketOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (g *jsiiProxy_GoogleDataLossPreventionDiscoveryConfigTargetsOtherCloudTargetFilterSingleResourceAmazonS3BucketOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (g *jsiiProxy_GoogleDataLossPreventionDiscoveryConfigTargetsOtherCloudTargetFilterSingleResourceAmazonS3BucketOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (g *jsiiProxy_GoogleDataLossPreventionDiscoveryConfigTargetsOtherCloudTargetFilterSingleResourceAmazonS3BucketOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (g *jsiiProxy_GoogleDataLossPreventionDiscoveryConfigTargetsOtherCloudTargetFilterSingleResourceAmazonS3BucketOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		g,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleDataLossPreventionDiscoveryConfigTargetsOtherCloudTargetFilterSingleResourceAmazonS3BucketOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (g *jsiiProxy_GoogleDataLossPreventionDiscoveryConfigTargetsOtherCloudTargetFilterSingleResourceAmazonS3BucketOutputReference) PutAwsAccount(value *GoogleDataLossPreventionDiscoveryConfigTargetsOtherCloudTargetFilterSingleResourceAmazonS3BucketAwsAccount) {
	if err := g.validatePutAwsAccountParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putAwsAccount",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleDataLossPreventionDiscoveryConfigTargetsOtherCloudTargetFilterSingleResourceAmazonS3BucketOutputReference) ResetAwsAccount() {
	_jsii_.InvokeVoid(
		g,
		"resetAwsAccount",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleDataLossPreventionDiscoveryConfigTargetsOtherCloudTargetFilterSingleResourceAmazonS3BucketOutputReference) ResetBucketName() {
	_jsii_.InvokeVoid(
		g,
		"resetBucketName",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleDataLossPreventionDiscoveryConfigTargetsOtherCloudTargetFilterSingleResourceAmazonS3BucketOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (g *jsiiProxy_GoogleDataLossPreventionDiscoveryConfigTargetsOtherCloudTargetFilterSingleResourceAmazonS3BucketOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

