package config

import (
	"time"

	"github.com/gofiber/fiber/v2/middleware/session"
	"github.com/gofiber/utils"
)

var Store = session.New(session.Config{
	Expiration:   120 * time.Minute,
	KeyGenerator: utils.UUID,
})
