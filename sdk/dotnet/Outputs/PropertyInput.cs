// *** WARNING: this file was generated by pulumi. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace JamieDavenport.V7Go.Outputs
{

    [OutputType]
    public sealed class PropertyInput
    {
        public readonly string PropertyId;

        [OutputConstructor]
        private PropertyInput(string propertyId)
        {
            PropertyId = propertyId;
        }
    }
}
