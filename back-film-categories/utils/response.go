package utils

import (
	"encoding/json"
	"net/http"

	"github.com/aws/aws-lambda-go/events"
)

func APISuccessResponse(data interface{}) (events.APIGatewayProxyResponse, error) {
	responseBody, err := json.Marshal(data)
	if err != nil {
		return APIErrorResponse(http.StatusInternalServerError, "Failed to marshal response"), nil
	}

	if data == nil {
		return events.APIGatewayProxyResponse{
			StatusCode: http.StatusNoContent,
			Headers: map[string]string{
				"Content-Type": "application/json",
			},
		}, nil
	}

	return events.APIGatewayProxyResponse{
		StatusCode: http.StatusOK,
		Headers: map[string]string{
			"Content-Type": "application/json",
		},
		Body: string(responseBody),
	}, nil
}

func APIErrorResponse(status int, message string) events.APIGatewayProxyResponse {
	body, err := json.Marshal(map[string]string{
		"message": message,
	})
	if err != nil {
		return events.APIGatewayProxyResponse{
			StatusCode: http.StatusInternalServerError,
			Headers: map[string]string{
				"Content-Type": "application/json",
			},
			Body: message,
		}
	}

	return events.APIGatewayProxyResponse{
		StatusCode: status,
		Headers: map[string]string{
			"Content-Type": "application/json",
		},
		Body: string(body),
	}
}
