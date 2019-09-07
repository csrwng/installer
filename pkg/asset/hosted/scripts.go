package hosted

import (
	"fmt"
	"path"

	"github.com/openshift/installer/pkg/asset"
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
	}
}

// Generate creates the manifests needed to launch a hosted cluster
func (o *RenderingScripts) Generate(dependencies asset.Parents) error {
	releaseImage := &releaseimage.Image{}
	dependencies.Get(releaseImage)

	templateData := map[string]string{
		"ReleaseImage": releaseImage.PullSpec,
	}
	for _, f := range []string{
		"render-operand-manifests.sh",
	} {
		data, err := getFileContents(path.Join("scripts/hosted", fmt.Sprintf("%s.template", f)))
		if err != nil {
			return err
		}
		o.files = append(o.files, &asset.File{
			Filename: f,
			Data:     applyTemplateData(data, templateData),
		})
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
