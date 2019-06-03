# AT001

The AT001 analyzer reports likely incorrect uses of `TestCase`
which do not define a `CheckDestroy` function. `CheckDestroy` is used to verify
that test infrastructure has been removed at the end of an acceptance test.

More information can be found at:
https://www.terraform.io/docs/extend/testing/acceptance-tests/testcase.html#checkdestroy

## Flagged Code

NOTE: This analyzer does not differentiate between resource acceptance tests and data source acceptance tests. This is by design to ensure authors add the equivalent resource `CheckDestroy` function to data source testing, if available.

```go
# Trivial Example
resource.Test(t, resource.TestCase{})

# Real World Example
func TestAccExampleThing_Attr1(t *testing.T) {
    resource.ParallelTest(t, resource.TestCase{
        PreCheck:  func() { testAccPreCheck(t) },
        Providers: testAccProviders,
        Steps: []resource.TestStep{
            {
                Config: testAccExampleThingConfig(),
                Check: resource.ComposeTestCheckFunc(
                    resource.TestCheckResourceAttrSet("example_thing.test", "attr1"),
                ),
            },
        },
    })
}
```

## Passing Code

```go
func testAccCheckExampleThingDestroy(s *terraform.State) error {
    for _, rs := range s.RootModule().Resources {
        if rs.Type != "example_thing" {
            continue
        }

        /* Code to check API for existence of Example Thing */
    }

    return nil
}

# Trivial Example
resource.Test(t, resource.TestCase{
    CheckDestroy: testAccCheckExampleThingDestroy,
})

# Real World Example
func TestAccExampleThing_Attr1(t *testing.T) {
    resource.ParallelTest(t, resource.TestCase{
        PreCheck:     func() { testAccPreCheck(t) },
        Providers:    testAccProviders,
        CheckDestroy: testAccCheckExampleThingDestroy,
        Steps: []resource.TestStep{
            {
                Config: testAccExampleThingConfig(),
                Check: resource.ComposeTestCheckFunc(
                    resource.TestCheckResourceAttrSet("example_thing.test", "attr1"),
                ),
            },
        },
    })
}
```
