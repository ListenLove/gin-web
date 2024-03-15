package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/locales/en"
	"github.com/go-playground/locales/zh"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	entranslations "github.com/go-playground/validator/v10/translations/en"
	zhtranslations "github.com/go-playground/validator/v10/translations/zh"
	"go.uber.org/zap"
	"reflect"
	"strings"
)

var transMap = make(map[string]ut.Translator)
var locales = []string{
	"en",
	"zh",
}

func init() {
	// 准备多语言翻译器
	uni := ut.New(en.New(), zh.New())
	for _, locale := range locales {
		trans, _ := uni.GetTranslator(locale)
		vEngine, ok := binding.Validator.Engine().(*validator.Validate)
		if !ok {
			zap.L().Error("binding.Validator.Engine() 出错", zap.String("locale", locale))
			return
		}
		// 注册一个函数，获取struct tag里自定义的json作为字段名
		vEngine.RegisterTagNameFunc(func(fld reflect.StructField) string {
			name := strings.SplitN(fld.Tag.Get("json"), ",", 2)[0]
			if name == "-" {
				return ""
			}
			return name
		})
		switch locale {
		case "zh":
			_ = zhtranslations.RegisterDefaultTranslations(vEngine, trans)
			transMap[locale] = trans
		case "en":
			_ = entranslations.RegisterDefaultTranslations(vEngine, trans)
			transMap[locale] = trans
		default:
			_ = zhtranslations.RegisterDefaultTranslations(vEngine, trans)
			transMap[locale] = trans
		}
	}
}

func getTransByLocale(locale string) ut.Translator {
	trans, ok := transMap[locale]
	if !ok {
		return transMap["zh"]
	}
	return trans
}

func Translations() gin.HandlerFunc {
	return func(c *gin.Context) {
		locale := c.GetHeader("locale")
		trans := getTransByLocale(locale)
		c.Set("trans", trans)
		c.Next()
	}
}
