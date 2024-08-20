package globals

import (
	"github.com/mari-muthu-k/gin-template/model/appschema"
	"gorm.io/gorm"
)

var RelationalDb *gorm.DB
var AppKeys appschema.CertificateKeys