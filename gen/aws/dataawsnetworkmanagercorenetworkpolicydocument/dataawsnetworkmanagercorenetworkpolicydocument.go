package dataawsnetworkmanagercorenetworkpolicydocument

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/aws/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/sourcegraph/controller-cdktf/gen/aws/dataawsnetworkmanagercorenetworkpolicydocument/internal"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/aws/4.54.0/docs/data-sources/networkmanager_core_network_policy_document aws_networkmanager_core_network_policy_document}.
type DataAwsNetworkmanagerCoreNetworkPolicyDocument interface {
	cdktf.TerraformDataSource
	AttachmentPolicies() DataAwsNetworkmanagerCoreNetworkPolicyDocumentAttachmentPoliciesList
	AttachmentPoliciesInput() any
	// Experimental.
	CdktfStack() cdktf.TerraformStack
	// Experimental.
	ConstructNodeMetadata() *map[string]any
	CoreNetworkConfiguration() DataAwsNetworkmanagerCoreNetworkPolicyDocumentCoreNetworkConfigurationList
	CoreNetworkConfigurationInput() any
	// Experimental.
	Count() any
	// Experimental.
	SetCount(val any)
	// Experimental.
	DependsOn() *[]*string
	// Experimental.
	SetDependsOn(val *[]*string)
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
	Json() *string
	// Experimental.
	Lifecycle() *cdktf.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	// The tree node.
	Node() constructs.Node
	// Experimental.
	Provider() cdktf.TerraformProvider
	// Experimental.
	SetProvider(val cdktf.TerraformProvider)
	// Experimental.
	RawOverrides() any
	SegmentActions() DataAwsNetworkmanagerCoreNetworkPolicyDocumentSegmentActionsList
	SegmentActionsInput() any
	Segments() DataAwsNetworkmanagerCoreNetworkPolicyDocumentSegmentsList
	SegmentsInput() any
	// Experimental.
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]any
	// Experimental.
	TerraformResourceType() *string
	Version() *string
	SetVersion(val *string)
	VersionInput() *string
	// Experimental.
	AddOverride(path *string, value any)
	// Experimental.
	GetAnyMapAttribute(terraformAttribute *string) *map[string]any
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
	PutAttachmentPolicies(value any)
	PutCoreNetworkConfiguration(value any)
	PutSegmentActions(value any)
	PutSegments(value any)
	ResetAttachmentPolicies()
	ResetId()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetSegmentActions()
	ResetVersion()
	SynthesizeAttributes() *map[string]any
	SynthesizeHclAttributes() *map[string]any
	// Adds this resource to the terraform JSON output.
	// Experimental.
	ToHclTerraform() any
	// Experimental.
	ToMetadata() any
	// Returns a string representation of this construct.
	ToString() *string
	// Adds this resource to the terraform JSON output.
	// Experimental.
	ToTerraform() any
}

// The jsii proxy struct for DataAwsNetworkmanagerCoreNetworkPolicyDocument
type jsiiProxy_DataAwsNetworkmanagerCoreNetworkPolicyDocument struct {
	internal.Type__cdktfTerraformDataSource
}

func (j *jsiiProxy_DataAwsNetworkmanagerCoreNetworkPolicyDocument) AttachmentPolicies() DataAwsNetworkmanagerCoreNetworkPolicyDocumentAttachmentPoliciesList {
	var returns DataAwsNetworkmanagerCoreNetworkPolicyDocumentAttachmentPoliciesList
	_jsii_.Get(
		j,
		"attachmentPolicies",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsNetworkmanagerCoreNetworkPolicyDocument) AttachmentPoliciesInput() any {
	var returns any
	_jsii_.Get(
		j,
		"attachmentPoliciesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsNetworkmanagerCoreNetworkPolicyDocument) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsNetworkmanagerCoreNetworkPolicyDocument) ConstructNodeMetadata() *map[string]any {
	var returns *map[string]any
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsNetworkmanagerCoreNetworkPolicyDocument) CoreNetworkConfiguration() DataAwsNetworkmanagerCoreNetworkPolicyDocumentCoreNetworkConfigurationList {
	var returns DataAwsNetworkmanagerCoreNetworkPolicyDocumentCoreNetworkConfigurationList
	_jsii_.Get(
		j,
		"coreNetworkConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsNetworkmanagerCoreNetworkPolicyDocument) CoreNetworkConfigurationInput() any {
	var returns any
	_jsii_.Get(
		j,
		"coreNetworkConfigurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsNetworkmanagerCoreNetworkPolicyDocument) Count() any {
	var returns any
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsNetworkmanagerCoreNetworkPolicyDocument) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsNetworkmanagerCoreNetworkPolicyDocument) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsNetworkmanagerCoreNetworkPolicyDocument) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsNetworkmanagerCoreNetworkPolicyDocument) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsNetworkmanagerCoreNetworkPolicyDocument) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsNetworkmanagerCoreNetworkPolicyDocument) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsNetworkmanagerCoreNetworkPolicyDocument) Json() *string {
	var returns *string
	_jsii_.Get(
		j,
		"json",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsNetworkmanagerCoreNetworkPolicyDocument) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsNetworkmanagerCoreNetworkPolicyDocument) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsNetworkmanagerCoreNetworkPolicyDocument) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsNetworkmanagerCoreNetworkPolicyDocument) RawOverrides() any {
	var returns any
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsNetworkmanagerCoreNetworkPolicyDocument) SegmentActions() DataAwsNetworkmanagerCoreNetworkPolicyDocumentSegmentActionsList {
	var returns DataAwsNetworkmanagerCoreNetworkPolicyDocumentSegmentActionsList
	_jsii_.Get(
		j,
		"segmentActions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsNetworkmanagerCoreNetworkPolicyDocument) SegmentActionsInput() any {
	var returns any
	_jsii_.Get(
		j,
		"segmentActionsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsNetworkmanagerCoreNetworkPolicyDocument) Segments() DataAwsNetworkmanagerCoreNetworkPolicyDocumentSegmentsList {
	var returns DataAwsNetworkmanagerCoreNetworkPolicyDocumentSegmentsList
	_jsii_.Get(
		j,
		"segments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsNetworkmanagerCoreNetworkPolicyDocument) SegmentsInput() any {
	var returns any
	_jsii_.Get(
		j,
		"segmentsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsNetworkmanagerCoreNetworkPolicyDocument) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsNetworkmanagerCoreNetworkPolicyDocument) TerraformMetaArguments() *map[string]any {
	var returns *map[string]any
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsNetworkmanagerCoreNetworkPolicyDocument) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsNetworkmanagerCoreNetworkPolicyDocument) Version() *string {
	var returns *string
	_jsii_.Get(
		j,
		"version",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsNetworkmanagerCoreNetworkPolicyDocument) VersionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"versionInput",
		&returns,
	)
	return returns
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/aws/4.54.0/docs/data-sources/networkmanager_core_network_policy_document aws_networkmanager_core_network_policy_document} Data Source.
func NewDataAwsNetworkmanagerCoreNetworkPolicyDocument(scope constructs.Construct, id *string, config *DataAwsNetworkmanagerCoreNetworkPolicyDocumentConfig) DataAwsNetworkmanagerCoreNetworkPolicyDocument {
	_init_.Initialize()

	if err := validateNewDataAwsNetworkmanagerCoreNetworkPolicyDocumentParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_DataAwsNetworkmanagerCoreNetworkPolicyDocument{}

	_jsii_.Create(
		"@cdktf/provider-aws.dataAwsNetworkmanagerCoreNetworkPolicyDocument.DataAwsNetworkmanagerCoreNetworkPolicyDocument",
		[]any{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/aws/4.54.0/docs/data-sources/networkmanager_core_network_policy_document aws_networkmanager_core_network_policy_document} Data Source.
func NewDataAwsNetworkmanagerCoreNetworkPolicyDocument_Override(d DataAwsNetworkmanagerCoreNetworkPolicyDocument, scope constructs.Construct, id *string, config *DataAwsNetworkmanagerCoreNetworkPolicyDocumentConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-aws.dataAwsNetworkmanagerCoreNetworkPolicyDocument.DataAwsNetworkmanagerCoreNetworkPolicyDocument",
		[]any{scope, id, config},
		d,
	)
}

func (j *jsiiProxy_DataAwsNetworkmanagerCoreNetworkPolicyDocument) SetCount(val any) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_DataAwsNetworkmanagerCoreNetworkPolicyDocument) SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_DataAwsNetworkmanagerCoreNetworkPolicyDocument) SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_DataAwsNetworkmanagerCoreNetworkPolicyDocument) SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_DataAwsNetworkmanagerCoreNetworkPolicyDocument) SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_DataAwsNetworkmanagerCoreNetworkPolicyDocument) SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_DataAwsNetworkmanagerCoreNetworkPolicyDocument) SetVersion(val *string) {
	if err := j.validateSetVersionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"version",
		val,
	)
}

// Generates CDKTF code for importing a DataAwsNetworkmanagerCoreNetworkPolicyDocument resource upon running "cdktf plan <stack-name>".
func DataAwsNetworkmanagerCoreNetworkPolicyDocument_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktf.TerraformProvider) cdktf.ImportableResource {
	_init_.Initialize()

	if err := validateDataAwsNetworkmanagerCoreNetworkPolicyDocument_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktf.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktf/provider-aws.dataAwsNetworkmanagerCoreNetworkPolicyDocument.DataAwsNetworkmanagerCoreNetworkPolicyDocument",
		"generateConfigForImport",
		[]any{scope, importToId, importFromId, provider},
		&returns,
	)

	return returns
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
func DataAwsNetworkmanagerCoreNetworkPolicyDocument_IsConstruct(x any) *bool {
	_init_.Initialize()

	if err := validateDataAwsNetworkmanagerCoreNetworkPolicyDocument_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-aws.dataAwsNetworkmanagerCoreNetworkPolicyDocument.DataAwsNetworkmanagerCoreNetworkPolicyDocument",
		"isConstruct",
		[]any{x},
		&returns,
	)

	return returns
}

// Experimental.
func DataAwsNetworkmanagerCoreNetworkPolicyDocument_IsTerraformDataSource(x any) *bool {
	_init_.Initialize()

	if err := validateDataAwsNetworkmanagerCoreNetworkPolicyDocument_IsTerraformDataSourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-aws.dataAwsNetworkmanagerCoreNetworkPolicyDocument.DataAwsNetworkmanagerCoreNetworkPolicyDocument",
		"isTerraformDataSource",
		[]any{x},
		&returns,
	)

	return returns
}

// Experimental.
func DataAwsNetworkmanagerCoreNetworkPolicyDocument_IsTerraformElement(x any) *bool {
	_init_.Initialize()

	if err := validateDataAwsNetworkmanagerCoreNetworkPolicyDocument_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-aws.dataAwsNetworkmanagerCoreNetworkPolicyDocument.DataAwsNetworkmanagerCoreNetworkPolicyDocument",
		"isTerraformElement",
		[]any{x},
		&returns,
	)

	return returns
}

func DataAwsNetworkmanagerCoreNetworkPolicyDocument_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-aws.dataAwsNetworkmanagerCoreNetworkPolicyDocument.DataAwsNetworkmanagerCoreNetworkPolicyDocument",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (d *jsiiProxy_DataAwsNetworkmanagerCoreNetworkPolicyDocument) AddOverride(path *string, value any) {
	if err := d.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"addOverride",
		[]any{path, value},
	)
}

func (d *jsiiProxy_DataAwsNetworkmanagerCoreNetworkPolicyDocument) GetAnyMapAttribute(terraformAttribute *string) *map[string]any {
	if err := d.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]any

	_jsii_.Invoke(
		d,
		"getAnyMapAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsNetworkmanagerCoreNetworkPolicyDocument) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := d.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"getBooleanAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsNetworkmanagerCoreNetworkPolicyDocument) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := d.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		d,
		"getBooleanMapAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsNetworkmanagerCoreNetworkPolicyDocument) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := d.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		d,
		"getListAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsNetworkmanagerCoreNetworkPolicyDocument) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := d.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		d,
		"getNumberAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsNetworkmanagerCoreNetworkPolicyDocument) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := d.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		d,
		"getNumberListAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsNetworkmanagerCoreNetworkPolicyDocument) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := d.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		d,
		"getNumberMapAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsNetworkmanagerCoreNetworkPolicyDocument) GetStringAttribute(terraformAttribute *string) *string {
	if err := d.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		d,
		"getStringAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsNetworkmanagerCoreNetworkPolicyDocument) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := d.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		d,
		"getStringMapAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsNetworkmanagerCoreNetworkPolicyDocument) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := d.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationForAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsNetworkmanagerCoreNetworkPolicyDocument) OverrideLogicalId(newLogicalId *string) {
	if err := d.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"overrideLogicalId",
		[]any{newLogicalId},
	)
}

func (d *jsiiProxy_DataAwsNetworkmanagerCoreNetworkPolicyDocument) PutAttachmentPolicies(value any) {
	if err := d.validatePutAttachmentPoliciesParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putAttachmentPolicies",
		[]any{value},
	)
}

func (d *jsiiProxy_DataAwsNetworkmanagerCoreNetworkPolicyDocument) PutCoreNetworkConfiguration(value any) {
	if err := d.validatePutCoreNetworkConfigurationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putCoreNetworkConfiguration",
		[]any{value},
	)
}

func (d *jsiiProxy_DataAwsNetworkmanagerCoreNetworkPolicyDocument) PutSegmentActions(value any) {
	if err := d.validatePutSegmentActionsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putSegmentActions",
		[]any{value},
	)
}

func (d *jsiiProxy_DataAwsNetworkmanagerCoreNetworkPolicyDocument) PutSegments(value any) {
	if err := d.validatePutSegmentsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putSegments",
		[]any{value},
	)
}

func (d *jsiiProxy_DataAwsNetworkmanagerCoreNetworkPolicyDocument) ResetAttachmentPolicies() {
	_jsii_.InvokeVoid(
		d,
		"resetAttachmentPolicies",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataAwsNetworkmanagerCoreNetworkPolicyDocument) ResetId() {
	_jsii_.InvokeVoid(
		d,
		"resetId",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataAwsNetworkmanagerCoreNetworkPolicyDocument) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		d,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataAwsNetworkmanagerCoreNetworkPolicyDocument) ResetSegmentActions() {
	_jsii_.InvokeVoid(
		d,
		"resetSegmentActions",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataAwsNetworkmanagerCoreNetworkPolicyDocument) ResetVersion() {
	_jsii_.InvokeVoid(
		d,
		"resetVersion",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataAwsNetworkmanagerCoreNetworkPolicyDocument) SynthesizeAttributes() *map[string]any {
	var returns *map[string]any

	_jsii_.Invoke(
		d,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsNetworkmanagerCoreNetworkPolicyDocument) SynthesizeHclAttributes() *map[string]any {
	var returns *map[string]any

	_jsii_.Invoke(
		d,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsNetworkmanagerCoreNetworkPolicyDocument) ToHclTerraform() any {
	var returns any

	_jsii_.Invoke(
		d,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsNetworkmanagerCoreNetworkPolicyDocument) ToMetadata() any {
	var returns any

	_jsii_.Invoke(
		d,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsNetworkmanagerCoreNetworkPolicyDocument) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsNetworkmanagerCoreNetworkPolicyDocument) ToTerraform() any {
	var returns any

	_jsii_.Invoke(
		d,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}
