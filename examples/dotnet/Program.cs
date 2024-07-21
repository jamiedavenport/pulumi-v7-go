using System.Collections.Generic;
using System.Linq;
using Pulumi;
using V7Go = Pulumi.V7Go;

return await Deployment.RunAsync(() => 
{
    var myProject = new V7Go.Project("myProject", new()
    {
        WorkspaceId = "01905686-d113-79ff-9b33-0fa7c3ff6664",
        Name = "My New Pulumi Project Updated",
    });

    var myInput = new V7Go.Property("myInput", new()
    {
        WorkspaceId = myProject.WorkspaceId,
        ProjectId = myProject.ProjectId,
        Name = "My Input",
        Tool = "manual",
        Type = "text",
        Inputs = new[] {},
    });

    var myOutput = new V7Go.Property("myOutput", new()
    {
        WorkspaceId = myProject.WorkspaceId,
        ProjectId = myProject.ProjectId,
        Name = "My Output",
        Description = "Return the text in uppercase",
        Tool = "gpt_4o",
        Type = "text",
        Inputs = new[]
        {
            new V7Go.Inputs.PropertyInputArgs
            {
                PropertyId = myInput.PropertyId,
            },
        },
    });

    return new Dictionary<string, object?>
    {
        ["output"] = new[] {},
    };
});

