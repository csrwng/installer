package hosted

import (
	"path"
	"strings"

	"github.com/openshift/installer/data"
	"github.com/openshift/installer/pkg/asset"
	"github.com/openshift/installer/pkg/asset/installconfig"
	"github.com/openshift/installer/pkg/asset/releaseimage"
)

type RenderingScripts struct {
	files []*asset.File
}

// Name returns a human friendly name for the hosted cluster asset
func (o *RenderingScripts) Name() string {
	return "Rendering Scripts"
}

// Dependencies returns all of the dependencies directly needed by the
// hosted cluster asset
func (o *RenderingScripts) Dependencies() []asset.Asset {
	return []asset.Asset{
		&releaseimage.Image{},
		&installconfig.InstallConfig{},
	}
}

// Generate creates the manifests needed to launch a hosted cluster
func (o *RenderingScripts) Generate(dependencies asset.Parents) error {
	releaseImage := &releaseimage.Image{}
	installConfig := &installconfig.InstallConfig{}
	dependencies.Get(releaseImage, installConfig)

	templateData := map[string]string{
		"ReleaseImage":  releaseImage.PullSpec,
		"BaseDomain":    installConfig.Config.BaseDomain,
		"ClusterDomain": installConfig.Config.ClusterDomain(),
	}

	dir, err := data.Assets.Open("scripts/hosted")
	if err != nil {
		return err
	}
	files, err := dir.Readdir(0)
	if err != nil {
		return err
	}
	for _, f := range files {
		name := path.Base(f.Name())
		data, err := getFileContents(path.Join("scripts/hosted", name))
		if err != nil {
			return err
		}
		if strings.HasSuffix(name, ".template") {
			o.files = append(o.files, &asset.File{
				Filename: strings.TrimSuffix(name, ".template"),
				Data:     applyTemplateData(data, templateData),
			})
		} else {
			o.files = append(o.files, &asset.File{
				Filename: name,
				Data:     data,
			})
		}
	}
	return nil
}

// Files returns the manifests generated for the cluster
func (o *RenderingScripts) Files() []*asset.File {
	return o.files
}

// Load returns the openshift asset from disk.
func (c *RenderingScripts) Load(f asset.FileFetcher) (bool, error) {
	return false, nil
}
