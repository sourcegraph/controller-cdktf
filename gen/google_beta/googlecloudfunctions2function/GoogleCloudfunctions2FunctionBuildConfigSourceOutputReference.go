package googlecloudfunctions2function

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/google_beta/jsii"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/sourcegraph/controller-cdktf/gen/google_beta/googlecloudfunctions2function/internal"
)

type GoogleCloudfunctions2FunctionBuildConfigSourceOutputReference interface {
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
	InternalValue() *GoogleCloudfunctions2FunctionBuildConfigSource
	SetInternalValue(val *GoogleCloudfunctions2FunctionBuildConfigSource)
	RepoSource() GoogleCloudfunctions2FunctionBuildConfigSourceRepoSourceOutputReference
	RepoSourceInput() *GoogleCloudfunctions2FunctionBuildConfigSourceRepoSource
	StorageSource() GoogleCloudfunctions2FunctionBuildConfigSourceStorageSourceOutputReference
	StorageSourceInput() *GoogleCloudfunctions2FunctionBuildConfigSourceStorageSource
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
	PutRepoSource(value *GoogleCloudfunctions2FunctionBuildConfigSourceRepoSource)
	PutStorageSource(value *GoogleCloudfunctions2FunctionBuildConfigSourceStorageSource)
	ResetRepoSource()
	ResetStorageSource()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for GoogleCloudfunctions2FunctionBuildConfigSourceOutputReference
type jsiiProxy_GoogleCloudfunctions2FunctionBuildConfigSourceOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_GoogleCloudfunctions2FunctionBuildConfigSourceOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCloudfunctions2FunctionBuildConfigSourceOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCloudfunctions2FunctionBuildConfigSourceOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCloudfunctions2FunctionBuildConfigSourceOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCloudfunctions2FunctionBuildConfigSourceOutputReference) InternalValue() *GoogleCloudfunctions2FunctionBuildConfigSource {
	var returns *GoogleCloudfunctions2FunctionBuildConfigSource
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCloudfunctions2FunctionBuildConfigSourceOutputReference) RepoSource() GoogleCloudfunctions2FunctionBuildConfigSourceRepoSourceOutputReference {
	var returns GoogleCloudfunctions2FunctionBuildConfigSourceRepoSourceOutputReference
	_jsii_.Get(
		j,
		"repoSource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCloudfunctions2FunctionBuildConfigSourceOutputReference) RepoSourceInput() *GoogleCloudfunctions2FunctionBuildConfigSourceRepoSource {
	var returns *GoogleCloudfunctions2FunctionBuildConfigSourceRepoSource
	_jsii_.Get(
		j,
		"repoSourceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCloudfunctions2FunctionBuildConfigSourceOutputReference) StorageSource() GoogleCloudfunctions2FunctionBuildConfigSourceStorageSourceOutputReference {
	var returns GoogleCloudfunctions2FunctionBuildConfigSourceStorageSourceOutputReference
	_jsii_.Get(
		j,
		"storageSource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCloudfunctions2FunctionBuildConfigSourceOutputReference) StorageSourceInput() *GoogleCloudfunctions2FunctionBuildConfigSourceStorageSource {
	var returns *GoogleCloudfunctions2FunctionBuildConfigSourceStorageSource
	_jsii_.Get(
		j,
		"storageSourceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCloudfunctions2FunctionBuildConfigSourceOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCloudfunctions2FunctionBuildConfigSourceOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewGoogleCloudfunctions2FunctionBuildConfigSourceOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) GoogleCloudfunctions2FunctionBuildConfigSourceOutputReference {
	_init_.Initialize()

	if err := validateNewGoogleCloudfunctions2FunctionBuildConfigSourceOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_GoogleCloudfunctions2FunctionBuildConfigSourceOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-google-beta.googleCloudfunctions2Function.GoogleCloudfunctions2FunctionBuildConfigSourceOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewGoogleCloudfunctions2FunctionBuildConfigSourceOutputReference_Override(g GoogleCloudfunctions2FunctionBuildConfigSourceOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google-beta.googleCloudfunctions2Function.GoogleCloudfunctions2FunctionBuildConfigSourceOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		g,
	)
}

func (j *jsiiProxy_GoogleCloudfunctions2FunctionBuildConfigSourceOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_GoogleCloudfunctions2FunctionBuildConfigSourceOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_GoogleCloudfunctions2FunctionBuildConfigSourceOutputReference)SetInternalValue(val *GoogleCloudfunctions2FunctionBuildConfigSource) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_GoogleCloudfunctions2FunctionBuildConfigSourceOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_GoogleCloudfunctions2FunctionBuildConfigSourceOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (g *jsiiProxy_GoogleCloudfunctions2FunctionBuildConfigSourceOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleCloudfunctions2FunctionBuildConfigSourceOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (g *jsiiProxy_GoogleCloudfunctions2FunctionBuildConfigSourceOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GoogleCloudfunctions2FunctionBuildConfigSourceOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (g *jsiiProxy_GoogleCloudfunctions2FunctionBuildConfigSourceOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (g *jsiiProxy_GoogleCloudfunctions2FunctionBuildConfigSourceOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (g *jsiiProxy_GoogleCloudfunctions2FunctionBuildConfigSourceOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (g *jsiiProxy_GoogleCloudfunctions2FunctionBuildConfigSourceOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (g *jsiiProxy_GoogleCloudfunctions2FunctionBuildConfigSourceOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (g *jsiiProxy_GoogleCloudfunctions2FunctionBuildConfigSourceOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (g *jsiiProxy_GoogleCloudfunctions2FunctionBuildConfigSourceOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		g,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleCloudfunctions2FunctionBuildConfigSourceOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GoogleCloudfunctions2FunctionBuildConfigSourceOutputReference) PutRepoSource(value *GoogleCloudfunctions2FunctionBuildConfigSourceRepoSource) {
	if err := g.validatePutRepoSourceParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putRepoSource",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleCloudfunctions2FunctionBuildConfigSourceOutputReference) PutStorageSource(value *GoogleCloudfunctions2FunctionBuildConfigSourceStorageSource) {
	if err := g.validatePutStorageSourceParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putStorageSource",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleCloudfunctions2FunctionBuildConfigSourceOutputReference) ResetRepoSource() {
	_jsii_.InvokeVoid(
		g,
		"resetRepoSource",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleCloudfunctions2FunctionBuildConfigSourceOutputReference) ResetStorageSource() {
	_jsii_.InvokeVoid(
		g,
		"resetStorageSource",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleCloudfunctions2FunctionBuildConfigSourceOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (g *jsiiProxy_GoogleCloudfunctions2FunctionBuildConfigSourceOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

