package hosted

import (
	"path"

	"github.com/vincent-petithory/dataurl"

	"github.com/openshift/installer/pkg/asset"
	"github.com/openshift/installer/pkg/asset/ignition/bootstrap"
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
	}
}

// Generate creates the manifests needed to launch a hosted cluster
func (o *ControlPlaneSecrets) Generate(dependencies asset.Parents) error {
	bt := &bootstrap.Bootstrap{}
	dependencies.Get(bt)

	secretFiles := map[string][]byte{}
	for _, file := range bt.Config.Storage.Files {
		if fileDir := path.Dir(file.Path); fileDir == "/opt/openshift/tls" || fileDir == "/opt/openshift/auth" {
			u, err := dataurl.DecodeString(file.Contents.Source)
			if err != nil {
				return err
			}
			secretFiles[path.Base(file.Path)] = u.Data
		}
	}
	templateData := map[string]interface{}{"Files": secretFiles}

	data, err := getFileContents(path.Join("manifests/hosted/control-plane-certificates.yaml.template"))
	if err != nil {
		return err
	}
	o.files = append(o.files, &asset.File{
		Filename: "hosted/control-plane-certificates.yaml",
		Data:     applyTemplateData(data, templateData),
	})

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
