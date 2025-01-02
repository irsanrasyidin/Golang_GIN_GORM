package utils

func GetOrder(s string) int {
	switch s {
	case "insertpostgresql":
		return 1
	case "insertjson":
		return 2
	case "getbyidpostgresql":
		return 3
	case "getbyidjson":
		return 4
	case "getbynamepostgresql":
		return 5
	case "getbynamejson":
		return 6
	case "getbyemailpostgresql":
		return 7
	case "getbyemailjson":
		return 8
	case "getbygenderpostgresql":
		return 9
	case "getbygenderjson":
		return 10
	case "getallpostgresql":
		return 11
	case "getalljson":
		return 12
	case "updatebyidpostgresql":
		return 13
	case "updatebyidjson":
		return 14
	case "deletebyidpostgresql":
		return 15
	case "deletebyidjson":
		return 16
	case "deleteallpostgresql":
		return 17
	case "deletealljson":
		return 18
	default:
		return 19
	}
}
