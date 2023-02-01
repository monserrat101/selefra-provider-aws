package aws_client

import (
	"github.com/selefra/selefra-provider-aws/constants"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/rs/zerolog"
)

type retryer struct {
	aws.Retryer
	logger	zerolog.Logger
}

func newRetryer(logger zerolog.Logger, maxRetries int, maxBackoff int) func() aws.Retryer {
	return func() aws.Retryer {
		return &retryer{
			Retryer: retry.NewStandard(func(o *retry.StandardOptions) {
				o.MaxAttempts = maxRetries
				o.MaxBackoff = time.Second * time.Duration(maxBackoff)
			}),
			logger:	logger,
		}
	}
}

func (r *retryer) RetryDelay(attempt int, err error) (time.Duration, error) {
	dur, retErr := r.Retryer.RetryDelay(attempt, err)

	logParams := []interface{}{
		constants.Duration, dur.String(),
		constants.Attempt, attempt,
		constants.Err, err,
	}
	if retErr != nil {
		logParams = append(logParams, constants.Retriererr, retErr)
		r.logger.Debug().Err(retErr).Interface(constants.LogParams, logParams).Msg(constants.RetryDelayreturnederr)
	} else {
		r.logger.Debug().Interface(constants.LogParams, logParams).Msg(constants.Waitingbeforeretry)
	}
	return dur, retErr
}
