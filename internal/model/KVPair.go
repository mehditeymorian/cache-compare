package model

type KVPair struct {
	Key   string `bson:"key"`
	Value string `bson:"value"`
}

func NewKVPair(key string, value string) *KVPair {
	return &KVPair{Key: key, Value: value}
}
