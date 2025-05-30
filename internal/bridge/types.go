package bridge

// Light represents ..... TODO
type Light struct {
}

// Light is a panel type which incluedes an array of lights and a panel cursor...
type Lights struct {
	items  []Light
	Cursor int
}

// LightID .... TODO
type LightID string

// Group is either a zone or .... TODO
type Group struct { // consider making it lowercase if you aint finna use it in main
}

// Groups is a panel type which includes an array of groups, a panel cursor, and a map from the lights to the groups
type Groups struct {
	Items  []Group
	LookUp map[LightID]Group
	Cursor int
}

// represents the user uesd for the API
type User struct {
	Username string `json:"username"`
}
