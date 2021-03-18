package autocode

func GetDbType(dt string) string {
	switch dt {
	case "varchar":
		return "string"
	case "datetime":
		return "time.Time"
	case "timestamp":
		return "time.Time"
	case "double":
		return "float64"
	case "float":
		return "float64"
	case "decimal":
		return "float64"
	case "int":
		return "int64"
	case "tinyint":
		return "int64"
	case "smallint":
		return "int64"
	case "mediumint":
		return "int64"
	case "integer":
		return "int64"
	case "bigint":
		return "int64"
	default:
		return "string"
	}

}
