/*
   binance.go
       Wrapper for the Binance Exchange API

   Authors:
       Pat DePippo  <patrick.depippo@dcrypt.io>
       Matthew Woop <matthew.woop@dcrypt.io>

   To Do:

*/
package binance

import "net/http"

//"errors"

const (
	BaseUrl = "https://www.binance.com"
)

type Binance struct {
	client *Client
}

/*
func handleErr(r jsonResponse) error {

    if !r.Success {
        return errors.New(r.Message)
    }

    return nil
}
*/
func New(httpClient *http.Client, key, secret string) *Binance {
	client := NewClient(httpClient, key, secret)
	return &Binance{client}
}
