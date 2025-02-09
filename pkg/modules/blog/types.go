package blog

type BlogTree struct {
	Tree []struct {
		Path string `json:"path"`
	} `json:"tree"`
}

type Blog struct {
	Path    string
	RawData string
}
