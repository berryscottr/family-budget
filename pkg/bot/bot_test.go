package bot

import (
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

// TestData_SetDir confirms ability to set the bot directory
func TestData_SetDir(t *testing.T) {
	assertion := assert.New(t)
	data := Data{Token: os.Getenv("BOT_TOKEN")}
	data = data.SetDir()
	assertion.NoError(data.Err, "failed to set the bot directory")
}

// TestData_Start confirms ability to start the discord listener
func TestData_Start(t *testing.T) {
	assertion := assert.New(t)
	data := Data{Token: os.Getenv("BOT_TOKEN")}
	data = data.SetDir()
	data.Start()
	assertion.NoError(data.Err, "failed to start discord listener")
}

// TestData_HandleSummary confirms ability to handle the finance summary post
func TestData_HandleSummary(t *testing.T) {
	assertion := assert.New(t)
	data := Data{Token: os.Getenv("BOT_TOKEN")}
	data = data.SetDir()
	assertion.NoError(data.Err, "failed to create a finance summary post")
}
