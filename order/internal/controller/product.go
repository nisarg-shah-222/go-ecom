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

func getProductPrice(env *dtos.Env, productId uint) (float64, error) {
	url := viper.GetString(constants.HOST_PRODUCT) + viper.GetString(constants.BASE_PRODUCT) + fmt.Sprintf(constants.URL_GET_PRODUCT_DETAILS, productId)
	req, _ := http.NewRequest("GET", url, nil)
	req.Header.Set("X-Api-Key", viper.GetString(constants.X_API_KEY_PRODUCT))
	client := http.Client{}
	response, err := client.Do(req)
	if err != nil {
		return 0, errors.New("unable to fetch product price")
	}
	if response.StatusCode != 200 {
		return 0, errors.New("unable to fetch product price")
	}

	responseBody, err := io.ReadAll(response.Body)
	getProductDetailsResponseDTOs := dtos.GetProductDetailsResponseDTOs{}
	json.Unmarshal(responseBody, &getProductDetailsResponseDTOs)
	if err := env.Validator.Struct(getProductDetailsResponseDTOs.Data); err != nil {
		return 0, errors.Join(errors.New("error validating product details response"), err)
	}
	if *getProductDetailsResponseDTOs.Data.Id != productId {
		return 0, errors.New("error requesting for product price")
	}

	return *getProductDetailsResponseDTOs.Data.Price, nil
}
