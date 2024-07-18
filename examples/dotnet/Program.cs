using System.Collections.Generic;
using System.Linq;
using Pulumi;
using v7-go = Pulumi.v7-go;

return await Deployment.RunAsync(() => 
{
    var myRandomResource = new v7-go.Random("myRandomResource", new()
    {
        Length = 24,
    });

    return new Dictionary<string, object?>
    {
        ["output"] = 
        {
            { "value", myRandomResource.Result },
        },
    };
});

