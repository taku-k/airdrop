package main

/*
#cgo CFLAGS: -x objective-c
#cgo LDFLAGS: -framework Foundation -framework Cocoa

#import <Cocoa/Cocoa.h>
#import "AirdropDelegate.h"

void run(char *url) {
  @autoreleasepool {
    NSApplication *application = [NSApplication sharedApplication];
    AirdropDelegate *airdropDelegate = [[AirdropDelegate alloc] init];
    [airdropDelegate setUrlName:[NSString stringWithCString:url encoding:NSUTF8StringEncoding]];
    [application setDelegate:airdropDelegate];
    [application run];
  }
}

*/
import "C"
import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"unsafe"
)

func printHelp() string {
	helpText := `
Usage: airdrop path

  Share files located at path by using AirDrop (so you need Mac OSX (>= 10.08))
`
	return strings.TrimSpace(helpText)
}

func main() {
	if len(os.Args) != 2 {
		fmt.Printf("%v\n", printHelp())
		os.Exit(1)
	}

	path, err := filepath.Abs(os.Args[1])
	if err != nil {
		fmt.Printf("%v\n", printHelp())
		os.Exit(1)
	}

	url := C.CString(fmt.Sprintf("file://%s", path))
	defer C.free(unsafe.Pointer(url))
	C.run(url)

	os.Exit(0)
}
