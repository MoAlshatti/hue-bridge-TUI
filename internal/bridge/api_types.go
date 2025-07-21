package bridge

type AuthSuccess struct {
	Success struct {
		UserName  string `json:"username"`
		ClientKey string `json:"clientkey"`
	} `json:"success"`
}

type ApiError struct {
	Error struct {
		Type        int    `json:"type"`
		Address     string `json:"address"`
		Description string `json:"description"`
	} `json:"error"`
}
