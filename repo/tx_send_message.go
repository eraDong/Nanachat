package repo

import "context"

type SendMessageParams struct {
}

type SendMessageResult struct {
}

func (store *Store) SendMessage(ctx context.Context, arg SendMessageParams) (SendMessageResult, error) {
	var result SendMessageResult

	return result, nil
}
