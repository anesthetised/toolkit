package pbtypes

import (
	"time"

	"google.golang.org/protobuf/types/known/timestamppb"
)

// TimestampToNillableTime converts *timestamppb.Timestamp to *time.Time.
func TimestampToNillableTime(v *timestamppb.Timestamp) *time.Time {
	var expiresAt *time.Time

	if v != nil {
		tmp := v.AsTime()
		if tmp.IsZero() {
			return nil
		}
		expiresAt = &tmp
	}

	return expiresAt
}

// TimeToNillableTimestamp converts time.Time to *timestamppb.Timestamp.
func TimeToNillableTimestamp(v time.Time) *timestamppb.Timestamp {
	var expiresAt *timestamppb.Timestamp

	if !v.IsZero() {
		expiresAt = timestamppb.New(v)
	}

	return expiresAt
}
