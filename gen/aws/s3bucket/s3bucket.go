package s3bucket

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/aws/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/sourcegraph/controller-cdktf/gen/aws/s3bucket/internal"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/aws/4.54.0/docs/resources/s3_bucket aws_s3_bucket}.
type S3Bucket interface {
	cdktf.TerraformResource
	AccelerationStatus() *string
	SetAccelerationStatus(val *string)
	AccelerationStatusInput() *string
	Acl() *string
	SetAcl(val *string)
	AclInput() *string
	Arn() *string
	Bucket() *string
	SetBucket(val *string)
	BucketDomainName() *string
	BucketInput() *string
	BucketPrefix() *string
	SetBucketPrefix(val *string)
	BucketPrefixInput() *string
	BucketRegionalDomainName() *string
	// Experimental.
	CdktfStack() cdktf.TerraformStack
	// Experimental.
	Connection() any
	// Experimental.
	SetConnection(val any)
	// Experimental.
	ConstructNodeMetadata() *map[string]any
	CorsRule() S3BucketCorsRuleList
	CorsRuleInput() any
	// Experimental.
	Count() any
	// Experimental.
	SetCount(val any)
	// Experimental.
	DependsOn() *[]*string
	// Experimental.
	SetDependsOn(val *[]*string)
	ForceDestroy() any
	SetForceDestroy(val any)
	ForceDestroyInput() any
	// Experimental.
	ForEach() cdktf.ITerraformIterator
	// Experimental.
	SetForEach(val cdktf.ITerraformIterator)
	// Experimental.
	Fqn() *string
	// Experimental.
	FriendlyUniqueId() *string
	Grant() S3BucketGrantList
	GrantInput() any
	HostedZoneId() *string
	Id() *string
	SetId(val *string)
	IdInput() *string
	// Experimental.
	Lifecycle() *cdktf.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	LifecycleRule() S3BucketLifecycleRuleList
	LifecycleRuleInput() any
	Logging() S3BucketLoggingOutputReference
	LoggingInput() *S3BucketLogging
	// The tree node.
	Node() constructs.Node
	ObjectLockConfiguration() S3BucketObjectLockConfigurationOutputReference
	ObjectLockConfigurationInput() *S3BucketObjectLockConfiguration
	ObjectLockEnabled() any
	SetObjectLockEnabled(val any)
	ObjectLockEnabledInput() any
	Policy() *string
	SetPolicy(val *string)
	PolicyInput() *string
	// Experimental.
	Provider() cdktf.TerraformProvider
	// Experimental.
	SetProvider(val cdktf.TerraformProvider)
	// Experimental.
	Provisioners() *[]any
	// Experimental.
	SetProvisioners(val *[]any)
	// Experimental.
	RawOverrides() any
	Region() *string
	ReplicationConfiguration() S3BucketReplicationConfigurationOutputReference
	ReplicationConfigurationInput() *S3BucketReplicationConfiguration
	RequestPayer() *string
	SetRequestPayer(val *string)
	RequestPayerInput() *string
	ServerSideEncryptionConfiguration() S3BucketServerSideEncryptionConfigurationOutputReference
	ServerSideEncryptionConfigurationInput() *S3BucketServerSideEncryptionConfiguration
	Tags() *map[string]*string
	SetTags(val *map[string]*string)
	TagsAll() *map[string]*string
	SetTagsAll(val *map[string]*string)
	TagsAllInput() *map[string]*string
	TagsInput() *map[string]*string
	// Experimental.
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]any
	// Experimental.
	TerraformResourceType() *string
	Timeouts() S3BucketTimeoutsOutputReference
	TimeoutsInput() any
	Versioning() S3BucketVersioningOutputReference
	VersioningInput() *S3BucketVersioning
	Website() S3BucketWebsiteOutputReference
	WebsiteDomain() *string
	WebsiteEndpoint() *string
	WebsiteInput() *S3BucketWebsite
	// Adds a user defined moveTarget string to this resource to be later used in .moveTo(moveTarget) to resolve the location of the move.
	// Experimental.
	AddMoveTarget(moveTarget *string)
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
	HasResourceMove() any
	// Experimental.
	ImportFrom(id *string, provider cdktf.TerraformProvider)
	// Experimental.
	InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable
	// Move the resource corresponding to "id" to this resource.
	//
	// Note that the resource being moved from must be marked as moved using it's instance function.
	// Experimental.
	MoveFromId(id *string)
	// Moves this resource to the target resource given by moveTarget.
	// Experimental.
	MoveTo(moveTarget *string, index any)
	// Moves this resource to the resource corresponding to "id".
	// Experimental.
	MoveToId(id *string)
	// Overrides the auto-generated logical ID with a specific ID.
	// Experimental.
	OverrideLogicalId(newLogicalId *string)
	PutCorsRule(value any)
	PutGrant(value any)
	PutLifecycleRule(value any)
	PutLogging(value *S3BucketLogging)
	PutObjectLockConfiguration(value *S3BucketObjectLockConfiguration)
	PutReplicationConfiguration(value *S3BucketReplicationConfiguration)
	PutServerSideEncryptionConfiguration(value *S3BucketServerSideEncryptionConfiguration)
	PutTimeouts(value *S3BucketTimeouts)
	PutVersioning(value *S3BucketVersioning)
	PutWebsite(value *S3BucketWebsite)
	ResetAccelerationStatus()
	ResetAcl()
	ResetBucket()
	ResetBucketPrefix()
	ResetCorsRule()
	ResetForceDestroy()
	ResetGrant()
	ResetId()
	ResetLifecycleRule()
	ResetLogging()
	ResetObjectLockConfiguration()
	ResetObjectLockEnabled()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetPolicy()
	ResetReplicationConfiguration()
	ResetRequestPayer()
	ResetServerSideEncryptionConfiguration()
	ResetTags()
	ResetTagsAll()
	ResetTimeouts()
	ResetVersioning()
	ResetWebsite()
	SynthesizeAttributes() *map[string]any
	SynthesizeHclAttributes() *map[string]any
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

// The jsii proxy struct for S3Bucket
type jsiiProxy_S3Bucket struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_S3Bucket) AccelerationStatus() *string {
	var returns *string
	_jsii_.Get(
		j,
		"accelerationStatus",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_S3Bucket) AccelerationStatusInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"accelerationStatusInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_S3Bucket) Acl() *string {
	var returns *string
	_jsii_.Get(
		j,
		"acl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_S3Bucket) AclInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"aclInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_S3Bucket) Arn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"arn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_S3Bucket) Bucket() *string {
	var returns *string
	_jsii_.Get(
		j,
		"bucket",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_S3Bucket) BucketDomainName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"bucketDomainName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_S3Bucket) BucketInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"bucketInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_S3Bucket) BucketPrefix() *string {
	var returns *string
	_jsii_.Get(
		j,
		"bucketPrefix",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_S3Bucket) BucketPrefixInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"bucketPrefixInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_S3Bucket) BucketRegionalDomainName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"bucketRegionalDomainName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_S3Bucket) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_S3Bucket) Connection() any {
	var returns any
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_S3Bucket) ConstructNodeMetadata() *map[string]any {
	var returns *map[string]any
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_S3Bucket) CorsRule() S3BucketCorsRuleList {
	var returns S3BucketCorsRuleList
	_jsii_.Get(
		j,
		"corsRule",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_S3Bucket) CorsRuleInput() any {
	var returns any
	_jsii_.Get(
		j,
		"corsRuleInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_S3Bucket) Count() any {
	var returns any
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_S3Bucket) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_S3Bucket) ForceDestroy() any {
	var returns any
	_jsii_.Get(
		j,
		"forceDestroy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_S3Bucket) ForceDestroyInput() any {
	var returns any
	_jsii_.Get(
		j,
		"forceDestroyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_S3Bucket) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_S3Bucket) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_S3Bucket) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_S3Bucket) Grant() S3BucketGrantList {
	var returns S3BucketGrantList
	_jsii_.Get(
		j,
		"grant",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_S3Bucket) GrantInput() any {
	var returns any
	_jsii_.Get(
		j,
		"grantInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_S3Bucket) HostedZoneId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"hostedZoneId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_S3Bucket) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_S3Bucket) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_S3Bucket) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_S3Bucket) LifecycleRule() S3BucketLifecycleRuleList {
	var returns S3BucketLifecycleRuleList
	_jsii_.Get(
		j,
		"lifecycleRule",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_S3Bucket) LifecycleRuleInput() any {
	var returns any
	_jsii_.Get(
		j,
		"lifecycleRuleInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_S3Bucket) Logging() S3BucketLoggingOutputReference {
	var returns S3BucketLoggingOutputReference
	_jsii_.Get(
		j,
		"logging",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_S3Bucket) LoggingInput() *S3BucketLogging {
	var returns *S3BucketLogging
	_jsii_.Get(
		j,
		"loggingInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_S3Bucket) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_S3Bucket) ObjectLockConfiguration() S3BucketObjectLockConfigurationOutputReference {
	var returns S3BucketObjectLockConfigurationOutputReference
	_jsii_.Get(
		j,
		"objectLockConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_S3Bucket) ObjectLockConfigurationInput() *S3BucketObjectLockConfiguration {
	var returns *S3BucketObjectLockConfiguration
	_jsii_.Get(
		j,
		"objectLockConfigurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_S3Bucket) ObjectLockEnabled() any {
	var returns any
	_jsii_.Get(
		j,
		"objectLockEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_S3Bucket) ObjectLockEnabledInput() any {
	var returns any
	_jsii_.Get(
		j,
		"objectLockEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_S3Bucket) Policy() *string {
	var returns *string
	_jsii_.Get(
		j,
		"policy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_S3Bucket) PolicyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"policyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_S3Bucket) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_S3Bucket) Provisioners() *[]any {
	var returns *[]any
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_S3Bucket) RawOverrides() any {
	var returns any
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_S3Bucket) Region() *string {
	var returns *string
	_jsii_.Get(
		j,
		"region",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_S3Bucket) ReplicationConfiguration() S3BucketReplicationConfigurationOutputReference {
	var returns S3BucketReplicationConfigurationOutputReference
	_jsii_.Get(
		j,
		"replicationConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_S3Bucket) ReplicationConfigurationInput() *S3BucketReplicationConfiguration {
	var returns *S3BucketReplicationConfiguration
	_jsii_.Get(
		j,
		"replicationConfigurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_S3Bucket) RequestPayer() *string {
	var returns *string
	_jsii_.Get(
		j,
		"requestPayer",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_S3Bucket) RequestPayerInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"requestPayerInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_S3Bucket) ServerSideEncryptionConfiguration() S3BucketServerSideEncryptionConfigurationOutputReference {
	var returns S3BucketServerSideEncryptionConfigurationOutputReference
	_jsii_.Get(
		j,
		"serverSideEncryptionConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_S3Bucket) ServerSideEncryptionConfigurationInput() *S3BucketServerSideEncryptionConfiguration {
	var returns *S3BucketServerSideEncryptionConfiguration
	_jsii_.Get(
		j,
		"serverSideEncryptionConfigurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_S3Bucket) Tags() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_S3Bucket) TagsAll() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tagsAll",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_S3Bucket) TagsAllInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tagsAllInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_S3Bucket) TagsInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tagsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_S3Bucket) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_S3Bucket) TerraformMetaArguments() *map[string]any {
	var returns *map[string]any
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_S3Bucket) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_S3Bucket) Timeouts() S3BucketTimeoutsOutputReference {
	var returns S3BucketTimeoutsOutputReference
	_jsii_.Get(
		j,
		"timeouts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_S3Bucket) TimeoutsInput() any {
	var returns any
	_jsii_.Get(
		j,
		"timeoutsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_S3Bucket) Versioning() S3BucketVersioningOutputReference {
	var returns S3BucketVersioningOutputReference
	_jsii_.Get(
		j,
		"versioning",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_S3Bucket) VersioningInput() *S3BucketVersioning {
	var returns *S3BucketVersioning
	_jsii_.Get(
		j,
		"versioningInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_S3Bucket) Website() S3BucketWebsiteOutputReference {
	var returns S3BucketWebsiteOutputReference
	_jsii_.Get(
		j,
		"website",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_S3Bucket) WebsiteDomain() *string {
	var returns *string
	_jsii_.Get(
		j,
		"websiteDomain",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_S3Bucket) WebsiteEndpoint() *string {
	var returns *string
	_jsii_.Get(
		j,
		"websiteEndpoint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_S3Bucket) WebsiteInput() *S3BucketWebsite {
	var returns *S3BucketWebsite
	_jsii_.Get(
		j,
		"websiteInput",
		&returns,
	)
	return returns
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/aws/4.54.0/docs/resources/s3_bucket aws_s3_bucket} Resource.
func NewS3Bucket(scope constructs.Construct, id *string, config *S3BucketConfig) S3Bucket {
	_init_.Initialize()

	if err := validateNewS3BucketParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_S3Bucket{}

	_jsii_.Create(
		"@cdktf/provider-aws.s3Bucket.S3Bucket",
		[]any{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/aws/4.54.0/docs/resources/s3_bucket aws_s3_bucket} Resource.
func NewS3Bucket_Override(s S3Bucket, scope constructs.Construct, id *string, config *S3BucketConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-aws.s3Bucket.S3Bucket",
		[]any{scope, id, config},
		s,
	)
}

func (j *jsiiProxy_S3Bucket) SetAccelerationStatus(val *string) {
	if err := j.validateSetAccelerationStatusParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"accelerationStatus",
		val,
	)
}

func (j *jsiiProxy_S3Bucket) SetAcl(val *string) {
	if err := j.validateSetAclParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"acl",
		val,
	)
}

func (j *jsiiProxy_S3Bucket) SetBucket(val *string) {
	if err := j.validateSetBucketParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"bucket",
		val,
	)
}

func (j *jsiiProxy_S3Bucket) SetBucketPrefix(val *string) {
	if err := j.validateSetBucketPrefixParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"bucketPrefix",
		val,
	)
}

func (j *jsiiProxy_S3Bucket) SetConnection(val any) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_S3Bucket) SetCount(val any) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_S3Bucket) SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_S3Bucket) SetForceDestroy(val any) {
	if err := j.validateSetForceDestroyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"forceDestroy",
		val,
	)
}

func (j *jsiiProxy_S3Bucket) SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_S3Bucket) SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_S3Bucket) SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_S3Bucket) SetObjectLockEnabled(val any) {
	if err := j.validateSetObjectLockEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"objectLockEnabled",
		val,
	)
}

func (j *jsiiProxy_S3Bucket) SetPolicy(val *string) {
	if err := j.validateSetPolicyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"policy",
		val,
	)
}

func (j *jsiiProxy_S3Bucket) SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_S3Bucket) SetProvisioners(val *[]any) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_S3Bucket) SetRequestPayer(val *string) {
	if err := j.validateSetRequestPayerParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"requestPayer",
		val,
	)
}

func (j *jsiiProxy_S3Bucket) SetTags(val *map[string]*string) {
	if err := j.validateSetTagsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tags",
		val,
	)
}

func (j *jsiiProxy_S3Bucket) SetTagsAll(val *map[string]*string) {
	if err := j.validateSetTagsAllParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tagsAll",
		val,
	)
}

// Generates CDKTF code for importing a S3Bucket resource upon running "cdktf plan <stack-name>".
func S3Bucket_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktf.TerraformProvider) cdktf.ImportableResource {
	_init_.Initialize()

	if err := validateS3Bucket_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktf.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktf/provider-aws.s3Bucket.S3Bucket",
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
func S3Bucket_IsConstruct(x any) *bool {
	_init_.Initialize()

	if err := validateS3Bucket_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-aws.s3Bucket.S3Bucket",
		"isConstruct",
		[]any{x},
		&returns,
	)

	return returns
}

// Experimental.
func S3Bucket_IsTerraformElement(x any) *bool {
	_init_.Initialize()

	if err := validateS3Bucket_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-aws.s3Bucket.S3Bucket",
		"isTerraformElement",
		[]any{x},
		&returns,
	)

	return returns
}

// Experimental.
func S3Bucket_IsTerraformResource(x any) *bool {
	_init_.Initialize()

	if err := validateS3Bucket_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-aws.s3Bucket.S3Bucket",
		"isTerraformResource",
		[]any{x},
		&returns,
	)

	return returns
}

func S3Bucket_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-aws.s3Bucket.S3Bucket",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (s *jsiiProxy_S3Bucket) AddMoveTarget(moveTarget *string) {
	if err := s.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"addMoveTarget",
		[]any{moveTarget},
	)
}

func (s *jsiiProxy_S3Bucket) AddOverride(path *string, value any) {
	if err := s.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"addOverride",
		[]any{path, value},
	)
}

func (s *jsiiProxy_S3Bucket) GetAnyMapAttribute(terraformAttribute *string) *map[string]any {
	if err := s.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]any

	_jsii_.Invoke(
		s,
		"getAnyMapAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_S3Bucket) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := s.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		s,
		"getBooleanAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_S3Bucket) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := s.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		s,
		"getBooleanMapAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_S3Bucket) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := s.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		s,
		"getListAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_S3Bucket) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := s.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		s,
		"getNumberAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_S3Bucket) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := s.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		s,
		"getNumberListAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_S3Bucket) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := s.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		s,
		"getNumberMapAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_S3Bucket) GetStringAttribute(terraformAttribute *string) *string {
	if err := s.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		s,
		"getStringAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_S3Bucket) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := s.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		s,
		"getStringMapAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_S3Bucket) HasResourceMove() any {
	var returns any

	_jsii_.Invoke(
		s,
		"hasResourceMove",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_S3Bucket) ImportFrom(id *string, provider cdktf.TerraformProvider) {
	if err := s.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"importFrom",
		[]any{id, provider},
	)
}

func (s *jsiiProxy_S3Bucket) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := s.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		s,
		"interpolationForAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_S3Bucket) MoveFromId(id *string) {
	if err := s.validateMoveFromIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"moveFromId",
		[]any{id},
	)
}

func (s *jsiiProxy_S3Bucket) MoveTo(moveTarget *string, index any) {
	if err := s.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"moveTo",
		[]any{moveTarget, index},
	)
}

func (s *jsiiProxy_S3Bucket) MoveToId(id *string) {
	if err := s.validateMoveToIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"moveToId",
		[]any{id},
	)
}

func (s *jsiiProxy_S3Bucket) OverrideLogicalId(newLogicalId *string) {
	if err := s.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"overrideLogicalId",
		[]any{newLogicalId},
	)
}

func (s *jsiiProxy_S3Bucket) PutCorsRule(value any) {
	if err := s.validatePutCorsRuleParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putCorsRule",
		[]any{value},
	)
}

func (s *jsiiProxy_S3Bucket) PutGrant(value any) {
	if err := s.validatePutGrantParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putGrant",
		[]any{value},
	)
}

func (s *jsiiProxy_S3Bucket) PutLifecycleRule(value any) {
	if err := s.validatePutLifecycleRuleParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putLifecycleRule",
		[]any{value},
	)
}

func (s *jsiiProxy_S3Bucket) PutLogging(value *S3BucketLogging) {
	if err := s.validatePutLoggingParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putLogging",
		[]any{value},
	)
}

func (s *jsiiProxy_S3Bucket) PutObjectLockConfiguration(value *S3BucketObjectLockConfiguration) {
	if err := s.validatePutObjectLockConfigurationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putObjectLockConfiguration",
		[]any{value},
	)
}

func (s *jsiiProxy_S3Bucket) PutReplicationConfiguration(value *S3BucketReplicationConfiguration) {
	if err := s.validatePutReplicationConfigurationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putReplicationConfiguration",
		[]any{value},
	)
}

func (s *jsiiProxy_S3Bucket) PutServerSideEncryptionConfiguration(value *S3BucketServerSideEncryptionConfiguration) {
	if err := s.validatePutServerSideEncryptionConfigurationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putServerSideEncryptionConfiguration",
		[]any{value},
	)
}

func (s *jsiiProxy_S3Bucket) PutTimeouts(value *S3BucketTimeouts) {
	if err := s.validatePutTimeoutsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putTimeouts",
		[]any{value},
	)
}

func (s *jsiiProxy_S3Bucket) PutVersioning(value *S3BucketVersioning) {
	if err := s.validatePutVersioningParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putVersioning",
		[]any{value},
	)
}

func (s *jsiiProxy_S3Bucket) PutWebsite(value *S3BucketWebsite) {
	if err := s.validatePutWebsiteParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putWebsite",
		[]any{value},
	)
}

func (s *jsiiProxy_S3Bucket) ResetAccelerationStatus() {
	_jsii_.InvokeVoid(
		s,
		"resetAccelerationStatus",
		nil, // no parameters
	)
}

func (s *jsiiProxy_S3Bucket) ResetAcl() {
	_jsii_.InvokeVoid(
		s,
		"resetAcl",
		nil, // no parameters
	)
}

func (s *jsiiProxy_S3Bucket) ResetBucket() {
	_jsii_.InvokeVoid(
		s,
		"resetBucket",
		nil, // no parameters
	)
}

func (s *jsiiProxy_S3Bucket) ResetBucketPrefix() {
	_jsii_.InvokeVoid(
		s,
		"resetBucketPrefix",
		nil, // no parameters
	)
}

func (s *jsiiProxy_S3Bucket) ResetCorsRule() {
	_jsii_.InvokeVoid(
		s,
		"resetCorsRule",
		nil, // no parameters
	)
}

func (s *jsiiProxy_S3Bucket) ResetForceDestroy() {
	_jsii_.InvokeVoid(
		s,
		"resetForceDestroy",
		nil, // no parameters
	)
}

func (s *jsiiProxy_S3Bucket) ResetGrant() {
	_jsii_.InvokeVoid(
		s,
		"resetGrant",
		nil, // no parameters
	)
}

func (s *jsiiProxy_S3Bucket) ResetId() {
	_jsii_.InvokeVoid(
		s,
		"resetId",
		nil, // no parameters
	)
}

func (s *jsiiProxy_S3Bucket) ResetLifecycleRule() {
	_jsii_.InvokeVoid(
		s,
		"resetLifecycleRule",
		nil, // no parameters
	)
}

func (s *jsiiProxy_S3Bucket) ResetLogging() {
	_jsii_.InvokeVoid(
		s,
		"resetLogging",
		nil, // no parameters
	)
}

func (s *jsiiProxy_S3Bucket) ResetObjectLockConfiguration() {
	_jsii_.InvokeVoid(
		s,
		"resetObjectLockConfiguration",
		nil, // no parameters
	)
}

func (s *jsiiProxy_S3Bucket) ResetObjectLockEnabled() {
	_jsii_.InvokeVoid(
		s,
		"resetObjectLockEnabled",
		nil, // no parameters
	)
}

func (s *jsiiProxy_S3Bucket) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		s,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (s *jsiiProxy_S3Bucket) ResetPolicy() {
	_jsii_.InvokeVoid(
		s,
		"resetPolicy",
		nil, // no parameters
	)
}

func (s *jsiiProxy_S3Bucket) ResetReplicationConfiguration() {
	_jsii_.InvokeVoid(
		s,
		"resetReplicationConfiguration",
		nil, // no parameters
	)
}

func (s *jsiiProxy_S3Bucket) ResetRequestPayer() {
	_jsii_.InvokeVoid(
		s,
		"resetRequestPayer",
		nil, // no parameters
	)
}

func (s *jsiiProxy_S3Bucket) ResetServerSideEncryptionConfiguration() {
	_jsii_.InvokeVoid(
		s,
		"resetServerSideEncryptionConfiguration",
		nil, // no parameters
	)
}

func (s *jsiiProxy_S3Bucket) ResetTags() {
	_jsii_.InvokeVoid(
		s,
		"resetTags",
		nil, // no parameters
	)
}

func (s *jsiiProxy_S3Bucket) ResetTagsAll() {
	_jsii_.InvokeVoid(
		s,
		"resetTagsAll",
		nil, // no parameters
	)
}

func (s *jsiiProxy_S3Bucket) ResetTimeouts() {
	_jsii_.InvokeVoid(
		s,
		"resetTimeouts",
		nil, // no parameters
	)
}

func (s *jsiiProxy_S3Bucket) ResetVersioning() {
	_jsii_.InvokeVoid(
		s,
		"resetVersioning",
		nil, // no parameters
	)
}

func (s *jsiiProxy_S3Bucket) ResetWebsite() {
	_jsii_.InvokeVoid(
		s,
		"resetWebsite",
		nil, // no parameters
	)
}

func (s *jsiiProxy_S3Bucket) SynthesizeAttributes() *map[string]any {
	var returns *map[string]any

	_jsii_.Invoke(
		s,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_S3Bucket) SynthesizeHclAttributes() *map[string]any {
	var returns *map[string]any

	_jsii_.Invoke(
		s,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_S3Bucket) ToHclTerraform() any {
	var returns any

	_jsii_.Invoke(
		s,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_S3Bucket) ToMetadata() any {
	var returns any

	_jsii_.Invoke(
		s,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_S3Bucket) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_S3Bucket) ToTerraform() any {
	var returns any

	_jsii_.Invoke(
		s,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}
