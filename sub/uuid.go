package sub

import (
	guuid "github.com/google/uuid"
)

func genUUID() string {
	id := guuid.New()
	return id.String()
}
