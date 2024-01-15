// *** WARNING: this file was generated by pulumi. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.CephRadosgw.Outputs
{

    [OutputType]
    public sealed class KeyEntry
    {
        public readonly string? AccessKey;
        public readonly string KeyType;
        public readonly string SecretKey;

        [OutputConstructor]
        private KeyEntry(
            string? accessKey,

            string keyType,

            string secretKey)
        {
            AccessKey = accessKey;
            KeyType = keyType;
            SecretKey = secretKey;
        }
    }
}
