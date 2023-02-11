package utils

import (
	"bytes"
	"encoding/hex"
	"lark/pkg/constant"
	"strconv"
	"strings"
)

var (
	imageSuffixes = map[string]interface{}{"jpeg": 1, "jpg": 1, "png": 1, "gif": 1, "tif": 1, "bmp": 1, "dwg": 1}
	audioSuffixes = map[string]interface{}{"mp3": 1, "wma": 1, "wav": 1, "mid": 1, "ape": 1, "flac": 1}
	mediaSuffixes = map[string]interface{}{"rmvb": 1, "flv": 1, "mp4": 1, "mpg": 1, "mpeg": 1, "avi": 1, "rm": 1, "mov": 1, "wmv": 1, "webm": 1}
	fileTypes     = map[string]string{}
)

func init() {
	setFileType("ffd8ffe000104a464946", "jpg")  //JPEG (jpg)
	setFileType("89504e470d0a1a0a0000", "png")  //PNG (png)
	setFileType("47494638396126026f01", "gif")  //GIF (gif)
	setFileType("49492a00227105008037", "tif")  //TIFF (tif)
	setFileType("424d228c010000000000", "bmp")  //16色位图(bmp)
	setFileType("424d8240090000000000", "bmp")  //24位位图(bmp)
	setFileType("424d8e1b030000000000", "bmp")  //256色位图(bmp)
	setFileType("41433130313500000000", "dwg")  //CAD (dwg)
	setFileType("3c21444f435459504520", "html") //HTML (html)   3c68746d6c3e0  3c68746d6c3e0
	setFileType("3c68746d6c3e0", "html")        //HTML (html)   3c68746d6c3e0  3c68746d6c3e0
	setFileType("3c21646f637479706520", "htm")  //HTM (htm)
	setFileType("48544d4c207b0d0a0942", "css")  //css
	setFileType("696b2e71623d696b2e71", "js")   //js
	setFileType("7b5c727466315c616e73", "rtf")  //Rich Text Format (rtf)
	setFileType("38425053000100000000", "psd")  //Photoshop (psd)
	setFileType("46726f6d3a203d3f6762", "eml")  //Email [Outlook Express 6] (eml)
	setFileType("d0cf11e0a1b11ae10000", "vsd")  //Visio 绘图
	setFileType("5374616E64617264204A", "mdb")  //MS Access (mdb)
	setFileType("252150532D41646F6265", "ps")

	setFileType("255044462d312e350d0a", "pdf")          //Adobe Acrobat (pdf)
	setFileType("D0CF11E0", "xls")                      //xls
	setFileType("504B030414000600080000002100", "xlsx") //xls
	setFileType("d0cf11e0a1b11ae10000", "doc")          //MS Excel 注意：word、msi 和 excel的文件头一样
	setFileType("504b0304140006000800", "docx")         //docx文件
	setFileType("d0cf11e0a1b11ae10000", "wps")          //WPS文字wps、表格et、演示dps都是一样的

	setFileType("2e524d46000000120001", "rmvb") //rmvb/rm相同
	setFileType("464c5601050000000900", "flv")  //flv与f4v相同
	setFileType("00000020667479706d70", "mp4")
	setFileType("49443303000000002176", "mp3")
	setFileType("000001ba210001000180", "mpg")
	setFileType("3026b2758e66cf11a6d9", "wmv") //wmv与asf相同
	setFileType("52494646e27807005741", "wav") //Wave (wav)
	setFileType("52494646246009005741", "wav") //Wave (wav)
	setFileType("52494646", "wav")             //Wave (wav)

	setFileType("52494646d07d60074156", "avi")
	setFileType("1a45dfa3a34286810142", "webm")

	setFileType("4d546864000000060001", "mid") //MIDI (mid)
	setFileType("504b0304140000000800", "zip")
	setFileType("526172211a0700cf9073", "rar")
	setFileType("235468697320636f6e66", "ini")
	setFileType("504b03040a0000000000", "jar")
	setFileType("4d5a9000030000000400", "exe")        //可执行文件
	setFileType("3c25402070616765206c", "jsp")        //jsp文件
	setFileType("4d616e69666573742d56", "mf")         //MF文件
	setFileType("3c3f786d6c2076657273", "xml")        //xml文件
	setFileType("494e5345525420494e54", "sql")        //xml文件
	setFileType("7061636b616765207765", "java")       //java文件
	setFileType("406563686f206f66660d", "bat")        //bat文件
	setFileType("1f8b0800000000000000", "gz")         //gz文件
	setFileType("6c6f67346a2e726f6f74", "properties") //bat文件
	setFileType("cafebabe0000002e0041", "class")      //bat文件
	setFileType("49545346030000006000", "chm")        //bat文件
	setFileType("04000000010000001300", "mxp")        //bat文件
	setFileType("6431303a637265617465", "torrent")
	setFileType("6D6F6F76", "mov")         //Quicktime (mov)
	setFileType("FF575043", "wpd")         //WordPerfect (wpd)
	setFileType("CFAD12FEC5FD746F", "dbx") //Outlook Express (dbx)
	setFileType("2142444E", "pst")         //Outlook (pst)
	setFileType("AC9EBD8F", "qdf")         //Quicken (qdf)
	setFileType("E3828596", "pwl")         //Windows Password (pwl)
	setFileType("2E7261FD", "ram")         //Real Audio (ram)
}

func setFileType(key string, val string) {
	fileTypes[key] = val
}

// 获取前面结果字节的二进制
func bytesToHexString(src []byte) string {
	res := bytes.Buffer{}
	if src == nil || len(src) <= 0 {
		return ""
	}
	temp := make([]byte, 0)
	for _, v := range src {
		sub := v & 0xFF
		hv := hex.EncodeToString(append(temp, sub))
		if len(hv) < 2 {
			res.WriteString(strconv.FormatInt(int64(0), 10))
		}
		res.WriteString(hv)
	}
	return res.String()
}

func GetFileType(buf []byte) (fileType string) {
	var (
		fileCode string
		key      string
		val      string
	)
	fileCode = bytesToHexString(buf)
	for key, val = range fileTypes {
		if strings.HasPrefix(fileCode, strings.ToLower(key)) ||
			strings.HasPrefix(key, strings.ToLower(fileCode)) {
			fileType = val
		}
	}
	return
}

func GetFileTypeBySuffix(suffix string) int32 {
	var (
		ok bool
	)
	if _, ok = imageSuffixes[suffix]; ok {
		return constant.FILE_TYPE_IMAGE
	}
	if _, ok = audioSuffixes[suffix]; ok {
		return constant.FILE_TYPE_AUDIO
	}
	if _, ok = mediaSuffixes[suffix]; ok {
		return constant.FILE_TYPE_MEDIA
	}
	return constant.FILE_TYPE_FILE
}
