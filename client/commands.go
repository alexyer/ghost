package client

import (
	"errors"
	"strconv"

	"github.com/alexyer/ghost/protocol"
)

type processor struct {
	process func(cmd *protocol.Command) (*protocol.Reply, error)
}

func (p *processor) Process(cmd *protocol.Command) {
	p.process(cmd)
}

// PING command.
func (p *processor) Ping() (*protocol.Reply, error) {
	cmdId := protocol.CommandId_PING

	cmd := &protocol.Command{
		CommandId: &cmdId,
	}

	reply, err := p.process(cmd)
	return reply, err
}

// SET command.
// SET <key> <val>
func (p *processor) Set(key, val string) {
	cmdId := protocol.CommandId_SET

	cmd := &protocol.Command{
		CommandId: &cmdId,
		Args:      []string{key, val},
	}

	p.process(cmd)
	return
}

// GET command.
// GET <key>
func (p *processor) Get(key string) (string, error) {
	cmdId := protocol.CommandId_GET

	cmd := &protocol.Command{
		CommandId: &cmdId,
		Args:      []string{key},
	}

	reply, err := p.process(cmd)

	if err != nil {
		return "", err
	}

	if *reply.Error != "" {
		return "", errors.New(*reply.Error)
	}

	return reply.Values[0], nil
}

// DEL command.
// DEL <key> <val>
func (p *processor) Del(key string) {
	cmdId := protocol.CommandId_DEL

	cmd := &protocol.Command{
		CommandId: &cmdId,
		Args:      []string{key},
	}

	p.process(cmd)
	return
}

// CGET command.
// CGET <collection name>
// Change user's collection.
func (p *processor) CGet(collectionName string) (string, error) {
	cmdId := protocol.CommandId_CGET

	cmd := &protocol.Command{
		CommandId: &cmdId,
		Args:      []string{collectionName},
	}

	reply, err := p.process(cmd)

	if err != nil {
		return "", err
	}

	if *reply.Error != "" {
		return "", errors.New(*reply.Error)
	}

	return "", nil
}

// CADD command.
// CADD <collection name>
// Add new collection.
func (p *processor) CAdd(collectionName string) (string, error) {
	cmdId := protocol.CommandId_CADD

	cmd := &protocol.Command{
		CommandId: &cmdId,
		Args:      []string{collectionName},
	}

	reply, err := p.process(cmd)

	if err != nil {
		return "", err
	}

	if *reply.Error != "" {
		return "", errors.New(*reply.Error)
	}

	return "", nil
}

// EXPIRE command.
// EXPIRE <key> <seconds>
// Set expiration time of the key.
func (p *processor) Expire(key string, ttl int) (string, error) {
	cmdId := protocol.CommandId_EXPIRE

	cmd := &protocol.Command{
		CommandId: &cmdId,
		Args:      []string{key, strconv.Itoa(ttl)},
	}

	reply, err := p.process(cmd)

	if err != nil {
		return "", err
	}

	if *reply.Error != "" {
		return "", errors.New(*reply.Error)
	}

	return "", nil
}

// TTL command.
// TTL <key>
// Get expiration time of the key.
func (p *processor) TTL(key string) (int, error) {
	cmdId := protocol.CommandId_TTL

	cmd := &protocol.Command{
		CommandId: &cmdId,
		Args:      []string{key},
	}

	reply, err := p.process(cmd)

	if err != nil {
		return -1, err
	}

	if *reply.Error != "" {
		return -1, errors.New(*reply.Error)
	}

	ttl, err := strconv.Atoi(reply.Values[0])
	if err != nil {
		return -1, err
	}

	return ttl, nil
}

// PERSIST command.
// PERSIST <key>
// Remove the existing timeout of the key.
func (p *processor) Persist(key string) error {
	cmdId := protocol.CommandId_PERSIST

	cmd := &protocol.Command{
		CommandId: &cmdId,
		Args:      []string{key},
	}

	reply, err := p.process(cmd)

	if err != nil {
		return err
	}

	if *reply.Error != "" {
		return errors.New(*reply.Error)
	}

	return nil
}
