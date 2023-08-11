package googlecloudbuildtrigger

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/google_beta/jsii"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/sourcegraph/controller-cdktf/gen/google_beta/googlecloudbuildtrigger/internal"
)

type GoogleCloudbuildTriggerBuildOutputReference interface {
	cdktf.ComplexObject
	Artifacts() GoogleCloudbuildTriggerBuildArtifactsOutputReference
	ArtifactsInput() *GoogleCloudbuildTriggerBuildArtifacts
	AvailableSecrets() GoogleCloudbuildTriggerBuildAvailableSecretsOutputReference
	AvailableSecretsInput() *GoogleCloudbuildTriggerBuildAvailableSecrets
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
	Images() *[]*string
	SetImages(val *[]*string)
	ImagesInput() *[]*string
	InternalValue() *GoogleCloudbuildTriggerBuild
	SetInternalValue(val *GoogleCloudbuildTriggerBuild)
	LogsBucket() *string
	SetLogsBucket(val *string)
	LogsBucketInput() *string
	Options() GoogleCloudbuildTriggerBuildOptionsOutputReference
	OptionsInput() *GoogleCloudbuildTriggerBuildOptions
	QueueTtl() *string
	SetQueueTtl(val *string)
	QueueTtlInput() *string
	Secret() GoogleCloudbuildTriggerBuildSecretList
	SecretInput() interface{}
	Source() GoogleCloudbuildTriggerBuildSourceOutputReference
	SourceInput() *GoogleCloudbuildTriggerBuildSource
	Step() GoogleCloudbuildTriggerBuildStepList
	StepInput() interface{}
	Substitutions() *map[string]*string
	SetSubstitutions(val *map[string]*string)
	SubstitutionsInput() *map[string]*string
	Tags() *[]*string
	SetTags(val *[]*string)
	TagsInput() *[]*string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	Timeout() *string
	SetTimeout(val *string)
	TimeoutInput() *string
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
	PutArtifacts(value *GoogleCloudbuildTriggerBuildArtifacts)
	PutAvailableSecrets(value *GoogleCloudbuildTriggerBuildAvailableSecrets)
	PutOptions(value *GoogleCloudbuildTriggerBuildOptions)
	PutSecret(value interface{})
	PutSource(value *GoogleCloudbuildTriggerBuildSource)
	PutStep(value interface{})
	ResetArtifacts()
	ResetAvailableSecrets()
	ResetImages()
	ResetLogsBucket()
	ResetOptions()
	ResetQueueTtl()
	ResetSecret()
	ResetSource()
	ResetSubstitutions()
	ResetTags()
	ResetTimeout()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for GoogleCloudbuildTriggerBuildOutputReference
type jsiiProxy_GoogleCloudbuildTriggerBuildOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_GoogleCloudbuildTriggerBuildOutputReference) Artifacts() GoogleCloudbuildTriggerBuildArtifactsOutputReference {
	var returns GoogleCloudbuildTriggerBuildArtifactsOutputReference
	_jsii_.Get(
		j,
		"artifacts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCloudbuildTriggerBuildOutputReference) ArtifactsInput() *GoogleCloudbuildTriggerBuildArtifacts {
	var returns *GoogleCloudbuildTriggerBuildArtifacts
	_jsii_.Get(
		j,
		"artifactsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCloudbuildTriggerBuildOutputReference) AvailableSecrets() GoogleCloudbuildTriggerBuildAvailableSecretsOutputReference {
	var returns GoogleCloudbuildTriggerBuildAvailableSecretsOutputReference
	_jsii_.Get(
		j,
		"availableSecrets",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCloudbuildTriggerBuildOutputReference) AvailableSecretsInput() *GoogleCloudbuildTriggerBuildAvailableSecrets {
	var returns *GoogleCloudbuildTriggerBuildAvailableSecrets
	_jsii_.Get(
		j,
		"availableSecretsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCloudbuildTriggerBuildOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCloudbuildTriggerBuildOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCloudbuildTriggerBuildOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCloudbuildTriggerBuildOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCloudbuildTriggerBuildOutputReference) Images() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"images",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCloudbuildTriggerBuildOutputReference) ImagesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"imagesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCloudbuildTriggerBuildOutputReference) InternalValue() *GoogleCloudbuildTriggerBuild {
	var returns *GoogleCloudbuildTriggerBuild
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCloudbuildTriggerBuildOutputReference) LogsBucket() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logsBucket",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCloudbuildTriggerBuildOutputReference) LogsBucketInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logsBucketInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCloudbuildTriggerBuildOutputReference) Options() GoogleCloudbuildTriggerBuildOptionsOutputReference {
	var returns GoogleCloudbuildTriggerBuildOptionsOutputReference
	_jsii_.Get(
		j,
		"options",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCloudbuildTriggerBuildOutputReference) OptionsInput() *GoogleCloudbuildTriggerBuildOptions {
	var returns *GoogleCloudbuildTriggerBuildOptions
	_jsii_.Get(
		j,
		"optionsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCloudbuildTriggerBuildOutputReference) QueueTtl() *string {
	var returns *string
	_jsii_.Get(
		j,
		"queueTtl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCloudbuildTriggerBuildOutputReference) QueueTtlInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"queueTtlInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCloudbuildTriggerBuildOutputReference) Secret() GoogleCloudbuildTriggerBuildSecretList {
	var returns GoogleCloudbuildTriggerBuildSecretList
	_jsii_.Get(
		j,
		"secret",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCloudbuildTriggerBuildOutputReference) SecretInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"secretInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCloudbuildTriggerBuildOutputReference) Source() GoogleCloudbuildTriggerBuildSourceOutputReference {
	var returns GoogleCloudbuildTriggerBuildSourceOutputReference
	_jsii_.Get(
		j,
		"source",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCloudbuildTriggerBuildOutputReference) SourceInput() *GoogleCloudbuildTriggerBuildSource {
	var returns *GoogleCloudbuildTriggerBuildSource
	_jsii_.Get(
		j,
		"sourceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCloudbuildTriggerBuildOutputReference) Step() GoogleCloudbuildTriggerBuildStepList {
	var returns GoogleCloudbuildTriggerBuildStepList
	_jsii_.Get(
		j,
		"step",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCloudbuildTriggerBuildOutputReference) StepInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"stepInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCloudbuildTriggerBuildOutputReference) Substitutions() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"substitutions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCloudbuildTriggerBuildOutputReference) SubstitutionsInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"substitutionsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCloudbuildTriggerBuildOutputReference) Tags() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCloudbuildTriggerBuildOutputReference) TagsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"tagsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCloudbuildTriggerBuildOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCloudbuildTriggerBuildOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCloudbuildTriggerBuildOutputReference) Timeout() *string {
	var returns *string
	_jsii_.Get(
		j,
		"timeout",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCloudbuildTriggerBuildOutputReference) TimeoutInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"timeoutInput",
		&returns,
	)
	return returns
}


func NewGoogleCloudbuildTriggerBuildOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) GoogleCloudbuildTriggerBuildOutputReference {
	_init_.Initialize()

	if err := validateNewGoogleCloudbuildTriggerBuildOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_GoogleCloudbuildTriggerBuildOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-google-beta.googleCloudbuildTrigger.GoogleCloudbuildTriggerBuildOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewGoogleCloudbuildTriggerBuildOutputReference_Override(g GoogleCloudbuildTriggerBuildOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google-beta.googleCloudbuildTrigger.GoogleCloudbuildTriggerBuildOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		g,
	)
}

func (j *jsiiProxy_GoogleCloudbuildTriggerBuildOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_GoogleCloudbuildTriggerBuildOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_GoogleCloudbuildTriggerBuildOutputReference)SetImages(val *[]*string) {
	if err := j.validateSetImagesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"images",
		val,
	)
}

func (j *jsiiProxy_GoogleCloudbuildTriggerBuildOutputReference)SetInternalValue(val *GoogleCloudbuildTriggerBuild) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_GoogleCloudbuildTriggerBuildOutputReference)SetLogsBucket(val *string) {
	if err := j.validateSetLogsBucketParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"logsBucket",
		val,
	)
}

func (j *jsiiProxy_GoogleCloudbuildTriggerBuildOutputReference)SetQueueTtl(val *string) {
	if err := j.validateSetQueueTtlParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"queueTtl",
		val,
	)
}

func (j *jsiiProxy_GoogleCloudbuildTriggerBuildOutputReference)SetSubstitutions(val *map[string]*string) {
	if err := j.validateSetSubstitutionsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"substitutions",
		val,
	)
}

func (j *jsiiProxy_GoogleCloudbuildTriggerBuildOutputReference)SetTags(val *[]*string) {
	if err := j.validateSetTagsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tags",
		val,
	)
}

func (j *jsiiProxy_GoogleCloudbuildTriggerBuildOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_GoogleCloudbuildTriggerBuildOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_GoogleCloudbuildTriggerBuildOutputReference)SetTimeout(val *string) {
	if err := j.validateSetTimeoutParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"timeout",
		val,
	)
}

func (g *jsiiProxy_GoogleCloudbuildTriggerBuildOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleCloudbuildTriggerBuildOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (g *jsiiProxy_GoogleCloudbuildTriggerBuildOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GoogleCloudbuildTriggerBuildOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (g *jsiiProxy_GoogleCloudbuildTriggerBuildOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (g *jsiiProxy_GoogleCloudbuildTriggerBuildOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (g *jsiiProxy_GoogleCloudbuildTriggerBuildOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (g *jsiiProxy_GoogleCloudbuildTriggerBuildOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (g *jsiiProxy_GoogleCloudbuildTriggerBuildOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (g *jsiiProxy_GoogleCloudbuildTriggerBuildOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (g *jsiiProxy_GoogleCloudbuildTriggerBuildOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		g,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleCloudbuildTriggerBuildOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GoogleCloudbuildTriggerBuildOutputReference) PutArtifacts(value *GoogleCloudbuildTriggerBuildArtifacts) {
	if err := g.validatePutArtifactsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putArtifacts",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleCloudbuildTriggerBuildOutputReference) PutAvailableSecrets(value *GoogleCloudbuildTriggerBuildAvailableSecrets) {
	if err := g.validatePutAvailableSecretsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putAvailableSecrets",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleCloudbuildTriggerBuildOutputReference) PutOptions(value *GoogleCloudbuildTriggerBuildOptions) {
	if err := g.validatePutOptionsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putOptions",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleCloudbuildTriggerBuildOutputReference) PutSecret(value interface{}) {
	if err := g.validatePutSecretParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putSecret",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleCloudbuildTriggerBuildOutputReference) PutSource(value *GoogleCloudbuildTriggerBuildSource) {
	if err := g.validatePutSourceParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putSource",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleCloudbuildTriggerBuildOutputReference) PutStep(value interface{}) {
	if err := g.validatePutStepParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putStep",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleCloudbuildTriggerBuildOutputReference) ResetArtifacts() {
	_jsii_.InvokeVoid(
		g,
		"resetArtifacts",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleCloudbuildTriggerBuildOutputReference) ResetAvailableSecrets() {
	_jsii_.InvokeVoid(
		g,
		"resetAvailableSecrets",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleCloudbuildTriggerBuildOutputReference) ResetImages() {
	_jsii_.InvokeVoid(
		g,
		"resetImages",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleCloudbuildTriggerBuildOutputReference) ResetLogsBucket() {
	_jsii_.InvokeVoid(
		g,
		"resetLogsBucket",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleCloudbuildTriggerBuildOutputReference) ResetOptions() {
	_jsii_.InvokeVoid(
		g,
		"resetOptions",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleCloudbuildTriggerBuildOutputReference) ResetQueueTtl() {
	_jsii_.InvokeVoid(
		g,
		"resetQueueTtl",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleCloudbuildTriggerBuildOutputReference) ResetSecret() {
	_jsii_.InvokeVoid(
		g,
		"resetSecret",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleCloudbuildTriggerBuildOutputReference) ResetSource() {
	_jsii_.InvokeVoid(
		g,
		"resetSource",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleCloudbuildTriggerBuildOutputReference) ResetSubstitutions() {
	_jsii_.InvokeVoid(
		g,
		"resetSubstitutions",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleCloudbuildTriggerBuildOutputReference) ResetTags() {
	_jsii_.InvokeVoid(
		g,
		"resetTags",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleCloudbuildTriggerBuildOutputReference) ResetTimeout() {
	_jsii_.InvokeVoid(
		g,
		"resetTimeout",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleCloudbuildTriggerBuildOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (g *jsiiProxy_GoogleCloudbuildTriggerBuildOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

