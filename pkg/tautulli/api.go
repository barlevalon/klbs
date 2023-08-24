package tautulli

import (
	"fmt"

	"github.com/imroc/req/v3"
	"github.com/spf13/viper"
)

type GetActivityError struct {
  Message string
  Response string
}

func (e *GetActivityError) Error() string {
  return fmt.Sprintf("Tautulli error: %s Raw response: %s", e.Message, e.Response)
}

func GetActivity() (*GetActivityResponse, *GetActivityError) {
	var getActivityResponse GetActivityResponse
	var errMsg ErrorMessage

	client := req.C().
		SetBaseURL(viper.GetString("tautulli_api_url")).
		SetCommonQueryParams(map[string]string{"apikey": viper.GetString("tautulli_api_key")})
	resp, err := client.R().
		SetSuccessResult(&getActivityResponse).
		SetErrorResult(&errMsg).
		Get("?cmd=get_activity")
	if err != nil {
    return nil, &GetActivityError{Message: err.Error(), Response: resp.Dump()}
	}
	if resp.IsErrorState() {
    return nil, &GetActivityError{Message: errMsg.Response.Message, Response: resp.Dump()}
	}

	if resp.IsSuccessState() {
		return &getActivityResponse, nil
	}

  return nil, &GetActivityError{Message: fmt.Sprintf("unknown status %s", resp.Status), Response: resp.Dump()}
}
