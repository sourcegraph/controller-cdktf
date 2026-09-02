package googlehypercomputeclustercluster

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/google_beta/jsii"

	"github.com/open-constructs/cdk-terrain-go/cdktn"
	"github.com/sourcegraph/controller-cdktf/gen/google_beta/googlehypercomputeclustercluster/internal"
)

type GoogleHypercomputeclusterClusterStorageResourcesConfigNewBucketOutputReference interface {
	cdktn.ComplexObject
	Autoclass() GoogleHypercomputeclusterClusterStorageResourcesConfigNewBucketAutoclassOutputReference
	AutoclassInput() *GoogleHypercomputeclusterClusterStorageResourcesConfigNewBucketAutoclass
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
	HierarchicalNamespace() GoogleHypercomputeclusterClusterStorageResourcesConfigNewBucketHierarchicalNamespaceOutputReference
	HierarchicalNamespaceInput() *GoogleHypercomputeclusterClusterStorageResourcesConfigNewBucketHierarchicalNamespace
	InternalValue() *GoogleHypercomputeclusterClusterStorageResourcesConfigNewBucket
	SetInternalValue(val *GoogleHypercomputeclusterClusterStorageResourcesConfigNewBucket)
	StorageClass() *string
	SetStorageClass(val *string)
	StorageClassInput() *string
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
	PutAutoclass(value *GoogleHypercomputeclusterClusterStorageResourcesConfigNewBucketAutoclass)
	PutHierarchicalNamespace(value *GoogleHypercomputeclusterClusterStorageResourcesConfigNewBucketHierarchicalNamespace)
	ResetAutoclass()
	ResetHierarchicalNamespace()
	ResetStorageClass()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for GoogleHypercomputeclusterClusterStorageResourcesConfigNewBucketOutputReference
type jsiiProxy_GoogleHypercomputeclusterClusterStorageResourcesConfigNewBucketOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_GoogleHypercomputeclusterClusterStorageResourcesConfigNewBucketOutputReference) Autoclass() GoogleHypercomputeclusterClusterStorageResourcesConfigNewBucketAutoclassOutputReference {
	var returns GoogleHypercomputeclusterClusterStorageResourcesConfigNewBucketAutoclassOutputReference
	_jsii_.Get(
		j,
		"autoclass",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleHypercomputeclusterClusterStorageResourcesConfigNewBucketOutputReference) AutoclassInput() *GoogleHypercomputeclusterClusterStorageResourcesConfigNewBucketAutoclass {
	var returns *GoogleHypercomputeclusterClusterStorageResourcesConfigNewBucketAutoclass
	_jsii_.Get(
		j,
		"autoclassInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleHypercomputeclusterClusterStorageResourcesConfigNewBucketOutputReference) Bucket() *string {
	var returns *string
	_jsii_.Get(
		j,
		"bucket",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleHypercomputeclusterClusterStorageResourcesConfigNewBucketOutputReference) BucketInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"bucketInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleHypercomputeclusterClusterStorageResourcesConfigNewBucketOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleHypercomputeclusterClusterStorageResourcesConfigNewBucketOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleHypercomputeclusterClusterStorageResourcesConfigNewBucketOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleHypercomputeclusterClusterStorageResourcesConfigNewBucketOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleHypercomputeclusterClusterStorageResourcesConfigNewBucketOutputReference) HierarchicalNamespace() GoogleHypercomputeclusterClusterStorageResourcesConfigNewBucketHierarchicalNamespaceOutputReference {
	var returns GoogleHypercomputeclusterClusterStorageResourcesConfigNewBucketHierarchicalNamespaceOutputReference
	_jsii_.Get(
		j,
		"hierarchicalNamespace",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleHypercomputeclusterClusterStorageResourcesConfigNewBucketOutputReference) HierarchicalNamespaceInput() *GoogleHypercomputeclusterClusterStorageResourcesConfigNewBucketHierarchicalNamespace {
	var returns *GoogleHypercomputeclusterClusterStorageResourcesConfigNewBucketHierarchicalNamespace
	_jsii_.Get(
		j,
		"hierarchicalNamespaceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleHypercomputeclusterClusterStorageResourcesConfigNewBucketOutputReference) InternalValue() *GoogleHypercomputeclusterClusterStorageResourcesConfigNewBucket {
	var returns *GoogleHypercomputeclusterClusterStorageResourcesConfigNewBucket
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleHypercomputeclusterClusterStorageResourcesConfigNewBucketOutputReference) StorageClass() *string {
	var returns *string
	_jsii_.Get(
		j,
		"storageClass",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleHypercomputeclusterClusterStorageResourcesConfigNewBucketOutputReference) StorageClassInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"storageClassInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleHypercomputeclusterClusterStorageResourcesConfigNewBucketOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleHypercomputeclusterClusterStorageResourcesConfigNewBucketOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewGoogleHypercomputeclusterClusterStorageResourcesConfigNewBucketOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) GoogleHypercomputeclusterClusterStorageResourcesConfigNewBucketOutputReference {
	_init_.Initialize()

	if err := validateNewGoogleHypercomputeclusterClusterStorageResourcesConfigNewBucketOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_GoogleHypercomputeclusterClusterStorageResourcesConfigNewBucketOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-google-beta.googleHypercomputeclusterCluster.GoogleHypercomputeclusterClusterStorageResourcesConfigNewBucketOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewGoogleHypercomputeclusterClusterStorageResourcesConfigNewBucketOutputReference_Override(g GoogleHypercomputeclusterClusterStorageResourcesConfigNewBucketOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-google-beta.googleHypercomputeclusterCluster.GoogleHypercomputeclusterClusterStorageResourcesConfigNewBucketOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		g,
	)
}

func (j *jsiiProxy_GoogleHypercomputeclusterClusterStorageResourcesConfigNewBucketOutputReference)SetBucket(val *string) {
	if err := j.validateSetBucketParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"bucket",
		val,
	)
}

func (j *jsiiProxy_GoogleHypercomputeclusterClusterStorageResourcesConfigNewBucketOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_GoogleHypercomputeclusterClusterStorageResourcesConfigNewBucketOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_GoogleHypercomputeclusterClusterStorageResourcesConfigNewBucketOutputReference)SetInternalValue(val *GoogleHypercomputeclusterClusterStorageResourcesConfigNewBucket) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_GoogleHypercomputeclusterClusterStorageResourcesConfigNewBucketOutputReference)SetStorageClass(val *string) {
	if err := j.validateSetStorageClassParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"storageClass",
		val,
	)
}

func (j *jsiiProxy_GoogleHypercomputeclusterClusterStorageResourcesConfigNewBucketOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_GoogleHypercomputeclusterClusterStorageResourcesConfigNewBucketOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (g *jsiiProxy_GoogleHypercomputeclusterClusterStorageResourcesConfigNewBucketOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleHypercomputeclusterClusterStorageResourcesConfigNewBucketOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (g *jsiiProxy_GoogleHypercomputeclusterClusterStorageResourcesConfigNewBucketOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (g *jsiiProxy_GoogleHypercomputeclusterClusterStorageResourcesConfigNewBucketOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (g *jsiiProxy_GoogleHypercomputeclusterClusterStorageResourcesConfigNewBucketOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (g *jsiiProxy_GoogleHypercomputeclusterClusterStorageResourcesConfigNewBucketOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (g *jsiiProxy_GoogleHypercomputeclusterClusterStorageResourcesConfigNewBucketOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (g *jsiiProxy_GoogleHypercomputeclusterClusterStorageResourcesConfigNewBucketOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (g *jsiiProxy_GoogleHypercomputeclusterClusterStorageResourcesConfigNewBucketOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (g *jsiiProxy_GoogleHypercomputeclusterClusterStorageResourcesConfigNewBucketOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (g *jsiiProxy_GoogleHypercomputeclusterClusterStorageResourcesConfigNewBucketOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		g,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleHypercomputeclusterClusterStorageResourcesConfigNewBucketOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (g *jsiiProxy_GoogleHypercomputeclusterClusterStorageResourcesConfigNewBucketOutputReference) PutAutoclass(value *GoogleHypercomputeclusterClusterStorageResourcesConfigNewBucketAutoclass) {
	if err := g.validatePutAutoclassParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putAutoclass",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleHypercomputeclusterClusterStorageResourcesConfigNewBucketOutputReference) PutHierarchicalNamespace(value *GoogleHypercomputeclusterClusterStorageResourcesConfigNewBucketHierarchicalNamespace) {
	if err := g.validatePutHierarchicalNamespaceParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putHierarchicalNamespace",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleHypercomputeclusterClusterStorageResourcesConfigNewBucketOutputReference) ResetAutoclass() {
	_jsii_.InvokeVoid(
		g,
		"resetAutoclass",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleHypercomputeclusterClusterStorageResourcesConfigNewBucketOutputReference) ResetHierarchicalNamespace() {
	_jsii_.InvokeVoid(
		g,
		"resetHierarchicalNamespace",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleHypercomputeclusterClusterStorageResourcesConfigNewBucketOutputReference) ResetStorageClass() {
	_jsii_.InvokeVoid(
		g,
		"resetStorageClass",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleHypercomputeclusterClusterStorageResourcesConfigNewBucketOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (g *jsiiProxy_GoogleHypercomputeclusterClusterStorageResourcesConfigNewBucketOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

