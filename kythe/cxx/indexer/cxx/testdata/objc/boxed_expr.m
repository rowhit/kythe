// Test that indexing works as expected when a boxed expression is present.

// Mock NSNumber so we don't have to include NS* classes.
@interface NSNumber
+(instancetype) numberWithInt:(int)i;
@end

//- @Box defines/binding BoxIface
@interface Box

//- @"foo" defines/binding FooDecl
-() foo;

@end

//- @Box defines/binding BoxImpl
@implementation Box

//- @"foo " defines/binding FooDefn
//- @"foo " completes/uniquely FooDecl
-() foo {
  return @8;
}

@end

//- @main defines/binding Main
int main(int argc, char **argv) {
  Box *box = [[Box alloc] init];

  //- @"[box foo]" ref/call FooDecl
  //- @"[box foo]" childof Main
  //- @"[box foo]".node/kind anchor
  [box foo];

  return 0;
}

