package ewelink

import (
	"fmt"
	"math/rand"
)

const (
	version    = "8"
	apkVersion = "1.8"
)

var applicationVersions = [2]string{"3.5.3", "3.5.4"}

// Application contains the application specific fields.
type Application struct {
	AppVersion string
	Version    string
	ApkVersion string
}

func (a Application) String() string {
	return fmt.Sprintf("%#v", a)
}

func newApplication() *Application {
	return &Application{AppVersion: getRandomApplicationVersion(), Version: version, ApkVersion: apkVersion}
}

func getRandomApplicationVersion() string {
	return applicationVersions[rand.Intn(len(applicationVersions))] // #nosec:G404
}
