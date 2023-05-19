package googlecomputebackendbucket

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/google_beta/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/sourcegraph/controller-cdktf/gen/google_beta/googlecomputebackendbucket/internal"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/google-beta/4.65.2/docs/resources/google_compute_backend_bucket google_compute_backend_bucket}.
type GoogleComputeBackendBucket interface {
	cdktf.TerraformResource
	BucketName() *string
	SetBucketName(val *string)
	BucketNameInput() *string
	// Experimental.
	CdktfStack() cdktf.TerraformStack
	CdnPolicy() GoogleComputeBackendBucketCdnPolicyOutputReference
	CdnPolicyInput() *GoogleComputeBackendBucketCdnPolicy
	CompressionMode() *string
	SetCompressionMode(val *string)
	CompressionModeInput() *string
	// Experimental.
	Connection() interface{}
	// Experimental.
	SetConnection(val interface{})
	// Experimental.
	ConstructNodeMetadata() *map[string]interface{}
	// Experimental.
	Count() interface{}
	// Experimental.
	SetCount(val interface{})
	CreationTimestamp() *string
	CustomResponseHeaders() *[]*string
	SetCustomResponseHeaders(val *[]*string)
	CustomResponseHeadersInput() *[]*string
	// Experimental.
	DependsOn() *[]*string
	// Experimental.
	SetDependsOn(val *[]*string)
	Description() *string
	SetDescription(val *string)
	DescriptionInput() *string
	EdgeSecurityPolicy() *string
	SetEdgeSecurityPolicy(val *string)
	EdgeSecurityPolicyInput() *string
	EnableCdn() interface{}
	SetEnableCdn(val interface{})
	EnableCdnInput() interface{}
	// Experimental.
	ForEach() cdktf.ITerraformIterator
	// Experimental.
	SetForEach(val cdktf.ITerraformIterator)
	// Experimental.
	Fqn() *string
	// Experimental.
	FriendlyUniqueId() *string
	Id() *string
	SetId(val *string)
	IdInput() *string
	// Experimental.
	Lifecycle() *cdktf.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	Name() *string
	SetName(val *string)
	NameInput() *string
	// The tree node.
	Node() constructs.Node
	Project() *string
	SetProject(val *string)
	ProjectInput() *string
	// Experimental.
	Provider() cdktf.TerraformProvider
	// Experimental.
	SetProvider(val cdktf.TerraformProvider)
	// Experimental.
	Provisioners() *[]interface{}
	// Experimental.
	SetProvisioners(val *[]interface{})
	// Experimental.
	RawOverrides() interface{}
	SelfLink() *string
	// Experimental.
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	Timeouts() GoogleComputeBackendBucketTimeoutsOutputReference
	TimeoutsInput() interface{}
	// Experimental.
	AddOverride(path *string, value interface{})
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
	InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable
	// Overrides the auto-generated logical ID with a specific ID.
	// Experimental.
	OverrideLogicalId(newLogicalId *string)
	PutCdnPolicy(value *GoogleComputeBackendBucketCdnPolicy)
	PutTimeouts(value *GoogleComputeBackendBucketTimeouts)
	ResetCdnPolicy()
	ResetCompressionMode()
	ResetCustomResponseHeaders()
	ResetDescription()
	ResetEdgeSecurityPolicy()
	ResetEnableCdn()
	ResetId()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetProject()
	ResetTimeouts()
	SynthesizeAttributes() *map[string]interface{}
	// Experimental.
	ToMetadata() interface{}
	// Returns a string representation of this construct.
	ToString() *string
	// Adds this resource to the terraform JSON output.
	// Experimental.
	ToTerraform() interface{}
}

// The jsii proxy struct for GoogleComputeBackendBucket
type jsiiProxy_GoogleComputeBackendBucket struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_GoogleComputeBackendBucket) BucketName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"bucketName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeBackendBucket) BucketNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"bucketNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeBackendBucket) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeBackendBucket) CdnPolicy() GoogleComputeBackendBucketCdnPolicyOutputReference {
	var returns GoogleComputeBackendBucketCdnPolicyOutputReference
	_jsii_.Get(
		j,
		"cdnPolicy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeBackendBucket) CdnPolicyInput() *GoogleComputeBackendBucketCdnPolicy {
	var returns *GoogleComputeBackendBucketCdnPolicy
	_jsii_.Get(
		j,
		"cdnPolicyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeBackendBucket) CompressionMode() *string {
	var returns *string
	_jsii_.Get(
		j,
		"compressionMode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeBackendBucket) CompressionModeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"compressionModeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeBackendBucket) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeBackendBucket) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeBackendBucket) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeBackendBucket) CreationTimestamp() *string {
	var returns *string
	_jsii_.Get(
		j,
		"creationTimestamp",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeBackendBucket) CustomResponseHeaders() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"customResponseHeaders",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeBackendBucket) CustomResponseHeadersInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"customResponseHeadersInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeBackendBucket) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeBackendBucket) Description() *string {
	var returns *string
	_jsii_.Get(
		j,
		"description",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeBackendBucket) DescriptionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"descriptionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeBackendBucket) EdgeSecurityPolicy() *string {
	var returns *string
	_jsii_.Get(
		j,
		"edgeSecurityPolicy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeBackendBucket) EdgeSecurityPolicyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"edgeSecurityPolicyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeBackendBucket) EnableCdn() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableCdn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeBackendBucket) EnableCdnInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableCdnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeBackendBucket) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeBackendBucket) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeBackendBucket) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeBackendBucket) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeBackendBucket) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeBackendBucket) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeBackendBucket) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeBackendBucket) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeBackendBucket) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeBackendBucket) Project() *string {
	var returns *string
	_jsii_.Get(
		j,
		"project",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeBackendBucket) ProjectInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"projectInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeBackendBucket) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeBackendBucket) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeBackendBucket) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeBackendBucket) SelfLink() *string {
	var returns *string
	_jsii_.Get(
		j,
		"selfLink",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeBackendBucket) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeBackendBucket) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeBackendBucket) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeBackendBucket) Timeouts() GoogleComputeBackendBucketTimeoutsOutputReference {
	var returns GoogleComputeBackendBucketTimeoutsOutputReference
	_jsii_.Get(
		j,
		"timeouts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeBackendBucket) TimeoutsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"timeoutsInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/hashicorp/google-beta/4.65.2/docs/resources/google_compute_backend_bucket google_compute_backend_bucket} Resource.
func NewGoogleComputeBackendBucket(scope constructs.Construct, id *string, config *GoogleComputeBackendBucketConfig) GoogleComputeBackendBucket {
	_init_.Initialize()

	if err := validateNewGoogleComputeBackendBucketParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_GoogleComputeBackendBucket{}

	_jsii_.Create(
		"google-beta.googleComputeBackendBucket.GoogleComputeBackendBucket",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/google-beta/4.65.2/docs/resources/google_compute_backend_bucket google_compute_backend_bucket} Resource.
func NewGoogleComputeBackendBucket_Override(g GoogleComputeBackendBucket, scope constructs.Construct, id *string, config *GoogleComputeBackendBucketConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"google-beta.googleComputeBackendBucket.GoogleComputeBackendBucket",
		[]interface{}{scope, id, config},
		g,
	)
}

func (j *jsiiProxy_GoogleComputeBackendBucket)SetBucketName(val *string) {
	if err := j.validateSetBucketNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"bucketName",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeBackendBucket)SetCompressionMode(val *string) {
	if err := j.validateSetCompressionModeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"compressionMode",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeBackendBucket)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeBackendBucket)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeBackendBucket)SetCustomResponseHeaders(val *[]*string) {
	if err := j.validateSetCustomResponseHeadersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"customResponseHeaders",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeBackendBucket)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeBackendBucket)SetDescription(val *string) {
	if err := j.validateSetDescriptionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"description",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeBackendBucket)SetEdgeSecurityPolicy(val *string) {
	if err := j.validateSetEdgeSecurityPolicyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"edgeSecurityPolicy",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeBackendBucket)SetEnableCdn(val interface{}) {
	if err := j.validateSetEnableCdnParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"enableCdn",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeBackendBucket)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeBackendBucket)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeBackendBucket)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeBackendBucket)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeBackendBucket)SetProject(val *string) {
	if err := j.validateSetProjectParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"project",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeBackendBucket)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeBackendBucket)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

// Checks if `x` is a construct.
//
// Use this method instead of `instanceof` to properly detect `Construct`
// instances, even when the construct library is symlinked.
//
// Explanation: in JavaScript, multiple copies of the `constructs` library on
// disk are seen as independent, completely different libraries. As a
// consequence, the class `Construct` in each copy of the `constructs` library
// is seen as a different class, and an instance of one class will not test as
// `instanceof` the other class. `npm install` will not create installations
// like this, but users may manually symlink construct libraries together or
// use a monorepo tool: in those cases, multiple copies of the `constructs`
// library can be accidentally installed, and `instanceof` will behave
// unpredictably. It is safest to avoid using `instanceof`, and using
// this type-testing method instead.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
func GoogleComputeBackendBucket_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateGoogleComputeBackendBucket_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"google-beta.googleComputeBackendBucket.GoogleComputeBackendBucket",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func GoogleComputeBackendBucket_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateGoogleComputeBackendBucket_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"google-beta.googleComputeBackendBucket.GoogleComputeBackendBucket",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func GoogleComputeBackendBucket_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateGoogleComputeBackendBucket_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"google-beta.googleComputeBackendBucket.GoogleComputeBackendBucket",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func GoogleComputeBackendBucket_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"google-beta.googleComputeBackendBucket.GoogleComputeBackendBucket",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (g *jsiiProxy_GoogleComputeBackendBucket) AddOverride(path *string, value interface{}) {
	if err := g.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (g *jsiiProxy_GoogleComputeBackendBucket) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (g *jsiiProxy_GoogleComputeBackendBucket) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GoogleComputeBackendBucket) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (g *jsiiProxy_GoogleComputeBackendBucket) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (g *jsiiProxy_GoogleComputeBackendBucket) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (g *jsiiProxy_GoogleComputeBackendBucket) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (g *jsiiProxy_GoogleComputeBackendBucket) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (g *jsiiProxy_GoogleComputeBackendBucket) GetStringAttribute(terraformAttribute *string) *string {
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

func (g *jsiiProxy_GoogleComputeBackendBucket) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (g *jsiiProxy_GoogleComputeBackendBucket) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := g.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		g,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleComputeBackendBucket) OverrideLogicalId(newLogicalId *string) {
	if err := g.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (g *jsiiProxy_GoogleComputeBackendBucket) PutCdnPolicy(value *GoogleComputeBackendBucketCdnPolicy) {
	if err := g.validatePutCdnPolicyParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putCdnPolicy",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleComputeBackendBucket) PutTimeouts(value *GoogleComputeBackendBucketTimeouts) {
	if err := g.validatePutTimeoutsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putTimeouts",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleComputeBackendBucket) ResetCdnPolicy() {
	_jsii_.InvokeVoid(
		g,
		"resetCdnPolicy",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleComputeBackendBucket) ResetCompressionMode() {
	_jsii_.InvokeVoid(
		g,
		"resetCompressionMode",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleComputeBackendBucket) ResetCustomResponseHeaders() {
	_jsii_.InvokeVoid(
		g,
		"resetCustomResponseHeaders",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleComputeBackendBucket) ResetDescription() {
	_jsii_.InvokeVoid(
		g,
		"resetDescription",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleComputeBackendBucket) ResetEdgeSecurityPolicy() {
	_jsii_.InvokeVoid(
		g,
		"resetEdgeSecurityPolicy",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleComputeBackendBucket) ResetEnableCdn() {
	_jsii_.InvokeVoid(
		g,
		"resetEnableCdn",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleComputeBackendBucket) ResetId() {
	_jsii_.InvokeVoid(
		g,
		"resetId",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleComputeBackendBucket) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		g,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleComputeBackendBucket) ResetProject() {
	_jsii_.InvokeVoid(
		g,
		"resetProject",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleComputeBackendBucket) ResetTimeouts() {
	_jsii_.InvokeVoid(
		g,
		"resetTimeouts",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleComputeBackendBucket) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		g,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleComputeBackendBucket) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		g,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleComputeBackendBucket) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleComputeBackendBucket) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		g,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

