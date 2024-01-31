package syncapi

import (
	"context"
	"fmt"
	"time"
)

type System struct{}
type SystemBrowse []string

func (a *ApiClient) GetSystemBrowse(ctx context.Context) (systemDirs SystemBrowse, err error) {
	systemDirs, err = getRequest[SystemBrowse](ctx, a.client, a.apiKey, fmt.Sprintf("%s/rest/%s", a.host, "system/browse"))
	return
}

type SystemConnectionsResponse struct {
	Connections struct{} `json:"connections"`
	Total       struct {
		At            time.Time `json:"at"`
		InBytesTotal  int       `json:"inBytesTotal"`
		OutBytesTotal int       `json:"outBytesTotal"`
	} `json:"total"`
}

func (a *ApiClient) GetSystemConnections(ctx context.Context) (c SystemConnectionsResponse, err error) {
	c, err = getRequest[SystemConnectionsResponse](ctx, a.client, a.apiKey, fmt.Sprintf("%s/rest/%s", a.host, "system/connections"))
	return
}

type SystemDebug struct {
	Enabled    []string `json:"enabled"`
	Facilities struct {
		Beacon      string `json:"beacon"`
		Config      string `json:"config"`
		Connections string `json:"connections"`
		Db          string `json:"db"`
		Dialer      string `json:"dialer"`
		Discover    string `json:"discover"`
		Events      string `json:"events"`
		Http        string `json:"http"`
		Main        string `json:"main"`
		Model       string `json:"model"`
		Protocol    string `json:"protocol"`
		Relay       string `json:"relay"`
		Scanner     string `json:"scanner"`
		Stats       string `json:"stats"`
		Sync        string `json:"sync"`
		Upgrade     string `json:"upgrade"`
		Upnp        string `json:"upnp"`
		Versioner   string `json:"versioner"`
	} `json:"facilities"`
}

func (a *ApiClient) GetSystemDebug(ctx context.Context) (sysDebug SystemDebug, err error) {
	sysDebug, err = getRequest[SystemDebug](ctx, a.client, a.apiKey, fmt.Sprintf("%s/rest/%s", a.host, "system/debug"))
	return
}

//func (a *ApiClient) PostSystemdebug(ctx context.Context) (b any, err error) {
//	getRequest[](ctx, a.client, a.apiKey, fmt.Sprintf("%s/rest/%s", a.host, "system/debug"))
//}
//func (a *ApiClient) GetSystemdiscovery(ctx context.Context) (b any, err error) {
//	getRequest[](ctx, a.client, a.apiKey, fmt.Sprintf("%s/rest/%s", a.host, "system/discovery"))
//}
//func (a *ApiClient) PostSystemDiscovery(ctx context.Context) (b any, err error) {
//	postRequest[](ctx, a.client, a.apiKey, fmt.Sprintf("%s/rest/%s", a.host, "system/discovery"))
//}
//func (a *ApiClient) PostSystemerrorclear(ctx context.Context) (b any, err error) {
//	getRequest[](ctx, a.client, a.apiKey, fmt.Sprintf("%s/rest/%s", "system/error/clear"))
//}

type SystemErrors struct {
	Errors []struct {
		When    time.Time `json:"when"`
		Message string    `json:"message"`
	} `json:"errors"`
}

func (a *ApiClient) GetSystemError(ctx context.Context) (sysErrs SystemErrors, err error) {
	sysErrs, err = getRequest[SystemErrors](ctx, a.client, a.apiKey, fmt.Sprintf("%s/rest/%s", a.host, "system/error"))
	return
}

//func (a *ApiClient) PostSystemErrorClear(ctx context.Context) (b any, err error) {
//	postRequest[](ctx, a.client, a.apiKey, fmt.Sprintf("%s/rest/%s", a.host, "system/error"))
//}

type SystemMessages struct {
	Messages []struct {
		When    time.Time `json:"when"`
		Message string    `json:"message"`
	} `json:"messages"`
}

func (a *ApiClient) GetSystemLog(ctx context.Context) (systemMessages SystemMessages, err error) {
	systemMessages, err = getRequest[SystemMessages](ctx, a.client, a.apiKey, fmt.Sprintf("%s/rest/%s", a.host, "system/log"))
	return
}

//func (a *ApiClient) GetSystemlogtxt(ctx context.Context) (b any, err error) {
//	getRequest[](ctx, a.client, a.apiKey, fmt.Sprintf("%s/rest/%s", "system/log.txt"))
//}

type SystemPaths struct {
	AuditLog        string `json:"auditLog"`
	BaseDirConfig   string `json:"baseDir-config"`
	BaseDirData     string `json:"baseDir-data"`
	BaseDirUserHome string `json:"baseDir-userHome"`
	CertFile        string `json:"certFile"`
	Config          string `json:"config"`
	CsrfTokens      string `json:"csrfTokens"`
	Database        string `json:"database"`
	DefFolder       string `json:"defFolder"`
	GuiAssets       string `json:"guiAssets"`
	HttpsCertFile   string `json:"httpsCertFile"`
	HttpsKeyFile    string `json:"httpsKeyFile"`
	KeyFile         string `json:"keyFile"`
	LogFile         string `json:"logFile"`
	PanicLog        string `json:"panicLog"`
}

func (a *ApiClient) GetSystemPaths(ctx context.Context) (sysPaths SystemPaths, err error) {
	sysPaths, err = getRequest[SystemPaths](ctx, a.client, a.apiKey, fmt.Sprintf("%s/rest/%s", a.host, "system/paths"))
	return
}

//func (a *ApiClient) PostSystemPause(ctx context.Context) (b any, err error) {
//	getRequest[](ctx, a.client, a.apiKey, fmt.Sprintf("%s/rest/%s", a.host, "system/pause"))
//}

type SystemPing struct {
	Ping string `json:"ping"`
}

func (a *ApiClient) GetSystemPing(ctx context.Context) (sysPing SystemPing, err error) {
	sysPing, err = getRequest[SystemPing](ctx, a.client, a.apiKey, fmt.Sprintf("%s/rest/%s", a.host, "system/ping"))
	return
}

//func (a *ApiClient) PostSystemPing(ctx context.Context) (b any, err error) {
//	getRequest[](ctx, a.client, a.apiKey, fmt.Sprintf("%s/rest/%s", a.host, "system/ping"))
//}
//func (a *ApiClient) PostSystemReset(ctx context.Context) (b any, err error) {
//	getRequest[](ctx, a.client, a.apiKey, fmt.Sprintf("%s/rest/%s", a.host, "system/reset"))
//}
//func (a *ApiClient) PostSystemRestart(ctx context.Context) (b any, err error) {
//	getRequest[](ctx, a.client, a.apiKey, fmt.Sprintf("%s/rest/%s", a.host, "system/restart"))
//}
//func (a *ApiClient) PostSystemResume(ctx context.Context) (b any, err error) {
//	getRequest[](ctx, a.client, a.apiKey, fmt.Sprintf("%s/rest/%s", a.host, "system/resume"))
//}
//func (a *ApiClient) PostSystemShutdown(ctx context.Context) (b any, err error) {
//	getRequest[](ctx, a.client, a.apiKey, fmt.Sprintf("%s/rest/%s", a.host, "system/shutdown"))
//}

type SystemStatus struct {
	Alloc                   int `json:"alloc"`
	ConnectionServiceStatus struct {
		DynamicHttpsRelaysSyncthingNetEndpoint struct {
			Error        interface{} `json:"error"`
			LanAddresses []string    `json:"lanAddresses"`
			WanAddresses []string    `json:"wanAddresses"`
		} `json:"dynamic+https://relays.syncthing.net/endpoint"`
		Tcp000022000 struct {
			Error        interface{} `json:"error"`
			LanAddresses []string    `json:"lanAddresses"`
			WanAddresses []string    `json:"wanAddresses"`
		} `json:"tcp://0.0.0.0:22000"`
	} `json:"connectionServiceStatus"`
	CpuPercent       int  `json:"cpuPercent"`
	DiscoveryEnabled bool `json:"discoveryEnabled"`
	DiscoveryErrors  struct {
		GlobalHttpsDiscoveryV41SyncthingNetV2 string `json:"global@https://discovery-v4-1.syncthing.net/v2/"`
		GlobalHttpsDiscoveryV42SyncthingNetV2 string `json:"global@https://discovery-v4-2.syncthing.net/v2/"`
		GlobalHttpsDiscoveryV43SyncthingNetV2 string `json:"global@https://discovery-v4-3.syncthing.net/v2/"`
		GlobalHttpsDiscoveryV61SyncthingNetV2 string `json:"global@https://discovery-v6-1.syncthing.net/v2/"`
		GlobalHttpsDiscoveryV62SyncthingNetV2 string `json:"global@https://discovery-v6-2.syncthing.net/v2/"`
		GlobalHttpsDiscoveryV63SyncthingNetV2 string `json:"global@https://discovery-v6-3.syncthing.net/v2/"`
	} `json:"discoveryErrors"`
	DiscoveryStatus struct {
		IPv4Local struct {
			Error interface{} `json:"error"`
		} `json:"IPv4 local"`
		IPv6Local struct {
			Error interface{} `json:"error"`
		} `json:"IPv6 local"`
		GlobalHttpsDiscoveryV41SyncthingNetV2 struct {
			Error string `json:"error"`
		} `json:"global@https://discovery-v4-1.syncthing.net/v2/"`
		GlobalHttpsDiscoveryV42SyncthingNetV2 struct {
			Error string `json:"error"`
		} `json:"global@https://discovery-v4-2.syncthing.net/v2/"`
		GlobalHttpsDiscoveryV43SyncthingNetV2 struct {
			Error string `json:"error"`
		} `json:"global@https://discovery-v4-3.syncthing.net/v2/"`
		GlobalHttpsDiscoveryV61SyncthingNetV2 struct {
			Error string `json:"error"`
		} `json:"global@https://discovery-v6-1.syncthing.net/v2/"`
		GlobalHttpsDiscoveryV62SyncthingNetV2 struct {
			Error string `json:"error"`
		} `json:"global@https://discovery-v6-2.syncthing.net/v2/"`
		GlobalHttpsDiscoveryV63SyncthingNetV2 struct {
			Error string `json:"error"`
		} `json:"global@https://discovery-v6-3.syncthing.net/v2/"`
	} `json:"discoveryStatus"`
	DiscoveryMethods int `json:"discoveryMethods"`
	Goroutines       int `json:"goroutines"`
	LastDialStatus   struct {
		Tcp10203040 struct {
			When  time.Time `json:"when"`
			Error string    `json:"error"`
		} `json:"tcp://10.20.30.40"`
		Tcp1721633322000 struct {
			When time.Time `json:"when"`
			Ok   bool      `json:"ok"`
		} `json:"tcp://172.16.33.3:22000"`
		Tcp8323312022122000 struct {
			When  time.Time `json:"when"`
			Error string    `json:"error"`
		} `json:"tcp://83.233.120.221:22000"`
	} `json:"lastDialStatus"`
	MyID          string    `json:"myID"`
	PathSeparator string    `json:"pathSeparator"`
	StartTime     time.Time `json:"startTime"`
	Sys           int       `json:"sys"`
	Themes        []string  `json:"themes"`
	Tilde         string    `json:"tilde"`
	Uptime        int       `json:"uptime"`
}

func (a *ApiClient) GetSystemStatus(ctx context.Context) (sysStatus SystemStatus, err error) {
	sysStatus, err = getRequest[SystemStatus](ctx, a.client, a.apiKey, fmt.Sprintf("%s/rest/%s", a.host, "system/status"))
	return
}

type SystemUpgrade struct {
	Latest     string `json:"latest"`
	MajorNewer bool   `json:"majorNewer"`
	Newer      bool   `json:"newer"`
	Running    string `json:"running"`
}

func (a *ApiClient) GetSystemUpgrade(ctx context.Context) (sysUpgrade SystemUpgrade, err error) {
	sysUpgrade, err = getRequest[SystemUpgrade](ctx, a.client, a.apiKey, fmt.Sprintf("%s/rest/%s", a.host, "system/upgrade"))
	return
}

//func (a *ApiClient) PostSystemUpgrade(ctx context.Context) (b any, err error) {
//	getRequest[](ctx, a.client, a.apiKey, fmt.Sprintf("%s/rest/%s", a.host, "system/upgrade"))
//}

type SystemVersion struct {
	Arch        string `json:"arch"`
	LongVersion string `json:"longVersion"`
	Os          string `json:"os"`
	Version     string `json:"version"`
}

func (a *ApiClient) GetSystemVersion(ctx context.Context) (sysVersion SystemVersion, err error) {
	sysVersion, err = getRequest[SystemVersion](ctx, a.client, a.apiKey, fmt.Sprintf("%s/rest/%s", a.host, "system/version"))
	return
}
