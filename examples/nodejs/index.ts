import * as pulumi from "@pulumi/pulumi";
import * as v7_go from "@pulumi/v7-go";

const myProject = new v7_go.Project("myProject", {
    workspaceId: "01905686-d113-79ff-9b33-0fa7c3ff6664",
    name: "My New Pulumi Project Updated",
});
const myInput = new v7_go.Property("myInput", {
    workspaceId: myProject.workspaceId,
    projectId: myProject.projectId,
    name: "My Input",
    tool: "manual",
    type: "text",
    inputs: [],
});
const myOutput = new v7_go.Property("myOutput", {
    workspaceId: myProject.workspaceId,
    projectId: myProject.projectId,
    name: "My Output",
    description: "Return the text in uppercase",
    tool: "gpt_4o",
    type: "text",
    inputs: [{
        propertyId: myInput.propertyId,
    }],
});
export const output = [];
