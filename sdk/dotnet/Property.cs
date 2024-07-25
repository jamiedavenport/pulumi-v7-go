// *** WARNING: this file was generated by pulumi. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace JamieDavenport.V7Go
{
    [V7GoResourceType("v7-go:index:Property")]
    public partial class Property : global::Pulumi.CustomResource
    {
        [Output("description")]
        public Output<string?> Description { get; private set; } = null!;

        [Output("inputs")]
        public Output<ImmutableArray<Outputs.PropertyInput>> Inputs { get; private set; } = null!;

        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        [Output("projectId")]
        public Output<string> ProjectId { get; private set; } = null!;

        [Output("propertyId")]
        public Output<string> PropertyId { get; private set; } = null!;

        [Output("tool")]
        public Output<string> Tool { get; private set; } = null!;

        [Output("type")]
        public Output<string> Type { get; private set; } = null!;

        [Output("workspaceId")]
        public Output<string> WorkspaceId { get; private set; } = null!;


        /// <summary>
        /// Create a Property resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Property(string name, PropertyArgs args, CustomResourceOptions? options = null)
            : base("v7-go:index:Property", name, args ?? new PropertyArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Property(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("v7-go:index:Property", name, null, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
                PluginDownloadURL = "github://api.github.com/jamiedavenport/pulumi-v7-go",
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing Property resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Property Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new Property(name, id, options);
        }
    }

    public sealed class PropertyArgs : global::Pulumi.ResourceArgs
    {
        [Input("description")]
        public Input<string>? Description { get; set; }

        [Input("inputs", required: true)]
        private InputList<Inputs.PropertyInputArgs>? _inputs;
        public InputList<Inputs.PropertyInputArgs> Inputs
        {
            get => _inputs ?? (_inputs = new InputList<Inputs.PropertyInputArgs>());
            set => _inputs = value;
        }

        [Input("name", required: true)]
        public Input<string> Name { get; set; } = null!;

        [Input("projectId", required: true)]
        public Input<string> ProjectId { get; set; } = null!;

        [Input("tool", required: true)]
        public Input<string> Tool { get; set; } = null!;

        [Input("type", required: true)]
        public Input<string> Type { get; set; } = null!;

        [Input("workspaceId", required: true)]
        public Input<string> WorkspaceId { get; set; } = null!;

        public PropertyArgs()
        {
        }
        public static new PropertyArgs Empty => new PropertyArgs();
    }
}
