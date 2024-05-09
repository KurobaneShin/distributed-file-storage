package main

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"io"
)

func newEncryptionKey() []byte {
	keybuf := make([]byte, 32)
	io.ReadFull(rand.Reader, keybuf)
	return keybuf
}

func copyDecrypt(key []byte, src io.Reader, dst io.Writer) (int, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return 0, err
	}

	// Read de IV from the given io.Reader whick, in our case should be the
	// the block.blockSize() bytes we read

	iv := make([]byte, block.BlockSize())

	if _, err := src.Read(iv); err != nil {
		return 0, nil
	}

	buf := make([]byte, 32*1024)
	stream := cipher.NewCTR(block, iv)

	for {
		n, err := src.Read(buf)

		if n > 0 {
			stream.XORKeyStream(buf, buf[:n])
			if _, err := dst.Write(buf[:n]); err != nil {
				return 0, err
			}
		}

		if err == io.EOF {
			break
		}

		if err != nil {
			return 0, err
		}
	}

	return 0, nil
}

func copyEncrypt(key []byte, src io.Reader, dst io.Writer) (int, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return 0, err
	}

	iv := make([]byte, block.BlockSize()) // 16 bytes
	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		return 0, err
	}

	// prepend the IV to the file.
	if _, err := dst.Write(iv); err != nil {
		return 0, err
	}

	buf := make([]byte, 32*1024)
	stream := cipher.NewCTR(block, iv)

	for {
		n, err := src.Read(buf)
		if n > 0 {
			stream.XORKeyStream(buf, buf[:n])
			if _, err := dst.Write(buf[:n]); err != nil {
				return 0, err
			}
		}

		if err == io.EOF {
			break
		}

		if err != nil {
			return 0, err
		}
	}

	return 0, nil
}
