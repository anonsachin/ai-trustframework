package ipfs

import (
	"io"

	api "github.com/ipfs/go-ipfs-api"
	"github.com/sirupsen/logrus"
)

type Shell struct {
	client *api.Shell
	log *logrus.Entry
}

func NewShell(address string, log *logrus.Entry) *Shell {
	log.Info("Creating the ipfs client")
	return &Shell{
		client: api.NewShell(address),
		log: log,
	}
}

func (s *Shell) Add(content io.Reader) (string, error) {
	s.log.Info("Adding content")
	return s.client.Add(content)
}

func (s *Shell) Cat(contentID string) (io.ReadCloser, error) {
	s.log.Infof("Getting content %s",contentID)
	return s.client.Cat(contentID)
}
