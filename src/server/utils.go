package server

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	"github.com/grpc-ecosystem/go-grpc-middleware/util/metautils"
)

func GetUserInfo(ctx context.Context) (map[string]string, error) {
	var userInfo map[string]string

	token := metautils.ExtractIncoming(ctx).Get("authorization")

	requestURL := fmt.Sprintf("%s/oidc/v1/userinfo", issuer)
	req, err := http.NewRequest(http.MethodGet, requestURL, nil)
	if err != nil {
		return userInfo, err
	}
	req.Header.Add("Authorization", token)

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return userInfo, err
	}
	defer res.Body.Close()

	json.NewDecoder(res.Body).Decode(&userInfo)
	return userInfo, nil
}

func GetUserInfoID(ctx context.Context) (string, error) {
	info, err := GetUserInfo(ctx)
	if err != nil {
		return "", err
	}
	id, ok := info["sub"]
	if ok {
		return id, nil
	}
	return "", errors.New("unexpected error")
}
