// *** WARNING: this file was generated by pulumi. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.CephRadosgw
{
    [CephRadosgwResourceType("ceph-radosgw:index:SubUser")]
    public partial class SubUser : global::Pulumi.CustomResource
    {
        [Output("_assimilated")]
        public Output<bool> _assimilated { get; private set; } = null!;

        [Output("fullname")]
        public Output<string> Fullname { get; private set; } = null!;

        [Output("generateKey")]
        public Output<bool?> GenerateKey { get; private set; } = null!;

        [Output("keyType")]
        public Output<Pulumi.CephRadosgw.KeyType?> KeyType { get; private set; } = null!;

        [Output("keys")]
        public Output<ImmutableArray<Outputs.KeyEntry>> Keys { get; private set; } = null!;

        [Output("permissions")]
        public Output<Pulumi.CephRadosgw.SubUserPermission> Permissions { get; private set; } = null!;

        [Output("purgeKeys")]
        public Output<bool?> PurgeKeys { get; private set; } = null!;

        [Output("secretKey")]
        public Output<string?> SecretKey { get; private set; } = null!;

        /// <summary>
        /// Name of this subuser
        /// </summary>
        [Output("subUserName")]
        public Output<string> SubUserName { get; private set; } = null!;

        /// <summary>
        /// User-ID of 'parent' user
        /// </summary>
        [Output("userId")]
        public Output<string> UserId { get; private set; } = null!;


        /// <summary>
        /// Create a SubUser resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public SubUser(string name, SubUserArgs args, CustomResourceOptions? options = null)
            : base("ceph-radosgw:index:SubUser", name, args ?? new SubUserArgs(), MakeResourceOptions(options, ""))
        {
        }

        private SubUser(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("ceph-radosgw:index:SubUser", name, null, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
                AdditionalSecretOutputs =
                {
                    "secretKey",
                },
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing SubUser resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static SubUser Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new SubUser(name, id, options);
        }
    }

    public sealed class SubUserArgs : global::Pulumi.ResourceArgs
    {
        [Input("generateKey")]
        public Input<bool>? GenerateKey { get; set; }

        [Input("keyType")]
        public Input<Pulumi.CephRadosgw.KeyType>? KeyType { get; set; }

        [Input("permissions", required: true)]
        public Input<Pulumi.CephRadosgw.SubUserPermission> Permissions { get; set; } = null!;

        [Input("purgeKeys")]
        public Input<bool>? PurgeKeys { get; set; }

        [Input("secretKey")]
        private Input<string>? _secretKey;
        public Input<string>? SecretKey
        {
            get => _secretKey;
            set
            {
                var emptySecret = Output.CreateSecret(0);
                _secretKey = Output.Tuple<Input<string>?, int>(value, emptySecret).Apply(t => t.Item1);
            }
        }

        /// <summary>
        /// Name of this subuser
        /// </summary>
        [Input("subUserName", required: true)]
        public Input<string> SubUserName { get; set; } = null!;

        /// <summary>
        /// User-ID of 'parent' user
        /// </summary>
        [Input("userId", required: true)]
        public Input<string> UserId { get; set; } = null!;

        public SubUserArgs()
        {
        }
        public static new SubUserArgs Empty => new SubUserArgs();
    }
}