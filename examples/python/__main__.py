import pulumi
import pulumi_v7_go as v7_go

my_project = v7_go.Project("myProject",
    workspace_id="01905686-d113-79ff-9b33-0fa7c3ff6664",
    name="My New Pulumi Project Updated")
my_input = v7_go.Property("myInput",
    workspace_id=my_project.workspace_id,
    project_id=my_project.project_id,
    name="My Input",
    tool="manual",
    type="text",
    inputs=[])
my_output = v7_go.Property("myOutput",
    workspace_id=my_project.workspace_id,
    project_id=my_project.project_id,
    name="My Output",
    description="Return the text in uppercase",
    tool="gpt_4o",
    type="text",
    inputs=[v7_go.PropertyInputArgs(
        property_id=my_input.property_id,
    )])
pulumi.export("output", [])
