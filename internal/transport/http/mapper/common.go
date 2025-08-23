package mapper

import "github.com/ent1k1377/wallet/internal/transport/http/dto"

func ToErrorResponse(err string) dto.ErrorResponse {
	return dto.ErrorResponse{
		Error: err,
	}
}

func ToSuccessResponse(message string) dto.SuccessResponse {
	return dto.SuccessResponse{
		Message: message,
	}
}
