import pulumi
import pulumi_v7-go as v7-go

my_random_resource = v7-go.Random("myRandomResource", length=24)
pulumi.export("output", {
    "value": my_random_resource.result,
})
