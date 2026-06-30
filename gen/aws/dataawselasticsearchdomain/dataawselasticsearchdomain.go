package dataawselasticsearchdomain

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/aws/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/sourcegraph/controller-cdktf/gen/aws/dataawselasticsearchdomain/internal"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/aws/4.54.0/docs/data-sources/elasticsearch_domain aws_elasticsearch_domain}.
type DataAwsElasticsearchDomain interface {
	cdktf.TerraformDataSource
	AccessPolicies() *string
	AdvancedOptions() cdktf.StringMap
	AdvancedSecurityOptions() DataAwsElasticsearchDomainAdvancedSecurityOptionsList
	Arn() *string
	AutoTuneOptions() DataAwsElasticsearchDomainAutoTuneOptionsList
	// Experimental.
	CdktfStack() cdktf.TerraformStack
	ClusterConfig() DataAwsElasticsearchDomainClusterConfigList
	CognitoOptions() DataAwsElasticsearchDomainCognitoOptionsList
	// Experimental.
	ConstructNodeMetadata() *map[string]any
	// Experimental.
	Count() any
	// Experimental.
	SetCount(val any)
	Created() cdktf.IResolvable
	Deleted() cdktf.IResolvable
	// Experimental.
	DependsOn() *[]*string
	// Experimental.
	SetDependsOn(val *[]*string)
	DomainId() *string
	DomainName() *string
	SetDomainName(val *string)
	DomainNameInput() *string
	EbsOptions() DataAwsElasticsearchDomainEbsOptionsList
	ElasticsearchVersion() *string
	EncryptionAtRest() DataAwsElasticsearchDomainEncryptionAtRestList
	Endpoint() *string
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
	KibanaEndpoint() *string
	// Experimental.
	Lifecycle() *cdktf.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	LogPublishingOptions() DataAwsElasticsearchDomainLogPublishingOptionsList
	// The tree node.
	Node() constructs.Node
	NodeToNodeEncryption() DataAwsElasticsearchDomainNodeToNodeEncryptionList
	Processing() cdktf.IResolvable
	// Experimental.
	Provider() cdktf.TerraformProvider
	// Experimental.
	SetProvider(val cdktf.TerraformProvider)
	// Experimental.
	RawOverrides() any
	SnapshotOptions() DataAwsElasticsearchDomainSnapshotOptionsList
	Tags() *map[string]*string
	SetTags(val *map[string]*string)
	TagsInput() *map[string]*string
	// Experimental.
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]any
	// Experimental.
	TerraformResourceType() *string
	VpcOptions() DataAwsElasticsearchDomainVpcOptionsList
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
	ResetId()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetTags()
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

// The jsii proxy struct for DataAwsElasticsearchDomain
type jsiiProxy_DataAwsElasticsearchDomain struct {
	internal.Type__cdktfTerraformDataSource
}

func (j *jsiiProxy_DataAwsElasticsearchDomain) AccessPolicies() *string {
	var returns *string
	_jsii_.Get(
		j,
		"accessPolicies",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsElasticsearchDomain) AdvancedOptions() cdktf.StringMap {
	var returns cdktf.StringMap
	_jsii_.Get(
		j,
		"advancedOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsElasticsearchDomain) AdvancedSecurityOptions() DataAwsElasticsearchDomainAdvancedSecurityOptionsList {
	var returns DataAwsElasticsearchDomainAdvancedSecurityOptionsList
	_jsii_.Get(
		j,
		"advancedSecurityOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsElasticsearchDomain) Arn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"arn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsElasticsearchDomain) AutoTuneOptions() DataAwsElasticsearchDomainAutoTuneOptionsList {
	var returns DataAwsElasticsearchDomainAutoTuneOptionsList
	_jsii_.Get(
		j,
		"autoTuneOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsElasticsearchDomain) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsElasticsearchDomain) ClusterConfig() DataAwsElasticsearchDomainClusterConfigList {
	var returns DataAwsElasticsearchDomainClusterConfigList
	_jsii_.Get(
		j,
		"clusterConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsElasticsearchDomain) CognitoOptions() DataAwsElasticsearchDomainCognitoOptionsList {
	var returns DataAwsElasticsearchDomainCognitoOptionsList
	_jsii_.Get(
		j,
		"cognitoOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsElasticsearchDomain) ConstructNodeMetadata() *map[string]any {
	var returns *map[string]any
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsElasticsearchDomain) Count() any {
	var returns any
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsElasticsearchDomain) Created() cdktf.IResolvable {
	var returns cdktf.IResolvable
	_jsii_.Get(
		j,
		"created",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsElasticsearchDomain) Deleted() cdktf.IResolvable {
	var returns cdktf.IResolvable
	_jsii_.Get(
		j,
		"deleted",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsElasticsearchDomain) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsElasticsearchDomain) DomainId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"domainId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsElasticsearchDomain) DomainName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"domainName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsElasticsearchDomain) DomainNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"domainNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsElasticsearchDomain) EbsOptions() DataAwsElasticsearchDomainEbsOptionsList {
	var returns DataAwsElasticsearchDomainEbsOptionsList
	_jsii_.Get(
		j,
		"ebsOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsElasticsearchDomain) ElasticsearchVersion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"elasticsearchVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsElasticsearchDomain) EncryptionAtRest() DataAwsElasticsearchDomainEncryptionAtRestList {
	var returns DataAwsElasticsearchDomainEncryptionAtRestList
	_jsii_.Get(
		j,
		"encryptionAtRest",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsElasticsearchDomain) Endpoint() *string {
	var returns *string
	_jsii_.Get(
		j,
		"endpoint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsElasticsearchDomain) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsElasticsearchDomain) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsElasticsearchDomain) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsElasticsearchDomain) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsElasticsearchDomain) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsElasticsearchDomain) KibanaEndpoint() *string {
	var returns *string
	_jsii_.Get(
		j,
		"kibanaEndpoint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsElasticsearchDomain) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsElasticsearchDomain) LogPublishingOptions() DataAwsElasticsearchDomainLogPublishingOptionsList {
	var returns DataAwsElasticsearchDomainLogPublishingOptionsList
	_jsii_.Get(
		j,
		"logPublishingOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsElasticsearchDomain) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsElasticsearchDomain) NodeToNodeEncryption() DataAwsElasticsearchDomainNodeToNodeEncryptionList {
	var returns DataAwsElasticsearchDomainNodeToNodeEncryptionList
	_jsii_.Get(
		j,
		"nodeToNodeEncryption",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsElasticsearchDomain) Processing() cdktf.IResolvable {
	var returns cdktf.IResolvable
	_jsii_.Get(
		j,
		"processing",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsElasticsearchDomain) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsElasticsearchDomain) RawOverrides() any {
	var returns any
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsElasticsearchDomain) SnapshotOptions() DataAwsElasticsearchDomainSnapshotOptionsList {
	var returns DataAwsElasticsearchDomainSnapshotOptionsList
	_jsii_.Get(
		j,
		"snapshotOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsElasticsearchDomain) Tags() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsElasticsearchDomain) TagsInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tagsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsElasticsearchDomain) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsElasticsearchDomain) TerraformMetaArguments() *map[string]any {
	var returns *map[string]any
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsElasticsearchDomain) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsElasticsearchDomain) VpcOptions() DataAwsElasticsearchDomainVpcOptionsList {
	var returns DataAwsElasticsearchDomainVpcOptionsList
	_jsii_.Get(
		j,
		"vpcOptions",
		&returns,
	)
	return returns
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/aws/4.54.0/docs/data-sources/elasticsearch_domain aws_elasticsearch_domain} Data Source.
func NewDataAwsElasticsearchDomain(scope constructs.Construct, id *string, config *DataAwsElasticsearchDomainConfig) DataAwsElasticsearchDomain {
	_init_.Initialize()

	if err := validateNewDataAwsElasticsearchDomainParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_DataAwsElasticsearchDomain{}

	_jsii_.Create(
		"@cdktf/provider-aws.dataAwsElasticsearchDomain.DataAwsElasticsearchDomain",
		[]any{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/aws/4.54.0/docs/data-sources/elasticsearch_domain aws_elasticsearch_domain} Data Source.
func NewDataAwsElasticsearchDomain_Override(d DataAwsElasticsearchDomain, scope constructs.Construct, id *string, config *DataAwsElasticsearchDomainConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-aws.dataAwsElasticsearchDomain.DataAwsElasticsearchDomain",
		[]any{scope, id, config},
		d,
	)
}

func (j *jsiiProxy_DataAwsElasticsearchDomain) SetCount(val any) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_DataAwsElasticsearchDomain) SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_DataAwsElasticsearchDomain) SetDomainName(val *string) {
	if err := j.validateSetDomainNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"domainName",
		val,
	)
}

func (j *jsiiProxy_DataAwsElasticsearchDomain) SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_DataAwsElasticsearchDomain) SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_DataAwsElasticsearchDomain) SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_DataAwsElasticsearchDomain) SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_DataAwsElasticsearchDomain) SetTags(val *map[string]*string) {
	if err := j.validateSetTagsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tags",
		val,
	)
}

// Generates CDKTF code for importing a DataAwsElasticsearchDomain resource upon running "cdktf plan <stack-name>".
func DataAwsElasticsearchDomain_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktf.TerraformProvider) cdktf.ImportableResource {
	_init_.Initialize()

	if err := validateDataAwsElasticsearchDomain_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktf.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktf/provider-aws.dataAwsElasticsearchDomain.DataAwsElasticsearchDomain",
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
func DataAwsElasticsearchDomain_IsConstruct(x any) *bool {
	_init_.Initialize()

	if err := validateDataAwsElasticsearchDomain_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-aws.dataAwsElasticsearchDomain.DataAwsElasticsearchDomain",
		"isConstruct",
		[]any{x},
		&returns,
	)

	return returns
}

// Experimental.
func DataAwsElasticsearchDomain_IsTerraformDataSource(x any) *bool {
	_init_.Initialize()

	if err := validateDataAwsElasticsearchDomain_IsTerraformDataSourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-aws.dataAwsElasticsearchDomain.DataAwsElasticsearchDomain",
		"isTerraformDataSource",
		[]any{x},
		&returns,
	)

	return returns
}

// Experimental.
func DataAwsElasticsearchDomain_IsTerraformElement(x any) *bool {
	_init_.Initialize()

	if err := validateDataAwsElasticsearchDomain_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-aws.dataAwsElasticsearchDomain.DataAwsElasticsearchDomain",
		"isTerraformElement",
		[]any{x},
		&returns,
	)

	return returns
}

func DataAwsElasticsearchDomain_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-aws.dataAwsElasticsearchDomain.DataAwsElasticsearchDomain",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (d *jsiiProxy_DataAwsElasticsearchDomain) AddOverride(path *string, value any) {
	if err := d.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"addOverride",
		[]any{path, value},
	)
}

func (d *jsiiProxy_DataAwsElasticsearchDomain) GetAnyMapAttribute(terraformAttribute *string) *map[string]any {
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

func (d *jsiiProxy_DataAwsElasticsearchDomain) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (d *jsiiProxy_DataAwsElasticsearchDomain) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (d *jsiiProxy_DataAwsElasticsearchDomain) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (d *jsiiProxy_DataAwsElasticsearchDomain) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (d *jsiiProxy_DataAwsElasticsearchDomain) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (d *jsiiProxy_DataAwsElasticsearchDomain) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (d *jsiiProxy_DataAwsElasticsearchDomain) GetStringAttribute(terraformAttribute *string) *string {
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

func (d *jsiiProxy_DataAwsElasticsearchDomain) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (d *jsiiProxy_DataAwsElasticsearchDomain) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (d *jsiiProxy_DataAwsElasticsearchDomain) OverrideLogicalId(newLogicalId *string) {
	if err := d.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"overrideLogicalId",
		[]any{newLogicalId},
	)
}

func (d *jsiiProxy_DataAwsElasticsearchDomain) ResetId() {
	_jsii_.InvokeVoid(
		d,
		"resetId",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataAwsElasticsearchDomain) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		d,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataAwsElasticsearchDomain) ResetTags() {
	_jsii_.InvokeVoid(
		d,
		"resetTags",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataAwsElasticsearchDomain) SynthesizeAttributes() *map[string]any {
	var returns *map[string]any

	_jsii_.Invoke(
		d,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsElasticsearchDomain) SynthesizeHclAttributes() *map[string]any {
	var returns *map[string]any

	_jsii_.Invoke(
		d,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsElasticsearchDomain) ToHclTerraform() any {
	var returns any

	_jsii_.Invoke(
		d,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsElasticsearchDomain) ToMetadata() any {
	var returns any

	_jsii_.Invoke(
		d,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsElasticsearchDomain) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsElasticsearchDomain) ToTerraform() any {
	var returns any

	_jsii_.Invoke(
		d,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}
