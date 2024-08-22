package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/spf13/cobra"
)

const MaxPollRetries = 100
const PollWaitSeconds = 20

func PollInstance(cmd *cobra.Command, instanceId string, waitingStatus InstanceStatus) (*GetInstanceResponse, error) {
	path := fmt.Sprintf("/instances/%s", instanceId)

	for i := 0; i < MaxPollRetries; i++ {
		resBody, statusCode, err := MakeRequest(cmd, http.MethodGet, path, nil)
		if err != nil {
			return nil, err
		}

		if statusCode == http.StatusOK {
			var response GetInstanceResponse
			if err := json.Unmarshal(resBody, &response); err != nil {
				return nil, err
			}

			if response.Data.Status == waitingStatus {
				time.Sleep(time.Second * PollWaitSeconds)
			} else {
				return &response, nil
			}
		}
	}

	return nil, fmt.Errorf("hit max retries for polling")
}

func PollCMK(cmd *cobra.Command, cmkId string, waitingStatus CMKStatus) (*GetCMKResponse, error) {
	path := fmt.Sprintf("/customer-managed-keys/%s", cmkId)

	for i := 0; i < MaxPollRetries; i++ {
		resBody, statusCode, err := MakeRequest(cmd, http.MethodGet, path, nil)
		if err != nil {
			return nil, err
		}

		if statusCode == http.StatusOK {
			var response GetCMKResponse
			if err := json.Unmarshal(resBody, &response); err != nil {
				return nil, err
			}

			if response.Data.Status == waitingStatus {
				time.Sleep(time.Second * PollWaitSeconds)
			} else {
				return &response, nil
			}
		}
	}

	return nil, fmt.Errorf("hit max retries for polling")
}
