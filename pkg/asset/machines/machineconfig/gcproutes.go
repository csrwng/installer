package machineconfig

import (
	"fmt"
	"io/ioutil"

	"github.com/coreos/ignition/config/util"
	igntypes "github.com/coreos/ignition/config/v2_2/types"
	"github.com/openshift/installer/data"
	mcfgv1 "github.com/openshift/machine-config-operator/pkg/apis/machineconfiguration.openshift.io/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	"github.com/openshift/installer/pkg/asset/ignition"
)

// GCPRoutesService adds hack gcp-routes service to control-plane/compute machines
func GCPRoutesService(role string) *mcfgv1.MachineConfig {

	scriptFile, err := data.Assets.Open("bootstrap/files/usr/local/bin/gcp-routes.sh")
	if err != nil {
		panic(err.Error())
	}
	defer scriptFile.Close()
	script, err := ioutil.ReadAll(scriptFile)
	if err != nil {
		panic(err.Error())
	}

	serviceFile, err := data.Assets.Open("bootstrap/systemd/units/gcp-routes2.service")
	if err != nil {
		panic(err.Error())
	}
	defer serviceFile.Close()
	service, err := ioutil.ReadAll(serviceFile)
	if err != nil {
		panic(err.Error())
	}

	return &mcfgv1.MachineConfig{
		TypeMeta: metav1.TypeMeta{
			APIVersion: "machineconfiguration.openshift.io/v1",
			Kind:       "MachineConfig",
		},
		ObjectMeta: metav1.ObjectMeta{
			Name: fmt.Sprintf("99-%s-gcp-routes-service", role),
			Labels: map[string]string{
				"machineconfiguration.openshift.io/role": role,
			},
		},
		Spec: mcfgv1.MachineConfigSpec{
			Config: igntypes.Config{
				Ignition: igntypes.Ignition{
					Version: igntypes.MaxVersion.String(),
				},
				Storage: igntypes.Storage{
					Files: []igntypes.File{
						ignition.FileFromBytes("/usr/local/bin/gcp-routes.sh", "root", 0555, script),
					},
				},
				Systemd: igntypes.Systemd{
					Units: []igntypes.Unit{
						{
							Name:     "gcp-routes2.service",
							Enabled:  util.BoolToPtr(true),
							Contents: string(service),
						},
					},
				},
			},
		},
	}
}
