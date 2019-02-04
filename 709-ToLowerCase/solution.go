package solution

func ToLowerCase(str string) string {
    lowerRunes := []rune(str)
    for i, val:= range str {
        if isLowerCase(val) {
            lowerRunes[i] = val
            continue
        }
        lowerRunes[i] = lowerCase(val)
    }
    return string(lowerRunes)
}

func isLowerCase(r rune) bool {
    return r<'A' || r >'Z'
}

func lowerCase(r rune) rune {
    return r+32
}
