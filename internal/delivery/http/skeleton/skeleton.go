package skeleton

import (
	"context"
	"encoding/json"
	"errors"
	sentity "go-skeleton/internal/entity/skeleton"
	"go-skeleton/pkg/response"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

// ISkeletonSvc is an interface to Skeleton Service
// Masukkan function dari service ke dalam interface ini
type ISkeletonSvc interface {
	GetAllUser(ctx context.Context) ([]sentity.Skeleton, error)
	GetAllNip(ctx context.Context, nip string) (sentity.Skeleton, error)
	InsertAllUser(ctx context.Context, sk sentity.Skeleton) error
	UpdateAllUser(ctx context.Context, sk sentity.Skeleton) (sentity.Skeleton, error)
	DeleteAllUser(ctx context.Context, sk sentity.Skeleton) error
}

type (
	// Handler ...
	Handler struct {
		skeletonSvc ISkeletonSvc
	}
)

// New for bridging product handler initialization
func New(is ISkeletonSvc) *Handler {
	return &Handler{
		skeletonSvc: is,
	}
}

// SkeletonHandler will receive request and return response
func (h *Handler) SkeletonHandler(w http.ResponseWriter, r *http.Request) {
	var (
		resp     *response.Response
		result   interface{}
		sktemp   sentity.Skeleton
		metadata interface{}
		err      error
		errRes   response.Error
	)
	resp = &response.Response{}
	body, _ := ioutil.ReadAll(r.Body)
	defer resp.RenderJSON(w, r)

	switch r.Method {
	// Check if request method is GET
	case http.MethodGet:

		paramMap := r.URL.Query()
		len := len(paramMap)
		switch len {

		case 1:
			_, nipOK := paramMap["nip"]
			if nipOK {
				nip := r.FormValue("nip")

				result, err = h.skeletonSvc.GetAllNip(context.Background(), nip)

			}
		case 0:
			result, err = h.skeletonSvc.GetAllUser(context.Background())
		}

	// Check if request method is POST
	case http.MethodPost:

		json.Unmarshal(body, &sktemp)
		err = h.skeletonSvc.InsertAllUser(context.Background(), sktemp)
	// Check if request method is PUT
	case http.MethodPut:
		json.NewDecoder(r.Body).Decode(&sktemp)
		result, err = h.skeletonSvc.UpdateAllUser(context.Background(), sktemp)
	// Check if request method is DELETE
	case http.MethodDelete:
		json.NewDecoder(r.Body).Decode(&sktemp)
		err = h.skeletonSvc.DeleteAllUser(context.Background(), sktemp)

	default:
		err = errors.New("404")
	}

	// If anything from service or data return an error
	if err != nil {
		// Error response handling
		errRes = response.Error{
			Code:   101,
			Msg:    "Data Not Found",
			Status: true,
		}
		// If service returns an error
		if strings.Contains(err.Error(), "service") {
			// Replace error with server error
			errRes = response.Error{
				Code:   201,
				Msg:    "Failed to process request due to server error",
				Status: true,
			}
		}

		log.Printf("[ERROR] %s %s - %v\n", r.Method, r.URL, err)
		resp.Error = errRes
		return
	}

	resp.Data = result
	resp.Metadata = metadata
	log.Printf("[INFO] %s %s\n", r.Method, r.URL)
	return
}
