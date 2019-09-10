package kubeconfig

import (
	"github.com/openshift/installer/pkg/asset"
	"github.com/openshift/installer/pkg/asset/installconfig"
	"github.com/openshift/installer/pkg/asset/tls"
)

// ServiceAdminClient is the asset for the admin kubeconfig.
type ServiceAdminClient struct {
	kubeconfig
}

var _ asset.WritableAsset = (*ServiceAdminClient)(nil)

// Dependencies returns the dependency of the kubeconfig.
func (k *ServiceAdminClient) Dependencies() []asset.Asset {
	return []asset.Asset{
		&tls.AdminKubeConfigClientCertKey{},
		&tls.KubeAPIServerCompleteCABundle{},
		&installconfig.InstallConfig{},
	}
}

// Generate generates the kubeconfig.
func (k *ServiceAdminClient) Generate(parents asset.Parents) error {
	ca := &tls.KubeAPIServerCompleteCABundle{}
	clientCertKey := &tls.AdminKubeConfigClientCertKey{}
	installConfig := &installconfig.InstallConfig{}
	parents.Get(ca, clientCertKey, installConfig)

	return k.kubeconfig.generate(
		ca,
		clientCertKey,
		"https://kube-apiserver:6443",
		installConfig.Config.GetName(),
		"admin",
		"internal-kubeconfig",
	)
}

// Name returns the human-friendly name of the asset.
func (k *ServiceAdminClient) Name() string {
	return "Kubeconfig Admin Client"
}

// Load returns the kubeconfig from disk.
func (k *ServiceAdminClient) Load(f asset.FileFetcher) (found bool, err error) {
	return false, nil
}
