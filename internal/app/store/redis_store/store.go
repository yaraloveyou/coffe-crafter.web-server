package redisstore

import (
	"context"

	"github.com/redis/go-redis/v9"
)

type Store struct {
	client *redis.Client
}

var (
	ctx = context.Background()
)

func New(addr string) *Store {
	rdb := redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: "",
		DB:       0,
	})
	defer rdb.Close()
	return &Store{
		client: rdb,
	}
}

func (s *Store) Get(k string) (string, error) {
	v, err := s.client.Get(ctx, k).Result()
	if err != nil {
		return "", err
	}

	return v, nil
}

func (s *Store) Set(k, v string) error {
	err := s.client.Set(ctx, k, v, 0).Err()
	if err != nil {
		return err
	}

	return nil
}

func (s *Store) Delete(k string) error {
	_, err := s.client.Del(ctx, k).Result()
	if err != nil {
		return err
	}
	return nil
}
