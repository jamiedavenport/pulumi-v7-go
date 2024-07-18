import * as pulumi from "@pulumi/pulumi";
import * as v7 from "@pulumi/v7-go";

const myRandomResource = new v7.Random("myRandomResource", { length: 24 });
export const output = {
  value: myRandomResource.result,
};
