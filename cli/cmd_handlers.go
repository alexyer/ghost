package cli

import (
	"errors"
	"fmt"

	"github.com/alexyer/ghost/client"
)

// these functions handle basic commands to ghost-server

// ping command
func pingServer(c *client.GhostClient, args []string) (string, error) {
	if len(args) != 0 {
		return "", errors.New(fmt.Sprintf("wrong number of arguments to PING: need 0, get %d", len(args)))
	}

	reply, err := c.Ping()
	return reply.Values[0], err
}

// set command
func setValue(c *client.GhostClient, args []string) error {
	if len(args) != 2 {
		return errors.New(fmt.Sprintf("wrong number of arguments to SET: need 2, get %d", len(args)))
	}

	c.Set(args[0], args[1])
	return nil
}

// get command
func getValue(c *client.GhostClient, args []string) (string, error) {
	if len(args) != 1 {
		return "", errors.New(fmt.Sprintf("wrong number of arguments to GET: need 1, get %d", len(args)))
	}

	return c.Get(args[0])
}

// del command
func delValue(c *client.GhostClient, args []string) error {
	if len(args) != 1 {
		return errors.New(fmt.Sprintf("wrong number of arguments to DEL: need 1, get %d", len(args)))
	}

	c.Del(args[0])
	return nil
}

// add collection command
func addColl(c *client.GhostClient, args []string) error {
	if len(args) != 1 {
		return errors.New(fmt.Sprintf("wrong number of arguments to CADD: need 1, get %d", len(args)))
	}

	if _, err := c.CAdd(args[0]); err != nil {
		return err
	}
	return nil
}

// get collection command
func getColl(c *client.GhostClient, args []string) error {
	if len(args) != 1 {
		return errors.New(fmt.Sprintf("wrong number of arguments to CGET: need 1, get %d", len(args)))
	}

	if _, err := c.CGet(args[0]); err != nil {
		return err
	}
	return nil
}
