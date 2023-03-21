package main

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"errors"
	"fmt"
	"io"
	"strconv"
	"strings"
	"testing"
)

func TestJiaMi(t *testing.T) {

	key := []byte("passphrasewhichneedstobe32bytest") // 密钥
	t.Log(key, "key")
	// encrypt 密文
	//hertext, err := hex.DecodeString("hsuehly")
	text := []byte("#EXTM3U\n#EXT-X-VERSION:3\n#EXT-X-MEDIA-SEQUENCE:0\n#EXT-X-ALLOW-CACHE:YES\n#EXT-X-TARGETDURATION:11\n#EXTINF:10.08,\nhttps://ae02.alicdn.com/kf/Hc9a63e7e42df489ea047616b30bbc5f1R.png\n#EXTINF:10,\nhttps://ae02.alicdn.com/kf/H3a2175e757c94c2e8e52db7bc2c0e684I.png\n#EXTINF:10,\nhttps://ae02.alicdn.com/kf/H7af6404414b647419726347e3a4fd8fbT.png\n#EXTINF:10,\nhttps://ae02.alicdn.com/kf/Ha519cfcee17e4a74856c708a4ce9df51P.png\n#EXTINF:10,\nhttps://ae02.alicdn.com/kf/H5bb6c57d52e7441f852e0f494db73434M.png\n#EXTINF:10,\nhttps://ae02.alicdn.com/kf/Hdd8a91d8fe304002ae5cd9608efd5e962.png\n#EXTINF:10,\nhttps://ae02.alicdn.com/kf/H1efeb834c26b4af6acde58183dcbd73e5.png\n#EXTINF:10,\nhttps://ae02.alicdn.com/kf/Hfd6ff0a601dd4e0c8989b5f36a4a6636F.png\n#EXTINF:10,\nhttps://ae02.alicdn.com/kf/Hae5bc8f19b0a474bb5b1831380bbdbfbS.png\n#EXTINF:10,\nhttps://ae02.alicdn.com/kf/H40d05d890f8249d995449b5459a8329aK.png\n#EXTINF:10,\nhttps://ae02.alicdn.com/kf/H293325c24e29461b8f8dceb31811ab2bk.png\n#EXTINF:10,\nhttps://ae02.alicdn.com/kf/H94d0737ef1c04350b07b3a5f3ccad50fC.png\n#EXTINF:10,\nhttps://ae02.alicdn.com/kf/H1003d5f323844ed588751f70f2499c62y.png\n#EXTINF:10,\nhttps://ae02.alicdn.com/kf/H9a0bfbb9d5cb4dc786de84d46be882667.png\n#EXTINF:10,\nhttps://ae02.alicdn.com/kf/H8ac2d198c8604b73b64c9bd6a766963dV.png\n#EXTINF:10,\nhttps://ae02.alicdn.com/kf/H4a0be082f0594c4db2acca002aa9421cG.png\n#EXTINF:10,\nhttps://ae02.alicdn.com/kf/H1a0fec4bb95c47c28aed93db650c406bd.png\n#EXTINF:10,\nhttps://ae02.alicdn.com/kf/Ha5c6852f7b1648028cea3ccf95c73961p.png\n#EXTINF:10,\nhttps://ae02.alicdn.com/kf/Hd138cc6c360047aab07403e9fad18b612.png\n#EXTINF:10,\nhttps://ae02.alicdn.com/kf/H812c5bcd839c4cee8c5dce4679c0dfd8n.png\n#EXTINF:10,\nhttps://ae02.alicdn.com/kf/H6fad18913de74a10abb1da9bdd20940fj.png\n#EXTINF:2.36,\nhttps://ae02.alicdn.com/kf/Haff7bc0e50eb4ff4a2874612a61cb066m.png\n#EXT-X-ENDLIST")
	encrypt, err := Encrypt([]byte(key), text)
	if err != nil {
		fmt.Println(err)
		return
	}
	//hertext, err := hex.DecodeString(string(encrypt))
	//var hext []byte
	//n, err := hex.
	//if err != nil {
	//	t.Log(err)
	//}
	//t.Log(n, "n")
	//f, err := os.Create("1.a")
	//f.Write(encrypt)
	//f.Close()
	var strArr []string
	for _, b := range encrypt {
		strArr = append(strArr, strconv.Itoa(int(b)))
	}

	s := strings.Join(strArr, ", ")
	fmt.Println(s)
	//fmt.Printf("str is %v\n", encrypt)
	//t.Log(encrypt[8], "encrypt")
	//fmt.Printf("hertext is %v\n", hext)

	//
	decrypt, err := Decrypt([]byte(key), encrypt)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("str is %v\n", string(decrypt))
}
func Decrypt2(ciphertext, key []byte) ([]byte, error) {
	pkey := PaddingLeft(key, '0', 16) //和js的key补码方法一致

	block, err := aes.NewCipher(pkey) //选择加密算法
	if err != nil {
		return nil, fmt.Errorf("key 长度必须 16/24/32长度: %s", err)
	}
	blockModel := cipher.NewCBCDecrypter(block, pkey) //和前端代码对应:   mode: CryptoJS.mode.CBC,// CBC算法
	plantText := make([]byte, len(ciphertext))
	blockModel.CryptBlocks(plantText, ciphertext)
	plantText = PKCS7UnPadding(plantText) //和前端代码对应:  padding: CryptoJS.pad.Pkcs7
	return plantText, nil
}

func PKCS7UnPadding(plantText []byte) []byte {
	length := len(plantText)
	unpadding := int(plantText[length-1])
	return plantText[:(length - unpadding)]
}

// 这个方案必须和js的方法是一样的
func PaddingLeft(ori []byte, pad byte, length int) []byte {
	if len(ori) >= length {
		return ori[:length]
	}
	pads := bytes.Repeat([]byte{pad}, length-len(ori))
	return append(pads, ori...)
}

func Test_Decry(t *testing.T) {
	key := []byte("passphrasewhichneedstobe32bytest") // 密钥
	//encrypt 10 进制 byte[]
	encrypt := []byte{143, 97, 5, 125, 170, 230, 231, 250, 238, 227, 60, 173, 32, 164, 249, 233, 42, 213, 159, 133, 76, 208, 153, 145, 99, 245, 48, 48, 140, 205, 242, 173}
	// 解密
	// 解密之后的 [104 115 117 101 104 108 121]
	decrypt, err := Decrypt([]byte(key), encrypt)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("str is %v\n", string(decrypt))
}

func Decrypt(key, src []byte) (data []byte, err error) {

	if len(src) < aes.BlockSize {
		return nil, errors.New("data length error")
	}
	// iv 值 在数据前16个字节
	iv := src[:aes.BlockSize]
	// 密文内容
	ciphertext := src[aes.BlockSize:]
	// 判断
	if len(ciphertext)%aes.BlockSize != 0 {
		return nil, errors.New("ciphertext is not a multiple of the block size")
	}

	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	bm := cipher.NewCBCDecrypter(block, iv)
	bm.CryptBlocks(ciphertext, ciphertext)
	ciphertext, err = pkcs7Unpad(ciphertext, aes.BlockSize)

	if err != nil {
		return nil, err
	}

	return ciphertext, nil
}

func pkcs7Unpad(src []byte, blockSize int) (dest []byte, err error) {

	if blockSize <= 0 {
		return nil, errors.New("block size is 0")
	} else if len(src)%blockSize != 0 {
		return nil, errors.New("src length error")
	} else if src == nil || len(src) == 0 {
		return nil, errors.New("src is nil")
	}

	c := src[len(src)-1]

	padLength := int(c)

	if padLength == 0 || padLength > len(src) {
		return nil, errors.New("pad length error")
	}

	for i := 0; i < padLength; i++ {
		if src[len(src)-padLength+i] != c {
			return nil, errors.New("pad content error")
		}
	}

	return src[:len(src)-padLength], nil

}

func Encrypt(key, src []byte) (data []byte, err error) {

	block, err := aes.NewCipher(key)

	if err != nil {
		return nil, err
	} else if len(src) == 0 {
		return nil, errors.New("src is empty")
	}

	plaintext, err := pkcs7Pad(src, block.BlockSize())

	if err != nil {
		return nil, err
	}

	ciphertext := make([]byte, aes.BlockSize+len(plaintext))

	iv := ciphertext[:aes.BlockSize]
	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		return nil, err
	}

	bm := cipher.NewCBCEncrypter(block, iv)
	bm.CryptBlocks(ciphertext[aes.BlockSize:], plaintext)

	return ciphertext, nil
}
func pkcs7Pad(src []byte, blockSize int) (dest []byte, err error) {
	if blockSize <= 0 {
		return nil, errors.New("block size is 0")
	} else if src == nil || len(src) == 0 {
		return nil, errors.New("src is nil")
	}
	n := blockSize - (len(src) % blockSize)
	pb := make([]byte, len(src)+n)
	copy(pb, src)
	copy(pb[len(src):], bytes.Repeat([]byte{byte(n)}, n))
	return pb, nil
}

// t.Logf(base64.StdEncoding.EncodeToString(encrypt))
// decrypt, err := Decrypt([]byte(key), encrypt)
//
//	if err != nil {
//		fmt.Println(err)
//		return
//	}
//
// fmt.Printf("str is %s\n", string(decrypt))
