package cli

import (
	"errors"
	"fmt"
	"log"
	"strconv"

	"github.com/alexyer/ghost/client"
)

// initialize endless cli-session with provided client
// as a connection to ghost-server
func StartCliSession(c *client.GhostClient) {
	log.Println("Cli-ghost session started")

	for {
		comm, args, err := processUserInput()
		if err != nil {
			log.Printf("Error on input processing: %s", err.Error())
			continue
		}

		result, err := makeRequest(c, comm, args)
		if err != nil {
			log.Printf("Error on request: %s", err.Error())
			continue
		}

		fmt.Println(result)
	}
}

func makeRequest(c *client.GhostClient, comm string, args []string) (string, error) {
	switch comm {
	case "PING":
		return pingServer(c, args)
	case "SET":
		if err := setValue(c, args); err != nil {
			return "", err
		}
	case "GET":
		return getValue(c, args)
	case "DEL":
		if err := delValue(c, args); err != nil {
			return "", err
		}
	case "CGET":
		if err := getColl(c, args); err != nil {
			return "", err
		}
	case "CADD":
		if err := addColl(c, args); err != nil {
			return "", err
		}
	case "EXPIRE":
		if err := expire(c, args); err != nil {
			return "", err
		}
	default:
		return "", errors.New("unknown command: " + comm)
	}

	return "OK", nil
}

func pingServer(c *client.GhostClient, args []string) (string, error) {
	if len(args) != 0 {
		return "", errors.New(fmt.Sprintf("wrong number of arguments to PING: need 0, get %d", len(args)))
	}

	reply, err := c.Ping()
	return reply.Values[0], err
}

func setValue(c *client.GhostClient, args []string) error {
	if len(args) != 2 {
		return errors.New(fmt.Sprintf("wrong number of arguments to SET: need 2, get %d", len(args)))
	}

	c.Set(args[0], args[1])
	return nil
}

func getValue(c *client.GhostClient, args []string) (string, error) {
	if len(args) != 1 {
		return "", errors.New(fmt.Sprintf("wrong number of arguments to GET: need 1, get %d", len(args)))
	}

	return c.Get(args[0])
}

func delValue(c *client.GhostClient, args []string) error {
	if len(args) != 1 {
		return errors.New(fmt.Sprintf("wrong number of arguments to DEL: need 1, get %d", len(args)))
	}

	c.Del(args[0])
	return nil
}

func addColl(c *client.GhostClient, args []string) error {
	if len(args) != 1 {
		return errors.New(fmt.Sprintf("wrong number of arguments to CADD: need 1, get %d", len(args)))
	}

	if _, err := c.CAdd(args[0]); err != nil {
		return err
	}
	return nil
}

func getColl(c *client.GhostClient, args []string) error {
	if len(args) != 1 {
		return errors.New(fmt.Sprintf("wrong number of arguments to CGET: need 1, get %d", len(args)))
	}

	if _, err := c.CGet(args[0]); err != nil {
		return err
	}
	return nil
}

func expire(c *client.GhostClient, args []string) error {
	if len(args) != 2 {
		return errors.New(fmt.Sprintf("wrong number of arguments to EXPIRE: need 2, get %d", len(args)))
	}

	ttl, err := strconv.Atoi(args[1])
	if err != nil {
		return err
	}

	if _, err := c.Expire(args[0], ttl); err != nil {
		return err
	}
	return nil
}
