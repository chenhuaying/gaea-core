package dev

import (
	"crypto/rand"
	"crypto/sha256"
	"io/ioutil"
	"os"
	"testing"

	"github.com/gaeanetwork/gaea-core/common"
	"github.com/stretchr/testify/assert"
)

func Test_New(t *testing.T) {
	c := New()
	assert.NotNil(t, c)
}

func Test_Create(t *testing.T) {
	c := New()
	err := c.Create()
	assert.NoError(t, err)
	assert.NotEmpty(t, c.address)
}

func Test_Upload(t *testing.T) {
	c := New()
	algorithm, err := ioutil.ReadFile("/home/rabbit/teetest/client/resume")
	assert.NoError(t, err)
	A, err := ioutil.ReadFile("/home/rabbit/teetest/A/A_resume.txt")
	assert.NoError(t, err)
	B, err := ioutil.ReadFile("/home/rabbit/teetest/B/B_resume.txt")
	assert.NoError(t, err)
	C, err := ioutil.ReadFile("/home/rabbit/teetest/C/C_resume.txt")
	assert.NoError(t, err)
	dataList := [][]byte{A, B, C}
	err = c.Upload(algorithm, dataList)
	assert.NoError(t, err)

	os.Remove("main")
	os.Remove("0")
	os.Remove("1")
	os.Remove("2")
}

func Test_Verify(t *testing.T) {
	c := New()
	err := c.Create()
	defer c.Destroy()
	assert.NoError(t, err)

	algorithmBytes := make([]byte, 32)
	rand.Read(algorithmBytes)
	hash := sha256.Sum256(algorithmBytes)
	algorithmHash := common.BytesToHex(hash[:])

	dataList, dataHashes := make([][]byte, 10), make([]string, 10)
	for index := 0; index < len(dataList); index++ {
		dataBytes := make([]byte, 32)
		rand.Read(dataBytes)
		hash = sha256.Sum256(dataBytes)
		dataList[index] = dataBytes
		dataHashes[index] = common.BytesToHex(hash[:])
	}

	err1 := c.Upload(algorithmBytes, dataList)
	assert.NoError(t, err1)
	err2 := c.Verify(algorithmHash, dataHashes)
	assert.NoError(t, err2)

	// again
	err3 := c.Upload(algorithmBytes, dataList)
	assert.NoError(t, err3)
	err4 := c.Verify(algorithmHash, dataHashes)
	assert.NoError(t, err4)
}

func Test_Execute(t *testing.T) {
	c := New()
	algorithm, err := ioutil.ReadFile("/home/rabbit/teetest/client/resume")
	assert.NoError(t, err)
	A, err := ioutil.ReadFile("/home/rabbit/teetest/A/A_resume.txt")
	assert.NoError(t, err)
	B, err := ioutil.ReadFile("/home/rabbit/teetest/B/B_resume.txt")
	assert.NoError(t, err)
	C, err := ioutil.ReadFile("/home/rabbit/teetest/C/C_resume.txt")
	assert.NoError(t, err)
	dataList := [][]byte{A, B, C}
	err = c.Upload(algorithm, dataList)
	assert.NoError(t, err)

	defer func() {
		os.Remove("main")
		os.Remove("0")
		os.Remove("1")
		os.Remove("2")
	}()
	dir, err := os.Getwd()
	os.Setenv("PATH", os.Getenv("PATH")+":"+dir)
	data, err := c.Execute()
	assert.NoError(t, err)
	assert.NotEmpty(t, data)
}