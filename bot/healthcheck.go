package bot

import (
	"time"

	"github.com/shomali11/slacker"
)

// HealthCheck ...
func HealthCheck(request *slacker.Request, response slacker.ResponseWriter) {
	response.Typing()
	time.Sleep(1 * time.Second)
	response.Reply("Olá")
}
