package license_server

type LicenseType int

const (
	File           LicenseType = iota // 文件授权
	Dongle                            // 电子狗授权
	PrivateServer                     // 私有云服务
	PublicServer                      // 公有云服务
	LocalContainer                    // 本地容器服务
)

func (lt LicenseType) ToInt() int {
	return int(lt)
}

func (lt LicenseType) String() string {
	switch lt {
	case File:
		return "FILE"
	case Dongle:
		return "DONGLE"
	case PrivateServer:
		return "PRIVATE_SERVER"
	case PublicServer:
		return "PUBLIC_SERVER"
	case LocalContainer:
		return "LOCAL_CONTAINER"
	default:
		return "UNKNOWN"
	}
}

func ParseLicenseType(value int) LicenseType {
	switch value {
	case 0:
		return File
	case 1:
		return Dongle
	case 2:
		return PrivateServer
	case 3:
		return PublicServer
	case 4:
		return LocalContainer
	default:
		return File
	}
}
