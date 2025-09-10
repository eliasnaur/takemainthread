#import <Foundation/Foundation.h>

#include "_cgo_export.h"

void gio_runOnMain(uintptr_t handle) {
	dispatch_async(dispatch_get_main_queue(), ^{
		gio_runFunc(handle);
	});
}
