package hooks

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/k8sdeploy/hooks-service/internal/config"
	keybuf "github.com/k8sdeploy/protos/generated/key/v1"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
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

type K8sDeployEvent struct {
	ImageHash        string `json:"imageHash"`
	ImageTag         string `json:"imageTag"`
	ServiceName      string `json:"serviceName"`
	ServiceNamespace string `json:"serviceNamespace"`
}

type HookEvent struct {
	GithubEvent
	K8sDeployEvent
}

func (h *Hooks) HandleHook(w http.ResponseWriter, r *http.Request) {
	var i HookEvent

	companyId := r.Header.Get("X-API-ID")
	key := r.Header.Get("X-API-KEY")
	secret := r.Header.Get("X-API-SECRET")

	if companyId == "" || key == "" || secret == "" {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}
	if ok, err := h.ValidateKey(companyId, key, secret); !ok {
		fmt.Printf("failed to validate key: %v", err)
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	if err := json.NewDecoder(r.Body).Decode(&i); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	fmt.Printf("\n\nevent: %+v\n\n", i)
}

func (h *Hooks) ValidateKey(companyId, key, secret string) (bool, error) {
	if h.Config.Development {
		return true, nil
	}

	conn, err := grpc.Dial(h.Config.Local.KeyService.Address, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return false, err
	}
	defer func() {
		if err := conn.Close(); err != nil {
			fmt.Printf("failed to close connection: %v", err)
		}
	}()

	c := keybuf.NewKeyServiceClient(conn)
	resp, err := c.ValidateHookKey(context.Background(), &keybuf.ValidateSystemKeyRequest{
		ServiceKey: h.Config.Local.KeyService.Key,
		CompanyId:  companyId,
		Key:        key,
		Secret:     secret,
	})
	if err != nil {
		return false, err
	}

	if resp.Status != nil {
		return false, errors.New(*resp.Status)
	}

	if resp.Valid {
		return true, nil
	}

	return false, nil
}
