package error

import (
	"encoding/json"
	"invs-common/src/model"
	"net/http"
)

func HttpException(w http.ResponseWriter, code int, msg string) {
	resp := model.BaseResponse[string]{
		Code:    code,
		Message: msg,
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(resp)
}
