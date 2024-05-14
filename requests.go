package art

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
)

func (a *ArtClient) GetSpecificArt(ctx context.Context, id string) (*ArtResp, error) {
	res, err := a.doRequest(ctx, nil, "GET", fmt.Sprintf("/%s", id))
	if err != nil {
		return nil, err
	}

	defer res.Body.Close()
	if err = checkApiError(res.StatusCode, res.Status); err != nil {
		return nil, err
	}

	var art *ArtResp
	if err = json.NewDecoder(res.Body).Decode(&art); err != nil {
		return nil, err
	}
	return art, nil
}

func (a *ArtClient) GetAllArt(ctx context.Context) (*[]ArtResp, error) {
	res, err := a.doRequest(ctx, nil, "GET", "/")
	if err != nil {
		return nil, err
	}

	defer res.Body.Close()
	if err = checkApiError(res.StatusCode, res.Status); err != nil {
		return nil, err
	}

	var arts *[]ArtResp

	if err = json.NewDecoder(res.Body).Decode(&arts); err != nil {
		return nil, err
	}
	return arts, nil
}

func (a *ArtClient) DeleteArt(ctx context.Context, id string) (string, error) {
	res, err := a.doRequest(ctx, nil, "DELETE", fmt.Sprintf("/%s", id))
	if err != nil {
		return fmt.Sprintf("Error Deleting Record with ID: %s", id), err
	}

	defer res.Body.Close()
	if err = checkApiError(res.StatusCode, res.Status); err != nil {
		return fmt.Sprintf("Error Deleting Record with ID: %s", id), err
	}
	return fmt.Sprintf("Record with ID: %s Deleted Successfully", id), nil
}

func (a *ArtClient) AddNewArt(ctx context.Context, body *ArtReq) (*ArtResp, error) {
	var buf bytes.Buffer

	if err := json.NewEncoder(&buf).Encode(&body); err != nil {
		return nil, err
	}
	res, err := a.doRequest(ctx, &buf, "POST", "/")
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	if err = checkApiError(res.StatusCode, res.Status); err != nil {
		return nil, err
	}

	var art *ArtResp

	if err = json.NewDecoder(res.Body).Decode(&art); err != nil {
		return nil, err
	}
	return art, nil
}

func (a *ArtClient) UpdateArt(ctx context.Context, body *ArtReq, id string) (string, error) {
	var buf bytes.Buffer

	if err := json.NewEncoder(&buf).Encode(&body); err != nil {
		return fmt.Sprintf("Error Updating Record with ID: %s", id), err
	}

	res, err := a.doRequest(ctx, &buf, "PUT", fmt.Sprintf("/%s", id))
	if err != nil {
		return fmt.Sprintf("Error Updating Record with ID: %s", id), err
	}

	defer res.Body.Close()

	if err = checkApiError(res.StatusCode, res.Status); err != nil {
		return fmt.Sprintf("Error Updating Record with ID: %s", id), err
	}
	return fmt.Sprintf("Record with ID: %s Updated Successfully", id), nil
}
