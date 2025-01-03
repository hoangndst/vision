package casbin

import (
	"context"
	"fmt"
	"reflect"

	pb "github.com/casbin/casbin-server/proto"
	cserver "github.com/casbin/casbin-server/server"
)

type Config struct {
	DriverName              string
	ConnectString           string
	ModelText               string
	DbSpecified             bool
	EnableAcceptJsonRequest bool
}

type Enforcer struct {
	handler int32
	client  *Client
}

func (c *Client) NewEnforcer(ctx context.Context, config Config) (*Enforcer, error) {
	var adapterHandler int32 = -1
	enforcer := &Enforcer{client: c}

	// Maybe it does not need NewAdapter.
	if config.DriverName != "" && config.ConnectString != "" {
		adapterReply, err := c.remoteClient.NewAdapter(ctx, &pb.NewAdapterRequest{
			DriverName:    config.DriverName,
			ConnectString: config.ConnectString,
			DbSpecified:   config.DbSpecified,
		})
		if err != nil {
			return enforcer, err
		}
		adapterHandler = adapterReply.Handler
	}

	e, err := c.remoteClient.NewEnforcer(ctx, &pb.NewEnforcerRequest{
		ModelText:               config.ModelText,
		AdapterHandle:           adapterHandler,
		EnableAcceptJsonRequest: config.EnableAcceptJsonRequest,
	})
	if err != nil {
		return enforcer, err
	}
	enforcer.handler = e.Handler

	return enforcer, nil
}

func (e *Enforcer) Enforce(ctx context.Context, params ...interface{}) (bool, error) {
	// (sub, obj, act)
	var data []string
	var err error
	for _, item := range params {
		var value string
		if reflect.TypeOf(item).Kind() == reflect.Struct {
			value, err = cserver.MakeABAC(item)
			if err != nil {
				return false, err
			}
		} else {
			value = fmt.Sprintf("%v", item)
		}
		data = append(data, value)
	}

	resp, err := e.client.remoteClient.Enforce(ctx, &pb.EnforceRequest{
		EnforcerHandler: e.handler,
		Params:          data,
	})
	if err != nil {
		return false, err
	}
	return resp.Res, nil
}

// LoadPolicy reloads the policy from file or database.
func (e *Enforcer) LoadPolicy(ctx context.Context) error {
	_, err := e.client.remoteClient.LoadPolicy(ctx, &pb.EmptyRequest{Handler: e.handler})
	return err
}

// SavePolicy saves the current policy (usually after changed with Casbin API) back to file or database.
func (e *Enforcer) SavePolicy(ctx context.Context) error {
	_, err := e.client.remoteClient.SavePolicy(ctx, &pb.EmptyRequest{Handler: e.handler})
	return err
}
