package helper

func MarkStringWithRedColor(stringToMark string) string {
	redColorPrefix := "\033[31m"
	returnCommonColor := "\033[0m"
	return redColorPrefix + stringToMark + returnCommonColor
}
