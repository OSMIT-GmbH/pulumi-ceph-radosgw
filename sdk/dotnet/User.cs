// *** WARNING: this file was generated by pulumi. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.CephRadosgw
{
    [CephRadosgwResourceType("ceph-radosgw:index:User")]
    public partial class User : global::Pulumi.CustomResource
    {
        [Output("_assimilated")]
        public Output<bool> _assimilated { get; private set; } = null!;

        [Output("capabilities")]
        public Output<Outputs.Capabilities?> Capabilities { get; private set; } = null!;

        [Output("displayName")]
        public Output<string?> DisplayName { get; private set; } = null!;

        [Output("email")]
        public Output<string?> Email { get; private set; } = null!;

        [Output("keys")]
        public Output<ImmutableArray<Outputs.KeyEntry>> Keys { get; private set; } = null!;

        [Output("maxBuckets")]
        public Output<int?> MaxBuckets { get; private set; } = null!;

        [Output("suspended")]
        public Output<bool?> Suspended { get; private set; } = null!;

        [Output("userId")]
        public Output<string> UserId { get; private set; } = null!;


        /// <summary>
        /// Create a User resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public User(string name, UserArgs args, CustomResourceOptions? options = null)
            : base("ceph-radosgw:index:User", name, args ?? new UserArgs(), MakeResourceOptions(options, ""))
        {
        }

        private User(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("ceph-radosgw:index:User", name, null, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing User resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static User Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new User(name, id, options);
        }
    }

    public sealed class UserArgs : global::Pulumi.ResourceArgs
    {
        [Input("capabilities")]
        public Input<Inputs.CapabilitiesArgs>? Capabilities { get; set; }

        [Input("displayName")]
        public Input<string>? DisplayName { get; set; }

        [Input("email")]
        public Input<string>? Email { get; set; }

        [Input("maxBuckets")]
        public Input<int>? MaxBuckets { get; set; }

        [Input("suspended")]
        public Input<bool>? Suspended { get; set; }

        [Input("userId", required: true)]
        public Input<string> UserId { get; set; } = null!;

        public UserArgs()
        {
        }
        public static new UserArgs Empty => new UserArgs();
    }
}
