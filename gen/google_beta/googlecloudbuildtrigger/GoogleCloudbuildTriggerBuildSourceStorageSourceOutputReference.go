package googlecloudbuildtrigger

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/google_beta/jsii"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/sourcegraph/controller-cdktf/gen/google_beta/googlecloudbuildtrigger/internal"
)

type GoogleCloudbuildTriggerBuildSourceStorageSourceOutputReference interface {
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
	Generation() *string
	SetGeneration(val *string)
	GenerationInput() *string
	InternalValue() *GoogleCloudbuildTriggerBuildSourceStorageSource
	SetInternalValue(val *GoogleCloudbuildTriggerBuildSourceStorageSource)
	Object() *string
	SetObject(val *string)
	ObjectInput() *string
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
	ResetGeneration()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for GoogleCloudbuildTriggerBuildSourceStorageSourceOutputReference
type jsiiProxy_GoogleCloudbuildTriggerBuildSourceStorageSourceOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_GoogleCloudbuildTriggerBuildSourceStorageSourceOutputReference) Bucket() *string {
	var returns *string
	_jsii_.Get(
		j,
		"bucket",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCloudbuildTriggerBuildSourceStorageSourceOutputReference) BucketInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"bucketInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCloudbuildTriggerBuildSourceStorageSourceOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCloudbuildTriggerBuildSourceStorageSourceOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCloudbuildTriggerBuildSourceStorageSourceOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCloudbuildTriggerBuildSourceStorageSourceOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCloudbuildTriggerBuildSourceStorageSourceOutputReference) Generation() *string {
	var returns *string
	_jsii_.Get(
		j,
		"generation",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCloudbuildTriggerBuildSourceStorageSourceOutputReference) GenerationInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"generationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCloudbuildTriggerBuildSourceStorageSourceOutputReference) InternalValue() *GoogleCloudbuildTriggerBuildSourceStorageSource {
	var returns *GoogleCloudbuildTriggerBuildSourceStorageSource
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCloudbuildTriggerBuildSourceStorageSourceOutputReference) Object() *string {
	var returns *string
	_jsii_.Get(
		j,
		"object",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCloudbuildTriggerBuildSourceStorageSourceOutputReference) ObjectInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"objectInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCloudbuildTriggerBuildSourceStorageSourceOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCloudbuildTriggerBuildSourceStorageSourceOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewGoogleCloudbuildTriggerBuildSourceStorageSourceOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) GoogleCloudbuildTriggerBuildSourceStorageSourceOutputReference {
	_init_.Initialize()

	if err := validateNewGoogleCloudbuildTriggerBuildSourceStorageSourceOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_GoogleCloudbuildTriggerBuildSourceStorageSourceOutputReference{}

	_jsii_.Create(
		"google-beta.googleCloudbuildTrigger.GoogleCloudbuildTriggerBuildSourceStorageSourceOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewGoogleCloudbuildTriggerBuildSourceStorageSourceOutputReference_Override(g GoogleCloudbuildTriggerBuildSourceStorageSourceOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"google-beta.googleCloudbuildTrigger.GoogleCloudbuildTriggerBuildSourceStorageSourceOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		g,
	)
}

func (j *jsiiProxy_GoogleCloudbuildTriggerBuildSourceStorageSourceOutputReference)SetBucket(val *string) {
	if err := j.validateSetBucketParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"bucket",
		val,
	)
}

func (j *jsiiProxy_GoogleCloudbuildTriggerBuildSourceStorageSourceOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_GoogleCloudbuildTriggerBuildSourceStorageSourceOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_GoogleCloudbuildTriggerBuildSourceStorageSourceOutputReference)SetGeneration(val *string) {
	if err := j.validateSetGenerationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"generation",
		val,
	)
}

func (j *jsiiProxy_GoogleCloudbuildTriggerBuildSourceStorageSourceOutputReference)SetInternalValue(val *GoogleCloudbuildTriggerBuildSourceStorageSource) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_GoogleCloudbuildTriggerBuildSourceStorageSourceOutputReference)SetObject(val *string) {
	if err := j.validateSetObjectParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"object",
		val,
	)
}

func (j *jsiiProxy_GoogleCloudbuildTriggerBuildSourceStorageSourceOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_GoogleCloudbuildTriggerBuildSourceStorageSourceOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (g *jsiiProxy_GoogleCloudbuildTriggerBuildSourceStorageSourceOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleCloudbuildTriggerBuildSourceStorageSourceOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (g *jsiiProxy_GoogleCloudbuildTriggerBuildSourceStorageSourceOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GoogleCloudbuildTriggerBuildSourceStorageSourceOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (g *jsiiProxy_GoogleCloudbuildTriggerBuildSourceStorageSourceOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (g *jsiiProxy_GoogleCloudbuildTriggerBuildSourceStorageSourceOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (g *jsiiProxy_GoogleCloudbuildTriggerBuildSourceStorageSourceOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (g *jsiiProxy_GoogleCloudbuildTriggerBuildSourceStorageSourceOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (g *jsiiProxy_GoogleCloudbuildTriggerBuildSourceStorageSourceOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (g *jsiiProxy_GoogleCloudbuildTriggerBuildSourceStorageSourceOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (g *jsiiProxy_GoogleCloudbuildTriggerBuildSourceStorageSourceOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		g,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleCloudbuildTriggerBuildSourceStorageSourceOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GoogleCloudbuildTriggerBuildSourceStorageSourceOutputReference) ResetGeneration() {
	_jsii_.InvokeVoid(
		g,
		"resetGeneration",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleCloudbuildTriggerBuildSourceStorageSourceOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (g *jsiiProxy_GoogleCloudbuildTriggerBuildSourceStorageSourceOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

