package googlehypercomputeclustercluster

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/google_beta/jsii"

	"github.com/open-constructs/cdk-terrain-go/cdktn"
	"github.com/sourcegraph/controller-cdktf/gen/google_beta/googlehypercomputeclustercluster/internal"
)

type GoogleHypercomputeclusterClusterStorageResourcesConfigOutputReference interface {
	cdktn.ComplexObject
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
	ExistingBucket() GoogleHypercomputeclusterClusterStorageResourcesConfigExistingBucketOutputReference
	ExistingBucketInput() *GoogleHypercomputeclusterClusterStorageResourcesConfigExistingBucket
	ExistingFilestore() GoogleHypercomputeclusterClusterStorageResourcesConfigExistingFilestoreOutputReference
	ExistingFilestoreInput() *GoogleHypercomputeclusterClusterStorageResourcesConfigExistingFilestore
	ExistingLustre() GoogleHypercomputeclusterClusterStorageResourcesConfigExistingLustreOutputReference
	ExistingLustreInput() *GoogleHypercomputeclusterClusterStorageResourcesConfigExistingLustre
	// Experimental.
	Fqn() *string
	InternalValue() *GoogleHypercomputeclusterClusterStorageResourcesConfig
	SetInternalValue(val *GoogleHypercomputeclusterClusterStorageResourcesConfig)
	NewBucket() GoogleHypercomputeclusterClusterStorageResourcesConfigNewBucketOutputReference
	NewBucketInput() *GoogleHypercomputeclusterClusterStorageResourcesConfigNewBucket
	NewFilestore() GoogleHypercomputeclusterClusterStorageResourcesConfigNewFilestoreOutputReference
	NewFilestoreInput() *GoogleHypercomputeclusterClusterStorageResourcesConfigNewFilestore
	NewLustre() GoogleHypercomputeclusterClusterStorageResourcesConfigNewLustreOutputReference
	NewLustreInput() *GoogleHypercomputeclusterClusterStorageResourcesConfigNewLustre
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
	PutExistingBucket(value *GoogleHypercomputeclusterClusterStorageResourcesConfigExistingBucket)
	PutExistingFilestore(value *GoogleHypercomputeclusterClusterStorageResourcesConfigExistingFilestore)
	PutExistingLustre(value *GoogleHypercomputeclusterClusterStorageResourcesConfigExistingLustre)
	PutNewBucket(value *GoogleHypercomputeclusterClusterStorageResourcesConfigNewBucket)
	PutNewFilestore(value *GoogleHypercomputeclusterClusterStorageResourcesConfigNewFilestore)
	PutNewLustre(value *GoogleHypercomputeclusterClusterStorageResourcesConfigNewLustre)
	ResetExistingBucket()
	ResetExistingFilestore()
	ResetExistingLustre()
	ResetNewBucket()
	ResetNewFilestore()
	ResetNewLustre()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for GoogleHypercomputeclusterClusterStorageResourcesConfigOutputReference
type jsiiProxy_GoogleHypercomputeclusterClusterStorageResourcesConfigOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_GoogleHypercomputeclusterClusterStorageResourcesConfigOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleHypercomputeclusterClusterStorageResourcesConfigOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleHypercomputeclusterClusterStorageResourcesConfigOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleHypercomputeclusterClusterStorageResourcesConfigOutputReference) ExistingBucket() GoogleHypercomputeclusterClusterStorageResourcesConfigExistingBucketOutputReference {
	var returns GoogleHypercomputeclusterClusterStorageResourcesConfigExistingBucketOutputReference
	_jsii_.Get(
		j,
		"existingBucket",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleHypercomputeclusterClusterStorageResourcesConfigOutputReference) ExistingBucketInput() *GoogleHypercomputeclusterClusterStorageResourcesConfigExistingBucket {
	var returns *GoogleHypercomputeclusterClusterStorageResourcesConfigExistingBucket
	_jsii_.Get(
		j,
		"existingBucketInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleHypercomputeclusterClusterStorageResourcesConfigOutputReference) ExistingFilestore() GoogleHypercomputeclusterClusterStorageResourcesConfigExistingFilestoreOutputReference {
	var returns GoogleHypercomputeclusterClusterStorageResourcesConfigExistingFilestoreOutputReference
	_jsii_.Get(
		j,
		"existingFilestore",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleHypercomputeclusterClusterStorageResourcesConfigOutputReference) ExistingFilestoreInput() *GoogleHypercomputeclusterClusterStorageResourcesConfigExistingFilestore {
	var returns *GoogleHypercomputeclusterClusterStorageResourcesConfigExistingFilestore
	_jsii_.Get(
		j,
		"existingFilestoreInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleHypercomputeclusterClusterStorageResourcesConfigOutputReference) ExistingLustre() GoogleHypercomputeclusterClusterStorageResourcesConfigExistingLustreOutputReference {
	var returns GoogleHypercomputeclusterClusterStorageResourcesConfigExistingLustreOutputReference
	_jsii_.Get(
		j,
		"existingLustre",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleHypercomputeclusterClusterStorageResourcesConfigOutputReference) ExistingLustreInput() *GoogleHypercomputeclusterClusterStorageResourcesConfigExistingLustre {
	var returns *GoogleHypercomputeclusterClusterStorageResourcesConfigExistingLustre
	_jsii_.Get(
		j,
		"existingLustreInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleHypercomputeclusterClusterStorageResourcesConfigOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleHypercomputeclusterClusterStorageResourcesConfigOutputReference) InternalValue() *GoogleHypercomputeclusterClusterStorageResourcesConfig {
	var returns *GoogleHypercomputeclusterClusterStorageResourcesConfig
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleHypercomputeclusterClusterStorageResourcesConfigOutputReference) NewBucket() GoogleHypercomputeclusterClusterStorageResourcesConfigNewBucketOutputReference {
	var returns GoogleHypercomputeclusterClusterStorageResourcesConfigNewBucketOutputReference
	_jsii_.Get(
		j,
		"newBucket",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleHypercomputeclusterClusterStorageResourcesConfigOutputReference) NewBucketInput() *GoogleHypercomputeclusterClusterStorageResourcesConfigNewBucket {
	var returns *GoogleHypercomputeclusterClusterStorageResourcesConfigNewBucket
	_jsii_.Get(
		j,
		"newBucketInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleHypercomputeclusterClusterStorageResourcesConfigOutputReference) NewFilestore() GoogleHypercomputeclusterClusterStorageResourcesConfigNewFilestoreOutputReference {
	var returns GoogleHypercomputeclusterClusterStorageResourcesConfigNewFilestoreOutputReference
	_jsii_.Get(
		j,
		"newFilestore",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleHypercomputeclusterClusterStorageResourcesConfigOutputReference) NewFilestoreInput() *GoogleHypercomputeclusterClusterStorageResourcesConfigNewFilestore {
	var returns *GoogleHypercomputeclusterClusterStorageResourcesConfigNewFilestore
	_jsii_.Get(
		j,
		"newFilestoreInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleHypercomputeclusterClusterStorageResourcesConfigOutputReference) NewLustre() GoogleHypercomputeclusterClusterStorageResourcesConfigNewLustreOutputReference {
	var returns GoogleHypercomputeclusterClusterStorageResourcesConfigNewLustreOutputReference
	_jsii_.Get(
		j,
		"newLustre",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleHypercomputeclusterClusterStorageResourcesConfigOutputReference) NewLustreInput() *GoogleHypercomputeclusterClusterStorageResourcesConfigNewLustre {
	var returns *GoogleHypercomputeclusterClusterStorageResourcesConfigNewLustre
	_jsii_.Get(
		j,
		"newLustreInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleHypercomputeclusterClusterStorageResourcesConfigOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleHypercomputeclusterClusterStorageResourcesConfigOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewGoogleHypercomputeclusterClusterStorageResourcesConfigOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) GoogleHypercomputeclusterClusterStorageResourcesConfigOutputReference {
	_init_.Initialize()

	if err := validateNewGoogleHypercomputeclusterClusterStorageResourcesConfigOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_GoogleHypercomputeclusterClusterStorageResourcesConfigOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-google-beta.googleHypercomputeclusterCluster.GoogleHypercomputeclusterClusterStorageResourcesConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewGoogleHypercomputeclusterClusterStorageResourcesConfigOutputReference_Override(g GoogleHypercomputeclusterClusterStorageResourcesConfigOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-google-beta.googleHypercomputeclusterCluster.GoogleHypercomputeclusterClusterStorageResourcesConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		g,
	)
}

func (j *jsiiProxy_GoogleHypercomputeclusterClusterStorageResourcesConfigOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_GoogleHypercomputeclusterClusterStorageResourcesConfigOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_GoogleHypercomputeclusterClusterStorageResourcesConfigOutputReference)SetInternalValue(val *GoogleHypercomputeclusterClusterStorageResourcesConfig) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_GoogleHypercomputeclusterClusterStorageResourcesConfigOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_GoogleHypercomputeclusterClusterStorageResourcesConfigOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (g *jsiiProxy_GoogleHypercomputeclusterClusterStorageResourcesConfigOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleHypercomputeclusterClusterStorageResourcesConfigOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (g *jsiiProxy_GoogleHypercomputeclusterClusterStorageResourcesConfigOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (g *jsiiProxy_GoogleHypercomputeclusterClusterStorageResourcesConfigOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (g *jsiiProxy_GoogleHypercomputeclusterClusterStorageResourcesConfigOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (g *jsiiProxy_GoogleHypercomputeclusterClusterStorageResourcesConfigOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (g *jsiiProxy_GoogleHypercomputeclusterClusterStorageResourcesConfigOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (g *jsiiProxy_GoogleHypercomputeclusterClusterStorageResourcesConfigOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (g *jsiiProxy_GoogleHypercomputeclusterClusterStorageResourcesConfigOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (g *jsiiProxy_GoogleHypercomputeclusterClusterStorageResourcesConfigOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (g *jsiiProxy_GoogleHypercomputeclusterClusterStorageResourcesConfigOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		g,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleHypercomputeclusterClusterStorageResourcesConfigOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (g *jsiiProxy_GoogleHypercomputeclusterClusterStorageResourcesConfigOutputReference) PutExistingBucket(value *GoogleHypercomputeclusterClusterStorageResourcesConfigExistingBucket) {
	if err := g.validatePutExistingBucketParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putExistingBucket",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleHypercomputeclusterClusterStorageResourcesConfigOutputReference) PutExistingFilestore(value *GoogleHypercomputeclusterClusterStorageResourcesConfigExistingFilestore) {
	if err := g.validatePutExistingFilestoreParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putExistingFilestore",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleHypercomputeclusterClusterStorageResourcesConfigOutputReference) PutExistingLustre(value *GoogleHypercomputeclusterClusterStorageResourcesConfigExistingLustre) {
	if err := g.validatePutExistingLustreParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putExistingLustre",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleHypercomputeclusterClusterStorageResourcesConfigOutputReference) PutNewBucket(value *GoogleHypercomputeclusterClusterStorageResourcesConfigNewBucket) {
	if err := g.validatePutNewBucketParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putNewBucket",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleHypercomputeclusterClusterStorageResourcesConfigOutputReference) PutNewFilestore(value *GoogleHypercomputeclusterClusterStorageResourcesConfigNewFilestore) {
	if err := g.validatePutNewFilestoreParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putNewFilestore",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleHypercomputeclusterClusterStorageResourcesConfigOutputReference) PutNewLustre(value *GoogleHypercomputeclusterClusterStorageResourcesConfigNewLustre) {
	if err := g.validatePutNewLustreParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putNewLustre",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleHypercomputeclusterClusterStorageResourcesConfigOutputReference) ResetExistingBucket() {
	_jsii_.InvokeVoid(
		g,
		"resetExistingBucket",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleHypercomputeclusterClusterStorageResourcesConfigOutputReference) ResetExistingFilestore() {
	_jsii_.InvokeVoid(
		g,
		"resetExistingFilestore",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleHypercomputeclusterClusterStorageResourcesConfigOutputReference) ResetExistingLustre() {
	_jsii_.InvokeVoid(
		g,
		"resetExistingLustre",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleHypercomputeclusterClusterStorageResourcesConfigOutputReference) ResetNewBucket() {
	_jsii_.InvokeVoid(
		g,
		"resetNewBucket",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleHypercomputeclusterClusterStorageResourcesConfigOutputReference) ResetNewFilestore() {
	_jsii_.InvokeVoid(
		g,
		"resetNewFilestore",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleHypercomputeclusterClusterStorageResourcesConfigOutputReference) ResetNewLustre() {
	_jsii_.InvokeVoid(
		g,
		"resetNewLustre",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleHypercomputeclusterClusterStorageResourcesConfigOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (g *jsiiProxy_GoogleHypercomputeclusterClusterStorageResourcesConfigOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

