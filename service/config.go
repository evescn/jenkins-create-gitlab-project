package service

const (
	GitLabUrl   = "https://gitlab.evescn.com"
	GitLabToken = "Nh-Gm6cZC2G4n-aF8X1t"
)

type Commond struct {
	GroupName   string `json:"groupname"`
	ProjectName string `json:"projectname"`
	Visibility  string `json:"visibility"`
	Description string `json:"desc"`
}

type TempInfo []struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type ProjectInfo struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Path        string `json:"path"`
	NameSpaceID int    `json:"namespace_id"`
	Visibility  string `json:"visibility"`
	ImportUrl   string `json:"import_url"`
}
