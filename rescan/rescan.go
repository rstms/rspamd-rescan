package rescan

import (
    "fmt"
)

func Rescan(file string) error {
    fmt.Printf("rescan: %s\n", file)
    rspamd := NewAPIClient()
    status, err := rspamd.get("stat")
    if err != nil {
	return err
    }
    fmt.Println(status);
    return nil
}
