package es

import (
	"context"
	"net/http"

	"github.com/rerost/es-cli/setting"
	"github.com/srvc/fail"
)

type Index struct {
	Name string
}
type Mapping struct{}
type Opt struct{}
type Alias struct{}
type Task struct{}
type Version struct{}

// Client is http wrapper
type BaseClient interface {
	// Index
	ListIndex(ctx context.Context, indexName string) ([]Index, error)
	CreateIndex(ctx context.Context, indexName string, mappingJSON string) error
	CopyIndex(ctx context.Context, srcIndexName string, dstIndexName string) error
	DeleteIndex(ctx context.Context, indexName string) error

	// Mapping
	GetMapping(ctx context.Context, indexOrAliasName string) (Mapping, error)
	// UpdateMapping(ctx context.Context, aliasName string, mappingJSON string) error

	// Alias
	CreateAlias(ctx context.Context, indexName string, aliasName string) error
	DropAlias(ctx context.Context, aliasName string, opts []Opt) error
	AddAlias(ctx context.Context, aliasName string, indexNames ...string) error
	RemoveAlias(ctx context.Context, aliasName string, indexNames ...string) error
	GetAlias(ctx context.Context, aliasName string) (Alias, error)

	// Task
	ListTask(ctx context.Context) ([]Task, error)
	GetTask(ctx context.Context, taskID string) (Task, error)

	Version(ctx context.Context) (Version, error)
	Ping(ctx context.Context) (bool, error)
}

type baseClientImp struct {
	Host       string
	Port       string
	Type       string
	httpClient http.Client
}

func NewBaseClient(ctx context.Context, httpClient http.Client) (BaseClient, error) {
	client := baseClientImp{}
	client.httpClient = httpClient

	_host, ok := ctx.Value(setting.SettingKey("Host")).(string)
	if !ok {
		return client, fail.New("Failed to extract host")
	}

	_port, ok := ctx.Value(setting.SettingKey("Port")).(string)
	if !ok {
		return client, fail.New("Failed to extract port")
	}

	_type, ok := ctx.Value(setting.SettingKey("Type")).(string)
	if !ok {
		return client, fail.New("Failed to extract type")
	}

	client.Host = _host
	client.Port = _port
	client.Type = _type

	return client, nil
}

func (client baseClientImp) ListIndex(ctx context.Context, indexName string) ([]Index, error) {
	return []Index{}, nil
}
func (client baseClientImp) CreateIndex(ctx context.Context, indexName string, mappingJSON string) error {
	return nil
}
func (client baseClientImp) CopyIndex(ctx context.Context, srcIndexName string, dstIndexName string) error {
	return nil
}
func (client baseClientImp) DeleteIndex(ctx context.Context, indexName string) error {
	return nil
}

// Mapping
func (client baseClientImp) GetMapping(ctx context.Context, indexOrAliasName string) (Mapping, error) {
	return Mapping{}, nil
}

// Alias
func (client baseClientImp) CreateAlias(ctx context.Context, indexName string, aliasName string) error {
	return nil
}
func (client baseClientImp) DropAlias(ctx context.Context, aliasName string, opts []Opt) error {
	return nil
}
func (client baseClientImp) AddAlias(ctx context.Context, aliasName string, indexNames ...string) error {
	return nil
}
func (client baseClientImp) RemoveAlias(ctx context.Context, aliasName string, indexNames ...string) error {
	return nil
}
func (client baseClientImp) GetAlias(ctx context.Context, aliasName string) (Alias, error) {
	return Alias{}, nil
}

// Task
func (client baseClientImp) ListTask(ctx context.Context) ([]Task, error) {
	return []Task{}, nil
}
func (client baseClientImp) GetTask(ctx context.Context, taskID string) (Task, error) {
	return Task{}, nil
}

// Util
func (client baseClientImp) Version(ctx context.Context) (Version, error) {
	return Version{}, nil
}
func (client baseClientImp) Ping(ctx context.Context) (bool, error) {
	return false, nil
}
