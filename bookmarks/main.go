package bookmarks

type bookmark struct {
	Title string
	Command string
}

type Bookmarks []bookmark

func (b *Bookmarks) AddBookmark(title, command string) {
	*b = append(*b, bookmark{title, command})
}

func (b *Bookmarks) RemoveBookmark(title string) {
	for i, bookmark := range *b {
		if bookmark.Title == title {
			*b = append((*b)[:i], (*b)[i+1:]...)
			break
		}
	}
}

func (b *Bookmarks) ListBookmarks() []string {
	var titles []string
	for _, bookmark := range *b {
		titles = append(titles, bookmark.Title)
	}
	return titles
}

func (b *Bookmarks) GetCommand(title string) string {
	for _, bookmark := range *b {
		if bookmark.Title == title {
			return bookmark.Command
		}
	}
	return ""
}

func (b *Bookmarks) EditCommand(title, command string) {
	for i, bookmark := range *b {
		if bookmark.Title == title {
			(*b)[i].Command = command
			break
		}
	}
}
