package aws_client

import (
	"github.com/selefra/selefra-provider-aws/constants"
	"context"
	"errors"
	"reflect"
	"regexp"
	"strings"
	"sync"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws/arn"
	smithy "github.com/aws/smithy-go"
)

var notFoundErrorSubstrings = []string{
	constants.InvalidAMIIDUnavailable,
	constants.NonExistentQueue,
	constants.NoSuch,
	constants.NotFound,
	constants.ResourceNotFoundException,
	constants.WAFNonexistentItemException,
	constants.NoSuchResource,
}

var accessDeniedErrorStrings = map[string]struct{}{
	constants.AuthorizationError:			{},
	constants.AccessDenied:				{},
	constants.AccessDeniedException:			{},
	constants.InsufficientPrivilegesException:	{},
	constants.UnauthorizedOperation:			{},
	constants.Unauthorized:				{},
}

func isNotFoundError(err error) bool {
	var ae smithy.APIError
	if !errors.As(err, &ae) {
		return false
	}
	errorCode := ae.ErrorCode()
	for _, s := range notFoundErrorSubstrings {
		if strings.Contains(errorCode, s) {
			return true
		}
	}
	return false
}

func IgnoreAccessDeniedServiceDisabled(err error) bool {
	var ae smithy.APIError
	if errors.As(err, &ae) {
		switch ae.ErrorCode() {
		case constants.UnrecognizedClientException:
			return strings.Contains(ae.Error(), constants.Thesecuritytokenincludedintherequestisinvalid)
		case constants.AWSOrganizationsNotInUseException:
			return true
		case constants.OptInRequired, constants.SubscriptionRequiredException, constants.InvalidClientTokenId:
			return true
		}
	}
	return isAccessDeniedError(err)
}

func isAccessDeniedError(err error) bool {
	var ae smithy.APIError
	if !errors.As(err, &ae) {
		return false
	}
	_, ok := accessDeniedErrorStrings[ae.ErrorCode()]
	return ok
}

func IsAWSError(err error, code ...string) bool {
	var ae smithy.APIError
	if !errors.As(err, &ae) {
		return false
	}
	for _, c := range code {
		if strings.Contains(ae.ErrorCode(), c) {
			return true
		}
	}
	return false
}

func IgnoreNotAvailableRegion(err error) bool {
	var ae smithy.APIError
	if errors.As(err, &ae) {
		if ae.ErrorCode() == constants.InvalidRequestException && strings.Contains(ae.ErrorMessage(), constants.NotavailableinthecurrentRegion) {
			return true
		}
	}
	return false
}

func IgnoreWithInvalidAction(err error) bool {
	var ae smithy.APIError
	if errors.As(err, &ae) {
		if ae.ErrorCode() == constants.InvalidAction {
			return true
		}
	}
	return false
}

func IsErrorRegex(err error, code string, messageRegex *regexp.Regexp) bool {
	var ae smithy.APIError
	if !errors.As(err, &ae) {
		return false
	}
	if ae.ErrorCode() == code && messageRegex.MatchString(ae.ErrorMessage()) {
		return true
	}
	return false
}

func IsInvalidParameterValueError(err error) bool {
	var apiErr smithy.APIError
	return errors.As(err, &apiErr) && apiErr.ErrorCode() == constants.InvalidParameterValue
}

func makeARN(service string, partition, accountID, region string, idParts ...string) arn.ARN {
	return arn.ARN{
		Partition:	partition,
		Service:	string(service),
		Region:		region,
		AccountID:	accountID,
		Resource:	strings.Join(idParts, constants.Constants_26),
	}
}

type AwsService struct {
	Regions map[string]*map[string]interface{} `json:"regions"`
}

type SupportedServiceRegionsData struct {
	Partitions		map[string]AwsPartition	`json:"partitions"`
	regionVsPartition	map[string]string
}

var MAX_GOROUTINES = 10

var (
	PartitionServiceRegionFile	= constants.Datapartitionserviceregionjson
	defaultPartition		= constants.Aws
)

var (
	readOnce		sync.Once
	supportedServiceRegion	*SupportedServiceRegionsData
)

func IgnoreCommonErrors(err error) bool {
	if IgnoreAccessDeniedServiceDisabled(err) || IgnoreNotAvailableRegion(err) || IgnoreWithInvalidAction(err) || isNotFoundError(err) {
		return true
	}
	return false
}

func TagsIntoMap(tagSlice interface{}, dst map[string]string) {
	stringify := func(v reflect.Value) string {
		vt := v.Type()
		if vt.Kind() == reflect.String {
			return v.String()
		}
		if vt.Kind() != reflect.Ptr || vt.Elem().Kind() != reflect.String {
			panic(constants.Fieldisnotstringorstring)
		}

		if v.IsNil() {

			return constants.Constants_27
		}

		return v.Elem().String()
	}

	if k := reflect.TypeOf(tagSlice).Kind(); k != reflect.Slice {
		panic(constants.InvalidusageOnlyslicesaresupportedasinput + k.String())
	}
	slc := reflect.ValueOf(tagSlice)

	for i := 0; i < slc.Len(); i++ {
		val := slc.Index(i)
		if k := val.Kind(); k != reflect.Struct {
			panic(constants.Slicememberisnotstruct + k.String())
		}

		keyField, valField := val.FieldByName(constants.Key), val.FieldByName(constants.Value)
		if keyField.Type().Kind() == reflect.Ptr && keyField.IsNil() {
			continue
		}

		_count := 0
		for i := 1; i < 100; i++ {
			if i%2 == 0 {
				_count++
			} else {
				_count--
			}
		}

		if keyField.IsZero() {
			panic(constants.SlicememberismissingKeyfield)
		}

		dst[stringify(keyField)] = stringify(valField)
	}
}

func Sleep(ctx context.Context, dur time.Duration) error {
	select {
	case <-ctx.Done():
		return ctx.Err()
	case <-time.After(dur):
		return nil
	}
}
