package exp

type CheckFunc func(string, map[string]string, string) (bool, string)
type ExploitFunc func(string, map[string]string, string, string) string
type UploadFunc func(string, map[string]string, string, string, string) string

type ExploitEntry struct {
	Check   CheckFunc
	Exploit ExploitFunc
	Upload  UploadFunc
}

var registry = make(map[string]ExploitEntry)

func RegisterExploit(name string, entry ExploitEntry) {
	registry[name] = entry
}

func GetExploit(name string) (ExploitEntry, bool) {
	e, ok := registry[name]
	return e, ok
}

func ListAll() map[string]ExploitEntry {
	return registry
}
