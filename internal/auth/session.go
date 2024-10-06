package auth

import (
	"time"

	"github.com/SunilKividor/drape/pkg/utils"
)

const (
	SessionStatusPending = SessionStatusT(0)
	SessionStatusActive  = SessionStatusT(1)
	SessionStatusExpired = SessionStatusT(2)
)

type (
	SessionStatusT uint

	Session struct {
		ID             uint64
		User           *User
		CreatedAt      time.Time
		LastAccessedAt time.Time
		SessionStatus  SessionStatusT
	}

	User struct {
		ID       string
		Username string
	}
)

func NewSession() (session *Session) {
	session = &Session{
		ID:             uint64(utils.GetCurrentTime().UTC().Unix()),
		CreatedAt:      utils.GetCurrentTime(),
		LastAccessedAt: utils.GetCurrentTime(),
		SessionStatus:  SessionStatusPending,
	}
	return
}
