package jsonrpc

import (
	"bytes"
	"context"
	"net/http"
	"sync"

	"github.com/golang/protobuf/jsonpb"
)

// Repo Репозиторий сервисов
type Repo struct {
	Marshaler *jsonpb.Marshaler
	methods   map[string]Service
}

// New Новый репозиторий
func New() *Repo {
	return &Repo{Marshaler: new(jsonpb.Marshaler), methods: make(map[string]Service)}
}

// RegisterService Зарегистрировать сервис
func (repo *Repo) RegisterService(service Service) {
	methods := service.Methods()
	for _, method := range methods {
		repo.methods[method] = service
	}
}

// takeService Получить сервис по имени метода
func (repo *Repo) takeService(methodName string) (Service, bool) {
	service, exist := repo.methods[methodName]
	return service, exist
}

// ServeHTTP .
func (repo *Repo) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	body, batch, err := prepareBody(r)
	if err != nil {
		sendResponse(w, batch, errorResponse(nil, ErrParseError(err.Error())))
		return
	}
	defer r.Body.Close()

	requests, err := prepareRequests(body, batch)
	if err != nil {
		sendResponse(w, batch, errorResponse(nil, ErrParseError(err.Error())))
		return
	}

	responses, err := repo.handleRequests(r.Context(), requests)
	if err != nil {
		sendResponse(w, batch, errorResponse(nil, ErrInternalError(err.Error())))
		return
	}

	sendResponse(w, batch, responses...)
}

func (repo *Repo) handleRequests(ctx context.Context, requests []*request) ([]*response, error) {
	responses := make([]*response, 0, len(requests))

	wg := new(sync.WaitGroup)
	responseChan := make(chan *response, len(requests))

	for _, req := range requests {
		wg.Add(1)
		go func(ctx context.Context, wg *sync.WaitGroup, req *request, result chan<- *response) {
			defer wg.Done()
			res := repo.handleRequest(ctx, req)
			result <- res
		}(ctx, wg, req, responseChan)
	}

	wg.Wait()
	close(responseChan)

	for res := range responseChan {
		if res != nil {
			responses = append(responses, res)
		}
	}

	return responses, nil
}

func (repo *Repo) handleRequest(ctx context.Context, req *request) *response {
	if err := req.validate(); err != nil {
		return errorResponse(req.ID, ErrInvalidRequest(err.Error()))
	}

	ctx = context.WithValue(ctx, requestID, req.ID)
	params := bytes.NewBuffer(req.Params)

	service, exist := repo.takeService(req.Method)
	if !exist {
		return errorResponse(req.ID, ErrMethodNotFound(nil))
	}

	out, methodErr := service.Exec(ctx, req.Method, params)
	if methodErr != nil {
		return errorResponse(req.ID, methodErr)
	}

	if req.isNotification() {
		return nil
	}

	return successResponse(repo.Marshaler, req.ID, out)
}
