// *** WARNING: this file was generated by pulumi. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.CephRadosgw
{
    [CephRadosgwResourceType("ceph-radosgw:index:Bucket")]
    public partial class Bucket : global::Pulumi.CustomResource
    {
        /// <summary>
        /// Bucket was 'assimilated' - managing an existing bucket.
        /// </summary>
        [Output("_assimilated")]
        public Output<bool> _assimilated { get; private set; } = null!;

        [Output("_location")]
        public Output<string> _location { get; private set; } = null!;

        /// <summary>
        /// Bucket name
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// Bucket object locking enabled
        /// </summary>
        [Output("objectLocking")]
        public Output<bool?> ObjectLocking { get; private set; } = null!;

        /// <summary>
        /// Purge bucket on delete.
        /// </summary>
        [Output("purgeOnDelete")]
        public Output<bool?> PurgeOnDelete { get; private set; } = null!;

        /// <summary>
        /// Bucket quota configuration
        /// </summary>
        [Output("quota")]
        public Output<Outputs.QuotaArgs?> Quota { get; private set; } = null!;

        /// <summary>
        /// The unique bucket id.
        /// </summary>
        [Output("ubid")]
        public Output<string> Ubid { get; private set; } = null!;

        [Output("versioning")]
        public Output<bool?> Versioning { get; private set; } = null!;


        /// <summary>
        /// Create a Bucket resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Bucket(string name, BucketArgs args, CustomResourceOptions? options = null)
            : base("ceph-radosgw:index:Bucket", name, args ?? new BucketArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Bucket(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("ceph-radosgw:index:Bucket", name, null, MakeResourceOptions(options, id))
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
        /// Get an existing Bucket resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Bucket Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new Bucket(name, id, options);
        }
    }

    public sealed class BucketArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Bucket name
        /// </summary>
        [Input("name", required: true)]
        public Input<string> Name { get; set; } = null!;

        /// <summary>
        /// Bucket object locking enabled
        /// </summary>
        [Input("objectLocking")]
        public Input<bool>? ObjectLocking { get; set; }

        /// <summary>
        /// Purge bucket on delete
        /// </summary>
        [Input("purgeOnDelete")]
        public Input<bool>? PurgeOnDelete { get; set; }

        /// <summary>
        /// Bucket quota configuration
        /// </summary>
        [Input("quota")]
        public Input<Inputs.QuotaArgsArgs>? Quota { get; set; }

        [Input("versioning")]
        public Input<bool>? Versioning { get; set; }

        public BucketArgs()
        {
        }
        public static new BucketArgs Empty => new BucketArgs();
    }
}