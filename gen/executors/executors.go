package executors

import (
	"reflect"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func init() {
	_jsii_.RegisterClass(
		"executors.Executors",
		reflect.TypeOf((*Executors)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addOverride", GoMethod: "AddOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addProvider", GoMethod: "AddProvider"},
			_jsii_.MemberProperty{JsiiProperty: "assignPublicIp", GoGetter: "AssignPublicIp"},
			_jsii_.MemberProperty{JsiiProperty: "bootDiskSize", GoGetter: "BootDiskSize"},
			_jsii_.MemberProperty{JsiiProperty: "cdktfStack", GoGetter: "CdktfStack"},
			_jsii_.MemberProperty{JsiiProperty: "constructNodeMetadata", GoGetter: "ConstructNodeMetadata"},
			_jsii_.MemberProperty{JsiiProperty: "dependsOn", GoGetter: "DependsOn"},
			_jsii_.MemberProperty{JsiiProperty: "dockerRegistryMirror", GoGetter: "DockerRegistryMirror"},
			_jsii_.MemberProperty{JsiiProperty: "dockerRegistryMirrorNodeExporterUrl", GoGetter: "DockerRegistryMirrorNodeExporterUrl"},
			_jsii_.MemberProperty{JsiiProperty: "firecrackerDiskSpace", GoGetter: "FirecrackerDiskSpace"},
			_jsii_.MemberProperty{JsiiProperty: "firecrackerMemory", GoGetter: "FirecrackerMemory"},
			_jsii_.MemberProperty{JsiiProperty: "firecrackerNumCpus", GoGetter: "FirecrackerNumCpus"},
			_jsii_.MemberProperty{JsiiProperty: "forEach", GoGetter: "ForEach"},
			_jsii_.MemberProperty{JsiiProperty: "fqn", GoGetter: "Fqn"},
			_jsii_.MemberProperty{JsiiProperty: "friendlyUniqueId", GoGetter: "FriendlyUniqueId"},
			_jsii_.MemberMethod{JsiiMethod: "getString", GoMethod: "GetString"},
			_jsii_.MemberProperty{JsiiProperty: "httpAccessCidrRanges", GoGetter: "HttpAccessCidrRanges"},
			_jsii_.MemberProperty{JsiiProperty: "instanceTag", GoGetter: "InstanceTag"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForOutput", GoMethod: "InterpolationForOutput"},
			_jsii_.MemberProperty{JsiiProperty: "jobMemory", GoGetter: "JobMemory"},
			_jsii_.MemberProperty{JsiiProperty: "jobNumCpus", GoGetter: "JobNumCpus"},
			_jsii_.MemberProperty{JsiiProperty: "jobsPerInstanceScaling", GoGetter: "JobsPerInstanceScaling"},
			_jsii_.MemberProperty{JsiiProperty: "machineImage", GoGetter: "MachineImage"},
			_jsii_.MemberProperty{JsiiProperty: "machineType", GoGetter: "MachineType"},
			_jsii_.MemberProperty{JsiiProperty: "maxActiveTime", GoGetter: "MaxActiveTime"},
			_jsii_.MemberProperty{JsiiProperty: "maximumNumJobs", GoGetter: "MaximumNumJobs"},
			_jsii_.MemberProperty{JsiiProperty: "maximumRuntimePerJob", GoGetter: "MaximumRuntimePerJob"},
			_jsii_.MemberProperty{JsiiProperty: "maxReplicas", GoGetter: "MaxReplicas"},
			_jsii_.MemberProperty{JsiiProperty: "metricsEnvironmentLabel", GoGetter: "MetricsEnvironmentLabel"},
			_jsii_.MemberProperty{JsiiProperty: "minReplicas", GoGetter: "MinReplicas"},
			_jsii_.MemberProperty{JsiiProperty: "networkId", GoGetter: "NetworkId"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "numTotalJobs", GoGetter: "NumTotalJobs"},
			_jsii_.MemberMethod{JsiiMethod: "overrideLogicalId", GoMethod: "OverrideLogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "preemptibleMachines", GoGetter: "PreemptibleMachines"},
			_jsii_.MemberProperty{JsiiProperty: "providers", GoGetter: "Providers"},
			_jsii_.MemberProperty{JsiiProperty: "queueName", GoGetter: "QueueName"},
			_jsii_.MemberProperty{JsiiProperty: "rawOverrides", GoGetter: "RawOverrides"},
			_jsii_.MemberMethod{JsiiMethod: "resetOverrideLogicalId", GoMethod: "ResetOverrideLogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "resourcePrefix", GoGetter: "ResourcePrefix"},
			_jsii_.MemberProperty{JsiiProperty: "skipAssetCreationFromLocalModules", GoGetter: "SkipAssetCreationFromLocalModules"},
			_jsii_.MemberProperty{JsiiProperty: "source", GoGetter: "Source"},
			_jsii_.MemberProperty{JsiiProperty: "sourcegraphExecutorProxyPassword", GoGetter: "SourcegraphExecutorProxyPassword"},
			_jsii_.MemberProperty{JsiiProperty: "sourcegraphExternalUrl", GoGetter: "SourcegraphExternalUrl"},
			_jsii_.MemberProperty{JsiiProperty: "subnetId", GoGetter: "SubnetId"},
			_jsii_.MemberMethod{JsiiMethod: "synthesizeAttributes", GoMethod: "SynthesizeAttributes"},
			_jsii_.MemberMethod{JsiiMethod: "toMetadata", GoMethod: "ToMetadata"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberMethod{JsiiMethod: "toTerraform", GoMethod: "ToTerraform"},
			_jsii_.MemberProperty{JsiiProperty: "useFirecracker", GoGetter: "UseFirecracker"},
			_jsii_.MemberProperty{JsiiProperty: "version", GoGetter: "Version"},
			_jsii_.MemberProperty{JsiiProperty: "zone", GoGetter: "Zone"},
		},
		func() interface{} {
			j := jsiiProxy_Executors{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfTerraformModule)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"executors.ExecutorsOptions",
		reflect.TypeOf((*ExecutorsOptions)(nil)).Elem(),
	)
}
