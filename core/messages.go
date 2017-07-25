package core

//SuccessMessage is return Zen message
type SuccessMessage struct {
	Message string `json:"message"`
}

//VersionMessage ...
type VersionMessage struct {
	AppID          string `json:"appID"`
	AppName        string `json:"appName"`
	ServerID       string `json:"serverID"`
	CreatedAt      string `json:"createdAt"`
	ReleaseVersion string `json:"version"`
	Commit         string `json:"commit"`
	Description    string `json:"description"`
}
