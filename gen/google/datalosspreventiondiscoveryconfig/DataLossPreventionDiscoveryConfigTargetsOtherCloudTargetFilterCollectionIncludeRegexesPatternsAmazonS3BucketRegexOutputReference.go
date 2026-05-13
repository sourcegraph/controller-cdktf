package datalosspreventiondiscoveryconfig

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/google/jsii"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/sourcegraph/controller-cdktf/gen/google/datalosspreventiondiscoveryconfig/internal"
)

type DataLossPreventionDiscoveryConfigTargetsOtherCloudTargetFilterCollectionIncludeRegexesPatternsAmazonS3BucketRegexOutputReference interface {
	cdktf.ComplexObject
	AwsAccountRegex() DataLossPreventionDiscoveryConfigTargetsOtherCloudTargetFilterCollectionIncludeRegexesPatternsAmazonS3BucketRegexAwsAccountRegexOutputReference
	AwsAccountRegexInput() *DataLossPreventionDiscoveryConfigTargetsOtherCloudTargetFilterCollectionIncludeRegexesPatternsAmazonS3BucketRegexAwsAccountRegex
	BucketNameRegex() *string
	SetBucketNameRegex(val *string)
	BucketNameRegexInput() *string
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
	InternalValue() *DataLossPreventionDiscoveryConfigTargetsOtherCloudTargetFilterCollectionIncludeRegexesPatternsAmazonS3BucketRegex
	SetInternalValue(val *DataLossPreventionDiscoveryConfigTargetsOtherCloudTargetFilterCollectionIncludeRegexesPatternsAmazonS3BucketRegex)
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
	PutAwsAccountRegex(value *DataLossPreventionDiscoveryConfigTargetsOtherCloudTargetFilterCollectionIncludeRegexesPatternsAmazonS3BucketRegexAwsAccountRegex)
	ResetAwsAccountRegex()
	ResetBucketNameRegex()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for DataLossPreventionDiscoveryConfigTargetsOtherCloudTargetFilterCollectionIncludeRegexesPatternsAmazonS3BucketRegexOutputReference
type jsiiProxy_DataLossPreventionDiscoveryConfigTargetsOtherCloudTargetFilterCollectionIncludeRegexesPatternsAmazonS3BucketRegexOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_DataLossPreventionDiscoveryConfigTargetsOtherCloudTargetFilterCollectionIncludeRegexesPatternsAmazonS3BucketRegexOutputReference) AwsAccountRegex() DataLossPreventionDiscoveryConfigTargetsOtherCloudTargetFilterCollectionIncludeRegexesPatternsAmazonS3BucketRegexAwsAccountRegexOutputReference {
	var returns DataLossPreventionDiscoveryConfigTargetsOtherCloudTargetFilterCollectionIncludeRegexesPatternsAmazonS3BucketRegexAwsAccountRegexOutputReference
	_jsii_.Get(
		j,
		"awsAccountRegex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataLossPreventionDiscoveryConfigTargetsOtherCloudTargetFilterCollectionIncludeRegexesPatternsAmazonS3BucketRegexOutputReference) AwsAccountRegexInput() *DataLossPreventionDiscoveryConfigTargetsOtherCloudTargetFilterCollectionIncludeRegexesPatternsAmazonS3BucketRegexAwsAccountRegex {
	var returns *DataLossPreventionDiscoveryConfigTargetsOtherCloudTargetFilterCollectionIncludeRegexesPatternsAmazonS3BucketRegexAwsAccountRegex
	_jsii_.Get(
		j,
		"awsAccountRegexInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataLossPreventionDiscoveryConfigTargetsOtherCloudTargetFilterCollectionIncludeRegexesPatternsAmazonS3BucketRegexOutputReference) BucketNameRegex() *string {
	var returns *string
	_jsii_.Get(
		j,
		"bucketNameRegex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataLossPreventionDiscoveryConfigTargetsOtherCloudTargetFilterCollectionIncludeRegexesPatternsAmazonS3BucketRegexOutputReference) BucketNameRegexInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"bucketNameRegexInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataLossPreventionDiscoveryConfigTargetsOtherCloudTargetFilterCollectionIncludeRegexesPatternsAmazonS3BucketRegexOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataLossPreventionDiscoveryConfigTargetsOtherCloudTargetFilterCollectionIncludeRegexesPatternsAmazonS3BucketRegexOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataLossPreventionDiscoveryConfigTargetsOtherCloudTargetFilterCollectionIncludeRegexesPatternsAmazonS3BucketRegexOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataLossPreventionDiscoveryConfigTargetsOtherCloudTargetFilterCollectionIncludeRegexesPatternsAmazonS3BucketRegexOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataLossPreventionDiscoveryConfigTargetsOtherCloudTargetFilterCollectionIncludeRegexesPatternsAmazonS3BucketRegexOutputReference) InternalValue() *DataLossPreventionDiscoveryConfigTargetsOtherCloudTargetFilterCollectionIncludeRegexesPatternsAmazonS3BucketRegex {
	var returns *DataLossPreventionDiscoveryConfigTargetsOtherCloudTargetFilterCollectionIncludeRegexesPatternsAmazonS3BucketRegex
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataLossPreventionDiscoveryConfigTargetsOtherCloudTargetFilterCollectionIncludeRegexesPatternsAmazonS3BucketRegexOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataLossPreventionDiscoveryConfigTargetsOtherCloudTargetFilterCollectionIncludeRegexesPatternsAmazonS3BucketRegexOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewDataLossPreventionDiscoveryConfigTargetsOtherCloudTargetFilterCollectionIncludeRegexesPatternsAmazonS3BucketRegexOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) DataLossPreventionDiscoveryConfigTargetsOtherCloudTargetFilterCollectionIncludeRegexesPatternsAmazonS3BucketRegexOutputReference {
	_init_.Initialize()

	if err := validateNewDataLossPreventionDiscoveryConfigTargetsOtherCloudTargetFilterCollectionIncludeRegexesPatternsAmazonS3BucketRegexOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_DataLossPreventionDiscoveryConfigTargetsOtherCloudTargetFilterCollectionIncludeRegexesPatternsAmazonS3BucketRegexOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-google.dataLossPreventionDiscoveryConfig.DataLossPreventionDiscoveryConfigTargetsOtherCloudTargetFilterCollectionIncludeRegexesPatternsAmazonS3BucketRegexOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewDataLossPreventionDiscoveryConfigTargetsOtherCloudTargetFilterCollectionIncludeRegexesPatternsAmazonS3BucketRegexOutputReference_Override(d DataLossPreventionDiscoveryConfigTargetsOtherCloudTargetFilterCollectionIncludeRegexesPatternsAmazonS3BucketRegexOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google.dataLossPreventionDiscoveryConfig.DataLossPreventionDiscoveryConfigTargetsOtherCloudTargetFilterCollectionIncludeRegexesPatternsAmazonS3BucketRegexOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		d,
	)
}

func (j *jsiiProxy_DataLossPreventionDiscoveryConfigTargetsOtherCloudTargetFilterCollectionIncludeRegexesPatternsAmazonS3BucketRegexOutputReference)SetBucketNameRegex(val *string) {
	if err := j.validateSetBucketNameRegexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"bucketNameRegex",
		val,
	)
}

func (j *jsiiProxy_DataLossPreventionDiscoveryConfigTargetsOtherCloudTargetFilterCollectionIncludeRegexesPatternsAmazonS3BucketRegexOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_DataLossPreventionDiscoveryConfigTargetsOtherCloudTargetFilterCollectionIncludeRegexesPatternsAmazonS3BucketRegexOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_DataLossPreventionDiscoveryConfigTargetsOtherCloudTargetFilterCollectionIncludeRegexesPatternsAmazonS3BucketRegexOutputReference)SetInternalValue(val *DataLossPreventionDiscoveryConfigTargetsOtherCloudTargetFilterCollectionIncludeRegexesPatternsAmazonS3BucketRegex) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_DataLossPreventionDiscoveryConfigTargetsOtherCloudTargetFilterCollectionIncludeRegexesPatternsAmazonS3BucketRegexOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DataLossPreventionDiscoveryConfigTargetsOtherCloudTargetFilterCollectionIncludeRegexesPatternsAmazonS3BucketRegexOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (d *jsiiProxy_DataLossPreventionDiscoveryConfigTargetsOtherCloudTargetFilterCollectionIncludeRegexesPatternsAmazonS3BucketRegexOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataLossPreventionDiscoveryConfigTargetsOtherCloudTargetFilterCollectionIncludeRegexesPatternsAmazonS3BucketRegexOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := d.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		d,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataLossPreventionDiscoveryConfigTargetsOtherCloudTargetFilterCollectionIncludeRegexesPatternsAmazonS3BucketRegexOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := d.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataLossPreventionDiscoveryConfigTargetsOtherCloudTargetFilterCollectionIncludeRegexesPatternsAmazonS3BucketRegexOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := d.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		d,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataLossPreventionDiscoveryConfigTargetsOtherCloudTargetFilterCollectionIncludeRegexesPatternsAmazonS3BucketRegexOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := d.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		d,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataLossPreventionDiscoveryConfigTargetsOtherCloudTargetFilterCollectionIncludeRegexesPatternsAmazonS3BucketRegexOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := d.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		d,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataLossPreventionDiscoveryConfigTargetsOtherCloudTargetFilterCollectionIncludeRegexesPatternsAmazonS3BucketRegexOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := d.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		d,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataLossPreventionDiscoveryConfigTargetsOtherCloudTargetFilterCollectionIncludeRegexesPatternsAmazonS3BucketRegexOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := d.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		d,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataLossPreventionDiscoveryConfigTargetsOtherCloudTargetFilterCollectionIncludeRegexesPatternsAmazonS3BucketRegexOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := d.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		d,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataLossPreventionDiscoveryConfigTargetsOtherCloudTargetFilterCollectionIncludeRegexesPatternsAmazonS3BucketRegexOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := d.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		d,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataLossPreventionDiscoveryConfigTargetsOtherCloudTargetFilterCollectionIncludeRegexesPatternsAmazonS3BucketRegexOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataLossPreventionDiscoveryConfigTargetsOtherCloudTargetFilterCollectionIncludeRegexesPatternsAmazonS3BucketRegexOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	if err := d.validateInterpolationForAttributeParameters(property); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataLossPreventionDiscoveryConfigTargetsOtherCloudTargetFilterCollectionIncludeRegexesPatternsAmazonS3BucketRegexOutputReference) PutAwsAccountRegex(value *DataLossPreventionDiscoveryConfigTargetsOtherCloudTargetFilterCollectionIncludeRegexesPatternsAmazonS3BucketRegexAwsAccountRegex) {
	if err := d.validatePutAwsAccountRegexParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putAwsAccountRegex",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DataLossPreventionDiscoveryConfigTargetsOtherCloudTargetFilterCollectionIncludeRegexesPatternsAmazonS3BucketRegexOutputReference) ResetAwsAccountRegex() {
	_jsii_.InvokeVoid(
		d,
		"resetAwsAccountRegex",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataLossPreventionDiscoveryConfigTargetsOtherCloudTargetFilterCollectionIncludeRegexesPatternsAmazonS3BucketRegexOutputReference) ResetBucketNameRegex() {
	_jsii_.InvokeVoid(
		d,
		"resetBucketNameRegex",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataLossPreventionDiscoveryConfigTargetsOtherCloudTargetFilterCollectionIncludeRegexesPatternsAmazonS3BucketRegexOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	if err := d.validateResolveParameters(_context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		d,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataLossPreventionDiscoveryConfigTargetsOtherCloudTargetFilterCollectionIncludeRegexesPatternsAmazonS3BucketRegexOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

