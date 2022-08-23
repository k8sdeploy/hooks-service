package hooks

import (
	"encoding/json"
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
	var i interface{}

	if err := json.NewDecoder(r.Body).Decode(&i); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	fmt.Printf(" %+v\n", i)
}
