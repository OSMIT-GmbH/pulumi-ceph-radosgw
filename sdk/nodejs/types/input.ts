// *** WARNING: this file was generated by pulumi-language-nodejs. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as enums from "../types/enums";

export interface CapabilitiesArgs {
    amzCache?: pulumi.Input<enums.CapabilityPermission>;
    bilog?: pulumi.Input<enums.CapabilityPermission>;
    buckets?: pulumi.Input<enums.CapabilityPermission>;
    datalog?: pulumi.Input<enums.CapabilityPermission>;
    info?: pulumi.Input<enums.CapabilityPermission>;
    mdlog?: pulumi.Input<enums.CapabilityPermission>;
    metadata?: pulumi.Input<enums.CapabilityPermission>;
    oidcProvider?: pulumi.Input<enums.CapabilityPermission>;
    ratelimit?: pulumi.Input<enums.CapabilityPermission>;
    roles?: pulumi.Input<enums.CapabilityPermission>;
    usage?: pulumi.Input<enums.CapabilityPermission>;
    userPolicy?: pulumi.Input<enums.CapabilityPermission>;
    users?: pulumi.Input<enums.CapabilityPermission>;
    zone?: pulumi.Input<enums.CapabilityPermission>;
}

export interface QuotaArgsArgs {
    /**
     * Enable or disable quota enforcement quota
     */
    enabled?: pulumi.Input<boolean>;
    /**
     * Maximum object count - numeric value (alternate to MaxObjectsHum)
     */
    maxObjects?: pulumi.Input<number>;
    /**
     * Maximum object count - human readable format (i.e. 10k) (alternate to MaxObjects)
     */
    maxObjectsHum?: pulumi.Input<string>;
    /**
     * Maximum size - numeric value (alternate to MaxSizeHum)
     */
    maxSize?: pulumi.Input<number>;
    /**
     * Maximum size - human readable format (alternate to MaxSize)
     */
    maxSizeHum?: pulumi.Input<string>;
}
