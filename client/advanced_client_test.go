package client

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/require"

	mokk "github.com/mikloslorinczi/mokk/server"
)

func Test_Mokk_Call(t *testing.T) {
	srv := mokk.NewTestServer(t)
	srv.Handle("/test", "GET", srv.Handler().WithResponseBody([]byte("OK")))
	srv.Init()
	defer srv.Close()

	resp, err := Call(srv.URL + "/test")

	require.NoError(t, err, "The Call shouldn't return any errors")

	require.True(t, strings.Contains(resp, "OK"), "Response should contain OK")
}
