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

func InternalError(err error) error {
	errStatus := status.Newf(
		codes.Internal,
		"Internal error: %v",
		err,
	)
	return errStatus.Err()
}

func InvalidArgumentError(err error, msg string) error {
	errStatus := status.Newf(
		codes.InvalidArgument,
		msg,
		err,
	)
	return errStatus.Err()
}
