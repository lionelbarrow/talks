package io

// Return the number of bytes written successfully, plus any errors
type Writer interface {
	Write([]byte) (int, error)
}



package http

type Response struct {
  // ..
}

func (r Response) Write([]byte) (int, error)



package os

type File struct {
  // ..
}

func (f File) Write([]byte) (int, error)



package bytes

type ByteBuffer struct {
  // ..
}

func (b ByteBuffer) Write([]byte) (int, error)
