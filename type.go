package botandb

import "context"

func (b *BotanDB) GetBool(ctx context.Context, key string) (bool, error) {
	value, err := b.Get(ctx, key)
	if err != nil {
		return false, err
	}

	if b, ok := value.(bool); ok {
		return b, nil
	}
	return false, ErrTypeConverted
}

func (b *BotanDB) GetInt(ctx context.Context, key string) (int, error) {
	value, err := b.Get(ctx, key)
	if err != nil {
		return 0, err
	}

	if i, ok := value.(int); ok {
		return i, nil
	}
	return 0, ErrTypeConverted
}

func (b *BotanDB) GetString(ctx context.Context, key string) (string, error) {
	value, err := b.Get(ctx, key)
	if err != nil {
		return "", err
	}

	if s, ok := value.(string); ok {
		return s, nil
	}
	return "", ErrTypeConverted
}
