package fileUtils

import (
	"go-web/utils/contants"
	"strings"
)

// ImageContentByFileName 根据图片名称获取http content
func ImageContentByFileName(imageFileName string) (content string) {
	if imageFileName == "" {
		return
	}
	index := strings.LastIndex(imageFileName, contants.FileNameSuffix)
	if index == -1 {
		return
	}
	s := imageFileName[index:]
	content = ImageContentByFilePrefix(s)
	return
}

// ImageContentByFilePrefix 根据后缀获取http content
func ImageContentByFilePrefix(filePrefix string) (content string) {
	if filePrefix == contants.FileNameSuffixPng {
		content = contants.FileContentTypePng
	} else if filePrefix == contants.FileNameSuffixJpeg || filePrefix == contants.FileNameSuffixJpg {
		content = contants.FileContentTypeJpg
	} else if filePrefix == contants.FileNameSuffixGif {
		content = contants.FileContentTypeGif
	} else if filePrefix == contants.FileNameSuffixWebp {
		content = contants.FileContentTypeWebp
	} else if filePrefix == contants.FileNameSuffixBmp {
		content = contants.FileContentTypeBmp
	} else {
		content = contants.FileContentTypeSvgXml
	}
	return
}
