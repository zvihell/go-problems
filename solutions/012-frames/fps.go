package fps

func frames(minutes, fps int) int {
	seconds := minutes * 60
	frames := seconds * fps
	return frames
}
