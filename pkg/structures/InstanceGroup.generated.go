package structures

import (
	"reflect"

	"github.com/eddycharly/terraform-provider-kops/pkg/api"
	"k8s.io/kops/pkg/apis/kops"
)

func ExpandInstanceGroup(in map[string]interface{}) api.InstanceGroup {
	if in == nil {
		panic("expand InstanceGroup failure, in is nil")
	}
	return api.InstanceGroup{
		Name: func(in interface{}) string {
			return string(ExpandString(in))
		}(in["name"]),
		Role: func(in interface{}) kops.InstanceGroupRole {
			return kops.InstanceGroupRole(ExpandString(in))
		}(in["role"]),
		Image: func(in interface{}) string {
			return string(ExpandString(in))
		}(in["image"]),
		MinSize: func(in interface{}) *int32 {
			return func(in interface{}) *int32 {
				if in == nil {
					return nil
				}
				if _, ok := in.([]interface{}); ok && len(in.([]interface{})) == 0 {
					return nil
				}
				return func(in int32) *int32 {
					return &in
				}(int32(ExpandInt(in)))
			}(in)
		}(in["min_size"]),
		MaxSize: func(in interface{}) *int32 {
			return func(in interface{}) *int32 {
				if in == nil {
					return nil
				}
				if _, ok := in.([]interface{}); ok && len(in.([]interface{})) == 0 {
					return nil
				}
				return func(in int32) *int32 {
					return &in
				}(int32(ExpandInt(in)))
			}(in)
		}(in["max_size"]),
		MachineType: func(in interface{}) string {
			return string(ExpandString(in))
		}(in["machine_type"]),
		RootVolumeSize: func(in interface{}) *int32 {
			if reflect.DeepEqual(in, reflect.Zero(reflect.TypeOf(in)).Interface()) {
				return nil
			}
			return func(in interface{}) *int32 {
				if in == nil {
					return nil
				}
				if _, ok := in.([]interface{}); ok && len(in.([]interface{})) == 0 {
					return nil
				}
				return func(in int32) *int32 {
					return &in
				}(int32(ExpandInt(in)))
			}(in)
		}(in["root_volume_size"]),
		RootVolumeType: func(in interface{}) *string {
			if reflect.DeepEqual(in, reflect.Zero(reflect.TypeOf(in)).Interface()) {
				return nil
			}
			return func(in interface{}) *string {
				if in == nil {
					return nil
				}
				if _, ok := in.([]interface{}); ok && len(in.([]interface{})) == 0 {
					return nil
				}
				return func(in string) *string {
					return &in
				}(string(ExpandString(in)))
			}(in)
		}(in["root_volume_type"]),
		RootVolumeIops: func(in interface{}) *int32 {
			if reflect.DeepEqual(in, reflect.Zero(reflect.TypeOf(in)).Interface()) {
				return nil
			}
			return func(in interface{}) *int32 {
				if in == nil {
					return nil
				}
				if _, ok := in.([]interface{}); ok && len(in.([]interface{})) == 0 {
					return nil
				}
				return func(in int32) *int32 {
					return &in
				}(int32(ExpandInt(in)))
			}(in)
		}(in["root_volume_iops"]),
		RootVolumeOptimization: func(in interface{}) *bool {
			if reflect.DeepEqual(in, reflect.Zero(reflect.TypeOf(in)).Interface()) {
				return nil
			}
			return func(in interface{}) *bool {
				if in == nil {
					return nil
				}
				if _, ok := in.([]interface{}); ok && len(in.([]interface{})) == 0 {
					return nil
				}
				return func(in bool) *bool {
					return &in
				}(bool(ExpandBool(in)))
			}(in)
		}(in["root_volume_optimization"]),
		RootVolumeDeleteOnTermination: func(in interface{}) *bool {
			if reflect.DeepEqual(in, reflect.Zero(reflect.TypeOf(in)).Interface()) {
				return nil
			}
			return func(in interface{}) *bool {
				if in == nil {
					return nil
				}
				if _, ok := in.([]interface{}); ok && len(in.([]interface{})) == 0 {
					return nil
				}
				return func(in bool) *bool {
					return &in
				}(bool(ExpandBool(in)))
			}(in)
		}(in["root_volume_delete_on_termination"]),
		RootVolumeEncryption: func(in interface{}) *bool {
			if reflect.DeepEqual(in, reflect.Zero(reflect.TypeOf(in)).Interface()) {
				return nil
			}
			return func(in interface{}) *bool {
				if in == nil {
					return nil
				}
				if _, ok := in.([]interface{}); ok && len(in.([]interface{})) == 0 {
					return nil
				}
				return func(in bool) *bool {
					return &in
				}(bool(ExpandBool(in)))
			}(in)
		}(in["root_volume_encryption"]),
		Volumes: func(in interface{}) []*kops.VolumeSpec {
			return func(in interface{}) []*kops.VolumeSpec {
				var out []*kops.VolumeSpec
				for _, in := range in.([]interface{}) {
					out = append(out, func(in interface{}) *kops.VolumeSpec {
						if in == nil {
							return nil
						}
						if _, ok := in.([]interface{}); ok && len(in.([]interface{})) == 0 {
							return nil
						}
						return func(in kops.VolumeSpec) *kops.VolumeSpec {
							return &in
						}(func(in interface{}) kops.VolumeSpec {
							if in == nil {
								return kops.VolumeSpec{}
							}
							return (ExpandVolumeSpec(in.(map[string]interface{})))
						}(in))
					}(in))
				}
				return out
			}(in)
		}(in["volumes"]),
		VolumeMounts: func(in interface{}) []*kops.VolumeMountSpec {
			return func(in interface{}) []*kops.VolumeMountSpec {
				var out []*kops.VolumeMountSpec
				for _, in := range in.([]interface{}) {
					out = append(out, func(in interface{}) *kops.VolumeMountSpec {
						if in == nil {
							return nil
						}
						if _, ok := in.([]interface{}); ok && len(in.([]interface{})) == 0 {
							return nil
						}
						return func(in kops.VolumeMountSpec) *kops.VolumeMountSpec {
							return &in
						}(func(in interface{}) kops.VolumeMountSpec {
							if in == nil {
								return kops.VolumeMountSpec{}
							}
							return (ExpandVolumeMountSpec(in.(map[string]interface{})))
						}(in))
					}(in))
				}
				return out
			}(in)
		}(in["volume_mounts"]),
		Subnets: func(in interface{}) []string {
			return func(in interface{}) []string {
				var out []string
				for _, in := range in.([]interface{}) {
					out = append(out, string(ExpandString(in)))
				}
				return out
			}(in)
		}(in["subnets"]),
		Zones: func(in interface{}) []string {
			return func(in interface{}) []string {
				var out []string
				for _, in := range in.([]interface{}) {
					out = append(out, string(ExpandString(in)))
				}
				return out
			}(in)
		}(in["zones"]),
		Hooks: func(in interface{}) []kops.HookSpec {
			return func(in interface{}) []kops.HookSpec {
				var out []kops.HookSpec
				for _, in := range in.([]interface{}) {
					out = append(out, func(in interface{}) kops.HookSpec {
						if in == nil {
							return kops.HookSpec{}
						}
						return (ExpandHookSpec(in.(map[string]interface{})))
					}(in))
				}
				return out
			}(in)
		}(in["hooks"]),
		MaxPrice: func(in interface{}) *string {
			if reflect.DeepEqual(in, reflect.Zero(reflect.TypeOf(in)).Interface()) {
				return nil
			}
			return func(in interface{}) *string {
				if in == nil {
					return nil
				}
				if _, ok := in.([]interface{}); ok && len(in.([]interface{})) == 0 {
					return nil
				}
				return func(in string) *string {
					return &in
				}(string(ExpandString(in)))
			}(in)
		}(in["max_price"]),
		SpotDurationInMinutes: func(in interface{}) *int64 {
			if reflect.DeepEqual(in, reflect.Zero(reflect.TypeOf(in)).Interface()) {
				return nil
			}
			return func(in interface{}) *int64 {
				if in == nil {
					return nil
				}
				if _, ok := in.([]interface{}); ok && len(in.([]interface{})) == 0 {
					return nil
				}
				return func(in int64) *int64 {
					return &in
				}(int64(ExpandInt(in)))
			}(in)
		}(in["spot_duration_in_minutes"]),
		AssociatePublicIP: func(in interface{}) *bool {
			if reflect.DeepEqual(in, reflect.Zero(reflect.TypeOf(in)).Interface()) {
				return nil
			}
			return func(in interface{}) *bool {
				if in == nil {
					return nil
				}
				if _, ok := in.([]interface{}); ok && len(in.([]interface{})) == 0 {
					return nil
				}
				return func(in bool) *bool {
					return &in
				}(bool(ExpandBool(in)))
			}(in)
		}(in["associate_public_ip"]),
		AdditionalSecurityGroups: func(in interface{}) []string {
			return func(in interface{}) []string {
				var out []string
				for _, in := range in.([]interface{}) {
					out = append(out, string(ExpandString(in)))
				}
				return out
			}(in)
		}(in["additional_security_groups"]),
		CloudLabels: func(in interface{}) map[string]string {
			return func(in interface{}) map[string]string {
				if in == nil {
					return nil
				}
				out := map[string]string{}
				for key, in := range in.(map[string]interface{}) {
					out[key] = string(ExpandString(in))
				}
				return out
			}(in)
		}(in["cloud_labels"]),
		NodeLabels: func(in interface{}) map[string]string {
			return func(in interface{}) map[string]string {
				if in == nil {
					return nil
				}
				out := map[string]string{}
				for key, in := range in.(map[string]interface{}) {
					out[key] = string(ExpandString(in))
				}
				return out
			}(in)
		}(in["node_labels"]),
		FileAssets: func(in interface{}) []kops.FileAssetSpec {
			return func(in interface{}) []kops.FileAssetSpec {
				var out []kops.FileAssetSpec
				for _, in := range in.([]interface{}) {
					out = append(out, func(in interface{}) kops.FileAssetSpec {
						if in == nil {
							return kops.FileAssetSpec{}
						}
						return (ExpandFileAssetSpec(in.(map[string]interface{})))
					}(in))
				}
				return out
			}(in)
		}(in["file_assets"]),
		Tenancy: func(in interface{}) string {
			return string(ExpandString(in))
		}(in["tenancy"]),
		Kubelet: func(in interface{}) *kops.KubeletConfigSpec {
			return func(in interface{}) *kops.KubeletConfigSpec {
				if in == nil {
					return nil
				}
				if _, ok := in.([]interface{}); ok && len(in.([]interface{})) == 0 {
					return nil
				}
				return func(in kops.KubeletConfigSpec) *kops.KubeletConfigSpec {
					return &in
				}(func(in interface{}) kops.KubeletConfigSpec {
					if in.([]interface{})[0] == nil {
						return kops.KubeletConfigSpec{}
					}
					return (ExpandKubeletConfigSpec(in.([]interface{})[0].(map[string]interface{})))
				}(in))
			}(in)
		}(in["kubelet"]),
		Taints: func(in interface{}) []string {
			return func(in interface{}) []string {
				var out []string
				for _, in := range in.([]interface{}) {
					out = append(out, string(ExpandString(in)))
				}
				return out
			}(in)
		}(in["taints"]),
		MixedInstancesPolicy: func(in interface{}) *kops.MixedInstancesPolicySpec {
			return func(in interface{}) *kops.MixedInstancesPolicySpec {
				if in == nil {
					return nil
				}
				if _, ok := in.([]interface{}); ok && len(in.([]interface{})) == 0 {
					return nil
				}
				return func(in kops.MixedInstancesPolicySpec) *kops.MixedInstancesPolicySpec {
					return &in
				}(func(in interface{}) kops.MixedInstancesPolicySpec {
					if in.([]interface{})[0] == nil {
						return kops.MixedInstancesPolicySpec{}
					}
					return (ExpandMixedInstancesPolicySpec(in.([]interface{})[0].(map[string]interface{})))
				}(in))
			}(in)
		}(in["mixed_instances_policy"]),
		AdditionalUserData: func(in interface{}) []kops.UserData {
			return func(in interface{}) []kops.UserData {
				var out []kops.UserData
				for _, in := range in.([]interface{}) {
					out = append(out, func(in interface{}) kops.UserData {
						if in == nil {
							return kops.UserData{}
						}
						return (ExpandUserData(in.(map[string]interface{})))
					}(in))
				}
				return out
			}(in)
		}(in["additional_user_data"]),
		SuspendProcesses: func(in interface{}) []string {
			return func(in interface{}) []string {
				var out []string
				for _, in := range in.([]interface{}) {
					out = append(out, string(ExpandString(in)))
				}
				return out
			}(in)
		}(in["suspend_processes"]),
		ExternalLoadBalancers: func(in interface{}) []kops.LoadBalancer {
			return func(in interface{}) []kops.LoadBalancer {
				var out []kops.LoadBalancer
				for _, in := range in.([]interface{}) {
					out = append(out, func(in interface{}) kops.LoadBalancer {
						if in == nil {
							return kops.LoadBalancer{}
						}
						return (ExpandLoadBalancer(in.(map[string]interface{})))
					}(in))
				}
				return out
			}(in)
		}(in["external_load_balancers"]),
		DetailedInstanceMonitoring: func(in interface{}) *bool {
			if reflect.DeepEqual(in, reflect.Zero(reflect.TypeOf(in)).Interface()) {
				return nil
			}
			return func(in interface{}) *bool {
				if in == nil {
					return nil
				}
				if _, ok := in.([]interface{}); ok && len(in.([]interface{})) == 0 {
					return nil
				}
				return func(in bool) *bool {
					return &in
				}(bool(ExpandBool(in)))
			}(in)
		}(in["detailed_instance_monitoring"]),
		IAM: func(in interface{}) *kops.IAMProfileSpec {
			return func(in interface{}) *kops.IAMProfileSpec {
				if in == nil {
					return nil
				}
				if _, ok := in.([]interface{}); ok && len(in.([]interface{})) == 0 {
					return nil
				}
				return func(in kops.IAMProfileSpec) *kops.IAMProfileSpec {
					return &in
				}(func(in interface{}) kops.IAMProfileSpec {
					if in.([]interface{})[0] == nil {
						return kops.IAMProfileSpec{}
					}
					return (ExpandIAMProfileSpec(in.([]interface{})[0].(map[string]interface{})))
				}(in))
			}(in)
		}(in["iam"]),
		SecurityGroupOverride: func(in interface{}) *string {
			if reflect.DeepEqual(in, reflect.Zero(reflect.TypeOf(in)).Interface()) {
				return nil
			}
			return func(in interface{}) *string {
				if in == nil {
					return nil
				}
				if _, ok := in.([]interface{}); ok && len(in.([]interface{})) == 0 {
					return nil
				}
				return func(in string) *string {
					return &in
				}(string(ExpandString(in)))
			}(in)
		}(in["security_group_override"]),
		InstanceProtection: func(in interface{}) *bool {
			if reflect.DeepEqual(in, reflect.Zero(reflect.TypeOf(in)).Interface()) {
				return nil
			}
			return func(in interface{}) *bool {
				if in == nil {
					return nil
				}
				if _, ok := in.([]interface{}); ok && len(in.([]interface{})) == 0 {
					return nil
				}
				return func(in bool) *bool {
					return &in
				}(bool(ExpandBool(in)))
			}(in)
		}(in["instance_protection"]),
		SysctlParameters: func(in interface{}) []string {
			return func(in interface{}) []string {
				var out []string
				for _, in := range in.([]interface{}) {
					out = append(out, string(ExpandString(in)))
				}
				return out
			}(in)
		}(in["sysctl_parameters"]),
		RollingUpdate: func(in interface{}) *kops.RollingUpdate {
			return func(in interface{}) *kops.RollingUpdate {
				if in == nil {
					return nil
				}
				if _, ok := in.([]interface{}); ok && len(in.([]interface{})) == 0 {
					return nil
				}
				return func(in kops.RollingUpdate) *kops.RollingUpdate {
					return &in
				}(func(in interface{}) kops.RollingUpdate {
					if in.([]interface{})[0] == nil {
						return kops.RollingUpdate{}
					}
					return (ExpandRollingUpdate(in.([]interface{})[0].(map[string]interface{})))
				}(in))
			}(in)
		}(in["rolling_update"]),
		InstanceInterruptionBehavior: func(in interface{}) *string {
			if reflect.DeepEqual(in, reflect.Zero(reflect.TypeOf(in)).Interface()) {
				return nil
			}
			return func(in interface{}) *string {
				if in == nil {
					return nil
				}
				if _, ok := in.([]interface{}); ok && len(in.([]interface{})) == 0 {
					return nil
				}
				return func(in string) *string {
					return &in
				}(string(ExpandString(in)))
			}(in)
		}(in["instance_interruption_behavior"]),
	}
}

func FlattenInstanceGroup(in api.InstanceGroup) map[string]interface{} {
	return map[string]interface{}{
		"name": func(in string) interface{} {
			return FlattenString(string(in))
		}(in.Name),
		"role": func(in kops.InstanceGroupRole) interface{} {
			return FlattenString(string(in))
		}(in.Role),
		"image": func(in string) interface{} {
			return FlattenString(string(in))
		}(in.Image),
		"min_size": func(in *int32) interface{} {
			return func(in *int32) interface{} {
				if in == nil {
					return nil
				}
				return func(in int32) interface{} {
					return FlattenInt(int(in))
				}(*in)
			}(in)
		}(in.MinSize),
		"max_size": func(in *int32) interface{} {
			return func(in *int32) interface{} {
				if in == nil {
					return nil
				}
				return func(in int32) interface{} {
					return FlattenInt(int(in))
				}(*in)
			}(in)
		}(in.MaxSize),
		"machine_type": func(in string) interface{} {
			return FlattenString(string(in))
		}(in.MachineType),
		"root_volume_size": func(in *int32) interface{} {
			return func(in *int32) interface{} {
				if in == nil {
					return nil
				}
				return func(in int32) interface{} {
					return FlattenInt(int(in))
				}(*in)
			}(in)
		}(in.RootVolumeSize),
		"root_volume_type": func(in *string) interface{} {
			return func(in *string) interface{} {
				if in == nil {
					return nil
				}
				return func(in string) interface{} {
					return FlattenString(string(in))
				}(*in)
			}(in)
		}(in.RootVolumeType),
		"root_volume_iops": func(in *int32) interface{} {
			return func(in *int32) interface{} {
				if in == nil {
					return nil
				}
				return func(in int32) interface{} {
					return FlattenInt(int(in))
				}(*in)
			}(in)
		}(in.RootVolumeIops),
		"root_volume_optimization": func(in *bool) interface{} {
			return func(in *bool) interface{} {
				if in == nil {
					return nil
				}
				return func(in bool) interface{} {
					return FlattenBool(bool(in))
				}(*in)
			}(in)
		}(in.RootVolumeOptimization),
		"root_volume_delete_on_termination": func(in *bool) interface{} {
			return func(in *bool) interface{} {
				if in == nil {
					return nil
				}
				return func(in bool) interface{} {
					return FlattenBool(bool(in))
				}(*in)
			}(in)
		}(in.RootVolumeDeleteOnTermination),
		"root_volume_encryption": func(in *bool) interface{} {
			return func(in *bool) interface{} {
				if in == nil {
					return nil
				}
				return func(in bool) interface{} {
					return FlattenBool(bool(in))
				}(*in)
			}(in)
		}(in.RootVolumeEncryption),
		"volumes": func(in []*kops.VolumeSpec) interface{} {
			return func(in []*kops.VolumeSpec) []interface{} {
				var out []interface{}
				for _, in := range in {
					out = append(out, func(in *kops.VolumeSpec) interface{} {
						if in == nil {
							return nil
						}
						return func(in kops.VolumeSpec) interface{} {
							return func(in kops.VolumeSpec) interface{} {
								return FlattenVolumeSpec(in)
							}(in)
						}(*in)
					}(in))
				}
				return out
			}(in)
		}(in.Volumes),
		"volume_mounts": func(in []*kops.VolumeMountSpec) interface{} {
			return func(in []*kops.VolumeMountSpec) []interface{} {
				var out []interface{}
				for _, in := range in {
					out = append(out, func(in *kops.VolumeMountSpec) interface{} {
						if in == nil {
							return nil
						}
						return func(in kops.VolumeMountSpec) interface{} {
							return func(in kops.VolumeMountSpec) interface{} {
								return FlattenVolumeMountSpec(in)
							}(in)
						}(*in)
					}(in))
				}
				return out
			}(in)
		}(in.VolumeMounts),
		"subnets": func(in []string) interface{} {
			return func(in []string) []interface{} {
				var out []interface{}
				for _, in := range in {
					out = append(out, FlattenString(string(in)))
				}
				return out
			}(in)
		}(in.Subnets),
		"zones": func(in []string) interface{} {
			return func(in []string) []interface{} {
				var out []interface{}
				for _, in := range in {
					out = append(out, FlattenString(string(in)))
				}
				return out
			}(in)
		}(in.Zones),
		"hooks": func(in []kops.HookSpec) interface{} {
			return func(in []kops.HookSpec) []interface{} {
				var out []interface{}
				for _, in := range in {
					out = append(out, func(in kops.HookSpec) interface{} {
						return FlattenHookSpec(in)
					}(in))
				}
				return out
			}(in)
		}(in.Hooks),
		"max_price": func(in *string) interface{} {
			return func(in *string) interface{} {
				if in == nil {
					return nil
				}
				return func(in string) interface{} {
					return FlattenString(string(in))
				}(*in)
			}(in)
		}(in.MaxPrice),
		"spot_duration_in_minutes": func(in *int64) interface{} {
			return func(in *int64) interface{} {
				if in == nil {
					return nil
				}
				return func(in int64) interface{} {
					return FlattenInt(int(in))
				}(*in)
			}(in)
		}(in.SpotDurationInMinutes),
		"associate_public_ip": func(in *bool) interface{} {
			return func(in *bool) interface{} {
				if in == nil {
					return nil
				}
				return func(in bool) interface{} {
					return FlattenBool(bool(in))
				}(*in)
			}(in)
		}(in.AssociatePublicIP),
		"additional_security_groups": func(in []string) interface{} {
			return func(in []string) []interface{} {
				var out []interface{}
				for _, in := range in {
					out = append(out, FlattenString(string(in)))
				}
				return out
			}(in)
		}(in.AdditionalSecurityGroups),
		"cloud_labels": func(in map[string]string) interface{} {
			return func(in map[string]string) map[string]interface{} {
				if in == nil {
					return nil
				}
				// TODO
				return nil
			}(in)
		}(in.CloudLabels),
		"node_labels": func(in map[string]string) interface{} {
			return func(in map[string]string) map[string]interface{} {
				if in == nil {
					return nil
				}
				// TODO
				return nil
			}(in)
		}(in.NodeLabels),
		"file_assets": func(in []kops.FileAssetSpec) interface{} {
			return func(in []kops.FileAssetSpec) []interface{} {
				var out []interface{}
				for _, in := range in {
					out = append(out, func(in kops.FileAssetSpec) interface{} {
						return FlattenFileAssetSpec(in)
					}(in))
				}
				return out
			}(in)
		}(in.FileAssets),
		"tenancy": func(in string) interface{} {
			return FlattenString(string(in))
		}(in.Tenancy),
		"kubelet": func(in *kops.KubeletConfigSpec) interface{} {
			return func(in *kops.KubeletConfigSpec) interface{} {
				if in == nil {
					return nil
				}
				return func(in kops.KubeletConfigSpec) interface{} {
					return func(in kops.KubeletConfigSpec) []map[string]interface{} {
						return []map[string]interface{}{FlattenKubeletConfigSpec(in)}
					}(in)
				}(*in)
			}(in)
		}(in.Kubelet),
		"taints": func(in []string) interface{} {
			return func(in []string) []interface{} {
				var out []interface{}
				for _, in := range in {
					out = append(out, FlattenString(string(in)))
				}
				return out
			}(in)
		}(in.Taints),
		"mixed_instances_policy": func(in *kops.MixedInstancesPolicySpec) interface{} {
			return func(in *kops.MixedInstancesPolicySpec) interface{} {
				if in == nil {
					return nil
				}
				return func(in kops.MixedInstancesPolicySpec) interface{} {
					return func(in kops.MixedInstancesPolicySpec) []map[string]interface{} {
						return []map[string]interface{}{FlattenMixedInstancesPolicySpec(in)}
					}(in)
				}(*in)
			}(in)
		}(in.MixedInstancesPolicy),
		"additional_user_data": func(in []kops.UserData) interface{} {
			return func(in []kops.UserData) []interface{} {
				var out []interface{}
				for _, in := range in {
					out = append(out, func(in kops.UserData) interface{} {
						return FlattenUserData(in)
					}(in))
				}
				return out
			}(in)
		}(in.AdditionalUserData),
		"suspend_processes": func(in []string) interface{} {
			return func(in []string) []interface{} {
				var out []interface{}
				for _, in := range in {
					out = append(out, FlattenString(string(in)))
				}
				return out
			}(in)
		}(in.SuspendProcesses),
		"external_load_balancers": func(in []kops.LoadBalancer) interface{} {
			return func(in []kops.LoadBalancer) []interface{} {
				var out []interface{}
				for _, in := range in {
					out = append(out, func(in kops.LoadBalancer) interface{} {
						return FlattenLoadBalancer(in)
					}(in))
				}
				return out
			}(in)
		}(in.ExternalLoadBalancers),
		"detailed_instance_monitoring": func(in *bool) interface{} {
			return func(in *bool) interface{} {
				if in == nil {
					return nil
				}
				return func(in bool) interface{} {
					return FlattenBool(bool(in))
				}(*in)
			}(in)
		}(in.DetailedInstanceMonitoring),
		"iam": func(in *kops.IAMProfileSpec) interface{} {
			return func(in *kops.IAMProfileSpec) interface{} {
				if in == nil {
					return nil
				}
				return func(in kops.IAMProfileSpec) interface{} {
					return func(in kops.IAMProfileSpec) []map[string]interface{} {
						return []map[string]interface{}{FlattenIAMProfileSpec(in)}
					}(in)
				}(*in)
			}(in)
		}(in.IAM),
		"security_group_override": func(in *string) interface{} {
			return func(in *string) interface{} {
				if in == nil {
					return nil
				}
				return func(in string) interface{} {
					return FlattenString(string(in))
				}(*in)
			}(in)
		}(in.SecurityGroupOverride),
		"instance_protection": func(in *bool) interface{} {
			return func(in *bool) interface{} {
				if in == nil {
					return nil
				}
				return func(in bool) interface{} {
					return FlattenBool(bool(in))
				}(*in)
			}(in)
		}(in.InstanceProtection),
		"sysctl_parameters": func(in []string) interface{} {
			return func(in []string) []interface{} {
				var out []interface{}
				for _, in := range in {
					out = append(out, FlattenString(string(in)))
				}
				return out
			}(in)
		}(in.SysctlParameters),
		"rolling_update": func(in *kops.RollingUpdate) interface{} {
			return func(in *kops.RollingUpdate) interface{} {
				if in == nil {
					return nil
				}
				return func(in kops.RollingUpdate) interface{} {
					return func(in kops.RollingUpdate) []map[string]interface{} {
						return []map[string]interface{}{FlattenRollingUpdate(in)}
					}(in)
				}(*in)
			}(in)
		}(in.RollingUpdate),
		"instance_interruption_behavior": func(in *string) interface{} {
			return func(in *string) interface{} {
				if in == nil {
					return nil
				}
				return func(in string) interface{} {
					return FlattenString(string(in))
				}(*in)
			}(in)
		}(in.InstanceInterruptionBehavior),
	}
}
