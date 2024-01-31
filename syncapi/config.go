package syncapi

import (
	"context"
	"fmt"
)

type Config struct {
	Version              int            `json:"version"`
	Folders              []ConfigFolder `json:"folders"`
	Devices              []ConfigDevice `json:"devices"`
	Gui                  ConfigGUI      `json:"gui"`
	Ldap                 ConfigLDAP     `json:"ldap"`
	Options              ConfigOptions  `json:"options"`
	RemoteIgnoredDevices []interface{}  `json:"remoteIgnoredDevices"`
	Defaults             struct {
		Folder  ConfigFolder  `json:"folder"`
		Device  ConfigDevice  `json:"device"`
		Ignores ConfigIgnores `json:"ignores"`
	} `json:"defaults"`
}

func (a *ApiClient) GetConfig(ctx context.Context) (c Config, err error) {
	c, err = getRequest[Config](ctx, a.client, a.apiKey, fmt.Sprintf("%s/rest/%s", a.host, "config"))
	return
}

type ConfigRestartRequired struct {
	RequiresRestart bool `json:"requiresRestart"`
}

func (a *ApiClient) GetConfigRestartRequired(ctx context.Context) (c ConfigRestartRequired, err error) {
	c, err = getRequest[ConfigRestartRequired](ctx, a.client, a.apiKey, fmt.Sprintf("%s/rest/%s", a.host, "config/restart-required"))
	return
}

type ConfigFolder struct {
	Id             string `json:"id"`
	Label          string `json:"label"`
	FilesystemType string `json:"filesystemType"`
	Path           string `json:"path"`
	Type           string `json:"type"`
	Devices        []struct {
		DeviceID           string `json:"deviceID"`
		IntroducedBy       string `json:"introducedBy"`
		EncryptionPassword string `json:"encryptionPassword"`
	} `json:"devices"`
	RescanIntervalS  int  `json:"rescanIntervalS"`
	FsWatcherEnabled bool `json:"fsWatcherEnabled"`
	FsWatcherDelayS  int  `json:"fsWatcherDelayS"`
	IgnorePerms      bool `json:"ignorePerms"`
	AutoNormalize    bool `json:"autoNormalize"`
	MinDiskFree      struct {
		Value int    `json:"value"`
		Unit  string `json:"unit"`
	} `json:"minDiskFree"`
	Versioning struct {
		Type   string `json:"type"`
		Params struct {
		} `json:"params"`
		CleanupIntervalS int    `json:"cleanupIntervalS"`
		FsPath           string `json:"fsPath"`
		FsType           string `json:"fsType"`
	} `json:"versioning"`
	Copiers                 int    `json:"copiers"`
	PullerMaxPendingKiB     int    `json:"pullerMaxPendingKiB"`
	Hashers                 int    `json:"hashers"`
	Order                   string `json:"order"`
	IgnoreDelete            bool   `json:"ignoreDelete"`
	ScanProgressIntervalS   int    `json:"scanProgressIntervalS"`
	PullerPauseS            int    `json:"pullerPauseS"`
	MaxConflicts            int    `json:"maxConflicts"`
	DisableSparseFiles      bool   `json:"disableSparseFiles"`
	DisableTempIndexes      bool   `json:"disableTempIndexes"`
	Paused                  bool   `json:"paused"`
	WeakHashThresholdPct    int    `json:"weakHashThresholdPct"`
	MarkerName              string `json:"markerName"`
	CopyOwnershipFromParent bool   `json:"copyOwnershipFromParent"`
	ModTimeWindowS          int    `json:"modTimeWindowS"`
	MaxConcurrentWrites     int    `json:"maxConcurrentWrites"`
	DisableFsync            bool   `json:"disableFsync"`
	BlockPullOrder          string `json:"blockPullOrder"`
	CopyRangeMethod         string `json:"copyRangeMethod"`
	CaseSensitiveFS         bool   `json:"caseSensitiveFS"`
	JunctionsAsDirs         bool   `json:"junctionsAsDirs"`
	SyncOwnership           bool   `json:"syncOwnership"`
	SendOwnership           bool   `json:"sendOwnership"`
	SyncXattrs              bool   `json:"syncXattrs"`
	SendXattrs              bool   `json:"sendXattrs"`
	XattrFilter             struct {
		Entries            []interface{} `json:"entries"`
		MaxSingleEntrySize int           `json:"maxSingleEntrySize"`
		MaxTotalSize       int           `json:"maxTotalSize"`
	} `json:"xattrFilter"`
}

// TODO PUT, DELETE, PATH
func (a *ApiClient) GetConfigFolders(ctx context.Context) (f []ConfigFolder, err error) {
	f, err = getRequest[[]ConfigFolder](ctx, a.client, a.apiKey, fmt.Sprintf("%s/rest/%s", a.host, "config/folders"))
	return
}

// TODO PUT, DELETE, PATH
func (a *ApiClient) GetConfigFolder(ctx context.Context, folderId string) (f ConfigFolder, err error) {
	f, err = getRequest[ConfigFolder](ctx, a.client, a.apiKey, fmt.Sprintf("%s/rest/%s%s", a.host, "config/folder/", folderId))
	return
}

type ConfigDevice struct {
	DeviceID                 string        `json:"deviceID"`
	Name                     string        `json:"name"`
	Addresses                []string      `json:"addresses"`
	Compression              string        `json:"compression"`
	CertName                 string        `json:"certName"`
	Introducer               bool          `json:"introducer"`
	SkipIntroductionRemovals bool          `json:"skipIntroductionRemovals"`
	IntroducedBy             string        `json:"introducedBy"`
	Paused                   bool          `json:"paused"`
	AllowedNetworks          []interface{} `json:"allowedNetworks"`
	AutoAcceptFolders        bool          `json:"autoAcceptFolders"`
	MaxSendKbps              int           `json:"maxSendKbps"`
	MaxRecvKbps              int           `json:"maxRecvKbps"`
	IgnoredFolders           []interface{} `json:"ignoredFolders"`
	MaxRequestKiB            int           `json:"maxRequestKiB"`
	Untrusted                bool          `json:"untrusted"`
	RemoteGUIPort            int           `json:"remoteGUIPort"`
	NumConnections           int           `json:"numConnections"`
}

// TODO PUT, DELETE, PATH
func (a *ApiClient) GetConfigDevices(ctx context.Context) (d []ConfigDevice, err error) {
	d, err = getRequest[[]ConfigDevice](ctx, a.client, a.apiKey, fmt.Sprintf("%s/rest/%s", a.host, "config/devices"))
	return
}

// TODO PUT, DELETE, PATH
func (a *ApiClient) GetConfigDevice(ctx context.Context, deviceId string) (f ConfigDevice, err error) {
	f, err = getRequest[ConfigDevice](ctx, a.client, a.apiKey, fmt.Sprintf("%s/rest/%s%s", a.host, "config/device/", deviceId))
	return
}

func (a *ApiClient) GetConfigDefaultsFolder(ctx context.Context) (c ConfigFolder, err error) {
	c, err = getRequest[ConfigFolder](ctx, a.client, a.apiKey, fmt.Sprintf("%s/rest/%s", a.host, "defaults/folder"))
	return
}
func (a *ApiClient) GetConfigDefaultsDevice(ctx context.Context) (c ConfigDevice, err error) {
	c, err = getRequest[ConfigDevice](ctx, a.client, a.apiKey, fmt.Sprintf("%s/rest/%s", a.host, "defaults/device"))
	return
}

type ConfigIgnores struct {
	Lines []interface{} `json:"lines"`
}

func (a *ApiClient) GetConfigDefaultsIgnores(ctx context.Context) (c ConfigIgnores, err error) {
	c, err = getRequest[ConfigIgnores](ctx, a.client, a.apiKey, fmt.Sprintf("%s/rest/%s", a.host, "config"))
	return
}

type ConfigOptions struct {
	ListenAddresses         []string `json:"listenAddresses"`
	GlobalAnnounceServers   []string `json:"globalAnnounceServers"`
	GlobalAnnounceEnabled   bool     `json:"globalAnnounceEnabled"`
	LocalAnnounceEnabled    bool     `json:"localAnnounceEnabled"`
	LocalAnnouncePort       int      `json:"localAnnouncePort"`
	LocalAnnounceMCAddr     string   `json:"localAnnounceMCAddr"`
	MaxSendKbps             int      `json:"maxSendKbps"`
	MaxRecvKbps             int      `json:"maxRecvKbps"`
	ReconnectionIntervalS   int      `json:"reconnectionIntervalS"`
	RelaysEnabled           bool     `json:"relaysEnabled"`
	RelayReconnectIntervalM int      `json:"relayReconnectIntervalM"`
	StartBrowser            bool     `json:"startBrowser"`
	NatEnabled              bool     `json:"natEnabled"`
	NatLeaseMinutes         int      `json:"natLeaseMinutes"`
	NatRenewalMinutes       int      `json:"natRenewalMinutes"`
	NatTimeoutSeconds       int      `json:"natTimeoutSeconds"`
	UrAccepted              int      `json:"urAccepted"`
	UrSeen                  int      `json:"urSeen"`
	UrUniqueId              string   `json:"urUniqueId"`
	UrURL                   string   `json:"urURL"`
	UrPostInsecurely        bool     `json:"urPostInsecurely"`
	UrInitialDelayS         int      `json:"urInitialDelayS"`
	AutoUpgradeIntervalH    int      `json:"autoUpgradeIntervalH"`
	UpgradeToPreReleases    bool     `json:"upgradeToPreReleases"`
	KeepTemporariesH        int      `json:"keepTemporariesH"`
	CacheIgnoredFiles       bool     `json:"cacheIgnoredFiles"`
	ProgressUpdateIntervalS int      `json:"progressUpdateIntervalS"`
	LimitBandwidthInLan     bool     `json:"limitBandwidthInLan"`
	MinHomeDiskFree         struct {
		Value int    `json:"value"`
		Unit  string `json:"unit"`
	} `json:"minHomeDiskFree"`
	ReleasesURL                         string        `json:"releasesURL"`
	AlwaysLocalNets                     []interface{} `json:"alwaysLocalNets"`
	OverwriteRemoteDeviceNamesOnConnect bool          `json:"overwriteRemoteDeviceNamesOnConnect"`
	TempIndexMinBlocks                  int           `json:"tempIndexMinBlocks"`
	UnackedNotificationIDs              []string      `json:"unackedNotificationIDs"`
	TrafficClass                        int           `json:"trafficClass"`
	SetLowPriority                      bool          `json:"setLowPriority"`
	MaxFolderConcurrency                int           `json:"maxFolderConcurrency"`
	CrURL                               string        `json:"crURL"`
	CrashReportingEnabled               bool          `json:"crashReportingEnabled"`
	StunKeepaliveStartS                 int           `json:"stunKeepaliveStartS"`
	StunKeepaliveMinS                   int           `json:"stunKeepaliveMinS"`
	StunServers                         []string      `json:"stunServers"`
	DatabaseTuning                      string        `json:"databaseTuning"`
	MaxConcurrentIncomingRequestKiB     int           `json:"maxConcurrentIncomingRequestKiB"`
	AnnounceLANAddresses                bool          `json:"announceLANAddresses"`
	SendFullIndexOnUpgrade              bool          `json:"sendFullIndexOnUpgrade"`
	FeatureFlags                        []interface{} `json:"featureFlags"`
	ConnectionLimitEnough               int           `json:"connectionLimitEnough"`
	ConnectionLimitMax                  int           `json:"connectionLimitMax"`
	InsecureAllowOldTLSVersions         bool          `json:"insecureAllowOldTLSVersions"`
	ConnectionPriorityTcpLan            int           `json:"connectionPriorityTcpLan"`
	ConnectionPriorityQuicLan           int           `json:"connectionPriorityQuicLan"`
	ConnectionPriorityTcpWan            int           `json:"connectionPriorityTcpWan"`
	ConnectionPriorityQuicWan           int           `json:"connectionPriorityQuicWan"`
	ConnectionPriorityRelay             int           `json:"connectionPriorityRelay"`
	ConnectionPriorityUpgradeThreshold  int           `json:"connectionPriorityUpgradeThreshold"`
}

// TODO PUT, PATH
func (a *ApiClient) GetConfigOptions(ctx context.Context) (c ConfigOptions, err error) {
	c, err = getRequest[ConfigOptions](ctx, a.client, a.apiKey, fmt.Sprintf("%s/rest/%s", a.host, "config/options"))
	return
}

type ConfigLDAP struct {
	Address            string `json:"address"`
	BindDN             string `json:"bindDN"`
	Transport          string `json:"transport"`
	InsecureSkipVerify bool   `json:"insecureSkipVerify"`
	SearchBaseDN       string `json:"searchBaseDN"`
	SearchFilter       string `json:"searchFilter"`
}

// TODO PUT, PATH
func (a *ApiClient) GetConfigLDAP(ctx context.Context) (c ConfigLDAP, err error) {
	c, err = getRequest[ConfigLDAP](ctx, a.client, a.apiKey, fmt.Sprintf("%s/rest/%s", a.host, "config/ldap"))
	return
}

type ConfigGUI struct {
	Enabled                   bool   `json:"enabled"`
	Address                   string `json:"address"`
	UnixSocketPermissions     string `json:"unixSocketPermissions"`
	User                      string `json:"user"`
	Password                  string `json:"password"`
	AuthMode                  string `json:"authMode"`
	UseTLS                    bool   `json:"useTLS"`
	ApiKey                    string `json:"apiKey"`
	InsecureAdminAccess       bool   `json:"insecureAdminAccess"`
	Theme                     string `json:"theme"`
	Debugging                 bool   `json:"debugging"`
	InsecureSkipHostcheck     bool   `json:"insecureSkipHostcheck"`
	InsecureAllowFrameLoading bool   `json:"insecureAllowFrameLoading"`
	SendBasicAuthPrompt       bool   `json:"sendBasicAuthPrompt"`
}

// TODO PUT, PATH
func (a *ApiClient) GetConfigGUI(ctx context.Context) (c ConfigGUI, err error) {
	c, err = getRequest[ConfigGUI](ctx, a.client, a.apiKey, fmt.Sprintf("%s/rest/%s", a.host, "config/gui"))
	return
}
