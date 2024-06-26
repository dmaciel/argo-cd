package helm

import (
	"os"
	"path"
)

// HelmV3 represents helm V3 specific settings
var HelmV3 = HelmVer{
	binaryName:                   "helm",
	templateNameArg:              "--name-template",
	kubeVersionSupported:         true,
	showCommand:                  "show",
	pullCommand:                  "pull",
	initSupported:                false,
	getPostTemplateCallback:      cleanupChartLockFile,
	includeCrds:                  true,
	insecureSkipVerifySupported:  true,
	helmPassCredentialsSupported: true,
}

// workaround for Helm3 bug. Remove after https://github.com/helm/helm/issues/6870 is fixed.
// The `helm template` command generates Chart.lock after which `helm dependency build` does not work
// As workaround removing lock file unless it exists before running helm template
func cleanupChartLockFile(chartPath string) (func(), error) {
	exists := true
	lockPath := path.Join(chartPath, "Chart.lock")
	if _, err := os.Stat(lockPath); err != nil {
		if os.IsNotExist(err) {
			exists = false
		} else {
			return nil, err
		}
	}
	return func() {
		if !exists {
			_ = os.Remove(lockPath)
		}
	}, nil
}

// HelmVer contains Helm version specific settings such as helm binary and command names
type HelmVer struct {
	binaryName                   string
	initSupported                bool
	templateNameArg              string
	showCommand                  string
	pullCommand                  string
	kubeVersionSupported         bool
	getPostTemplateCallback      func(chartPath string) (func(), error)
	includeCrds                  bool
	insecureSkipVerifySupported  bool
	helmPassCredentialsSupported bool
}
