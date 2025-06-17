package bridge

import (
	"crypto/tls"
	"crypto/x509"
	"embed"
	"errors"
	"net/http"
	"time"

	tea "github.com/charmbracelet/bubbletea"
)

var client *http.Client

//go:embed certs
var certs embed.FS

type ClientCreatedMsg string
type NoClientCreatedMsg ErrMsg

func Init_client() tea.Msg {

	activeCert, err := certs.ReadFile("certs/cert1.pem")
	if err != nil {
		return NoClientCreatedMsg(ErrMsg{err})
	}
	futureCert, err := certs.ReadFile("certs/cert2.pem")
	if err != nil {
	} //since we currently dont really need this cert anyways

	certPool, err := x509.SystemCertPool()
	if err != nil {
		return NoClientCreatedMsg(ErrMsg{err})
	}

	if ok := certPool.AppendCertsFromPEM(activeCert); !ok {
		return NoClientCreatedMsg(ErrMsg{errors.New("Invalid Certificate")})
	}
	if ok := certPool.AppendCertsFromPEM(futureCert); !ok {
		// we dont really care
	}

	tlsConf := &tls.Config{
		RootCAs: certPool,
	}

	tr := &http.Transport{
		TLSClientConfig: tlsConf,
	}

	client = &http.Client{
		Timeout:   15 * time.Second,
		Transport: tr,
	}
	return ClientCreatedMsg("")
}
