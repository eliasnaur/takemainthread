#import <AppKit/AppKit.h>

#include "_cgo_export.h"

void guipkg2_runOnMain(uintptr_t handle) {
	dispatch_async(dispatch_get_main_queue(), ^{
		guipkg2_runFunc(handle);
	});
}

void guipkg2_createWindow(CGFloat width, CGFloat height) {
	@autoreleasepool {
		NSRect rect = NSMakeRect(0, 0, width, height);
		NSUInteger styleMask = NSTitledWindowMask |
			NSResizableWindowMask |
			NSMiniaturizableWindowMask |
			NSClosableWindowMask;

		NSWindow* window = [[NSWindow alloc] initWithContentRect:rect
													   styleMask:styleMask
														 backing:NSBackingStoreBuffered
														   defer:NO];
		window.title = @"guipkg2 window";
		window.releasedWhenClosed = NO;
		[window makeKeyAndOrderFront:nil];
	}
}

@interface AppListener : NSObject
@end

// Hold on to the app listener because NSNotificationCenter
// doesn't.
static AppListener *appListener;

@implementation AppListener
- (void)launchFinished:(NSNotification *)notification {
	appListener = nil;
	guipkg2_onFinishLaunching();
}
@end

void guipkg2_init() {
	@autoreleasepool {
		appListener = [[AppListener alloc] init];
		[[NSNotificationCenter defaultCenter] addObserver:appListener
												 selector:@selector(launchFinished:)
													 name:NSApplicationDidFinishLaunchingNotification
												   object:nil];
	}
}

// The following is app global initialization, which is optional.

@interface GioAppDelegate : NSObject<NSApplicationDelegate>
@end

@implementation GioAppDelegate
- (void)applicationDidFinishLaunching:(NSNotification *)aNotification {
	[NSApp setActivationPolicy:NSApplicationActivationPolicyRegular];
	[NSApp activateIgnoringOtherApps:YES];
}
@end

void guipkg2_main() {
	@autoreleasepool {
		[NSApplication sharedApplication];
		GioAppDelegate *del = [[GioAppDelegate alloc] init];
		[NSApp setDelegate:del];

		NSMenuItem *mainMenu = [NSMenuItem new];

		NSMenu *menu = [NSMenu new];
		NSMenuItem *hideMenuItem = [[NSMenuItem alloc] initWithTitle:@"Hide"
															  action:@selector(hide:)
													   keyEquivalent:@"h"];
		[menu addItem:hideMenuItem];
		NSMenuItem *quitMenuItem = [[NSMenuItem alloc] initWithTitle:@"Quit"
															  action:@selector(terminate:)
													   keyEquivalent:@"q"];
		[menu addItem:quitMenuItem];
		[mainMenu setSubmenu:menu];
		NSMenu *menuBar = [NSMenu new];
		[menuBar addItem:mainMenu];
		[NSApp setMainMenu:menuBar];

		[NSApp run];
	}
}
