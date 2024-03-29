// *** WARNING: this file was generated by pulumi-language-nodejs. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

export class Key extends pulumi.CustomResource {
    /**
     * Get an existing Key resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): Key {
        return new Key(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'ceph-radosgw:index:Key';

    /**
     * Returns true if the given object is an instance of Key.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Key {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Key.__pulumiType;
    }

    public /*out*/ readonly _assimilated!: pulumi.Output<boolean>;
    /**
     * The access key
     */
    public readonly accessKey!: pulumi.Output<string | undefined>;
    /**
     * key type - either 's3' or 'swift'
     */
    public readonly keyType!: pulumi.Output<string | undefined>;
    /**
     * The secret key
     */
    public readonly secretKey!: pulumi.Output<string | undefined>;
    /**
     * Name of sub-user (optional)
     */
    public readonly subUserName!: pulumi.Output<string | undefined>;
    /**
     * User-ID of 'parent' user
     */
    public readonly userId!: pulumi.Output<string>;

    /**
     * Create a Key resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: KeyArgs, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (!opts.id) {
            if ((!args || args.userId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'userId'");
            }
            resourceInputs["accessKey"] = args ? args.accessKey : undefined;
            resourceInputs["keyType"] = args ? args.keyType : undefined;
            resourceInputs["secretKey"] = args ? args.secretKey : undefined;
            resourceInputs["subUserName"] = args ? args.subUserName : undefined;
            resourceInputs["userId"] = args ? args.userId : undefined;
            resourceInputs["_assimilated"] = undefined /*out*/;
        } else {
            resourceInputs["_assimilated"] = undefined /*out*/;
            resourceInputs["accessKey"] = undefined /*out*/;
            resourceInputs["keyType"] = undefined /*out*/;
            resourceInputs["secretKey"] = undefined /*out*/;
            resourceInputs["subUserName"] = undefined /*out*/;
            resourceInputs["userId"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(Key.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * The set of arguments for constructing a Key resource.
 */
export interface KeyArgs {
    /**
     * The access key
     */
    accessKey?: pulumi.Input<string>;
    /**
     * key type - either 's3' or 'swift'
     */
    keyType?: pulumi.Input<string>;
    /**
     * The secret key
     */
    secretKey?: pulumi.Input<string>;
    /**
     * Name of sub-user (optional)
     */
    subUserName?: pulumi.Input<string>;
    /**
     * User-ID of 'parent' user
     */
    userId: pulumi.Input<string>;
}
