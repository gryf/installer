package installconfig

import (
	"context"
	"fmt"
	"sort"

	survey "github.com/AlecAivazis/survey/v2"
	"github.com/AlecAivazis/survey/v2/core"
	"github.com/pkg/errors"

	"github.com/openshift/installer/pkg/asset"
	awsconfig "github.com/openshift/installer/pkg/asset/installconfig/aws"
	azureconfig "github.com/openshift/installer/pkg/asset/installconfig/azure"
	baremetalconfig "github.com/openshift/installer/pkg/asset/installconfig/baremetal"
	gcpconfig "github.com/openshift/installer/pkg/asset/installconfig/gcp"
	ibmcloudconfig "github.com/openshift/installer/pkg/asset/installconfig/ibmcloud"
	nutanixconfig "github.com/openshift/installer/pkg/asset/installconfig/nutanix"
	openstackconfig "github.com/openshift/installer/pkg/asset/installconfig/openstack"
	powervsconfig "github.com/openshift/installer/pkg/asset/installconfig/powervs"
	vsphereconfig "github.com/openshift/installer/pkg/asset/installconfig/vsphere"
	"github.com/openshift/installer/pkg/types"
	"github.com/openshift/installer/pkg/types/aws"
	"github.com/openshift/installer/pkg/types/azure"
	"github.com/openshift/installer/pkg/types/baremetal"
	"github.com/openshift/installer/pkg/types/external"
	"github.com/openshift/installer/pkg/types/gcp"
	"github.com/openshift/installer/pkg/types/ibmcloud"
	"github.com/openshift/installer/pkg/types/none"
	"github.com/openshift/installer/pkg/types/nutanix"
	"github.com/openshift/installer/pkg/types/openstack"
	"github.com/openshift/installer/pkg/types/ovirt"
	"github.com/openshift/installer/pkg/types/powervs"
	"github.com/openshift/installer/pkg/types/vsphere"
)

// Platform is an asset that queries the user for the platform on which to install
// the cluster.
type platform struct {
	types.Platform
}

var _ asset.Asset = (*platform)(nil)

// Dependencies returns no dependencies.
func (a *platform) Dependencies() []asset.Asset {
	return []asset.Asset{}
}

// GenerateWithContext queries for input from the user.
func (a *platform) GenerateWithContext(ctx context.Context, _ asset.Parents) error {
	platform, err := a.queryUserForPlatform()
	if err != nil {
		return err
	}

	switch platform {
	case aws.Name:
		a.AWS, err = awsconfig.Platform()
		if err != nil {
			return err
		}
	case azure.Name:
		a.Azure, err = azureconfig.Platform()
		if err != nil {
			return err
		}
	case baremetal.Name:
		a.BareMetal, err = baremetalconfig.Platform()
		if err != nil {
			return err
		}
	case gcp.Name:
		a.GCP, err = gcpconfig.Platform()
		if err != nil {
			return err
		}
	case ibmcloud.Name:
		a.IBMCloud, err = ibmcloudconfig.Platform()
		if err != nil {
			return err
		}
	case external.Name:
		a.External = &external.Platform{}
	case none.Name:
		a.None = &none.Platform{}
	case openstack.Name:
		a.OpenStack, err = openstackconfig.Platform(ctx)
		if err != nil {
			return err
		}
	case ovirt.Name:
		return fmt.Errorf("platform oVirt is no longer supported")
	case vsphere.Name:
		a.VSphere, err = vsphereconfig.Platform()
		if err != nil {
			return err
		}
	case powervs.Name:
		a.PowerVS, err = powervsconfig.Platform()
		if err != nil {
			return err
		}
	case nutanix.Name:
		a.Nutanix, err = nutanixconfig.Platform()
		if err != nil {
			return err
		}
	default:
		return fmt.Errorf("unknown platform type %q", platform)
	}

	return nil
}

// Name returns the human-friendly name of the asset.
func (a *platform) Name() string {
	return "Platform"
}

func (a *platform) queryUserForPlatform() (platform string, err error) {
	err = survey.Ask([]*survey.Question{
		{
			Prompt: &survey.Select{
				Message: "Platform",
				Options: types.PlatformNames,
				Help:    "The platform on which the cluster will run.  For a full list of platforms, including those not supported by this wizard, see https://github.com/openshift/installer",
			},
			Validate: survey.ComposeValidators(survey.Required, func(ans interface{}) error {
				choice := ans.(core.OptionAnswer).Value
				i := sort.SearchStrings(types.PlatformNames, choice)
				if i == len(types.PlatformNames) || types.PlatformNames[i] != choice {
					return errors.Errorf("invalid platform %q", choice)
				}
				return nil
			}),
		},
	}, &platform)
	return
}

func (a *platform) CurrentName() string {
	return a.Platform.Name()
}

// Generate is implemented so this asset maintains compatibility with the Asset
// interface. It should never be called.
func (*platform) Generate(_ asset.Parents) (err error) {
	panic("platform.Generate was called instead of platform.GenerateWithContext")
}
