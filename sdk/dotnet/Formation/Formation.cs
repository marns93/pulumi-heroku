// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace Pulumiverse.Heroku.Formation
{
    [HerokuResourceType("heroku:formation/formation:Formation")]
    public partial class Formation : global::Pulumi.CustomResource
    {
        [Output("appId")]
        public Output<string> AppId { get; private set; } = null!;

        [Output("quantity")]
        public Output<int> Quantity { get; private set; } = null!;

        [Output("size")]
        public Output<string> Size { get; private set; } = null!;

        [Output("type")]
        public Output<string> Type { get; private set; } = null!;


        /// <summary>
        /// Create a Formation resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Formation(string name, FormationArgs args, CustomResourceOptions? options = null)
            : base("heroku:formation/formation:Formation", name, args ?? new FormationArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Formation(string name, Input<string> id, FormationState? state = null, CustomResourceOptions? options = null)
            : base("heroku:formation/formation:Formation", name, state, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
                PluginDownloadURL = "github://api.github.com/pulumiverse",
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing Formation resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Formation Get(string name, Input<string> id, FormationState? state = null, CustomResourceOptions? options = null)
        {
            return new Formation(name, id, state, options);
        }
    }

    public sealed class FormationArgs : global::Pulumi.ResourceArgs
    {
        [Input("appId", required: true)]
        public Input<string> AppId { get; set; } = null!;

        [Input("quantity", required: true)]
        public Input<int> Quantity { get; set; } = null!;

        [Input("size", required: true)]
        public Input<string> Size { get; set; } = null!;

        [Input("type", required: true)]
        public Input<string> Type { get; set; } = null!;

        public FormationArgs()
        {
        }
        public static new FormationArgs Empty => new FormationArgs();
    }

    public sealed class FormationState : global::Pulumi.ResourceArgs
    {
        [Input("appId")]
        public Input<string>? AppId { get; set; }

        [Input("quantity")]
        public Input<int>? Quantity { get; set; }

        [Input("size")]
        public Input<string>? Size { get; set; }

        [Input("type")]
        public Input<string>? Type { get; set; }

        public FormationState()
        {
        }
        public static new FormationState Empty => new FormationState();
    }
}