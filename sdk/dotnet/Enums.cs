// *** WARNING: this file was generated by pulumi. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.ComponentModel;
using Pulumi;

namespace Pulumi.CephRadosgw
{
    [EnumType]
    public readonly struct CapabilityPermission : IEquatable<CapabilityPermission>
    {
        private readonly string _value;

        private CapabilityPermission(string value)
        {
            _value = value ?? throw new ArgumentNullException(nameof(value));
        }

        public static CapabilityPermission Read { get; } = new CapabilityPermission("read");
        public static CapabilityPermission Write { get; } = new CapabilityPermission("write");
        public static CapabilityPermission Asterisk { get; } = new CapabilityPermission("*");

        public static bool operator ==(CapabilityPermission left, CapabilityPermission right) => left.Equals(right);
        public static bool operator !=(CapabilityPermission left, CapabilityPermission right) => !left.Equals(right);

        public static explicit operator string(CapabilityPermission value) => value._value;

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override bool Equals(object? obj) => obj is CapabilityPermission other && Equals(other);
        public bool Equals(CapabilityPermission other) => string.Equals(_value, other._value, StringComparison.Ordinal);

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override int GetHashCode() => _value?.GetHashCode() ?? 0;

        public override string ToString() => _value;
    }

    [EnumType]
    public readonly struct KeyType : IEquatable<KeyType>
    {
        private readonly string _value;

        private KeyType(string value)
        {
            _value = value ?? throw new ArgumentNullException(nameof(value));
        }

        public static KeyType S3 { get; } = new KeyType("s3");
        public static KeyType Swift { get; } = new KeyType("swift");

        public static bool operator ==(KeyType left, KeyType right) => left.Equals(right);
        public static bool operator !=(KeyType left, KeyType right) => !left.Equals(right);

        public static explicit operator string(KeyType value) => value._value;

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override bool Equals(object? obj) => obj is KeyType other && Equals(other);
        public bool Equals(KeyType other) => string.Equals(_value, other._value, StringComparison.Ordinal);

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override int GetHashCode() => _value?.GetHashCode() ?? 0;

        public override string ToString() => _value;
    }

    [EnumType]
    public readonly struct SubUserPermission : IEquatable<SubUserPermission>
    {
        private readonly string _value;

        private SubUserPermission(string value)
        {
            _value = value ?? throw new ArgumentNullException(nameof(value));
        }

        public static SubUserPermission None { get; } = new SubUserPermission("none");
        public static SubUserPermission Read { get; } = new SubUserPermission("read");
        public static SubUserPermission Write { get; } = new SubUserPermission("write");
        public static SubUserPermission ReadWrite { get; } = new SubUserPermission("readWrite");
        public static SubUserPermission FullControl { get; } = new SubUserPermission("fullControl");

        public static bool operator ==(SubUserPermission left, SubUserPermission right) => left.Equals(right);
        public static bool operator !=(SubUserPermission left, SubUserPermission right) => !left.Equals(right);

        public static explicit operator string(SubUserPermission value) => value._value;

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override bool Equals(object? obj) => obj is SubUserPermission other && Equals(other);
        public bool Equals(SubUserPermission other) => string.Equals(_value, other._value, StringComparison.Ordinal);

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override int GetHashCode() => _value?.GetHashCode() ?? 0;

        public override string ToString() => _value;
    }
}
