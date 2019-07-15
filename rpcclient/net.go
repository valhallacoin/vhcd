// Copyright (c) 2014-2015 The btcsuite developers
// Copyright (c) 2015-2017 The Decred developers
// Use of this source code is governed by an ISC
// license that can be found in the LICENSE file.

package rpcclient

import (
	"encoding/json"

	"github.com/valhallacoin/vhcd/vhcjson"
)

// AddNodeCommand enumerates the available commands that the AddNode function
// accepts.
type AddNodeCommand string

// Constants used to indicate the command for the AddNode function.
const (
	// ANAdd indicates the specified host should be added as a persistent
	// peer.
	ANAdd AddNodeCommand = "add"

	// ANRemove indicates the specified peer should be removed.
	ANRemove AddNodeCommand = "remove"

	// ANOneTry indicates the specified host should try to connect once,
	// but it should not be made persistent.
	ANOneTry AddNodeCommand = "onetry"
)

// String returns the AddNodeCommand in human-readable form.
func (cmd AddNodeCommand) String() string {
	return string(cmd)
}

// FutureAddNodeResult is a future promise to deliver the result of an
// AddNodeAsync RPC invocation (or an applicable error).
type FutureAddNodeResult chan *response

// Receive waits for the response promised by the future and returns an error if
// any occurred when performing the specified command.
func (r FutureAddNodeResult) Receive() error {
	_, err := receiveFuture(r)
	return err
}

// AddNodeAsync returns an instance of a type that can be used to get the result
// of the RPC at some future time by invoking the Receive function on the
// returned instance.
//
// See AddNode for the blocking version and more details.
func (c *Client) AddNodeAsync(host string, command AddNodeCommand) FutureAddNodeResult {
	cmd := vhcjson.NewAddNodeCmd(host, vhcjson.AddNodeSubCmd(command))
	return c.sendCmd(cmd)
}

// AddNode attempts to perform the passed command on the passed persistent peer.
// For example, it can be used to add or a remove a persistent peer, or to do
// a one time connection to a peer.
//
// It may not be used to remove non-persistent peers.
func (c *Client) AddNode(host string, command AddNodeCommand) error {
	return c.AddNodeAsync(host, command).Receive()
}

// FutureGetAddedNodeInfoResult is a future promise to deliver the result of a
// GetAddedNodeInfoAsync RPC invocation (or an applicable error).
type FutureGetAddedNodeInfoResult chan *response

// Receive waits for the response promised by the future and returns information
// about manually added (persistent) peers.
func (r FutureGetAddedNodeInfoResult) Receive() ([]vhcjson.GetAddedNodeInfoResult, error) {
	res, err := receiveFuture(r)
	if err != nil {
		return nil, err
	}

	// Unmarshal as an array of getaddednodeinfo result objects.
	var nodeInfo []vhcjson.GetAddedNodeInfoResult
	err = json.Unmarshal(res, &nodeInfo)
	if err != nil {
		return nil, err
	}

	return nodeInfo, nil
}

// GetAddedNodeInfoAsync returns an instance of a type that can be used to get
// the result of the RPC at some future time by invoking the Receive function on
// the returned instance.
//
// See GetAddedNodeInfo for the blocking version and more details.
func (c *Client) GetAddedNodeInfoAsync(peer string) FutureGetAddedNodeInfoResult {
	cmd := vhcjson.NewGetAddedNodeInfoCmd(true, &peer)
	return c.sendCmd(cmd)
}

// GetAddedNodeInfo returns information about manually added (persistent) peers.
//
// See GetAddedNodeInfoNoDNS to retrieve only a list of the added (persistent)
// peers.
func (c *Client) GetAddedNodeInfo(peer string) ([]vhcjson.GetAddedNodeInfoResult, error) {
	return c.GetAddedNodeInfoAsync(peer).Receive()
}

// FutureGetAddedNodeInfoNoDNSResult is a future promise to deliver the result
// of a GetAddedNodeInfoNoDNSAsync RPC invocation (or an applicable error).
type FutureGetAddedNodeInfoNoDNSResult chan *response

// Receive waits for the response promised by the future and returns a list of
// manually added (persistent) peers.
func (r FutureGetAddedNodeInfoNoDNSResult) Receive() ([]string, error) {
	res, err := receiveFuture(r)
	if err != nil {
		return nil, err
	}

	// Unmarshal result as an array of strings.
	var nodes []string
	err = json.Unmarshal(res, &nodes)
	if err != nil {
		return nil, err
	}

	return nodes, nil
}

// GetAddedNodeInfoNoDNSAsync returns an instance of a type that can be used to
// get the result of the RPC at some future time by invoking the Receive
// function on the returned instance.
//
// See GetAddedNodeInfoNoDNS for the blocking version and more details.
func (c *Client) GetAddedNodeInfoNoDNSAsync(peer string) FutureGetAddedNodeInfoNoDNSResult {
	cmd := vhcjson.NewGetAddedNodeInfoCmd(false, &peer)
	return c.sendCmd(cmd)
}

// GetAddedNodeInfoNoDNS returns a list of manually added (persistent) peers.
// This works by setting the dns flag to false in the underlying RPC.
//
// See GetAddedNodeInfo to obtain more information about each added (persistent)
// peer.
func (c *Client) GetAddedNodeInfoNoDNS(peer string) ([]string, error) {
	return c.GetAddedNodeInfoNoDNSAsync(peer).Receive()
}

// FutureGetConnectionCountResult is a future promise to deliver the result
// of a GetConnectionCountAsync RPC invocation (or an applicable error).
type FutureGetConnectionCountResult chan *response

// Receive waits for the response promised by the future and returns the number
// of active connections to other peers.
func (r FutureGetConnectionCountResult) Receive() (int64, error) {
	res, err := receiveFuture(r)
	if err != nil {
		return 0, err
	}

	// Unmarshal result as an int64.
	var count int64
	err = json.Unmarshal(res, &count)
	if err != nil {
		return 0, err
	}

	return count, nil
}

// GetConnectionCountAsync returns an instance of a type that can be used to get
// the result of the RPC at some future time by invoking the Receive function on
// the returned instance.
//
// See GetConnectionCount for the blocking version and more details.
func (c *Client) GetConnectionCountAsync() FutureGetConnectionCountResult {
	cmd := vhcjson.NewGetConnectionCountCmd()
	return c.sendCmd(cmd)
}

// GetConnectionCount returns the number of active connections to other peers.
func (c *Client) GetConnectionCount() (int64, error) {
	return c.GetConnectionCountAsync().Receive()
}

// FuturePingResult is a future promise to deliver the result of a PingAsync RPC
// invocation (or an applicable error).
type FuturePingResult chan *response

// Receive waits for the response promised by the future and returns the result
// of queueing a ping to be sent to each connected peer.
func (r FuturePingResult) Receive() error {
	_, err := receiveFuture(r)
	return err
}

// PingAsync returns an instance of a type that can be used to get the result of
// the RPC at some future time by invoking the Receive function on the returned
// instance.
//
// See Ping for the blocking version and more details.
func (c *Client) PingAsync() FuturePingResult {
	cmd := vhcjson.NewPingCmd()
	return c.sendCmd(cmd)
}

// Ping queues a ping to be sent to each connected peer.
//
// Use the GetPeerInfo function and examine the PingTime and PingWait fields to
// access the ping times.
func (c *Client) Ping() error {
	return c.PingAsync().Receive()
}

// FutureGetPeerInfoResult is a future promise to deliver the result of a
// GetPeerInfoAsync RPC invocation (or an applicable error).
type FutureGetPeerInfoResult chan *response

// Receive waits for the response promised by the future and returns  data about
// each connected network peer.
func (r FutureGetPeerInfoResult) Receive() ([]vhcjson.GetPeerInfoResult, error) {
	res, err := receiveFuture(r)
	if err != nil {
		return nil, err
	}

	// Unmarshal result as an array of getpeerinfo result objects.
	var peerInfo []vhcjson.GetPeerInfoResult
	err = json.Unmarshal(res, &peerInfo)
	if err != nil {
		return nil, err
	}

	return peerInfo, nil
}

// GetPeerInfoAsync returns an instance of a type that can be used to get the
// result of the RPC at some future time by invoking the Receive function on the
// returned instance.
//
// See GetPeerInfo for the blocking version and more details.
func (c *Client) GetPeerInfoAsync() FutureGetPeerInfoResult {
	cmd := vhcjson.NewGetPeerInfoCmd()
	return c.sendCmd(cmd)
}

// GetPeerInfo returns data about each connected network peer.
func (c *Client) GetPeerInfo() ([]vhcjson.GetPeerInfoResult, error) {
	return c.GetPeerInfoAsync().Receive()
}

// FutureGetNetTotalsResult is a future promise to deliver the result of a
// GetNetTotalsAsync RPC invocation (or an applicable error).
type FutureGetNetTotalsResult chan *response

// Receive waits for the response promised by the future and returns network
// traffic statistics.
func (r FutureGetNetTotalsResult) Receive() (*vhcjson.GetNetTotalsResult, error) {
	res, err := receiveFuture(r)
	if err != nil {
		return nil, err
	}

	// Unmarshal result as a getnettotals result object.
	var totals vhcjson.GetNetTotalsResult
	err = json.Unmarshal(res, &totals)
	if err != nil {
		return nil, err
	}

	return &totals, nil
}

// GetNetTotalsAsync returns an instance of a type that can be used to get the
// result of the RPC at some future time by invoking the Receive function on the
// returned instance.
//
// See GetNetTotals for the blocking version and more details.
func (c *Client) GetNetTotalsAsync() FutureGetNetTotalsResult {
	cmd := vhcjson.NewGetNetTotalsCmd()
	return c.sendCmd(cmd)
}

// GetNetTotals returns network traffic statistics.
func (c *Client) GetNetTotals() (*vhcjson.GetNetTotalsResult, error) {
	return c.GetNetTotalsAsync().Receive()
}
