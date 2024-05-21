package googleartifactregistryrepository

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/google_beta/jsii"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/sourcegraph/controller-cdktf/gen/google_beta/googleartifactregistryrepository/internal"
)

type GoogleArtifactRegistryRepositoryRemoteRepositoryConfigDockerRepositoryOutputReference interface {
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
	CustomRepository() GoogleArtifactRegistryRepositoryRemoteRepositoryConfigDockerRepositoryCustomRepositoryOutputReference
	CustomRepositoryInput() *GoogleArtifactRegistryRepositoryRemoteRepositoryConfigDockerRepositoryCustomRepository
	// Experimental.
	Fqn() *string
	InternalValue() *GoogleArtifactRegistryRepositoryRemoteRepositoryConfigDockerRepository
	SetInternalValue(val *GoogleArtifactRegistryRepositoryRemoteRepositoryConfigDockerRepository)
	PublicRepository() *string
	SetPublicRepository(val *string)
	PublicRepositoryInput() *string
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
	PutCustomRepository(value *GoogleArtifactRegistryRepositoryRemoteRepositoryConfigDockerRepositoryCustomRepository)
	ResetCustomRepository()
	ResetPublicRepository()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for GoogleArtifactRegistryRepositoryRemoteRepositoryConfigDockerRepositoryOutputReference
type jsiiProxy_GoogleArtifactRegistryRepositoryRemoteRepositoryConfigDockerRepositoryOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_GoogleArtifactRegistryRepositoryRemoteRepositoryConfigDockerRepositoryOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleArtifactRegistryRepositoryRemoteRepositoryConfigDockerRepositoryOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleArtifactRegistryRepositoryRemoteRepositoryConfigDockerRepositoryOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleArtifactRegistryRepositoryRemoteRepositoryConfigDockerRepositoryOutputReference) CustomRepository() GoogleArtifactRegistryRepositoryRemoteRepositoryConfigDockerRepositoryCustomRepositoryOutputReference {
	var returns GoogleArtifactRegistryRepositoryRemoteRepositoryConfigDockerRepositoryCustomRepositoryOutputReference
	_jsii_.Get(
		j,
		"customRepository",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleArtifactRegistryRepositoryRemoteRepositoryConfigDockerRepositoryOutputReference) CustomRepositoryInput() *GoogleArtifactRegistryRepositoryRemoteRepositoryConfigDockerRepositoryCustomRepository {
	var returns *GoogleArtifactRegistryRepositoryRemoteRepositoryConfigDockerRepositoryCustomRepository
	_jsii_.Get(
		j,
		"customRepositoryInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleArtifactRegistryRepositoryRemoteRepositoryConfigDockerRepositoryOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleArtifactRegistryRepositoryRemoteRepositoryConfigDockerRepositoryOutputReference) InternalValue() *GoogleArtifactRegistryRepositoryRemoteRepositoryConfigDockerRepository {
	var returns *GoogleArtifactRegistryRepositoryRemoteRepositoryConfigDockerRepository
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleArtifactRegistryRepositoryRemoteRepositoryConfigDockerRepositoryOutputReference) PublicRepository() *string {
	var returns *string
	_jsii_.Get(
		j,
		"publicRepository",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleArtifactRegistryRepositoryRemoteRepositoryConfigDockerRepositoryOutputReference) PublicRepositoryInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"publicRepositoryInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleArtifactRegistryRepositoryRemoteRepositoryConfigDockerRepositoryOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleArtifactRegistryRepositoryRemoteRepositoryConfigDockerRepositoryOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewGoogleArtifactRegistryRepositoryRemoteRepositoryConfigDockerRepositoryOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) GoogleArtifactRegistryRepositoryRemoteRepositoryConfigDockerRepositoryOutputReference {
	_init_.Initialize()

	if err := validateNewGoogleArtifactRegistryRepositoryRemoteRepositoryConfigDockerRepositoryOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_GoogleArtifactRegistryRepositoryRemoteRepositoryConfigDockerRepositoryOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-google-beta.googleArtifactRegistryRepository.GoogleArtifactRegistryRepositoryRemoteRepositoryConfigDockerRepositoryOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewGoogleArtifactRegistryRepositoryRemoteRepositoryConfigDockerRepositoryOutputReference_Override(g GoogleArtifactRegistryRepositoryRemoteRepositoryConfigDockerRepositoryOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google-beta.googleArtifactRegistryRepository.GoogleArtifactRegistryRepositoryRemoteRepositoryConfigDockerRepositoryOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		g,
	)
}

func (j *jsiiProxy_GoogleArtifactRegistryRepositoryRemoteRepositoryConfigDockerRepositoryOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_GoogleArtifactRegistryRepositoryRemoteRepositoryConfigDockerRepositoryOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_GoogleArtifactRegistryRepositoryRemoteRepositoryConfigDockerRepositoryOutputReference)SetInternalValue(val *GoogleArtifactRegistryRepositoryRemoteRepositoryConfigDockerRepository) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_GoogleArtifactRegistryRepositoryRemoteRepositoryConfigDockerRepositoryOutputReference)SetPublicRepository(val *string) {
	if err := j.validateSetPublicRepositoryParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"publicRepository",
		val,
	)
}

func (j *jsiiProxy_GoogleArtifactRegistryRepositoryRemoteRepositoryConfigDockerRepositoryOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_GoogleArtifactRegistryRepositoryRemoteRepositoryConfigDockerRepositoryOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (g *jsiiProxy_GoogleArtifactRegistryRepositoryRemoteRepositoryConfigDockerRepositoryOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleArtifactRegistryRepositoryRemoteRepositoryConfigDockerRepositoryOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (g *jsiiProxy_GoogleArtifactRegistryRepositoryRemoteRepositoryConfigDockerRepositoryOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GoogleArtifactRegistryRepositoryRemoteRepositoryConfigDockerRepositoryOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (g *jsiiProxy_GoogleArtifactRegistryRepositoryRemoteRepositoryConfigDockerRepositoryOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (g *jsiiProxy_GoogleArtifactRegistryRepositoryRemoteRepositoryConfigDockerRepositoryOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (g *jsiiProxy_GoogleArtifactRegistryRepositoryRemoteRepositoryConfigDockerRepositoryOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (g *jsiiProxy_GoogleArtifactRegistryRepositoryRemoteRepositoryConfigDockerRepositoryOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (g *jsiiProxy_GoogleArtifactRegistryRepositoryRemoteRepositoryConfigDockerRepositoryOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (g *jsiiProxy_GoogleArtifactRegistryRepositoryRemoteRepositoryConfigDockerRepositoryOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (g *jsiiProxy_GoogleArtifactRegistryRepositoryRemoteRepositoryConfigDockerRepositoryOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		g,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleArtifactRegistryRepositoryRemoteRepositoryConfigDockerRepositoryOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GoogleArtifactRegistryRepositoryRemoteRepositoryConfigDockerRepositoryOutputReference) PutCustomRepository(value *GoogleArtifactRegistryRepositoryRemoteRepositoryConfigDockerRepositoryCustomRepository) {
	if err := g.validatePutCustomRepositoryParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putCustomRepository",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleArtifactRegistryRepositoryRemoteRepositoryConfigDockerRepositoryOutputReference) ResetCustomRepository() {
	_jsii_.InvokeVoid(
		g,
		"resetCustomRepository",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleArtifactRegistryRepositoryRemoteRepositoryConfigDockerRepositoryOutputReference) ResetPublicRepository() {
	_jsii_.InvokeVoid(
		g,
		"resetPublicRepository",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleArtifactRegistryRepositoryRemoteRepositoryConfigDockerRepositoryOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (g *jsiiProxy_GoogleArtifactRegistryRepositoryRemoteRepositoryConfigDockerRepositoryOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

