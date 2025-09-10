#import <AppKit/AppKit.h>

#include "_cgo_export.h"

void guipkg1_runOnMain(uintptr_t handle) {
	dispatch_async(dispatch_get_main_queue(), ^{
		guipkg1_runFunc(handle);
	});
}

void guipkg1_createWindow(CGFloat width, CGFloat height) {
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
		window.title = @"guipkg1 window";
		window.releasedWhenClosed = NO;
		[window makeKeyAndOrderFront:nil];
	}
}

@interface GUIPkg1AppListener : NSObject
@end

// Hold on to the app listener because NSNotificationCenter
// doesn't.
static GUIPkg1AppListener *appListener;

@implementation GUIPkg1AppListener
- (void)launchFinished:(NSNotification *)notification {
	appListener = nil;
	guipkg1_onFinishLaunching();
}
@end

void guipkg1_init() {
	@autoreleasepool {
		appListener = [[GUIPkg1AppListener alloc] init];
		[[NSNotificationCenter defaultCenter] addObserver:appListener
												 selector:@selector(launchFinished:)
													 name:NSApplicationDidFinishLaunchingNotification
												   object:nil];
	}
}

// The following is app global initialization, which is optional.

@interface GUIPkg1AppDelegate : NSObject<NSApplicationDelegate>
@end

@implementation GUIPkg1AppDelegate
- (void)applicationDidFinishLaunching:(NSNotification *)aNotification {
	[NSApp setActivationPolicy:NSApplicationActivationPolicyRegular];
	[NSApp activateIgnoringOtherApps:YES];
}
@end

void guipkg1_main() {
	@autoreleasepool {
		[NSApplication sharedApplication];
		GUIPkg1AppDelegate *del = [[GUIPkg1AppDelegate alloc] init];
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
