package writer

func fixLen(line string, lenline int) string {
	for {

		if len(line) < lenline {
			line = " " + line
		} else {
			return line
		}
		if len(line) < lenline {
			line = line + " "
		} else {
			return line
		}
	}
}
