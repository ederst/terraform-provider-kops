package structures

import (
	"reflect"

	"github.com/eddycharly/terraform-provider-kops/pkg/api"
	"k8s.io/kops/pkg/apis/kops"
)

func ExpandCluster(in map[string]interface{}) api.Cluster {
	if in == nil {
		panic("expand Cluster failure, in is nil")
	}
	return api.Cluster{
		Name: func(in interface{}) string {
			return string(ExpandString(in))
		}(in["name"]),
		Channel: func(in interface{}) string {
			return string(ExpandString(in))
		}(in["channel"]),
		Addons: func(in interface{}) []kops.AddonSpec {
			return func(in interface{}) []kops.AddonSpec {
				var out []kops.AddonSpec
				for _, in := range in.([]interface{}) {
					out = append(out, func(in interface{}) kops.AddonSpec {
						if in == nil {
							return kops.AddonSpec{}
						}
						return (ExpandAddonSpec(in.(map[string]interface{})))
					}(in))
				}
				return out
			}(in)
		}(in["addons"]),
		ConfigBase: func(in interface{}) string {
			return string(ExpandString(in))
		}(in["config_base"]),
		CloudProvider: func(in interface{}) string {
			return string(ExpandString(in))
		}(in["cloud_provider"]),
		ContainerRuntime: func(in interface{}) string {
			return string(ExpandString(in))
		}(in["container_runtime"]),
		KubernetesVersion: func(in interface{}) string {
			return string(ExpandString(in))
		}(in["kubernetes_version"]),
		Subnet: func(in interface{}) []kops.ClusterSubnetSpec {
			return func(in interface{}) []kops.ClusterSubnetSpec {
				var out []kops.ClusterSubnetSpec
				for _, in := range in.([]interface{}) {
					out = append(out, func(in interface{}) kops.ClusterSubnetSpec {
						if in == nil {
							return kops.ClusterSubnetSpec{}
						}
						return (ExpandClusterSubnetSpec(in.(map[string]interface{})))
					}(in))
				}
				return out
			}(in)
		}(in["subnet"]),
		Project: func(in interface{}) string {
			return string(ExpandString(in))
		}(in["project"]),
		MasterPublicName: func(in interface{}) string {
			return string(ExpandString(in))
		}(in["master_public_name"]),
		MasterInternalName: func(in interface{}) string {
			return string(ExpandString(in))
		}(in["master_internal_name"]),
		NetworkCIDR: func(in interface{}) string {
			return string(ExpandString(in))
		}(in["network_cidr"]),
		AdditionalNetworkCIDRs: func(in interface{}) []string {
			return func(in interface{}) []string {
				var out []string
				for _, in := range in.([]interface{}) {
					out = append(out, string(ExpandString(in)))
				}
				return out
			}(in)
		}(in["additional_network_cidrs"]),
		NetworkID: func(in interface{}) string {
			return string(ExpandString(in))
		}(in["network_id"]),
		Topology: func(in interface{}) *kops.TopologySpec {
			return func(in interface{}) *kops.TopologySpec {
				if in == nil {
					return nil
				}
				if _, ok := in.([]interface{}); ok && len(in.([]interface{})) == 0 {
					return nil
				}
				return func(in kops.TopologySpec) *kops.TopologySpec {
					return &in
				}(func(in interface{}) kops.TopologySpec {
					if in.([]interface{})[0] == nil {
						return kops.TopologySpec{}
					}
					return (ExpandTopologySpec(in.([]interface{})[0].(map[string]interface{})))
				}(in))
			}(in)
		}(in["topology"]),
		SecretStore: func(in interface{}) string {
			return string(ExpandString(in))
		}(in["secret_store"]),
		KeyStore: func(in interface{}) string {
			return string(ExpandString(in))
		}(in["key_store"]),
		ConfigStore: func(in interface{}) string {
			return string(ExpandString(in))
		}(in["config_store"]),
		DNSZone: func(in interface{}) string {
			return string(ExpandString(in))
		}(in["dns_zone"]),
		AdditionalSANs: func(in interface{}) []string {
			return func(in interface{}) []string {
				var out []string
				for _, in := range in.([]interface{}) {
					out = append(out, string(ExpandString(in)))
				}
				return out
			}(in)
		}(in["additional_sans"]),
		ClusterDNSDomain: func(in interface{}) string {
			return string(ExpandString(in))
		}(in["cluster_dns_domain"]),
		ServiceClusterIPRange: func(in interface{}) string {
			return string(ExpandString(in))
		}(in["service_cluster_ip_range"]),
		PodCIDR: func(in interface{}) string {
			return string(ExpandString(in))
		}(in["pod_cidr"]),
		NonMasqueradeCIDR: func(in interface{}) string {
			return string(ExpandString(in))
		}(in["non_masquerade_cidr"]),
		SSHAccess: func(in interface{}) []string {
			return func(in interface{}) []string {
				var out []string
				for _, in := range in.([]interface{}) {
					out = append(out, string(ExpandString(in)))
				}
				return out
			}(in)
		}(in["ssh_access"]),
		NodePortAccess: func(in interface{}) []string {
			return func(in interface{}) []string {
				var out []string
				for _, in := range in.([]interface{}) {
					out = append(out, string(ExpandString(in)))
				}
				return out
			}(in)
		}(in["node_port_access"]),
		EgressProxy: func(in interface{}) *kops.EgressProxySpec {
			return func(in interface{}) *kops.EgressProxySpec {
				if in == nil {
					return nil
				}
				if _, ok := in.([]interface{}); ok && len(in.([]interface{})) == 0 {
					return nil
				}
				return func(in kops.EgressProxySpec) *kops.EgressProxySpec {
					return &in
				}(func(in interface{}) kops.EgressProxySpec {
					if in.([]interface{})[0] == nil {
						return kops.EgressProxySpec{}
					}
					return (ExpandEgressProxySpec(in.([]interface{})[0].(map[string]interface{})))
				}(in))
			}(in)
		}(in["egress_proxy"]),
		SSHKeyName: func(in interface{}) *string {
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
		}(in["ssh_key_name"]),
		KubernetesAPIAccess: func(in interface{}) []string {
			return func(in interface{}) []string {
				var out []string
				for _, in := range in.([]interface{}) {
					out = append(out, string(ExpandString(in)))
				}
				return out
			}(in)
		}(in["kubernetes_api_access"]),
		IsolateMasters: func(in interface{}) *bool {
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
		}(in["isolate_masters"]),
		UpdatePolicy: func(in interface{}) *string {
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
		}(in["update_policy"]),
		ExternalPolicies: func(in interface{}) *map[string][]string {
			return func(in interface{}) *map[string][]string {
				if in == nil {
					return nil
				}
				if _, ok := in.([]interface{}); ok && len(in.([]interface{})) == 0 {
					return nil
				}
				return func(in map[string][]string) *map[string][]string {
					return &in
				}(func(in interface{}) map[string][]string {
					if in == nil {
						return nil
					}
					out := map[string][]string{}
					for key, in := range in.(map[string]interface{}) {
						out[key] = func(in interface{}) []string {
							var out []string
							for _, in := range in.([]interface{}) {
								out = append(out, string(ExpandString(in)))
							}
							return out
						}(in)
					}
					return out
				}(in))
			}(in)
		}(in["external_policies"]),
		AdditionalPolicies: func(in interface{}) *map[string]string {
			return func(in interface{}) *map[string]string {
				if in == nil {
					return nil
				}
				if _, ok := in.([]interface{}); ok && len(in.([]interface{})) == 0 {
					return nil
				}
				return func(in map[string]string) *map[string]string {
					return &in
				}(func(in interface{}) map[string]string {
					if in == nil {
						return nil
					}
					out := map[string]string{}
					for key, in := range in.(map[string]interface{}) {
						out[key] = string(ExpandString(in))
					}
					return out
				}(in))
			}(in)
		}(in["additional_policies"]),
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
		EtcdCluster: func(in interface{}) []*kops.EtcdClusterSpec {
			return func(in interface{}) []*kops.EtcdClusterSpec {
				var out []*kops.EtcdClusterSpec
				for _, in := range in.([]interface{}) {
					out = append(out, func(in interface{}) *kops.EtcdClusterSpec {
						if in == nil {
							return nil
						}
						if _, ok := in.([]interface{}); ok && len(in.([]interface{})) == 0 {
							return nil
						}
						return func(in kops.EtcdClusterSpec) *kops.EtcdClusterSpec {
							return &in
						}(func(in interface{}) kops.EtcdClusterSpec {
							if in == nil {
								return kops.EtcdClusterSpec{}
							}
							return (ExpandEtcdClusterSpec(in.(map[string]interface{})))
						}(in))
					}(in))
				}
				return out
			}(in)
		}(in["etcd_cluster"]),
		Containerd: func(in interface{}) *kops.ContainerdConfig {
			return func(in interface{}) *kops.ContainerdConfig {
				if in == nil {
					return nil
				}
				if _, ok := in.([]interface{}); ok && len(in.([]interface{})) == 0 {
					return nil
				}
				return func(in kops.ContainerdConfig) *kops.ContainerdConfig {
					return &in
				}(func(in interface{}) kops.ContainerdConfig {
					if in.([]interface{})[0] == nil {
						return kops.ContainerdConfig{}
					}
					return (ExpandContainerdConfig(in.([]interface{})[0].(map[string]interface{})))
				}(in))
			}(in)
		}(in["containerd"]),
		Docker: func(in interface{}) *kops.DockerConfig {
			return func(in interface{}) *kops.DockerConfig {
				if in == nil {
					return nil
				}
				if _, ok := in.([]interface{}); ok && len(in.([]interface{})) == 0 {
					return nil
				}
				return func(in kops.DockerConfig) *kops.DockerConfig {
					return &in
				}(func(in interface{}) kops.DockerConfig {
					if in.([]interface{})[0] == nil {
						return kops.DockerConfig{}
					}
					return (ExpandDockerConfig(in.([]interface{})[0].(map[string]interface{})))
				}(in))
			}(in)
		}(in["docker"]),
		KubeDNS: func(in interface{}) *kops.KubeDNSConfig {
			return func(in interface{}) *kops.KubeDNSConfig {
				if in == nil {
					return nil
				}
				if _, ok := in.([]interface{}); ok && len(in.([]interface{})) == 0 {
					return nil
				}
				return func(in kops.KubeDNSConfig) *kops.KubeDNSConfig {
					return &in
				}(func(in interface{}) kops.KubeDNSConfig {
					if in.([]interface{})[0] == nil {
						return kops.KubeDNSConfig{}
					}
					return (ExpandKubeDNSConfig(in.([]interface{})[0].(map[string]interface{})))
				}(in))
			}(in)
		}(in["kube_dns"]),
		KubeAPIServer: func(in interface{}) *kops.KubeAPIServerConfig {
			return func(in interface{}) *kops.KubeAPIServerConfig {
				if in == nil {
					return nil
				}
				if _, ok := in.([]interface{}); ok && len(in.([]interface{})) == 0 {
					return nil
				}
				return func(in kops.KubeAPIServerConfig) *kops.KubeAPIServerConfig {
					return &in
				}(func(in interface{}) kops.KubeAPIServerConfig {
					if in.([]interface{})[0] == nil {
						return kops.KubeAPIServerConfig{}
					}
					return (ExpandKubeAPIServerConfig(in.([]interface{})[0].(map[string]interface{})))
				}(in))
			}(in)
		}(in["kube_api_server"]),
		KubeControllerManager: func(in interface{}) *kops.KubeControllerManagerConfig {
			return func(in interface{}) *kops.KubeControllerManagerConfig {
				if in == nil {
					return nil
				}
				if _, ok := in.([]interface{}); ok && len(in.([]interface{})) == 0 {
					return nil
				}
				return func(in kops.KubeControllerManagerConfig) *kops.KubeControllerManagerConfig {
					return &in
				}(func(in interface{}) kops.KubeControllerManagerConfig {
					if in.([]interface{})[0] == nil {
						return kops.KubeControllerManagerConfig{}
					}
					return (ExpandKubeControllerManagerConfig(in.([]interface{})[0].(map[string]interface{})))
				}(in))
			}(in)
		}(in["kube_controller_manager"]),
		ExternalCloudControllerManager: func(in interface{}) *kops.CloudControllerManagerConfig {
			return func(in interface{}) *kops.CloudControllerManagerConfig {
				if in == nil {
					return nil
				}
				if _, ok := in.([]interface{}); ok && len(in.([]interface{})) == 0 {
					return nil
				}
				return func(in kops.CloudControllerManagerConfig) *kops.CloudControllerManagerConfig {
					return &in
				}(func(in interface{}) kops.CloudControllerManagerConfig {
					if in.([]interface{})[0] == nil {
						return kops.CloudControllerManagerConfig{}
					}
					return (ExpandCloudControllerManagerConfig(in.([]interface{})[0].(map[string]interface{})))
				}(in))
			}(in)
		}(in["external_cloud_controller_manager"]),
		KubeScheduler: func(in interface{}) *kops.KubeSchedulerConfig {
			return func(in interface{}) *kops.KubeSchedulerConfig {
				if in == nil {
					return nil
				}
				if _, ok := in.([]interface{}); ok && len(in.([]interface{})) == 0 {
					return nil
				}
				return func(in kops.KubeSchedulerConfig) *kops.KubeSchedulerConfig {
					return &in
				}(func(in interface{}) kops.KubeSchedulerConfig {
					if in.([]interface{})[0] == nil {
						return kops.KubeSchedulerConfig{}
					}
					return (ExpandKubeSchedulerConfig(in.([]interface{})[0].(map[string]interface{})))
				}(in))
			}(in)
		}(in["kube_scheduler"]),
		KubeProxy: func(in interface{}) *kops.KubeProxyConfig {
			return func(in interface{}) *kops.KubeProxyConfig {
				if in == nil {
					return nil
				}
				if _, ok := in.([]interface{}); ok && len(in.([]interface{})) == 0 {
					return nil
				}
				return func(in kops.KubeProxyConfig) *kops.KubeProxyConfig {
					return &in
				}(func(in interface{}) kops.KubeProxyConfig {
					if in.([]interface{})[0] == nil {
						return kops.KubeProxyConfig{}
					}
					return (ExpandKubeProxyConfig(in.([]interface{})[0].(map[string]interface{})))
				}(in))
			}(in)
		}(in["kube_proxy"]),
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
		MasterKubelet: func(in interface{}) *kops.KubeletConfigSpec {
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
		}(in["master_kubelet"]),
		CloudConfig: func(in interface{}) *kops.CloudConfiguration {
			return func(in interface{}) *kops.CloudConfiguration {
				if in == nil {
					return nil
				}
				if _, ok := in.([]interface{}); ok && len(in.([]interface{})) == 0 {
					return nil
				}
				return func(in kops.CloudConfiguration) *kops.CloudConfiguration {
					return &in
				}(func(in interface{}) kops.CloudConfiguration {
					if in.([]interface{})[0] == nil {
						return kops.CloudConfiguration{}
					}
					return (ExpandCloudConfiguration(in.([]interface{})[0].(map[string]interface{})))
				}(in))
			}(in)
		}(in["cloud_config"]),
		ExternalDNS: func(in interface{}) *kops.ExternalDNSConfig {
			return func(in interface{}) *kops.ExternalDNSConfig {
				if in == nil {
					return nil
				}
				if _, ok := in.([]interface{}); ok && len(in.([]interface{})) == 0 {
					return nil
				}
				return func(in kops.ExternalDNSConfig) *kops.ExternalDNSConfig {
					return &in
				}(func(in interface{}) kops.ExternalDNSConfig {
					if in.([]interface{})[0] == nil {
						return kops.ExternalDNSConfig{}
					}
					return (ExpandExternalDNSConfig(in.([]interface{})[0].(map[string]interface{})))
				}(in))
			}(in)
		}(in["external_dns"]),
		Networking: func(in interface{}) *kops.NetworkingSpec {
			return func(in interface{}) *kops.NetworkingSpec {
				if in == nil {
					return nil
				}
				if _, ok := in.([]interface{}); ok && len(in.([]interface{})) == 0 {
					return nil
				}
				return func(in kops.NetworkingSpec) *kops.NetworkingSpec {
					return &in
				}(func(in interface{}) kops.NetworkingSpec {
					if in.([]interface{})[0] == nil {
						return kops.NetworkingSpec{}
					}
					return (ExpandNetworkingSpec(in.([]interface{})[0].(map[string]interface{})))
				}(in))
			}(in)
		}(in["networking"]),
		API: func(in interface{}) *kops.AccessSpec {
			return func(in interface{}) *kops.AccessSpec {
				if in == nil {
					return nil
				}
				if _, ok := in.([]interface{}); ok && len(in.([]interface{})) == 0 {
					return nil
				}
				return func(in kops.AccessSpec) *kops.AccessSpec {
					return &in
				}(func(in interface{}) kops.AccessSpec {
					if in.([]interface{})[0] == nil {
						return kops.AccessSpec{}
					}
					return (ExpandAccessSpec(in.([]interface{})[0].(map[string]interface{})))
				}(in))
			}(in)
		}(in["api"]),
		Authentication: func(in interface{}) *kops.AuthenticationSpec {
			return func(in interface{}) *kops.AuthenticationSpec {
				if in == nil {
					return nil
				}
				if _, ok := in.([]interface{}); ok && len(in.([]interface{})) == 0 {
					return nil
				}
				return func(in kops.AuthenticationSpec) *kops.AuthenticationSpec {
					return &in
				}(func(in interface{}) kops.AuthenticationSpec {
					if in.([]interface{})[0] == nil {
						return kops.AuthenticationSpec{}
					}
					return (ExpandAuthenticationSpec(in.([]interface{})[0].(map[string]interface{})))
				}(in))
			}(in)
		}(in["authentication"]),
		Authorization: func(in interface{}) *kops.AuthorizationSpec {
			return func(in interface{}) *kops.AuthorizationSpec {
				if in == nil {
					return nil
				}
				if _, ok := in.([]interface{}); ok && len(in.([]interface{})) == 0 {
					return nil
				}
				return func(in kops.AuthorizationSpec) *kops.AuthorizationSpec {
					return &in
				}(func(in interface{}) kops.AuthorizationSpec {
					if in.([]interface{})[0] == nil {
						return kops.AuthorizationSpec{}
					}
					return (ExpandAuthorizationSpec(in.([]interface{})[0].(map[string]interface{})))
				}(in))
			}(in)
		}(in["authorization"]),
		NodeAuthorization: func(in interface{}) *kops.NodeAuthorizationSpec {
			return func(in interface{}) *kops.NodeAuthorizationSpec {
				if in == nil {
					return nil
				}
				if _, ok := in.([]interface{}); ok && len(in.([]interface{})) == 0 {
					return nil
				}
				return func(in kops.NodeAuthorizationSpec) *kops.NodeAuthorizationSpec {
					return &in
				}(func(in interface{}) kops.NodeAuthorizationSpec {
					if in.([]interface{})[0] == nil {
						return kops.NodeAuthorizationSpec{}
					}
					return (ExpandNodeAuthorizationSpec(in.([]interface{})[0].(map[string]interface{})))
				}(in))
			}(in)
		}(in["node_authorization"]),
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
		Assets: func(in interface{}) *kops.Assets {
			return func(in interface{}) *kops.Assets {
				if in == nil {
					return nil
				}
				if _, ok := in.([]interface{}); ok && len(in.([]interface{})) == 0 {
					return nil
				}
				return func(in kops.Assets) *kops.Assets {
					return &in
				}(func(in interface{}) kops.Assets {
					if in.([]interface{})[0] == nil {
						return kops.Assets{}
					}
					return (ExpandAssets(in.([]interface{})[0].(map[string]interface{})))
				}(in))
			}(in)
		}(in["assets"]),
		IAM: func(in interface{}) *kops.IAMSpec {
			return func(in interface{}) *kops.IAMSpec {
				if in == nil {
					return nil
				}
				if _, ok := in.([]interface{}); ok && len(in.([]interface{})) == 0 {
					return nil
				}
				return func(in kops.IAMSpec) *kops.IAMSpec {
					return &in
				}(func(in interface{}) kops.IAMSpec {
					if in.([]interface{})[0] == nil {
						return kops.IAMSpec{}
					}
					return (ExpandIAMSpec(in.([]interface{})[0].(map[string]interface{})))
				}(in))
			}(in)
		}(in["iam"]),
		EncryptionConfig: func(in interface{}) *bool {
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
		}(in["encryption_config"]),
		DisableSubnetTags: func(in interface{}) bool {
			return bool(ExpandBool(in))
		}(in["disable_subnet_tags"]),
		UseHostCertificates: func(in interface{}) *bool {
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
		}(in["use_host_certificates"]),
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
		KubeServer: func(in interface{}) string {
			return string(ExpandString(in))
		}(in["kube_server"]),
		KubeCertificateAuthority: func(in interface{}) string {
			return string(ExpandString(in))
		}(in["kube_certificate_authority"]),
		KubeClientCertificate: func(in interface{}) string {
			return string(ExpandString(in))
		}(in["kube_client_certificate"]),
		KubeClientKey: func(in interface{}) string {
			return string(ExpandString(in))
		}(in["kube_client_key"]),
		KubeUsername: func(in interface{}) string {
			return string(ExpandString(in))
		}(in["kube_username"]),
		KubePassword: func(in interface{}) string {
			return string(ExpandString(in))
		}(in["kube_password"]),
		InstanceGroup: func(in interface{}) []*api.InstanceGroup {
			return func(in interface{}) []*api.InstanceGroup {
				var out []*api.InstanceGroup
				for _, in := range in.([]interface{}) {
					out = append(out, func(in interface{}) *api.InstanceGroup {
						if in == nil {
							return nil
						}
						if _, ok := in.([]interface{}); ok && len(in.([]interface{})) == 0 {
							return nil
						}
						return func(in api.InstanceGroup) *api.InstanceGroup {
							return &in
						}(func(in interface{}) api.InstanceGroup {
							if in == nil {
								return api.InstanceGroup{}
							}
							return (ExpandInstanceGroup(in.(map[string]interface{})))
						}(in))
					}(in))
				}
				return out
			}(in)
		}(in["instance_group"]),
	}
}

func FlattenCluster(in api.Cluster) map[string]interface{} {
	return map[string]interface{}{
		"name": func(in string) interface{} {
			return FlattenString(string(in))
		}(in.Name),
		"channel": func(in string) interface{} {
			return FlattenString(string(in))
		}(in.Channel),
		"addons": func(in []kops.AddonSpec) interface{} {
			return func(in []kops.AddonSpec) []interface{} {
				var out []interface{}
				for _, in := range in {
					out = append(out, func(in kops.AddonSpec) interface{} {
						return FlattenAddonSpec(in)
					}(in))
				}
				return out
			}(in)
		}(in.Addons),
		"config_base": func(in string) interface{} {
			return FlattenString(string(in))
		}(in.ConfigBase),
		"cloud_provider": func(in string) interface{} {
			return FlattenString(string(in))
		}(in.CloudProvider),
		"container_runtime": func(in string) interface{} {
			return FlattenString(string(in))
		}(in.ContainerRuntime),
		"kubernetes_version": func(in string) interface{} {
			return FlattenString(string(in))
		}(in.KubernetesVersion),
		"subnet": func(in []kops.ClusterSubnetSpec) interface{} {
			return func(in []kops.ClusterSubnetSpec) []interface{} {
				var out []interface{}
				for _, in := range in {
					out = append(out, func(in kops.ClusterSubnetSpec) interface{} {
						return FlattenClusterSubnetSpec(in)
					}(in))
				}
				return out
			}(in)
		}(in.Subnet),
		"project": func(in string) interface{} {
			return FlattenString(string(in))
		}(in.Project),
		"master_public_name": func(in string) interface{} {
			return FlattenString(string(in))
		}(in.MasterPublicName),
		"master_internal_name": func(in string) interface{} {
			return FlattenString(string(in))
		}(in.MasterInternalName),
		"network_cidr": func(in string) interface{} {
			return FlattenString(string(in))
		}(in.NetworkCIDR),
		"additional_network_cidrs": func(in []string) interface{} {
			return func(in []string) []interface{} {
				var out []interface{}
				for _, in := range in {
					out = append(out, FlattenString(string(in)))
				}
				return out
			}(in)
		}(in.AdditionalNetworkCIDRs),
		"network_id": func(in string) interface{} {
			return FlattenString(string(in))
		}(in.NetworkID),
		"topology": func(in *kops.TopologySpec) interface{} {
			return func(in *kops.TopologySpec) interface{} {
				if in == nil {
					return nil
				}
				return func(in kops.TopologySpec) interface{} {
					return func(in kops.TopologySpec) []map[string]interface{} {
						return []map[string]interface{}{FlattenTopologySpec(in)}
					}(in)
				}(*in)
			}(in)
		}(in.Topology),
		"secret_store": func(in string) interface{} {
			return FlattenString(string(in))
		}(in.SecretStore),
		"key_store": func(in string) interface{} {
			return FlattenString(string(in))
		}(in.KeyStore),
		"config_store": func(in string) interface{} {
			return FlattenString(string(in))
		}(in.ConfigStore),
		"dns_zone": func(in string) interface{} {
			return FlattenString(string(in))
		}(in.DNSZone),
		"additional_sans": func(in []string) interface{} {
			return func(in []string) []interface{} {
				var out []interface{}
				for _, in := range in {
					out = append(out, FlattenString(string(in)))
				}
				return out
			}(in)
		}(in.AdditionalSANs),
		"cluster_dns_domain": func(in string) interface{} {
			return FlattenString(string(in))
		}(in.ClusterDNSDomain),
		"service_cluster_ip_range": func(in string) interface{} {
			return FlattenString(string(in))
		}(in.ServiceClusterIPRange),
		"pod_cidr": func(in string) interface{} {
			return FlattenString(string(in))
		}(in.PodCIDR),
		"non_masquerade_cidr": func(in string) interface{} {
			return FlattenString(string(in))
		}(in.NonMasqueradeCIDR),
		"ssh_access": func(in []string) interface{} {
			return func(in []string) []interface{} {
				var out []interface{}
				for _, in := range in {
					out = append(out, FlattenString(string(in)))
				}
				return out
			}(in)
		}(in.SSHAccess),
		"node_port_access": func(in []string) interface{} {
			return func(in []string) []interface{} {
				var out []interface{}
				for _, in := range in {
					out = append(out, FlattenString(string(in)))
				}
				return out
			}(in)
		}(in.NodePortAccess),
		"egress_proxy": func(in *kops.EgressProxySpec) interface{} {
			return func(in *kops.EgressProxySpec) interface{} {
				if in == nil {
					return nil
				}
				return func(in kops.EgressProxySpec) interface{} {
					return func(in kops.EgressProxySpec) []map[string]interface{} {
						return []map[string]interface{}{FlattenEgressProxySpec(in)}
					}(in)
				}(*in)
			}(in)
		}(in.EgressProxy),
		"ssh_key_name": func(in *string) interface{} {
			return func(in *string) interface{} {
				if in == nil {
					return nil
				}
				return func(in string) interface{} {
					return FlattenString(string(in))
				}(*in)
			}(in)
		}(in.SSHKeyName),
		"kubernetes_api_access": func(in []string) interface{} {
			return func(in []string) []interface{} {
				var out []interface{}
				for _, in := range in {
					out = append(out, FlattenString(string(in)))
				}
				return out
			}(in)
		}(in.KubernetesAPIAccess),
		"isolate_masters": func(in *bool) interface{} {
			return func(in *bool) interface{} {
				if in == nil {
					return nil
				}
				return func(in bool) interface{} {
					return FlattenBool(bool(in))
				}(*in)
			}(in)
		}(in.IsolateMasters),
		"update_policy": func(in *string) interface{} {
			return func(in *string) interface{} {
				if in == nil {
					return nil
				}
				return func(in string) interface{} {
					return FlattenString(string(in))
				}(*in)
			}(in)
		}(in.UpdatePolicy),
		"external_policies": func(in *map[string][]string) interface{} {
			return func(in *map[string][]string) interface{} {
				if in == nil {
					return nil
				}
				return func(in map[string][]string) interface{} {
					return func(in map[string][]string) map[string]interface{} {
						if in == nil {
							return nil
						}
						// TODO
						return nil
					}(in)
				}(*in)
			}(in)
		}(in.ExternalPolicies),
		"additional_policies": func(in *map[string]string) interface{} {
			return func(in *map[string]string) interface{} {
				if in == nil {
					return nil
				}
				return func(in map[string]string) interface{} {
					return func(in map[string]string) map[string]interface{} {
						if in == nil {
							return nil
						}
						// TODO
						return nil
					}(in)
				}(*in)
			}(in)
		}(in.AdditionalPolicies),
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
		"etcd_cluster": func(in []*kops.EtcdClusterSpec) interface{} {
			return func(in []*kops.EtcdClusterSpec) []interface{} {
				var out []interface{}
				for _, in := range in {
					out = append(out, func(in *kops.EtcdClusterSpec) interface{} {
						if in == nil {
							return nil
						}
						return func(in kops.EtcdClusterSpec) interface{} {
							return func(in kops.EtcdClusterSpec) interface{} {
								return FlattenEtcdClusterSpec(in)
							}(in)
						}(*in)
					}(in))
				}
				return out
			}(in)
		}(in.EtcdCluster),
		"containerd": func(in *kops.ContainerdConfig) interface{} {
			return func(in *kops.ContainerdConfig) interface{} {
				if in == nil {
					return nil
				}
				return func(in kops.ContainerdConfig) interface{} {
					return func(in kops.ContainerdConfig) []map[string]interface{} {
						return []map[string]interface{}{FlattenContainerdConfig(in)}
					}(in)
				}(*in)
			}(in)
		}(in.Containerd),
		"docker": func(in *kops.DockerConfig) interface{} {
			return func(in *kops.DockerConfig) interface{} {
				if in == nil {
					return nil
				}
				return func(in kops.DockerConfig) interface{} {
					return func(in kops.DockerConfig) []map[string]interface{} {
						return []map[string]interface{}{FlattenDockerConfig(in)}
					}(in)
				}(*in)
			}(in)
		}(in.Docker),
		"kube_dns": func(in *kops.KubeDNSConfig) interface{} {
			return func(in *kops.KubeDNSConfig) interface{} {
				if in == nil {
					return nil
				}
				return func(in kops.KubeDNSConfig) interface{} {
					return func(in kops.KubeDNSConfig) []map[string]interface{} {
						return []map[string]interface{}{FlattenKubeDNSConfig(in)}
					}(in)
				}(*in)
			}(in)
		}(in.KubeDNS),
		"kube_api_server": func(in *kops.KubeAPIServerConfig) interface{} {
			return func(in *kops.KubeAPIServerConfig) interface{} {
				if in == nil {
					return nil
				}
				return func(in kops.KubeAPIServerConfig) interface{} {
					return func(in kops.KubeAPIServerConfig) []map[string]interface{} {
						return []map[string]interface{}{FlattenKubeAPIServerConfig(in)}
					}(in)
				}(*in)
			}(in)
		}(in.KubeAPIServer),
		"kube_controller_manager": func(in *kops.KubeControllerManagerConfig) interface{} {
			return func(in *kops.KubeControllerManagerConfig) interface{} {
				if in == nil {
					return nil
				}
				return func(in kops.KubeControllerManagerConfig) interface{} {
					return func(in kops.KubeControllerManagerConfig) []map[string]interface{} {
						return []map[string]interface{}{FlattenKubeControllerManagerConfig(in)}
					}(in)
				}(*in)
			}(in)
		}(in.KubeControllerManager),
		"external_cloud_controller_manager": func(in *kops.CloudControllerManagerConfig) interface{} {
			return func(in *kops.CloudControllerManagerConfig) interface{} {
				if in == nil {
					return nil
				}
				return func(in kops.CloudControllerManagerConfig) interface{} {
					return func(in kops.CloudControllerManagerConfig) []map[string]interface{} {
						return []map[string]interface{}{FlattenCloudControllerManagerConfig(in)}
					}(in)
				}(*in)
			}(in)
		}(in.ExternalCloudControllerManager),
		"kube_scheduler": func(in *kops.KubeSchedulerConfig) interface{} {
			return func(in *kops.KubeSchedulerConfig) interface{} {
				if in == nil {
					return nil
				}
				return func(in kops.KubeSchedulerConfig) interface{} {
					return func(in kops.KubeSchedulerConfig) []map[string]interface{} {
						return []map[string]interface{}{FlattenKubeSchedulerConfig(in)}
					}(in)
				}(*in)
			}(in)
		}(in.KubeScheduler),
		"kube_proxy": func(in *kops.KubeProxyConfig) interface{} {
			return func(in *kops.KubeProxyConfig) interface{} {
				if in == nil {
					return nil
				}
				return func(in kops.KubeProxyConfig) interface{} {
					return func(in kops.KubeProxyConfig) []map[string]interface{} {
						return []map[string]interface{}{FlattenKubeProxyConfig(in)}
					}(in)
				}(*in)
			}(in)
		}(in.KubeProxy),
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
		"master_kubelet": func(in *kops.KubeletConfigSpec) interface{} {
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
		}(in.MasterKubelet),
		"cloud_config": func(in *kops.CloudConfiguration) interface{} {
			return func(in *kops.CloudConfiguration) interface{} {
				if in == nil {
					return nil
				}
				return func(in kops.CloudConfiguration) interface{} {
					return func(in kops.CloudConfiguration) []map[string]interface{} {
						return []map[string]interface{}{FlattenCloudConfiguration(in)}
					}(in)
				}(*in)
			}(in)
		}(in.CloudConfig),
		"external_dns": func(in *kops.ExternalDNSConfig) interface{} {
			return func(in *kops.ExternalDNSConfig) interface{} {
				if in == nil {
					return nil
				}
				return func(in kops.ExternalDNSConfig) interface{} {
					return func(in kops.ExternalDNSConfig) []map[string]interface{} {
						return []map[string]interface{}{FlattenExternalDNSConfig(in)}
					}(in)
				}(*in)
			}(in)
		}(in.ExternalDNS),
		"networking": func(in *kops.NetworkingSpec) interface{} {
			return func(in *kops.NetworkingSpec) interface{} {
				if in == nil {
					return nil
				}
				return func(in kops.NetworkingSpec) interface{} {
					return func(in kops.NetworkingSpec) []map[string]interface{} {
						return []map[string]interface{}{FlattenNetworkingSpec(in)}
					}(in)
				}(*in)
			}(in)
		}(in.Networking),
		"api": func(in *kops.AccessSpec) interface{} {
			return func(in *kops.AccessSpec) interface{} {
				if in == nil {
					return nil
				}
				return func(in kops.AccessSpec) interface{} {
					return func(in kops.AccessSpec) []map[string]interface{} {
						return []map[string]interface{}{FlattenAccessSpec(in)}
					}(in)
				}(*in)
			}(in)
		}(in.API),
		"authentication": func(in *kops.AuthenticationSpec) interface{} {
			return func(in *kops.AuthenticationSpec) interface{} {
				if in == nil {
					return nil
				}
				return func(in kops.AuthenticationSpec) interface{} {
					return func(in kops.AuthenticationSpec) []map[string]interface{} {
						return []map[string]interface{}{FlattenAuthenticationSpec(in)}
					}(in)
				}(*in)
			}(in)
		}(in.Authentication),
		"authorization": func(in *kops.AuthorizationSpec) interface{} {
			return func(in *kops.AuthorizationSpec) interface{} {
				if in == nil {
					return nil
				}
				return func(in kops.AuthorizationSpec) interface{} {
					return func(in kops.AuthorizationSpec) []map[string]interface{} {
						return []map[string]interface{}{FlattenAuthorizationSpec(in)}
					}(in)
				}(*in)
			}(in)
		}(in.Authorization),
		"node_authorization": func(in *kops.NodeAuthorizationSpec) interface{} {
			return func(in *kops.NodeAuthorizationSpec) interface{} {
				if in == nil {
					return nil
				}
				return func(in kops.NodeAuthorizationSpec) interface{} {
					return func(in kops.NodeAuthorizationSpec) []map[string]interface{} {
						return []map[string]interface{}{FlattenNodeAuthorizationSpec(in)}
					}(in)
				}(*in)
			}(in)
		}(in.NodeAuthorization),
		"cloud_labels": func(in map[string]string) interface{} {
			return func(in map[string]string) map[string]interface{} {
				if in == nil {
					return nil
				}
				// TODO
				return nil
			}(in)
		}(in.CloudLabels),
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
		"assets": func(in *kops.Assets) interface{} {
			return func(in *kops.Assets) interface{} {
				if in == nil {
					return nil
				}
				return func(in kops.Assets) interface{} {
					return func(in kops.Assets) []map[string]interface{} {
						return []map[string]interface{}{FlattenAssets(in)}
					}(in)
				}(*in)
			}(in)
		}(in.Assets),
		"iam": func(in *kops.IAMSpec) interface{} {
			return func(in *kops.IAMSpec) interface{} {
				if in == nil {
					return nil
				}
				return func(in kops.IAMSpec) interface{} {
					return func(in kops.IAMSpec) []map[string]interface{} {
						return []map[string]interface{}{FlattenIAMSpec(in)}
					}(in)
				}(*in)
			}(in)
		}(in.IAM),
		"encryption_config": func(in *bool) interface{} {
			return func(in *bool) interface{} {
				if in == nil {
					return nil
				}
				return func(in bool) interface{} {
					return FlattenBool(bool(in))
				}(*in)
			}(in)
		}(in.EncryptionConfig),
		"disable_subnet_tags": func(in bool) interface{} {
			return FlattenBool(bool(in))
		}(in.DisableSubnetTags),
		"use_host_certificates": func(in *bool) interface{} {
			return func(in *bool) interface{} {
				if in == nil {
					return nil
				}
				return func(in bool) interface{} {
					return FlattenBool(bool(in))
				}(*in)
			}(in)
		}(in.UseHostCertificates),
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
		"kube_server": func(in string) interface{} {
			return FlattenString(string(in))
		}(in.KubeServer),
		"kube_certificate_authority": func(in string) interface{} {
			return FlattenString(string(in))
		}(in.KubeCertificateAuthority),
		"kube_client_certificate": func(in string) interface{} {
			return FlattenString(string(in))
		}(in.KubeClientCertificate),
		"kube_client_key": func(in string) interface{} {
			return FlattenString(string(in))
		}(in.KubeClientKey),
		"kube_username": func(in string) interface{} {
			return FlattenString(string(in))
		}(in.KubeUsername),
		"kube_password": func(in string) interface{} {
			return FlattenString(string(in))
		}(in.KubePassword),
		"instance_group": func(in []*api.InstanceGroup) interface{} {
			return func(in []*api.InstanceGroup) []interface{} {
				var out []interface{}
				for _, in := range in {
					out = append(out, func(in *api.InstanceGroup) interface{} {
						if in == nil {
							return nil
						}
						return func(in api.InstanceGroup) interface{} {
							return func(in api.InstanceGroup) interface{} {
								return FlattenInstanceGroup(in)
							}(in)
						}(*in)
					}(in))
				}
				return out
			}(in)
		}(in.InstanceGroup),
	}
}
