#import "AirdropDelegate.h"

@interface AirdropDelegate () <NSSharingServiceDelegate>
@end

@implementation AirdropDelegate

- (void)setUrlName:(NSString *)url {
    _urlName = url;
  }

- (void)applicationDidFinishLaunching:(NSNotification *)notification {
    NSString *sharingServiceName = NSSharingServiceNameSendViaAirDrop;
    NSSharingService *sharingService = [NSSharingService sharingServiceNamed:sharingServiceName];
    sharingService.delegate = self;
    NSMutableArray *mutableItems = [NSMutableArray array];
    NSURL *url = [NSURL URLWithString:_urlName];
    [mutableItems addObject:url];
    [sharingService performWithItems:mutableItems];
}

#pragma mark - NSSharingServiceDelegate

- (void)sharingService:(NSSharingService *)sharingService
         didShareItems:(NSArray *)items
{
    exit(EXIT_SUCCESS);
}

- (void)sharingService:(NSSharingService *)sharingService
   didFailToShareItems:(NSArray *)items
                 error:(NSError *)error
{
    exit(EXIT_FAILURE);
}


@end
