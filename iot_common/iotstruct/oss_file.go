package iotstruct

type FileResponse struct {
	Size     int64  `json:"size"`
	Path     string `json:"path"`
	FullPath string `json:"fullPath"`
	Name     string `json:"name"`
	Type     string `json:"type"`
	Key      string `json:"key"`
}
