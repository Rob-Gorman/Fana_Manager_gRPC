package utils

import (
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func NotFoundError(err error) error {
	errStatus := status.Newf(
		codes.NotFound,
		"Error retrieving resource: %v",
		err,
	)
	return errStatus.Err()
}
