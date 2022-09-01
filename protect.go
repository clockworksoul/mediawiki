package mediawiki

import (
	"context"
	"fmt"
)

func (w *Client) Protect(ctx context.Context, title, reason string) (Response, error) {
	if err := w.checkKeepAlive(ctx); err != nil {
		return Response{}, err
	}

	token, err := w.GetToken(ctx, CSRFToken)
	if err != nil {
		return Response{}, err
	}

	// Specify parameters to send.
	parameters := map[string]string{
		"action":      "protect",
		"title":       title,
		"protections": "edit=sysop",
		"reason":      reason,
		"token":       token,
	}

	// Make the request.
	r, err := w.Post(ctx, parameters)
	if err != nil {
		return r, fmt.Errorf("failed to post: %w", err)
	}

	if e := r.Error; e != nil {
		return r, fmt.Errorf("%s: %s", e.Code, e.Info)
	}

	return r, nil
}
