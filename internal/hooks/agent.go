package hooks

import (
	"fmt"
	"github.com/k8sdeploy/hooks-service/internal/config"
	"net/http"
)

type Hooks struct {
	Config *config.Config
}

func NewHooks(cfg *config.Config) *Hooks {
	return &Hooks{
		Config: cfg,
	}
}

func (h *Hooks) HandleHook(w http.ResponseWriter, r *http.Request) {
	_ = fmt.Sprintf("%s", r.Body)
}
