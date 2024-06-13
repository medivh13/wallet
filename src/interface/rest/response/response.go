package response

import (
	"encoding/json"
	"math"
	"net/http"

	infra_errors "wallet/src/infra/errors"
)

// HealthCheckMessage displays information about the service
type HealthCheckMessage struct {
	ServiceName string `json:"serviceName"`
	Version     string `json:"version"`
	CommitId    string `json:"commitId"`
	UpdatedAt   string `json:"updatedAt"`
	Status      string `json:"status"`
}

// Meta consist of pagination details
type Meta struct {
	Skip  int     `json:"skip,omitempty"`
	Limit int     `json:"limit,omitempty"`
	Total float64 `json:"total,omitempty"`
}

// ResponseMessage consist of payload details
// Data -> Payload
// Meta -> Pagination etc
type ResponseMessage struct {
	Message string      `json:"message,omitempty"`
	Data    interface{} `json:"data,omitempty"`
	Meta    *Meta       `json:"meta,omitempty"`
}

type ErrorMessage struct {
	Message      string `json:"message"`
	ErrorMessage string `json:"errorMessage"`
	Type         string `json:"type"`
	Code         int    `json:"code"`
}

type ValidationErrorMessage struct {
	Message      string      `json:"message"`
	ErrorMessage interface{} `json:"errorMessage"`
	Type         string      `json:"type"`
	Code         int         `json:"code"`
}

type IResponseClient interface {
	JSON(w http.ResponseWriter, message string, data interface{}, meta *Meta) error
	BuildMeta(page int, perPage int, count int64) *Meta
	HttpError(w http.ResponseWriter, err error) error
}

type responseClient struct{}

func NewResponseClient() IResponseClient {
	return &responseClient{}
}

func (r *responseClient) JSON(
	w http.ResponseWriter,
	message string,
	data interface{},
	meta *Meta,
) error {
	response := ResponseMessage{
		Message: message,
		Data:    data,
		Meta:    meta,
	}
	if resp, err := json.Marshal(response); err != nil {
		return err
	} else {
		w.Header().Set("Content-Type", "application/json")
		w.Write(resp)
		return nil
	}
}

func (r *responseClient) HttpError(
	w http.ResponseWriter,
	err error,
) error {
	var respError infra_errors.HttpError

	if cerr, ok := err.(*infra_errors.CommonError); ok {
		respError = cerr.ToHttpError()
	} else {
		respError = infra_errors.NewError(infra_errors.UNKNOWN_ERROR, err).ToHttpError()
	}

	if resp, err := json.Marshal(respError); err != nil {
		return err
	} else {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(respError.GetHttpStatus())
		w.Write(resp)
		return nil
	}
}

func (r *responseClient) BuildMeta(skip int, limit int, count int64) *Meta {
	return &Meta{
		Skip:  skip,
		Limit: limit,
		Total: math.Ceil(float64(count) / float64(limit)),
	}
}
