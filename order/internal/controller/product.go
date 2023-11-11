package controller

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"order/internal/constants"
	"order/internal/models/dtos"

	"github.com/spf13/viper"
)

func getProductDetails(env *dtos.Env, productId uint) (*dtos.ProductDetails, error) {
	url := viper.GetString(constants.HOST_PRODUCT) + viper.GetString(constants.BASE_PRODUCT) + fmt.Sprintf(constants.URL_GET_PRODUCT_DETAILS, productId)
	req, _ := http.NewRequest("GET", url, nil)
	req.Header.Set("X-Api-Key", viper.GetString(constants.X_API_KEY_PRODUCT))
	client := http.Client{}
	response, err := client.Do(req)
	if err != nil {
		return nil, errors.Join(constants.ErrInternalServerError, constants.ErrProductAPIFailed, err)
	}
	if response.StatusCode != 200 {
		return nil, errors.Join(constants.ErrInternalServerError, constants.ErrProductAPIFailed)
	}

	responseBody, err := io.ReadAll(response.Body)
	getProductDetailsResponseDTOs := dtos.GetProductDetailsResponseDTOs{}
	json.Unmarshal(responseBody, &getProductDetailsResponseDTOs)
	if err := env.Validator.Struct(getProductDetailsResponseDTOs.Data); err != nil {
		return nil, errors.Join(constants.ErrProductAPIFailed, err)
	}
	if *getProductDetailsResponseDTOs.Data.Id != productId {
		return nil, constants.ErrProductAPIFailed
	}

	return getProductDetailsResponseDTOs.Data, nil
}
