package googlecontaineranalysisnote

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/googlebeta/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/sourcegraph/controller-cdktf/gen/googlebeta/googlecontaineranalysisnote/internal"
)

// Represents a {@link https://www.terraform.io/docs/providers/google-beta/r/google_container_analysis_note google_container_analysis_note}.
type GoogleContainerAnalysisNote interface {
	cdktf.TerraformResource
	AttestationAuthority() GoogleContainerAnalysisNoteAttestationAuthorityOutputReference
	AttestationAuthorityInput() *GoogleContainerAnalysisNoteAttestationAuthority
	// Experimental.
	CdktfStack() cdktf.TerraformStack
	// Experimental.
	Connection() interface{}
	// Experimental.
	SetConnection(val interface{})
	// Experimental.
	ConstructNodeMetadata() *map[string]interface{}
	// Experimental.
	Count() *float64
	// Experimental.
	SetCount(val *float64)
	CreateTime() *string
	// Experimental.
	DependsOn() *[]*string
	// Experimental.
	SetDependsOn(val *[]*string)
	ExpirationTime() *string
	SetExpirationTime(val *string)
	ExpirationTimeInput() *string
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
	Kind() *string
	// Experimental.
	Lifecycle() *cdktf.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	LongDescription() *string
	SetLongDescription(val *string)
	LongDescriptionInput() *string
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
	RelatedNoteNames() *[]*string
	SetRelatedNoteNames(val *[]*string)
	RelatedNoteNamesInput() *[]*string
	RelatedUrl() GoogleContainerAnalysisNoteRelatedUrlList
	RelatedUrlInput() interface{}
	ShortDescription() *string
	SetShortDescription(val *string)
	ShortDescriptionInput() *string
	// Experimental.
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	Timeouts() GoogleContainerAnalysisNoteTimeoutsOutputReference
	TimeoutsInput() interface{}
	UpdateTime() *string
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
	PutAttestationAuthority(value *GoogleContainerAnalysisNoteAttestationAuthority)
	PutRelatedUrl(value interface{})
	PutTimeouts(value *GoogleContainerAnalysisNoteTimeouts)
	ResetExpirationTime()
	ResetId()
	ResetLongDescription()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetProject()
	ResetRelatedNoteNames()
	ResetRelatedUrl()
	ResetShortDescription()
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

// The jsii proxy struct for GoogleContainerAnalysisNote
type jsiiProxy_GoogleContainerAnalysisNote struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_GoogleContainerAnalysisNote) AttestationAuthority() GoogleContainerAnalysisNoteAttestationAuthorityOutputReference {
	var returns GoogleContainerAnalysisNoteAttestationAuthorityOutputReference
	_jsii_.Get(
		j,
		"attestationAuthority",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContainerAnalysisNote) AttestationAuthorityInput() *GoogleContainerAnalysisNoteAttestationAuthority {
	var returns *GoogleContainerAnalysisNoteAttestationAuthority
	_jsii_.Get(
		j,
		"attestationAuthorityInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContainerAnalysisNote) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContainerAnalysisNote) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContainerAnalysisNote) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContainerAnalysisNote) Count() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContainerAnalysisNote) CreateTime() *string {
	var returns *string
	_jsii_.Get(
		j,
		"createTime",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContainerAnalysisNote) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContainerAnalysisNote) ExpirationTime() *string {
	var returns *string
	_jsii_.Get(
		j,
		"expirationTime",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContainerAnalysisNote) ExpirationTimeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"expirationTimeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContainerAnalysisNote) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContainerAnalysisNote) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContainerAnalysisNote) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContainerAnalysisNote) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContainerAnalysisNote) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContainerAnalysisNote) Kind() *string {
	var returns *string
	_jsii_.Get(
		j,
		"kind",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContainerAnalysisNote) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContainerAnalysisNote) LongDescription() *string {
	var returns *string
	_jsii_.Get(
		j,
		"longDescription",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContainerAnalysisNote) LongDescriptionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"longDescriptionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContainerAnalysisNote) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContainerAnalysisNote) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContainerAnalysisNote) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContainerAnalysisNote) Project() *string {
	var returns *string
	_jsii_.Get(
		j,
		"project",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContainerAnalysisNote) ProjectInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"projectInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContainerAnalysisNote) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContainerAnalysisNote) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContainerAnalysisNote) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContainerAnalysisNote) RelatedNoteNames() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"relatedNoteNames",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContainerAnalysisNote) RelatedNoteNamesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"relatedNoteNamesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContainerAnalysisNote) RelatedUrl() GoogleContainerAnalysisNoteRelatedUrlList {
	var returns GoogleContainerAnalysisNoteRelatedUrlList
	_jsii_.Get(
		j,
		"relatedUrl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContainerAnalysisNote) RelatedUrlInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"relatedUrlInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContainerAnalysisNote) ShortDescription() *string {
	var returns *string
	_jsii_.Get(
		j,
		"shortDescription",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContainerAnalysisNote) ShortDescriptionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"shortDescriptionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContainerAnalysisNote) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContainerAnalysisNote) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContainerAnalysisNote) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContainerAnalysisNote) Timeouts() GoogleContainerAnalysisNoteTimeoutsOutputReference {
	var returns GoogleContainerAnalysisNoteTimeoutsOutputReference
	_jsii_.Get(
		j,
		"timeouts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContainerAnalysisNote) TimeoutsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"timeoutsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContainerAnalysisNote) UpdateTime() *string {
	var returns *string
	_jsii_.Get(
		j,
		"updateTime",
		&returns,
	)
	return returns
}


// Create a new {@link https://www.terraform.io/docs/providers/google-beta/r/google_container_analysis_note google_container_analysis_note} Resource.
func NewGoogleContainerAnalysisNote(scope constructs.Construct, id *string, config *GoogleContainerAnalysisNoteConfig) GoogleContainerAnalysisNote {
	_init_.Initialize()

	if err := validateNewGoogleContainerAnalysisNoteParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_GoogleContainerAnalysisNote{}

	_jsii_.Create(
		"@cdktf/provider-googlebeta.googleContainerAnalysisNote.GoogleContainerAnalysisNote",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://www.terraform.io/docs/providers/google-beta/r/google_container_analysis_note google_container_analysis_note} Resource.
func NewGoogleContainerAnalysisNote_Override(g GoogleContainerAnalysisNote, scope constructs.Construct, id *string, config *GoogleContainerAnalysisNoteConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-googlebeta.googleContainerAnalysisNote.GoogleContainerAnalysisNote",
		[]interface{}{scope, id, config},
		g,
	)
}

func (j *jsiiProxy_GoogleContainerAnalysisNote)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_GoogleContainerAnalysisNote)SetCount(val *float64) {
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_GoogleContainerAnalysisNote)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_GoogleContainerAnalysisNote)SetExpirationTime(val *string) {
	if err := j.validateSetExpirationTimeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"expirationTime",
		val,
	)
}

func (j *jsiiProxy_GoogleContainerAnalysisNote)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_GoogleContainerAnalysisNote)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_GoogleContainerAnalysisNote)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_GoogleContainerAnalysisNote)SetLongDescription(val *string) {
	if err := j.validateSetLongDescriptionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"longDescription",
		val,
	)
}

func (j *jsiiProxy_GoogleContainerAnalysisNote)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_GoogleContainerAnalysisNote)SetProject(val *string) {
	if err := j.validateSetProjectParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"project",
		val,
	)
}

func (j *jsiiProxy_GoogleContainerAnalysisNote)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_GoogleContainerAnalysisNote)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_GoogleContainerAnalysisNote)SetRelatedNoteNames(val *[]*string) {
	if err := j.validateSetRelatedNoteNamesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"relatedNoteNames",
		val,
	)
}

func (j *jsiiProxy_GoogleContainerAnalysisNote)SetShortDescription(val *string) {
	if err := j.validateSetShortDescriptionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"shortDescription",
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
func GoogleContainerAnalysisNote_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateGoogleContainerAnalysisNote_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-googlebeta.googleContainerAnalysisNote.GoogleContainerAnalysisNote",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func GoogleContainerAnalysisNote_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateGoogleContainerAnalysisNote_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-googlebeta.googleContainerAnalysisNote.GoogleContainerAnalysisNote",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func GoogleContainerAnalysisNote_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateGoogleContainerAnalysisNote_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-googlebeta.googleContainerAnalysisNote.GoogleContainerAnalysisNote",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func GoogleContainerAnalysisNote_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-googlebeta.googleContainerAnalysisNote.GoogleContainerAnalysisNote",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (g *jsiiProxy_GoogleContainerAnalysisNote) AddOverride(path *string, value interface{}) {
	if err := g.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (g *jsiiProxy_GoogleContainerAnalysisNote) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (g *jsiiProxy_GoogleContainerAnalysisNote) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GoogleContainerAnalysisNote) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (g *jsiiProxy_GoogleContainerAnalysisNote) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (g *jsiiProxy_GoogleContainerAnalysisNote) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (g *jsiiProxy_GoogleContainerAnalysisNote) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (g *jsiiProxy_GoogleContainerAnalysisNote) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (g *jsiiProxy_GoogleContainerAnalysisNote) GetStringAttribute(terraformAttribute *string) *string {
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

func (g *jsiiProxy_GoogleContainerAnalysisNote) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (g *jsiiProxy_GoogleContainerAnalysisNote) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GoogleContainerAnalysisNote) OverrideLogicalId(newLogicalId *string) {
	if err := g.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (g *jsiiProxy_GoogleContainerAnalysisNote) PutAttestationAuthority(value *GoogleContainerAnalysisNoteAttestationAuthority) {
	if err := g.validatePutAttestationAuthorityParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putAttestationAuthority",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleContainerAnalysisNote) PutRelatedUrl(value interface{}) {
	if err := g.validatePutRelatedUrlParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putRelatedUrl",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleContainerAnalysisNote) PutTimeouts(value *GoogleContainerAnalysisNoteTimeouts) {
	if err := g.validatePutTimeoutsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putTimeouts",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleContainerAnalysisNote) ResetExpirationTime() {
	_jsii_.InvokeVoid(
		g,
		"resetExpirationTime",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleContainerAnalysisNote) ResetId() {
	_jsii_.InvokeVoid(
		g,
		"resetId",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleContainerAnalysisNote) ResetLongDescription() {
	_jsii_.InvokeVoid(
		g,
		"resetLongDescription",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleContainerAnalysisNote) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		g,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleContainerAnalysisNote) ResetProject() {
	_jsii_.InvokeVoid(
		g,
		"resetProject",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleContainerAnalysisNote) ResetRelatedNoteNames() {
	_jsii_.InvokeVoid(
		g,
		"resetRelatedNoteNames",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleContainerAnalysisNote) ResetRelatedUrl() {
	_jsii_.InvokeVoid(
		g,
		"resetRelatedUrl",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleContainerAnalysisNote) ResetShortDescription() {
	_jsii_.InvokeVoid(
		g,
		"resetShortDescription",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleContainerAnalysisNote) ResetTimeouts() {
	_jsii_.InvokeVoid(
		g,
		"resetTimeouts",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleContainerAnalysisNote) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		g,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleContainerAnalysisNote) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		g,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleContainerAnalysisNote) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleContainerAnalysisNote) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		g,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

