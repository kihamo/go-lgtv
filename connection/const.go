package connection

const (
	// Request types
	reqTypeRegister = "register"
	reqTypeRequest  = "request"

	// Response types
	respTypeRegistered = "registered"
	respTypeResponse   = "response"
	respTypeError      = "error"

	// Pairing type
	pairTypePrompt = "PROMPT"

	// Permissions
	permissionLaunch                 = "LAUNCH"
	permissionControlAudio           = "CONTROL_AUDIO"
	permissionControlPower           = "CONTROL_POWER"
	permissionControlInputTv         = "CONTROL_INPUT_TV"
	permissionControlPlayback        = "CONTROL_INPUT_MEDIA_PLAYBACK"
	permissionReadChannelList        = "READ_TV_CHANNEL_LIST"
	permissionReadCurrentChannel     = "READ_CURRENT_CHANNEL"
	permissionReadRunningApps        = "READ_RUNNING_APPS"
	permissionReadInstalledApps      = "READ_INSTALLED_APPS"
	permissionReadInputList          = "READ_INPUT_DEVICE_LIST"
	permissionWriteNotificationToast = "WRITE_NOTIFICATION_TOAST"
)
