// *** WARNING: this file was generated by pulumi. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Immutable;

namespace Pulumi.CephRadosgw
{
    public static class Config
    {
        [global::System.Diagnostics.CodeAnalysis.SuppressMessage("Microsoft.Design", "IDE1006", Justification = 
        "Double underscore prefix used to avoid conflicts with variable names.")]
        private sealed class __Value<T>
        {
            private readonly Func<T> _getter;
            private T _value = default!;
            private bool _set;

            public __Value(Func<T> getter)
            {
                _getter = getter;
            }

            public T Get() => _set ? _value : _getter();

            public void Set(T value)
            {
                _value = value;
                _set = true;
            }
        }

        private static readonly global::Pulumi.Config __config = new global::Pulumi.Config("ceph-radosgw");

        private static readonly __Value<string?> _accessKeyID = new __Value<string?>(() => __config.Get("accessKeyID"));
        /// <summary>
        /// The username. It's important but not secret.
        /// </summary>
        public static string? AccessKeyID
        {
            get => _accessKeyID.Get();
            set => _accessKeyID.Set(value);
        }

        private static readonly __Value<string?> _assimilate = new __Value<string?>(() => __config.Get("assimilate"));
        /// <summary>
        /// Assimilate an existing object during create
        /// </summary>
        public static string? Assimilate
        {
            get => _assimilate.Get();
            set => _assimilate.Set(value);
        }

        private static readonly __Value<string?> _deleteAssimilated = new __Value<string?>(() => __config.Get("deleteAssimilated"));
        /// <summary>
        /// Delete assimilated objects during delete (otherwise they would be kept on OpenZiti)
        /// </summary>
        public static string? DeleteAssimilated
        {
            get => _deleteAssimilated.Get();
            set => _deleteAssimilated.Set(value);
        }

        private static readonly __Value<string?> _endpoint = new __Value<string?>(() => __config.Get("endpoint"));
        /// <summary>
        /// The URI to the API
        /// </summary>
        public static string? Endpoint
        {
            get => _endpoint.Get();
            set => _endpoint.Set(value);
        }

        private static readonly __Value<string?> _insecure = new __Value<string?>(() => __config.Get("insecure"));
        /// <summary>
        /// Don't validate server SSL certificate
        /// </summary>
        public static string? Insecure
        {
            get => _insecure.Get();
            set => _insecure.Set(value);
        }

        private static readonly __Value<string?> _secretAccessKey = new __Value<string?>(() => __config.Get("secretAccessKey"));
        /// <summary>
        /// The password. It is very secret.
        /// </summary>
        public static string? SecretAccessKey
        {
            get => _secretAccessKey.Get();
            set => _secretAccessKey.Set(value);
        }

        private static readonly __Value<string?> _version = new __Value<string?>(() => __config.Get("version"));
        public static string? Version
        {
            get => _version.Get();
            set => _version.Set(value);
        }

    }
}
