package hosted

import (
	"encoding/base64"
	"fmt"
	"path"

	"github.com/vincent-petithory/dataurl"

	"github.com/openshift/installer/pkg/asset"
	"github.com/openshift/installer/pkg/asset/ignition/bootstrap"
	"github.com/openshift/installer/pkg/asset/kubeconfig"
)

type ControlPlaneSecrets struct {
	files []*asset.File
}

// Name returns a human friendly name for the hosted cluster asset
func (o *ControlPlaneSecrets) Name() string {
	return "Control Plane Secrets"
}

// Dependencies returns all of the dependencies directly needed by the
// hosted cluster asset
func (o *ControlPlaneSecrets) Dependencies() []asset.Asset {
	return []asset.Asset{
		&bootstrap.Bootstrap{},
		&KubeAPIServerInternalServiceCertKey{},
		&kubeconfig.ServiceAdminClient{},
	}
}

// Generate creates the manifests needed to launch a hosted cluster
func (o *ControlPlaneSecrets) Generate(dependencies asset.Parents) error {
	bt := &bootstrap.Bootstrap{}
	cpSvcCerts := &KubeAPIServerInternalServiceCertKey{}
	svcKubeconfig := &kubeconfig.ServiceAdminClient{}
	dependencies.Get(bt, cpSvcCerts, svcKubeconfig)

	secretFiles := map[string]string{}
	for _, file := range bt.Config.Storage.Files {
		if fileDir := path.Dir(file.Path); fileDir == "/opt/openshift/tls" || fileDir == "/opt/openshift/auth" {
			u, err := dataurl.DecodeString(file.Contents.Source)
			if err != nil {
				return err
			}
			secretFiles[path.Base(file.Path)] = base64.StdEncoding.EncodeToString(u.Data)
		}
	}
	for _, file := range cpSvcCerts.Files() {
		secretFiles[path.Base(file.Filename)] = base64.StdEncoding.EncodeToString(file.Data)
	}

	svcKubeconfigEncoded := base64.StdEncoding.EncodeToString(svcKubeconfig.Files()[0].Data)

	templateData := map[string]interface{}{"Files": secretFiles, "ServiceKubeconfig": svcKubeconfigEncoded}
	for _, file := range []string{
		"control-plane-certificates.yaml",
		"internal-kubeconfig.yaml",
	} {
		data, err := getFileContents(path.Join("manifests/hosted", fmt.Sprintf("%s.template", file)))
		if err != nil {
			return err
		}
		o.files = append(o.files, &asset.File{
			Filename: path.Join("hosted", file),
			Data:     applyTemplateData(data, templateData),
		})
	}

	return nil
}

// Files returns the manifests generated for the cluster
func (o *ControlPlaneSecrets) Files() []*asset.File {
	return o.files
}

// Load returns the openshift asset from disk.
func (c *ControlPlaneSecrets) Load(f asset.FileFetcher) (bool, error) {
	return false, nil
}
