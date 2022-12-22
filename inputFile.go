package appwrite

import (
	"bytes"
	"io"
	"os"
)

type InputFile struct {
	Stream   io.ReadCloser
	Size     int
	Filename string
}

type Blob interface {
	ArrayBuffer() ([]byte, error)
}

func NewInputFile(stream io.ReadCloser, filename string, size int) *InputFile {
	return &InputFile{
		Stream:   stream,
		Filename: filename,
		Size:     size,
	}
}

func NewInputFileFromPath(filePath, filename string) (*InputFile, error) {
	stream, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer stream.Close()

	info, err := stream.Stat()
	if err != nil {
		return nil, err
	}

	return NewInputFile(stream, filename, int(info.Size())), nil
}

type BufferReadCloser struct {
	*bytes.Buffer
}

func (brc *BufferReadCloser) Close() error {
	return nil
}

func NewInputFileFromBuffer(buffer []byte, filename string) *InputFile {
	bufferReader := bytes.NewBuffer(buffer)
	stream := &BufferReadCloser{bufferReader}
	return NewInputFile(stream, filename, len(buffer))
}

func NewInputFileFromBlob(blob Blob, filename string) (*InputFile, error) {
	buffer, err := blob.ArrayBuffer()
	if err != nil {
		return nil, err
	}

	return NewInputFileFromBuffer(buffer, filename), nil
}

func NewInputFileFromStream(stream io.ReadCloser, filename string, size int) *InputFile {
	return NewInputFile(stream, filename, size)
}

func NewInputFileFromPlainText(content, filename string) *InputFile {
	buffer := []byte(content)
	return NewInputFileFromBuffer(buffer, filename)
}