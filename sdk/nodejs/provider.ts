// *** WARNING: this file was generated by pulumi-language-nodejs. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

export class Provider extends pulumi.ProviderResource {
    /** @internal */
    public static readonly __pulumiType = 'ceph-radosgw';

    /**
     * Returns true if the given object is an instance of Provider.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Provider {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === "pulumi:providers:" + Provider.__pulumiType;
    }

    /**
     * The username. It's important but not secret.
     */
    public readonly accessKeyID!: pulumi.Output<string>;
    /**
     * Assimilate an existing object during create
     */
    public readonly assimilate!: pulumi.Output<string | undefined>;
    /**
     * Delete assimilated objects during delete (otherwise they would be kept on OpenZiti)
     */
    public readonly deleteAssimilated!: pulumi.Output<string | undefined>;
    /**
     * The URI to the API
     */
    public readonly endpoint!: pulumi.Output<string>;
    /**
     * Don't validate server SSL certificate
     */
    public readonly insecure!: pulumi.Output<string | undefined>;
    /**
     * The password. It is very secret.
     */
    public readonly secretAccessKey!: pulumi.Output<string>;
    public readonly version!: pulumi.Output<string | undefined>;

    /**
     * Create a Provider resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: ProviderArgs, opts?: pulumi.ResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        {
            if ((!args || args.accessKeyID === undefined) && !opts.urn) {
                throw new Error("Missing required property 'accessKeyID'");
            }
            if ((!args || args.endpoint === undefined) && !opts.urn) {
                throw new Error("Missing required property 'endpoint'");
            }
            if ((!args || args.secretAccessKey === undefined) && !opts.urn) {
                throw new Error("Missing required property 'secretAccessKey'");
            }
            resourceInputs["accessKeyID"] = args ? args.accessKeyID : undefined;
            resourceInputs["assimilate"] = args ? args.assimilate : undefined;
            resourceInputs["deleteAssimilated"] = args ? args.deleteAssimilated : undefined;
            resourceInputs["endpoint"] = args ? args.endpoint : undefined;
            resourceInputs["insecure"] = args ? args.insecure : undefined;
            resourceInputs["secretAccessKey"] = args?.secretAccessKey ? pulumi.secret(args.secretAccessKey) : undefined;
            resourceInputs["version"] = args ? args.version : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        const secretOpts = { additionalSecretOutputs: ["secretAccessKey"] };
        opts = pulumi.mergeOptions(opts, secretOpts);
        super(Provider.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * The set of arguments for constructing a Provider resource.
 */
export interface ProviderArgs {
    /**
     * The username. It's important but not secret.
     */
    accessKeyID: pulumi.Input<string>;
    /**
     * Assimilate an existing object during create
     */
    assimilate?: pulumi.Input<string>;
    /**
     * Delete assimilated objects during delete (otherwise they would be kept on OpenZiti)
     */
    deleteAssimilated?: pulumi.Input<string>;
    /**
     * The URI to the API
     */
    endpoint: pulumi.Input<string>;
    /**
     * Don't validate server SSL certificate
     */
    insecure?: pulumi.Input<string>;
    /**
     * The password. It is very secret.
     */
    secretAccessKey: pulumi.Input<string>;
    version?: pulumi.Input<string>;
}
