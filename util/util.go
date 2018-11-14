package util

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/hex"
	"errors"
	"fmt"
	"io"
	"os"
	"strconv"
	"sync"
	"time"
)

var commonIV = []byte{0x00, 0x01, 0x02, 0x03, 0x04, 0x05, 0x06, 0x07, 0x08, 0x09, 0x0a, 0x0b, 0x0c, 0x0d, 0x0e, 0x0f}

//aes加密,返回16进制数据
func AesEn(s string) string {
	//需要去加密的字符串
	plaintext := []byte(s)
	//如果传入加密串的话，plaint就是传入的字符串
	if len(os.Args) > 1 {
		plaintext = []byte(os.Args[1])
	}

	//aes的加密字符串,经测试,任意32位字符
	key_text := "astaxie12798akljzmknm.ahkjkljl;k"
	if len(os.Args) > 2 {
		key_text = os.Args[2]
	}

	fmt.Println(len(key_text))

	// 创建加密算法aes
	c, err := aes.NewCipher([]byte(key_text))
	if err != nil {
		fmt.Printf("Error: NewCipher(%d bytes) = %s", len(key_text), err)
		os.Exit(-1)
	}

	//加密字符串
	cfb := cipher.NewCFBEncrypter(c, commonIV)
	ciphertext := make([]byte, len(plaintext))
	cfb.XORKeyStream(ciphertext, plaintext)
	fmt.Printf("%s=>%x\n", plaintext, ciphertext)
	return string(hex.EncodeToString(ciphertext)) //16进制转换

	/*// 解密字符串
	cfbdec := cipher.NewCFBDecrypter(c, commonIV)
	plaintextCopy := make([]byte, len(plaintext))
	cfbdec.XORKeyStream(plaintextCopy, ciphertext)
	fmt.Printf("%x=>%s\n", ciphertext, plaintextCopy)*/
}

//日期差计算,年月日计算
func SubDate(date1, date2 time.Time) string {
	var y, m, d int
	y = date1.Year() - date2.Year()
	if date1.Month() < date2.Month() {
		y--
		m = 12 - int(date2.Month()) + int(date1.Month())
	} else {
		m = int(date1.Month()) - int(date2.Month())
	}
	//天数模糊计算
	if date1.Day() < date2.Day() {
		m--
		//闰年,29天
		day := []int{31, 28, 31, 30, 31, 30, 31, 31, 30, 31, 30, 31}

		if date2.Year()%4 == 0 && date2.Year()%100 != 0 || date2.Year()%400 == 0 {
			d = day[date2.Month()-1] + 1 - date2.Day() + date1.Day()
		} else {
			d = day[date2.Month()-1] - date2.Day() + date1.Day()
		}
	} else {
		d = date1.Day() - date2.Day()
	}
	return strconv.Itoa(y) + "年" + strconv.Itoa(m) + "月" + strconv.Itoa(d) + "日"
}

//时间格式化2006-01-02 15:04:05
type JsonTime time.Time

//实现它的json序列化方法
func (this JsonTime) MarshalJSON() ([]byte, error) {
	var stamp = fmt.Sprintf("\"%s\"", time.Time(this).Format("2006-01-02 15:04:05"))
	return []byte(stamp), nil
}

//时间格式化2006-01-02
type JsonDate time.Time

//实现它的json序列化方法
func (this JsonDate) MarshalJSON() ([]byte, error) {
	var stamp = fmt.Sprintf("\"%s\"", time.Time(this).Format("2006-01-02"))
	return []byte(stamp), nil
}

// key
const aesTable = "abcdefghijklmnopkrstuvwsyz012345"

var (
	block cipher.Block
	mutex sync.Mutex
)

// AES加密
func Encrypt(src []byte) ([]byte, error) {
	// 验证输入参数
	// 必须为aes.Blocksize的倍数
	if len(src)%aes.BlockSize != 0 {
		return nil, errors.New("crypto/cipher: input not full blocks")
	}

	encryptText := make([]byte, aes.BlockSize+len(src))

	iv := encryptText[:aes.BlockSize]
	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		return nil, err
	}

	mode := cipher.NewCBCEncrypter(block, iv)

	mode.CryptBlocks(encryptText[aes.BlockSize:], src)

	return encryptText, nil
}

// AES解密
func Decrypt(src []byte) ([]byte, error) {
	// hex
	decryptText, err := hex.DecodeString(fmt.Sprintf("%x", string(src)))
	if err != nil {
		return nil, err
	}

	// 长度不能小于aes.Blocksize
	if len(decryptText) < aes.BlockSize {
		return nil, errors.New("crypto/cipher: ciphertext too short")
	}

	iv := decryptText[:aes.BlockSize]
	decryptText = decryptText[aes.BlockSize:]

	// 验证输入参数
	// 必须为aes.Blocksize的倍数
	if len(decryptText)%aes.BlockSize != 0 {
		return nil, errors.New("crypto/cipher: ciphertext is not a multiple of the block size")
	}

	mode := cipher.NewCBCDecrypter(block, iv)

	mode.CryptBlocks(decryptText, decryptText)

	return decryptText, nil
}

func init() {
	mutex.Lock()
	defer mutex.Unlock()

	if block != nil {
		return
	}

	cblock, err := aes.NewCipher([]byte(aesTable))
	if err != nil {
		panic("aes.NewCipher: " + err.Error())
	}

	block = cblock
}