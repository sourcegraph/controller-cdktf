package googlestoragebatchoperationsjob

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/google_beta/jsii"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/sourcegraph/controller-cdktf/gen/google_beta/googlestoragebatchoperationsjob/internal"
)

type GoogleStorageBatchOperationsJobBucketListBucketsOutputReference interface {
	cdktf.ComplexObject
	Bucket() *string
	SetBucket(val *string)
	BucketInput() *string
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
	InternalValue() *GoogleStorageBatchOperationsJobBucketListBuckets
	SetInternalValue(val *GoogleStorageBatchOperationsJobBucketListBuckets)
	Manifest() GoogleStorageBatchOperationsJobBucketListBucketsManifestOutputReference
	ManifestInput() *GoogleStorageBatchOperationsJobBucketListBucketsManifest
	PrefixList() GoogleStorageBatchOperationsJobBucketListBucketsPrefixListStructOutputReference
	PrefixListInput() *GoogleStorageBatchOperationsJobBucketListBucketsPrefixListStruct
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
	PutManifest(value *GoogleStorageBatchOperationsJobBucketListBucketsManifest)
	PutPrefixList(value *GoogleStorageBatchOperationsJobBucketListBucketsPrefixListStruct)
	ResetManifest()
	ResetPrefixList()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for GoogleStorageBatchOperationsJobBucketListBucketsOutputReference
type jsiiProxy_GoogleStorageBatchOperationsJobBucketListBucketsOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_GoogleStorageBatchOperationsJobBucketListBucketsOutputReference) Bucket() *string {
	var returns *string
	_jsii_.Get(
		j,
		"bucket",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleStorageBatchOperationsJobBucketListBucketsOutputReference) BucketInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"bucketInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleStorageBatchOperationsJobBucketListBucketsOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleStorageBatchOperationsJobBucketListBucketsOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleStorageBatchOperationsJobBucketListBucketsOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleStorageBatchOperationsJobBucketListBucketsOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleStorageBatchOperationsJobBucketListBucketsOutputReference) InternalValue() *GoogleStorageBatchOperationsJobBucketListBuckets {
	var returns *GoogleStorageBatchOperationsJobBucketListBuckets
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleStorageBatchOperationsJobBucketListBucketsOutputReference) Manifest() GoogleStorageBatchOperationsJobBucketListBucketsManifestOutputReference {
	var returns GoogleStorageBatchOperationsJobBucketListBucketsManifestOutputReference
	_jsii_.Get(
		j,
		"manifest",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleStorageBatchOperationsJobBucketListBucketsOutputReference) ManifestInput() *GoogleStorageBatchOperationsJobBucketListBucketsManifest {
	var returns *GoogleStorageBatchOperationsJobBucketListBucketsManifest
	_jsii_.Get(
		j,
		"manifestInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleStorageBatchOperationsJobBucketListBucketsOutputReference) PrefixList() GoogleStorageBatchOperationsJobBucketListBucketsPrefixListStructOutputReference {
	var returns GoogleStorageBatchOperationsJobBucketListBucketsPrefixListStructOutputReference
	_jsii_.Get(
		j,
		"prefixList",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleStorageBatchOperationsJobBucketListBucketsOutputReference) PrefixListInput() *GoogleStorageBatchOperationsJobBucketListBucketsPrefixListStruct {
	var returns *GoogleStorageBatchOperationsJobBucketListBucketsPrefixListStruct
	_jsii_.Get(
		j,
		"prefixListInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleStorageBatchOperationsJobBucketListBucketsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleStorageBatchOperationsJobBucketListBucketsOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewGoogleStorageBatchOperationsJobBucketListBucketsOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) GoogleStorageBatchOperationsJobBucketListBucketsOutputReference {
	_init_.Initialize()

	if err := validateNewGoogleStorageBatchOperationsJobBucketListBucketsOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_GoogleStorageBatchOperationsJobBucketListBucketsOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-google-beta.googleStorageBatchOperationsJob.GoogleStorageBatchOperationsJobBucketListBucketsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewGoogleStorageBatchOperationsJobBucketListBucketsOutputReference_Override(g GoogleStorageBatchOperationsJobBucketListBucketsOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google-beta.googleStorageBatchOperationsJob.GoogleStorageBatchOperationsJobBucketListBucketsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		g,
	)
}

func (j *jsiiProxy_GoogleStorageBatchOperationsJobBucketListBucketsOutputReference)SetBucket(val *string) {
	if err := j.validateSetBucketParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"bucket",
		val,
	)
}

func (j *jsiiProxy_GoogleStorageBatchOperationsJobBucketListBucketsOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_GoogleStorageBatchOperationsJobBucketListBucketsOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_GoogleStorageBatchOperationsJobBucketListBucketsOutputReference)SetInternalValue(val *GoogleStorageBatchOperationsJobBucketListBuckets) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_GoogleStorageBatchOperationsJobBucketListBucketsOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_GoogleStorageBatchOperationsJobBucketListBucketsOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (g *jsiiProxy_GoogleStorageBatchOperationsJobBucketListBucketsOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleStorageBatchOperationsJobBucketListBucketsOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (g *jsiiProxy_GoogleStorageBatchOperationsJobBucketListBucketsOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GoogleStorageBatchOperationsJobBucketListBucketsOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (g *jsiiProxy_GoogleStorageBatchOperationsJobBucketListBucketsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (g *jsiiProxy_GoogleStorageBatchOperationsJobBucketListBucketsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (g *jsiiProxy_GoogleStorageBatchOperationsJobBucketListBucketsOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (g *jsiiProxy_GoogleStorageBatchOperationsJobBucketListBucketsOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (g *jsiiProxy_GoogleStorageBatchOperationsJobBucketListBucketsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (g *jsiiProxy_GoogleStorageBatchOperationsJobBucketListBucketsOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (g *jsiiProxy_GoogleStorageBatchOperationsJobBucketListBucketsOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		g,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleStorageBatchOperationsJobBucketListBucketsOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GoogleStorageBatchOperationsJobBucketListBucketsOutputReference) PutManifest(value *GoogleStorageBatchOperationsJobBucketListBucketsManifest) {
	if err := g.validatePutManifestParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putManifest",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleStorageBatchOperationsJobBucketListBucketsOutputReference) PutPrefixList(value *GoogleStorageBatchOperationsJobBucketListBucketsPrefixListStruct) {
	if err := g.validatePutPrefixListParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putPrefixList",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleStorageBatchOperationsJobBucketListBucketsOutputReference) ResetManifest() {
	_jsii_.InvokeVoid(
		g,
		"resetManifest",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleStorageBatchOperationsJobBucketListBucketsOutputReference) ResetPrefixList() {
	_jsii_.InvokeVoid(
		g,
		"resetPrefixList",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleStorageBatchOperationsJobBucketListBucketsOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (g *jsiiProxy_GoogleStorageBatchOperationsJobBucketListBucketsOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

