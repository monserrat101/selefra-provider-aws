package apigatewayv2fix

import (
	"github.com/selefra/selefra-provider-aws/constants"
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"strings"

	"github.com/aws/aws-sdk-go-v2/aws/protocol/restjson"
	"github.com/aws/aws-sdk-go-v2/service/apigatewayv2"
	"github.com/aws/aws-sdk-go-v2/service/apigatewayv2/types"
	smithy "github.com/aws/smithy-go"
	smithyio "github.com/aws/smithy-go/io"
	"github.com/aws/smithy-go/middleware"
	"github.com/aws/smithy-go/ptr"
	smithytime "github.com/aws/smithy-go/time"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

func SwapGetDomainNamesOperationDeserializer(stack *middleware.Stack) error {
	m := &awsRestjson1_deserializeOpGetDomainNames{}
	_, err := stack.Deserialize.Swap(m.ID(), m)
	return err
}

type awsRestjson1_deserializeOpGetDomainNames struct {
}

func (*awsRestjson1_deserializeOpGetDomainNames) ID() string {
	return constants.OperationDeserializer
}

func (m *awsRestjson1_deserializeOpGetDomainNames) HandleDeserialize(ctx context.Context, in middleware.DeserializeInput, next middleware.DeserializeHandler) (
	out middleware.DeserializeOutput, metadata middleware.Metadata, err error,
) {
	out, metadata, err = next.HandleDeserialize(ctx, in)
	if err != nil {
		return out, metadata, err
	}

	response, ok := out.RawResponse.(*smithyhttp.Response)
	if !ok {
		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf(constants.UnknowntransporttypeT, out.RawResponse)}
	}

	if response.StatusCode < 200 || response.StatusCode >= 300 {
		return out, metadata, awsRestjson1_deserializeOpErrorGetDomainNames(response, &metadata)
	}
	output := &apigatewayv2.GetDomainNamesOutput{}
	out.Result = output

	var buff [1024]byte
	ringBuffer := smithyio.NewRingBuffer(buff[:])

	body := io.TeeReader(response.Body, ringBuffer)

	decoder := json.NewDecoder(body)
	decoder.UseNumber()
	var shape interface{}
	if err := decoder.Decode(&shape); err != nil && err != io.EOF {
		var snapshot bytes.Buffer
		io.Copy(&snapshot, ringBuffer)
		err = &smithy.DeserializationError{
			Err:		fmt.Errorf(constants.Failedtodecoderesponsebodyw, err),
			Snapshot:	snapshot.Bytes(),
		}
		return out, metadata, err
	}

	err = awsRestjson1_deserializeOpDocumentGetDomainNamesOutput(&output, shape)
	if err != nil {
		var snapshot bytes.Buffer
		io.Copy(&snapshot, ringBuffer)
		return out, metadata, &smithy.DeserializationError{
			Err:		fmt.Errorf(constants.FailedtodecoderesponsebodywithinvalidJSONw, err),
			Snapshot:	snapshot.Bytes(),
		}
	}

	return out, metadata, err
}

func awsRestjson1_deserializeOpErrorGetDomainNames(response *smithyhttp.Response, metadata *middleware.Metadata) error {
	var errorBuffer bytes.Buffer
	if _, err := io.Copy(&errorBuffer, response.Body); err != nil {
		return &smithy.DeserializationError{Err: fmt.Errorf(constants.Failedtocopyerrorresponsebodyw, err)}
	}
	errorBody := bytes.NewReader(errorBuffer.Bytes())

	errorCode := constants.UnknownError
	errorMessage := errorCode

	code := response.Header.Get(constants.XAmznErrorType)
	if len(code) != 0 {
		errorCode = restjson.SanitizeErrorCode(code)
	}

	var buff [1024]byte
	ringBuffer := smithyio.NewRingBuffer(buff[:])

	body := io.TeeReader(errorBody, ringBuffer)
	decoder := json.NewDecoder(body)
	decoder.UseNumber()
	code, message, err := restjson.GetErrorInfo(decoder)
	if err != nil {
		var snapshot bytes.Buffer
		io.Copy(&snapshot, ringBuffer)
		err = &smithy.DeserializationError{
			Err:		fmt.Errorf(constants.Failedtodecoderesponsebodyw, err),
			Snapshot:	snapshot.Bytes(),
		}
		return err
	}

	errorBody.Seek(0, io.SeekStart)
	if len(code) != 0 {
		errorCode = restjson.SanitizeErrorCode(code)
	}
	if len(message) != 0 {
		errorMessage = message
	}

	switch {
	case strings.EqualFold(constants.BadRequestException, errorCode):
		return awsRestjson1_deserializeErrorBadRequestException(response, errorBody)

	case strings.EqualFold(constants.NotFoundException, errorCode):
		return awsRestjson1_deserializeErrorNotFoundException(response, errorBody)

	case strings.EqualFold(constants.TooManyRequestsException, errorCode):
		return awsRestjson1_deserializeErrorTooManyRequestsException(response, errorBody)

	default:
		genericError := &smithy.GenericAPIError{
			Code:		errorCode,
			Message:	errorMessage,
		}
		return genericError

	}
}

func awsRestjson1_deserializeOpDocumentGetDomainNamesOutput(v **apigatewayv2.GetDomainNamesOutput, value interface{}) error {
	if v == nil {
		return fmt.Errorf(constants.UnexpectedniloftypeT, v)
	}
	if value == nil {
		return nil
	}

	shape, ok := value.(map[string]interface{})
	if !ok {
		return fmt.Errorf(constants.UnexpectedJSONtypev, value)
	}

	var sv *apigatewayv2.GetDomainNamesOutput
	if *v == nil {
		sv = &apigatewayv2.GetDomainNamesOutput{}
	} else {
		sv = *v
	}

	for key, value := range shape {
		switch key {
		case constants.Items:
			if err := awsRestjson1_deserializeDocument__listOfDomainName(&sv.Items, value); err != nil {
				return err
			}

		case constants.NextToken:
			if value != nil {
				jtv, ok := value.(string)
				if !ok {
					return fmt.Errorf(constants.ExpectedNextTokentobeoftypestringgotTinstead, value)
				}
				sv.NextToken = ptr.String(jtv)
			}

		default:
			_, _ = key, value

		}
	}
	*v = sv
	return nil
}

func awsRestjson1_deserializeErrorBadRequestException(response *smithyhttp.Response, errorBody *bytes.Reader) error {
	output := &types.BadRequestException{}
	var buff [1024]byte
	ringBuffer := smithyio.NewRingBuffer(buff[:])

	body := io.TeeReader(errorBody, ringBuffer)
	decoder := json.NewDecoder(body)
	decoder.UseNumber()
	var shape interface{}
	if err := decoder.Decode(&shape); err != nil && err != io.EOF {
		var snapshot bytes.Buffer
		io.Copy(&snapshot, ringBuffer)
		err = &smithy.DeserializationError{
			Err:		fmt.Errorf(constants.Failedtodecoderesponsebodyw, err),
			Snapshot:	snapshot.Bytes(),
		}
		return err
	}

	err := awsRestjson1_deserializeDocumentBadRequestException(&output, shape)

	if err != nil {
		var snapshot bytes.Buffer
		io.Copy(&snapshot, ringBuffer)
		err = &smithy.DeserializationError{
			Err:		fmt.Errorf(constants.Failedtodecoderesponsebodyw, err),
			Snapshot:	snapshot.Bytes(),
		}
		return err
	}

	errorBody.Seek(0, io.SeekStart)

	return output
}

func awsRestjson1_deserializeErrorNotFoundException(response *smithyhttp.Response, errorBody *bytes.Reader) error {
	output := &types.NotFoundException{}
	var buff [1024]byte
	ringBuffer := smithyio.NewRingBuffer(buff[:])

	body := io.TeeReader(errorBody, ringBuffer)
	decoder := json.NewDecoder(body)
	decoder.UseNumber()
	var shape interface{}
	if err := decoder.Decode(&shape); err != nil && err != io.EOF {
		var snapshot bytes.Buffer
		io.Copy(&snapshot, ringBuffer)
		err = &smithy.DeserializationError{
			Err:		fmt.Errorf(constants.Failedtodecoderesponsebodyw, err),
			Snapshot:	snapshot.Bytes(),
		}
		return err
	}

	err := awsRestjson1_deserializeDocumentNotFoundException(&output, shape)

	if err != nil {
		var snapshot bytes.Buffer
		io.Copy(&snapshot, ringBuffer)
		err = &smithy.DeserializationError{
			Err:		fmt.Errorf(constants.Failedtodecoderesponsebodyw, err),
			Snapshot:	snapshot.Bytes(),
		}
		return err
	}

	errorBody.Seek(0, io.SeekStart)

	return output
}

func awsRestjson1_deserializeErrorTooManyRequestsException(response *smithyhttp.Response, errorBody *bytes.Reader) error {
	output := &types.TooManyRequestsException{}
	var buff [1024]byte
	ringBuffer := smithyio.NewRingBuffer(buff[:])

	body := io.TeeReader(errorBody, ringBuffer)
	decoder := json.NewDecoder(body)
	decoder.UseNumber()
	var shape interface{}
	if err := decoder.Decode(&shape); err != nil && err != io.EOF {
		var snapshot bytes.Buffer
		io.Copy(&snapshot, ringBuffer)
		err = &smithy.DeserializationError{
			Err:		fmt.Errorf(constants.Failedtodecoderesponsebodyw, err),
			Snapshot:	snapshot.Bytes(),
		}
		return err
	}

	err := awsRestjson1_deserializeDocumentTooManyRequestsException(&output, shape)

	if err != nil {
		var snapshot bytes.Buffer
		io.Copy(&snapshot, ringBuffer)
		err = &smithy.DeserializationError{
			Err:		fmt.Errorf(constants.Failedtodecoderesponsebodyw, err),
			Snapshot:	snapshot.Bytes(),
		}
		return err
	}

	errorBody.Seek(0, io.SeekStart)

	return output
}

func awsRestjson1_deserializeDocument__listOfDomainName(v *[]types.DomainName, value interface{}) error {
	if v == nil {
		return fmt.Errorf(constants.UnexpectedniloftypeT, v)
	}
	if value == nil {
		return nil
	}

	shape, ok := value.([]interface{})
	if !ok {
		return fmt.Errorf(constants.UnexpectedJSONtypev, value)
	}

	var cv []types.DomainName
	if *v == nil {
		cv = []types.DomainName{}
	} else {
		cv = *v
	}

	for _, value := range shape {
		var col types.DomainName
		destAddr := &col
		if err := awsRestjson1_deserializeDocumentDomainName(&destAddr, value); err != nil {
			return err
		}
		col = *destAddr
		cv = append(cv, col)

	}
	*v = cv
	return nil
}

func awsRestjson1_deserializeDocumentBadRequestException(v **types.BadRequestException, value interface{}) error {
	if v == nil {
		return fmt.Errorf(constants.UnexpectedniloftypeT, v)
	}
	if value == nil {
		return nil
	}

	shape, ok := value.(map[string]interface{})
	if !ok {
		return fmt.Errorf(constants.UnexpectedJSONtypev, value)
	}

	var sv *types.BadRequestException
	if *v == nil {
		sv = &types.BadRequestException{}
	} else {
		sv = *v
	}

	for key, value := range shape {
		switch key {
		case constants.Message:
			if value != nil {
				jtv, ok := value.(string)
				if !ok {
					return fmt.Errorf(constants.ExpectedstringtobeoftypestringgotTinstead, value)
				}
				sv.Message = ptr.String(jtv)
			}

		default:
			_, _ = key, value

		}
	}
	*v = sv
	return nil
}

func awsRestjson1_deserializeDocumentTooManyRequestsException(v **types.TooManyRequestsException, value interface{}) error {
	if v == nil {
		return fmt.Errorf(constants.UnexpectedniloftypeT, v)
	}
	if value == nil {
		return nil
	}

	shape, ok := value.(map[string]interface{})
	if !ok {
		return fmt.Errorf(constants.UnexpectedJSONtypev, value)
	}

	var sv *types.TooManyRequestsException
	if *v == nil {
		sv = &types.TooManyRequestsException{}
	} else {
		sv = *v
	}

	for key, value := range shape {
		switch key {
		case constants.LimitType:
			if value != nil {
				jtv, ok := value.(string)
				if !ok {
					return fmt.Errorf(constants.ExpectedstringtobeoftypestringgotTinstead, value)
				}
				sv.LimitType = ptr.String(jtv)
			}

		case constants.Message:
			if value != nil {
				jtv, ok := value.(string)
				if !ok {
					return fmt.Errorf(constants.ExpectedstringtobeoftypestringgotTinstead, value)
				}
				sv.Message = ptr.String(jtv)
			}

		default:
			_, _ = key, value

		}
	}
	*v = sv
	return nil
}

func awsRestjson1_deserializeDocumentDomainName(v **types.DomainName, value interface{}) error {
	if v == nil {
		return fmt.Errorf(constants.UnexpectedniloftypeT, v)
	}
	if value == nil {
		return nil
	}

	shape, ok := value.(map[string]interface{})
	if !ok {
		return fmt.Errorf(constants.UnexpectedJSONtypev, value)
	}

	var sv *types.DomainName
	if *v == nil {
		sv = &types.DomainName{}
	} else {
		sv = *v
	}

	for key, value := range shape {
		switch key {
		case constants.ApiMappingSelectionExpression:
			if value != nil {
				jtv, ok := value.(string)
				if !ok {
					return fmt.Errorf(constants.ExpectedSelectionExpressiontobeoftypestringgotTinstead, value)
				}
				sv.ApiMappingSelectionExpression = ptr.String(jtv)
			}

		case constants.DomainName:
			if value != nil {
				jtv, ok := value.(string)
				if !ok {
					return fmt.Errorf(constants.ExpectedStringWithLengthBetweenAndtobeoftypestringgotTinstead, value)
				}
				sv.DomainName = ptr.String(jtv)
			}

		case constants.DomainNameConfigurations:
			if err := awsRestjson1_deserializeDocumentDomainNameConfigurations(&sv.DomainNameConfigurations, value); err != nil {
				return err
			}

		case constants.MutualTlsAuthentication:
			if err := awsRestjson1_deserializeDocumentMutualTlsAuthentication(&sv.MutualTlsAuthentication, value); err != nil {
				return err
			}

		case constants.Tags:
			if err := awsRestjson1_deserializeDocumentTags(&sv.Tags, value); err != nil {
				return err
			}

		default:
			_, _ = key, value

		}
	}
	*v = sv
	return nil
}

func awsRestjson1_deserializeDocumentDomainNameConfigurations(v *[]types.DomainNameConfiguration, value interface{}) error {
	if v == nil {
		return fmt.Errorf(constants.UnexpectedniloftypeT, v)
	}
	if value == nil {
		return nil
	}

	shape, ok := value.([]interface{})
	if !ok {
		return fmt.Errorf(constants.UnexpectedJSONtypev, value)
	}

	var cv []types.DomainNameConfiguration
	if *v == nil {
		cv = []types.DomainNameConfiguration{}
	} else {
		cv = *v
	}

	for _, value := range shape {
		var col types.DomainNameConfiguration
		destAddr := &col
		if err := awsRestjson1_deserializeDocumentDomainNameConfiguration(&destAddr, value); err != nil {
			return err
		}
		col = *destAddr
		cv = append(cv, col)

	}
	*v = cv
	return nil
}

func awsRestjson1_deserializeDocumentMutualTlsAuthentication(v **types.MutualTlsAuthentication, value interface{}) error {
	if v == nil {
		return fmt.Errorf(constants.UnexpectedniloftypeT, v)
	}
	if value == nil {
		return nil
	}

	shape, ok := value.(map[string]interface{})
	if !ok {
		return fmt.Errorf(constants.UnexpectedJSONtypev, value)
	}

	var sv *types.MutualTlsAuthentication
	if *v == nil {
		sv = &types.MutualTlsAuthentication{}
	} else {
		sv = *v
	}

	for key, value := range shape {
		switch key {
		case constants.TruststoreUri:
			if value != nil {
				jtv, ok := value.(string)
				if !ok {
					return fmt.Errorf(constants.ExpectedUriWithLengthBetweenAndtobeoftypestringgotTinstead, value)
				}
				sv.TruststoreUri = ptr.String(jtv)
			}

		case constants.TruststoreVersion:
			if value != nil {
				jtv, ok := value.(string)
				if !ok {
					return fmt.Errorf("expected StringWithLengthBetween1And64 to be of type string, got %T instead", value)
				}
				sv.TruststoreVersion = ptr.String(jtv)
			}

		case constants.TruststoreWarnings:
			if err := awsRestjson1_deserializeDocument__listOf__string(&sv.TruststoreWarnings, value); err != nil {
				return err
			}

		default:
			_, _ = key, value

		}
	}
	*v = sv
	return nil
}

func awsRestjson1_deserializeDocumentTags(v *map[string]string, value interface{}) error {
	if v == nil {
		return fmt.Errorf(constants.UnexpectedniloftypeT, v)
	}
	if value == nil {
		return nil
	}

	shape, ok := value.(map[string]interface{})
	if !ok {
		return fmt.Errorf(constants.UnexpectedJSONtypev, value)
	}

	var mv map[string]string
	if *v == nil {
		mv = map[string]string{}
	} else {
		mv = *v
	}

	for key, value := range shape {
		var parsedVal string
		if value != nil {
			jtv, ok := value.(string)
			if !ok {
				return fmt.Errorf("expected StringWithLengthBetween1And1600 to be of type string, got %T instead", value)
			}
			parsedVal = jtv
		}
		mv[key] = parsedVal

	}
	*v = mv
	return nil
}

func awsRestjson1_deserializeDocumentNotFoundException(v **types.NotFoundException, value interface{}) error {
	if v == nil {
		return fmt.Errorf(constants.UnexpectedniloftypeT, v)
	}
	if value == nil {
		return nil
	}

	shape, ok := value.(map[string]interface{})
	if !ok {
		return fmt.Errorf(constants.UnexpectedJSONtypev, value)
	}

	var sv *types.NotFoundException
	if *v == nil {
		sv = &types.NotFoundException{}
	} else {
		sv = *v
	}

	for key, value := range shape {
		switch key {
		case constants.Message:
			if value != nil {
				jtv, ok := value.(string)
				if !ok {
					return fmt.Errorf(constants.ExpectedstringtobeoftypestringgotTinstead, value)
				}
				sv.Message = ptr.String(jtv)
			}

		case constants.ResourceType:
			if value != nil {
				jtv, ok := value.(string)
				if !ok {
					return fmt.Errorf(constants.ExpectedstringtobeoftypestringgotTinstead, value)
				}
				sv.ResourceType = ptr.String(jtv)
			}

		default:
			_, _ = key, value

		}
	}
	*v = sv
	return nil
}

func awsRestjson1_deserializeDocumentDomainNameConfiguration(v **types.DomainNameConfiguration, value interface{}) error {
	if v == nil {
		return fmt.Errorf(constants.UnexpectedniloftypeT, v)
	}
	if value == nil {
		return nil
	}

	shape, ok := value.(map[string]interface{})
	if !ok {
		return fmt.Errorf(constants.UnexpectedJSONtypev, value)
	}

	var sv *types.DomainNameConfiguration
	if *v == nil {
		sv = &types.DomainNameConfiguration{}
	} else {
		sv = *v
	}

	for key, value := range shape {
		switch key {
		case constants.ApiGatewayDomainName:
			if value != nil {
				jtv, ok := value.(string)
				if !ok {
					return fmt.Errorf(constants.ExpectedstringtobeoftypestringgotTinstead, value)
				}
				sv.ApiGatewayDomainName = ptr.String(jtv)
			}

		case constants.CertificateArn:
			if value != nil {
				jtv, ok := value.(string)
				if !ok {
					return fmt.Errorf(constants.ExpectedArntobeoftypestringgotTinstead, value)
				}
				sv.CertificateArn = ptr.String(jtv)
			}

		case constants.CertificateName:
			if value != nil {
				jtv, ok := value.(string)
				if !ok {
					return fmt.Errorf("expected StringWithLengthBetween1And128 to be of type string, got %T instead", value)
				}
				sv.CertificateName = ptr.String(jtv)
			}

		case constants.CertificateUploadDate:
			if value != nil {

				jtv, ok := value.(json.Number)
				if !ok {
					return fmt.Errorf(constants.ExpectedjsonNumbertobeoftypestringgotTinstead, value)
				}
				f, err := jtv.Float64()
				if err != nil {
					return err
				}
				t := smithytime.ParseEpochSeconds(f)
				sv.CertificateUploadDate = ptr.Time(t)
			}

		case constants.DomainNameStatus:
			if value != nil {
				jtv, ok := value.(string)
				if !ok {
					return fmt.Errorf(constants.ExpectedDomainNameStatustobeoftypestringgotTinstead, value)
				}
				sv.DomainNameStatus = types.DomainNameStatus(jtv)
			}

		case constants.DomainNameStatusMessage:
			if value != nil {
				jtv, ok := value.(string)
				if !ok {
					return fmt.Errorf(constants.ExpectedstringtobeoftypestringgotTinstead, value)
				}
				sv.DomainNameStatusMessage = ptr.String(jtv)
			}

		case constants.EndpointType:
			if value != nil {
				jtv, ok := value.(string)
				if !ok {
					return fmt.Errorf(constants.ExpectedEndpointTypetobeoftypestringgotTinstead, value)
				}
				sv.EndpointType = types.EndpointType(jtv)
			}

		case constants.HostedZoneId:
			if value != nil {
				jtv, ok := value.(string)
				if !ok {
					return fmt.Errorf(constants.ExpectedstringtobeoftypestringgotTinstead, value)
				}
				sv.HostedZoneId = ptr.String(jtv)
			}

		case constants.SecurityPolicy:
			if value != nil {
				jtv, ok := value.(string)
				if !ok {
					return fmt.Errorf(constants.ExpectedSecurityPolicytobeoftypestringgotTinstead, value)
				}
				sv.SecurityPolicy = types.SecurityPolicy(jtv)
			}

		default:
			_, _ = key, value

		}
	}
	*v = sv
	return nil
}

func awsRestjson1_deserializeDocument__listOf__string(v *[]string, value interface{}) error {
	if v == nil {
		return fmt.Errorf(constants.UnexpectedniloftypeT, v)
	}
	if value == nil {
		return nil
	}

	shape, ok := value.([]interface{})
	if !ok {
		return fmt.Errorf(constants.UnexpectedJSONtypev, value)
	}

	var cv []string
	if *v == nil {
		cv = []string{}
	} else {
		cv = *v
	}

	for _, value := range shape {
		var col string
		if value != nil {
			jtv, ok := value.(string)
			if !ok {
				return fmt.Errorf(constants.ExpectedstringtobeoftypestringgotTinstead, value)
			}
			col = jtv
		}
		cv = append(cv, col)

	}
	*v = cv
	return nil
}
