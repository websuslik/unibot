package tg

import (
	"bytes"
	"encoding/json"
	"io"
	"mime/multipart"
	"os"
	"path/filepath"
	"reflect"
	"strconv"
	"strings"
)

func buildJSONRequestArgs(params MethodArgs) (*RequestArgs, error) {
	body, err := json.Marshal(params)
	if err != nil {
		return nil, err
	}
	headers := map[string]string{"Content-Type": "application/json"}
	return &RequestArgs{
		Body:    bytes.NewBuffer(body),
		Headers: headers,
	}, nil
}

func buildMultipartRequestArgs(args map[string]string, uploadFiles []*InputFile) (*RequestArgs, error) {
	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)
	for _, uploadFile := range uploadFiles {
		file, err := os.Open(uploadFile.FileName)
		if err != nil {
			return nil, err
		}
		part, err := writer.CreateFormFile(uploadFile.Name, filepath.Base(uploadFile.FileName))
		if err != nil {
			return nil, err
		}
		if _, err = io.Copy(part, file); err != nil {
			return nil, err
		}
		if err = file.Close(); err != nil {
			return nil, err
		}
	}
	for key, val := range args {
		_ = writer.WriteField(key, val)
	}
	if err := writer.Close(); err != nil {
		return nil, err
	}
	headers := map[string]string{"Content-Type": writer.FormDataContentType()}
	return &RequestArgs{
		Body:    body,
		Headers: headers,
	}, nil
}

func getTagKey(tag string) string {
	if idx := strings.Index(tag, ","); idx != -1 {
		return tag[:idx]
	}
	return tag
}

func marshallToMap(args interface{}) map[string]string {
	result := make(map[string]string)
	t := reflect.ValueOf(args).Elem()
	for i := 0; i < t.NumField(); i++ {
		value := t.Field(i).Interface()
		tag := t.Type().Field(i).Tag.Get("json")
		if tag == "-" {
			continue
		}
		key := getTagKey(tag)
		switch v := value.(type) {
		case string:
			if v != "" {
				result[key] = v
			}
		case ChatID:
			if v.ID != 0 {
				result[key] = strconv.Itoa(v.ID)
			} else {
				result[key] = v.Username
			}
		case int:
			if v != 0 {
				result[key] = strconv.Itoa(v)
			}
		case bool:
			result[key] = strconv.FormatBool(v)
		case []interface{}:
			if len(v) > 0 {
				val, _ := json.Marshal(v)
				result[key] = string(val)
			}
		default:
			if v != nil {
				val, _ := json.Marshal(v)
				result[key] = string(val)
			}
		}
	}
	return result
}

func buildOptionalMessage(response *json.RawMessage) (*OptionalMessage, error) {
	var success *bool
	var message *Message
	optMessage := OptionalMessage{Successful: true}
	if err := json.Unmarshal(*response, &success); err != nil {
		if err = json.Unmarshal(*response, &message); err != nil {
			return nil, err
		}
		optMessage.Message = message
		return &optMessage, nil
	}
	return &optMessage, nil
}
